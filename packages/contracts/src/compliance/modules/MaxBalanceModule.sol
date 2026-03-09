// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IComplianceModule} from "../../interfaces/compliance/IComplianceModule.sol";
import {IDiamondCompliance} from "../../interfaces/compliance/IDiamondCompliance.sol";
import {LibReasonCodes} from "../../libraries/LibReasonCodes.sol";

/**
 * @title MaxBalanceModule
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice Enforces a maximum holding limit per wallet per tokenId.
 *         Prevents concentration risk by rejecting transfers that would
 *         push the receiver's total balance above the configured cap.
 * @dev Total balance = free + locked + custody + pendingSettlement.
 *      A limit of 0 means unlimited (no cap enforced).
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract MaxBalanceModule is IComplianceModule {
    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event MaxBalanceSet(uint256 indexed tokenId, uint256 maxBalance);

    /*//////////////////////////////////////////////////////////////
                                ERRORS
    //////////////////////////////////////////////////////////////*/

    error MaxBalanceModule__OnlyOwner();
    error MaxBalanceModule__ZeroDiamond();

    /*//////////////////////////////////////////////////////////////
                            STATE
    //////////////////////////////////////////////////////////////*/

    address public immutable DIAMOND;
    address public owner;

    /// @dev tokenId → max balance per wallet (0 = unlimited)
    mapping(uint256 => uint256) public maxBalance;

    /*//////////////////////////////////////////////////////////////
                        CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    constructor(address diamond_, address owner_) {
        if (diamond_ == address(0)) revert MaxBalanceModule__ZeroDiamond();
        DIAMOND = diamond_;
        owner = owner_;
    }

    /*//////////////////////////////////////////////////////////////
                    EXTERNAL STATE-CHANGING
    //////////////////////////////////////////////////////////////*/

    /// @notice Sets the maximum balance a single wallet can hold for a tokenId
    /// @param tokenId The asset class
    /// @param limit Maximum total balance (0 = unlimited)
    function setMaxBalance(uint256 tokenId, uint256 limit) external {
        _enforceOwner();
        maxBalance[tokenId] = limit;
        emit MaxBalanceSet(tokenId, limit);
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

    function canTransfer(uint256 tokenId, address, address to, uint256 amount, bytes calldata)
        external
        view
        returns (bool ok, bytes32 reason)
    {
        uint256 limit = maxBalance[tokenId];
        if (limit == 0) return (true, LibReasonCodes.REASON_OK);

        IDiamondCompliance diamond = IDiamondCompliance(DIAMOND);
        (uint256 free, uint256 locked, uint256 custody, uint256 pending) = diamond.partitionBalanceOf(to, tokenId);
        uint256 totalAfter = free + locked + custody + pending + amount;

        if (totalAfter > limit) {
            return (false, LibReasonCodes.REASON_HOLDING_LIMIT);
        }

        return (true, LibReasonCodes.REASON_OK);
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL
    //////////////////////////////////////////////////////////////*/

    function _enforceOwner() internal view {
        if (msg.sender != owner) revert MaxBalanceModule__OnlyOwner();
    }
}
