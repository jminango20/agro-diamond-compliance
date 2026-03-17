// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IHookablePlugin} from "../../../interfaces/plugins/IHookablePlugin.sol";
import {IDiamondCompliance} from "../../../interfaces/compliance/IDiamondCompliance.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC1155} from "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import {IERC1155Receiver} from "@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title YieldDistributorModule
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice Distributes real yield to token holders proportionally using the
 *         Synthetix/MasterChef accumulator pattern. O(1) per operation.
 *
 *         Supports three reward asset types:
 *         - **ERC-20** — USDC, WETH, any standard token
 *         - **ERC-1155 (own Diamond)** — another tokenId from the same protocol
 *         - **ERC-1155 (external)** — tokens from any external ERC-1155 contract
 *
 *         Flow:
 *         1. Admin registers a reward asset for a staked tokenId
 *         2. Admin deposits yield via depositYield()
 *         3. On every transfer/mint/burn, hooks update user checkpoints
 *         4. Holders call claimYield() to withdraw accumulated rewards
 *
 * @dev Accumulator math:
 *      accRewardPerShare += depositAmount * PRECISION / totalSupply
 *      pending(user) = balance(user) * accRewardPerShare / PRECISION - rewardDebt(user) + pendingRewards(user)
 *
 *      PRECISION is 1e36 to handle tokens with low decimals (e.g. USDC 6 decimals)
 *      without precision loss. Works equally well for ERC-1155 amounts (no decimals).
 *
 *      Uses SafeERC20 for ERC-20 compatibility with non-standard tokens (USDT, etc.).
 *      Uses IERC1155.safeTransferFrom for ERC-1155 assets.
 *
 *      Hooks fire AFTER balance mutation, so we reconstruct pre-mutation balances:
 *      - onTransfer: from had (currentBalance + amount), to had (currentBalance - amount)
 *      - onMint: to had (currentBalance - amount)
 *      - onBurn: from had (currentBalance + amount)
 *
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract YieldDistributorModule is
    IHookablePlugin,
    IERC1155Receiver,
    ReentrancyGuard
{
    using SafeERC20 for IERC20;

    /*//////////////////////////////////////////////////////////////
                            TYPES
    //////////////////////////////////////////////////////////////*/

    enum RewardType {
        ERC20,
        ERC1155
    }

    /// @param token The ERC-20 address or ERC-1155 contract address
    /// @param id Ignored for ERC-20; the tokenId for ERC-1155
    /// @param assetType Whether this is an ERC-20 or ERC-1155 reward
    struct RewardAsset {
        address token;
        uint256 id;
        RewardType assetType;
    }

    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event RewardAssetAdded(
        uint256 indexed tokenId,
        address indexed token,
        uint256 id,
        RewardType assetType
    );
    event RewardAssetRemoved(
        uint256 indexed tokenId,
        address indexed token,
        uint256 id,
        RewardType assetType
    );
    event YieldDeposited(
        uint256 indexed tokenId,
        bytes32 indexed rewardKey,
        uint256 amount
    );
    event YieldClaimed(
        uint256 indexed tokenId,
        bytes32 indexed rewardKey,
        address indexed holder,
        uint256 amount
    );

    /*//////////////////////////////////////////////////////////////
                                ERRORS
    //////////////////////////////////////////////////////////////*/

    error YieldDistributorModule__OnlyOwner();
    error YieldDistributorModule__OnlyDiamond();
    error YieldDistributorModule__ZeroDiamond();
    error YieldDistributorModule__ZeroAddress();
    error YieldDistributorModule__ZeroAmount();
    error YieldDistributorModule__ZeroSupply(uint256 tokenId);
    error YieldDistributorModule__RewardAssetAlreadyAdded(
        uint256 tokenId,
        bytes32 rewardKey
    );
    error YieldDistributorModule__RewardAssetNotFound(
        uint256 tokenId,
        bytes32 rewardKey
    );
    error YieldDistributorModule__NothingToClaim();
    error YieldDistributorModule__TooManyRewardAssets();
    error YieldDistributorModule__InvalidERC20Id();

    /*//////////////////////////////////////////////////////////////
                            CONSTANTS
    //////////////////////////////////////////////////////////////*/

    /// @dev High precision scalar to avoid truncation with low-decimal tokens (USDC=6, WBTC=8).
    uint256 internal constant PRECISION = 1e36;
    uint256 internal constant MAX_REWARD_ASSETS = 5;

    /*//////////////////////////////////////////////////////////////
                            STATE
    //////////////////////////////////////////////////////////////*/

    address public immutable DIAMOND;
    address public owner;

    /// @dev tokenId → list of reward assets (max MAX_REWARD_ASSETS)
    mapping(uint256 => RewardAsset[]) internal _rewardAssetsList;

    /// @dev tokenId → rewardKey → is registered
    mapping(uint256 => mapping(bytes32 => bool)) public isRewardAsset;

    /// @dev tokenId → rewardKey → accumulated reward per share (scaled by PRECISION)
    mapping(uint256 => mapping(bytes32 => uint256)) public accRewardPerShare;

    /// @dev tokenId → rewardKey → user → reward debt (scaled by PRECISION, divided on claim)
    mapping(uint256 => mapping(bytes32 => mapping(address => uint256)))
        public rewardDebt;

    /// @dev tokenId → rewardKey → user → crystallized pending rewards (in raw token units)
    mapping(uint256 => mapping(bytes32 => mapping(address => uint256)))
        public pendingRewards;

    /*//////////////////////////////////////////////////////////////
                        CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    constructor(address diamond_, address owner_) {
        if (diamond_ == address(0))
            revert YieldDistributorModule__ZeroDiamond();
        DIAMOND = diamond_;
        owner = owner_;
    }

    /*//////////////////////////////////////////////////////////////
                    ADMIN — REWARD ASSET MANAGEMENT
    //////////////////////////////////////////////////////////////*/

    /// @notice Registers a reward asset for a staked tokenId
    /// @param tokenId The staked asset class whose holders receive yield
    /// @param asset The reward asset (ERC-20 or ERC-1155)
    function addRewardAsset(
        uint256 tokenId,
        RewardAsset calldata asset
    ) external {
        _enforceOwner();
        if (asset.token == address(0))
            revert YieldDistributorModule__ZeroAddress();
        if (asset.assetType == RewardType.ERC20 && asset.id != 0) {
            revert YieldDistributorModule__InvalidERC20Id();
        }

        bytes32 key = _rewardKey(asset.token, asset.id, asset.assetType);
        if (isRewardAsset[tokenId][key]) {
            revert YieldDistributorModule__RewardAssetAlreadyAdded(
                tokenId,
                key
            );
        }
        if (_rewardAssetsList[tokenId].length >= MAX_REWARD_ASSETS) {
            revert YieldDistributorModule__TooManyRewardAssets();
        }

        isRewardAsset[tokenId][key] = true;
        _rewardAssetsList[tokenId].push(asset);

        emit RewardAssetAdded(tokenId, asset.token, asset.id, asset.assetType);
    }

    /// @notice Removes a reward asset from a staked tokenId
    /// @param tokenId The staked asset class
    /// @param asset The reward asset to remove
    function removeRewardAsset(
        uint256 tokenId,
        RewardAsset calldata asset
    ) external {
        _enforceOwner();
        bytes32 key = _rewardKey(asset.token, asset.id, asset.assetType);
        if (!isRewardAsset[tokenId][key]) {
            revert YieldDistributorModule__RewardAssetNotFound(tokenId, key);
        }

        isRewardAsset[tokenId][key] = false;

        RewardAsset[] storage assets = _rewardAssetsList[tokenId];
        uint256 len = assets.length;
        for (uint256 i; i < len; ) {
            if (
                _rewardKey(
                    assets[i].token,
                    assets[i].id,
                    assets[i].assetType
                ) == key
            ) {
                assets[i] = assets[len - 1];
                assets.pop();
                break;
            }
            unchecked {
                ++i;
            }
        }

        emit RewardAssetRemoved(
            tokenId,
            asset.token,
            asset.id,
            asset.assetType
        );
    }

    /*//////////////////////////////////////////////////////////////
                    ADMIN — YIELD DEPOSIT
    //////////////////////////////////////////////////////////////*/

    /// @notice Deposits yield to be distributed proportionally to holders
    /// @dev Caller must have approved this contract for `amount` of the reward asset.
    ///      For ERC-20: `IERC20.approve(thisModule, amount)`
    ///      For ERC-1155: `IERC1155.setApprovalForAll(thisModule, true)`
    ///      Reverts if totalSupply is 0 (no holders to distribute to).
    /// @param tokenId The staked asset class whose holders receive the yield
    /// @param asset The reward asset being distributed
    /// @param amount Amount of reward tokens to distribute (raw units)
    function depositYield(
        uint256 tokenId,
        RewardAsset calldata asset,
        uint256 amount
    ) external nonReentrant {
        _enforceOwner();
        if (amount == 0) revert YieldDistributorModule__ZeroAmount();

        bytes32 key = _rewardKey(asset.token, asset.id, asset.assetType);
        if (!isRewardAsset[tokenId][key]) {
            revert YieldDistributorModule__RewardAssetNotFound(tokenId, key);
        }

        uint256 supply = IDiamondCompliance(DIAMOND).totalSupply(tokenId);
        if (supply == 0) revert YieldDistributorModule__ZeroSupply(tokenId);

        _pullReward(asset, msg.sender, amount);

        accRewardPerShare[tokenId][key] += (amount * PRECISION) / supply;

        emit YieldDeposited(tokenId, key, amount);
    }

    /*//////////////////////////////////////////////////////////////
                    HOLDER — CLAIM
    //////////////////////////////////////////////////////////////*/

    /// @notice Claims accumulated yield for a specific reward asset
    /// @param tokenId The staked asset class
    /// @param asset The reward asset to claim
    function claimYield(
        uint256 tokenId,
        RewardAsset calldata asset
    ) external nonReentrant {
        bytes32 key = _rewardKey(asset.token, asset.id, asset.assetType);
        uint256 claimable = _computePending(tokenId, key, msg.sender);
        if (claimable == 0) revert YieldDistributorModule__NothingToClaim();

        // Reset pending and update debt BEFORE transfer (CEI pattern)
        pendingRewards[tokenId][key][msg.sender] = 0;
        uint256 balance = IDiamondCompliance(DIAMOND).balanceOf(
            msg.sender,
            tokenId
        );
        rewardDebt[tokenId][key][msg.sender] =
            (balance * accRewardPerShare[tokenId][key]) /
            PRECISION;

        _pushReward(asset, msg.sender, claimable);

        emit YieldClaimed(tokenId, key, msg.sender, claimable);
    }

    /// @notice Claims all accumulated yield across all reward assets for a tokenId
    /// @param tokenId The staked asset class
    function claimAllYield(uint256 tokenId) external nonReentrant {
        RewardAsset[] storage assets = _rewardAssetsList[tokenId];
        uint256 len = assets.length;
        uint256 balance = IDiamondCompliance(DIAMOND).balanceOf(
            msg.sender,
            tokenId
        );

        for (uint256 i; i < len; ) {
            RewardAsset storage asset = assets[i];
            bytes32 key = _rewardKey(asset.token, asset.id, asset.assetType);
            uint256 claimable = _computePending(tokenId, key, msg.sender);

            if (claimable > 0) {
                // Reset BEFORE transfer (CEI pattern)
                pendingRewards[tokenId][key][msg.sender] = 0;
                rewardDebt[tokenId][key][msg.sender] =
                    (balance * accRewardPerShare[tokenId][key]) /
                    PRECISION;

                _pushReward(asset, msg.sender, claimable);

                emit YieldClaimed(tokenId, key, msg.sender, claimable);
            }
            unchecked {
                ++i;
            }
        }
    }

    /*//////////////////////////////////////////////////////////////
                    IHookablePlugin — HOOK
    //////////////////////////////////////////////////////////////*/

    /// @inheritdoc IHookablePlugin
    function onAction(ActionParams calldata params) external {
        _enforceDiamond();
        if (_rewardAssetsList[params.tokenId].length == 0) return;

        if (params.actionType == ActionType.Transfer) {
            _onTransfer(params.tokenId, params.from, params.to, params.amount);
        } else if (params.actionType == ActionType.Mint) {
            _onMint(params.tokenId, params.to, params.amount);
        } else {
            _onBurn(params.tokenId, params.from, params.amount);
        }
    }

    /// @notice Returns the human-readable name of this plugin
    function name() external pure returns (string memory) {
        return "YieldDistributor";
    }

    /*//////////////////////////////////////////////////////////////
                    IERC1155Receiver — REQUIRED TO HOLD ERC-1155
    //////////////////////////////////////////////////////////////*/

    /// @inheritdoc IERC1155Receiver
    function onERC1155Received(
        address,
        address,
        uint256,
        uint256,
        bytes calldata
    ) external pure returns (bytes4) {
        return IERC1155Receiver.onERC1155Received.selector;
    }

    /// @inheritdoc IERC1155Receiver
    function onERC1155BatchReceived(
        address,
        address,
        uint256[] calldata,
        uint256[] calldata,
        bytes calldata
    ) external pure returns (bytes4) {
        return IERC1155Receiver.onERC1155BatchReceived.selector;
    }

    /// @inheritdoc IERC165
    function supportsInterface(
        bytes4 interfaceId
    ) external pure returns (bool) {
        return
            interfaceId == type(IERC1155Receiver).interfaceId ||
            interfaceId == type(IERC165).interfaceId;
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — HOOK HANDLERS
    //////////////////////////////////////////////////////////////*/

    function _onTransfer(
        uint256 tokenId,
        address from,
        address to,
        uint256 amount
    ) internal {
        IDiamondCompliance diamond = IDiamondCompliance(DIAMOND);
        uint256 fromBalance = diamond.balanceOf(from, tokenId);
        uint256 toBalance = diamond.balanceOf(to, tokenId);

        RewardAsset[] storage assets = _rewardAssetsList[tokenId];
        uint256 len = assets.length;
        for (uint256 i; i < len; ) {
            bytes32 key = _rewardKey(
                assets[i].token,
                assets[i].id,
                assets[i].assetType
            );
            _crystallize(tokenId, key, from, fromBalance + amount, fromBalance);
            _crystallize(
                tokenId,
                key,
                to,
                toBalance >= amount ? toBalance - amount : 0,
                toBalance
            );
            unchecked {
                ++i;
            }
        }
    }

    function _onMint(uint256 tokenId, address to, uint256 amount) internal {
        uint256 toBalance = IDiamondCompliance(DIAMOND).balanceOf(to, tokenId);
        uint256 preBal = toBalance >= amount ? toBalance - amount : 0;

        RewardAsset[] storage assets = _rewardAssetsList[tokenId];
        uint256 len = assets.length;
        for (uint256 i; i < len; ) {
            bytes32 key = _rewardKey(
                assets[i].token,
                assets[i].id,
                assets[i].assetType
            );
            _crystallize(tokenId, key, to, preBal, toBalance);
            unchecked {
                ++i;
            }
        }
    }

    function _onBurn(uint256 tokenId, address from, uint256 amount) internal {
        uint256 fromBalance = IDiamondCompliance(DIAMOND).balanceOf(
            from,
            tokenId
        );

        RewardAsset[] storage assets = _rewardAssetsList[tokenId];
        uint256 len = assets.length;
        for (uint256 i; i < len; ) {
            bytes32 key = _rewardKey(
                assets[i].token,
                assets[i].id,
                assets[i].assetType
            );
            _crystallize(tokenId, key, from, fromBalance + amount, fromBalance);
            unchecked {
                ++i;
            }
        }
    }

    /// @dev Crystallizes pending rewards for `user` and updates rewardDebt.
    /// @param preBalance The balance BEFORE the mutation
    /// @param postBalance The balance AFTER the mutation (current on-chain balance)
    function _crystallize(
        uint256 tokenId,
        bytes32 key,
        address user,
        uint256 preBalance,
        uint256 postBalance
    ) internal {
        uint256 acc = accRewardPerShare[tokenId][key];
        if (preBalance > 0) {
            pendingRewards[tokenId][key][user] +=
                (preBalance * acc) /
                PRECISION -
                rewardDebt[tokenId][key][user];
        }
        rewardDebt[tokenId][key][user] = (postBalance * acc) / PRECISION;
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — REWARD TRANSFERS
    //////////////////////////////////////////////////////////////*/

    /// @dev Pulls reward tokens from `from` into this contract
    function _pullReward(
        RewardAsset calldata asset,
        address from,
        uint256 amount
    ) internal {
        if (asset.assetType == RewardType.ERC20) {
            IERC20(asset.token).safeTransferFrom(from, address(this), amount);
        } else {
            IERC1155(asset.token).safeTransferFrom(
                from,
                address(this),
                asset.id,
                amount,
                ""
            );
        }
    }

    /// @dev Pushes reward tokens from this contract to `to`
    function _pushReward(
        RewardAsset storage asset,
        address to,
        uint256 amount
    ) internal {
        if (asset.assetType == RewardType.ERC20) {
            IERC20(asset.token).safeTransfer(to, amount);
        } else {
            IERC1155(asset.token).safeTransferFrom(
                address(this),
                to,
                asset.id,
                amount,
                ""
            );
        }
    }

    /// @dev Overload for calldata parameter (used in claimYield with single asset)
    function _pushReward(
        RewardAsset calldata asset,
        address to,
        uint256 amount
    ) internal {
        if (asset.assetType == RewardType.ERC20) {
            IERC20(asset.token).safeTransfer(to, amount);
        } else {
            IERC1155(asset.token).safeTransferFrom(
                address(this),
                to,
                asset.id,
                amount,
                ""
            );
        }
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL VIEWS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the claimable yield for a holder
    /// @param tokenId The staked asset class
    /// @param asset The reward asset
    /// @param holder The holder address
    /// @return claimable Amount of reward tokens claimable (raw units)
    function claimableYield(
        uint256 tokenId,
        RewardAsset calldata asset,
        address holder
    ) external view returns (uint256 claimable) {
        claimable = _computePending(
            tokenId,
            _rewardKey(asset.token, asset.id, asset.assetType),
            holder
        );
    }

    /// @notice Returns all registered reward assets for a tokenId
    /// @param tokenId The staked asset class
    /// @return assets Array of reward assets
    function getRewardAssets(
        uint256 tokenId
    ) external view returns (RewardAsset[] memory assets) {
        assets = _rewardAssetsList[tokenId];
    }

    /// @notice Returns the reward key for a given asset (useful for off-chain queries)
    /// @param asset The reward asset
    /// @return key The bytes32 key used in storage mappings
    function rewardKey(
        RewardAsset calldata asset
    ) external pure returns (bytes32 key) {
        key = _rewardKey(asset.token, asset.id, asset.assetType);
    }

    /*//////////////////////////////////////////////////////////////
                        INTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    function _computePending(
        uint256 tokenId,
        bytes32 key,
        address holder
    ) internal view returns (uint256) {
        uint256 balance = IDiamondCompliance(DIAMOND).balanceOf(
            holder,
            tokenId
        );
        uint256 acc = accRewardPerShare[tokenId][key];
        uint256 accumulated = (balance * acc) / PRECISION;
        uint256 debt = rewardDebt[tokenId][key][holder];
        uint256 pending = pendingRewards[tokenId][key][holder];

        if (accumulated >= debt) {
            return accumulated - debt + pending;
        }
        return pending;
    }

    function _rewardKey(
        address token,
        uint256 id,
        RewardType assetType
    ) internal pure returns (bytes32) {
        return keccak256(abi.encode(token, id, assetType));
    }

    function _enforceOwner() internal view {
        if (msg.sender != owner) revert YieldDistributorModule__OnlyOwner();
    }

    function _enforceDiamond() internal view {
        if (msg.sender != DIAMOND) revert YieldDistributorModule__OnlyDiamond();
    }
}
