// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IComplianceModule} from "../../interfaces/compliance/IComplianceModule.sol";
import {IDiamondCompliance} from "../../interfaces/compliance/IDiamondCompliance.sol";
import {LibReasonCodes} from "../../libraries/LibReasonCodes.sol";

bytes32 constant REASON_MAX_HOLDERS = keccak256("MAX_HOLDERS_EXCEEDED");

/**
 * @title MaxHoldersModule
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice Enforces a maximum number of unique holders per tokenId.
 *         Required by some jurisdictions that cap the number of investors
 *         in a given security offering (e.g., Reg D 506(b) = 35 non-accredited).
 * @dev Reads holderCount from the Diamond via IDiamondCompliance.
 *      A new holder is counted when receiver's current balance is zero.
 *      A limit of 0 means unlimited.
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract MaxHoldersModule is IComplianceModule {
    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event MaxHoldersSet(uint256 indexed tokenId, uint256 maxHolders);

    /*//////////////////////////////////////////////////////////////
                                ERRORS
    //////////////////////////////////////////////////////////////*/

    error MaxHoldersModule__OnlyOwner();
    error MaxHoldersModule__ZeroDiamond();

    /*//////////////////////////////////////////////////////////////
                            STATE
    //////////////////////////////////////////////////////////////*/

    address public immutable DIAMOND;
    address public owner;

    /// @dev tokenId → max unique holders (0 = unlimited)
    mapping(uint256 => uint256) public maxHolders;

    /*//////////////////////////////////////////////////////////////
                        CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    constructor(address diamond_, address owner_) {
        if (diamond_ == address(0)) revert MaxHoldersModule__ZeroDiamond();
        DIAMOND = diamond_;
        owner = owner_;
    }

    /*//////////////////////////////////////////////////////////////
                    EXTERNAL STATE-CHANGING
    //////////////////////////////////////////////////////////////*/

    /// @notice Sets the maximum number of unique holders for a tokenId
    /// @param tokenId The asset class
    /// @param limit Maximum holders (0 = unlimited)
    function setMaxHolders(uint256 tokenId, uint256 limit) external {
        _enforceOwner();
        maxHolders[tokenId] = limit;
        emit MaxHoldersSet(tokenId, limit);
    }

    /*//////////////////////////////////////////////////////////////
                    IComplianceModule — HOOKS
    //////////////////////////////////////////////////////////////*/

    function transferred(uint256, address, address, uint256) external {}

    function minted(uint256, address, uint256) external {}

    function burned(uint256, address, uint256) external {}

    /*//////////////////////////////////////////////////////////////
                    IComplianceModule — VALIDATION
    //////////////////////////////////////////////////////////////*/

    function canTransfer(uint256 tokenId, address, address to, uint256, bytes calldata)
        external
        view
        returns (bool ok, bytes32 reason)
    {
        uint256 limit = maxHolders[tokenId];
        if (limit == 0) return (true, LibReasonCodes.REASON_OK);

        IDiamondCompliance diamond = IDiamondCompliance(DIAMOND);

        // If receiver already holds tokens, no new holder is added
        uint256 receiverBalance = diamond.balanceOf(to, tokenId);
        if (receiverBalance > 0) return (true, LibReasonCodes.REASON_OK);

        // New holder would be added — check against limit
        uint256 current = diamond.holderCount(tokenId);
        if (current >= limit) {
            return (false, REASON_MAX_HOLDERS);
        }

        return (true, LibReasonCodes.REASON_OK);
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL
    //////////////////////////////////////////////////////////////*/

    function _enforceOwner() internal view {
        if (msg.sender != owner) revert MaxHoldersModule__OnlyOwner();
    }
}
