// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IComplianceModule} from "../../interfaces/compliance/IComplianceModule.sol";
import {IDiamondCompliance} from "../../interfaces/compliance/IDiamondCompliance.sol";
import {LibReasonCodes} from "../../libraries/LibReasonCodes.sol";

/**
 * @title CountryRestrictModule
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice Blocks transfers to/from wallets whose country is on the
 *         per-tokenId blocklist. Configurable by the module owner.
 * @dev Calls back to the Diamond via IDiamondCompliance.getCountry().
 *      Uses a blocklist pattern (deny-list) rather than allowlist to
 *      minimize gas for jurisdictions with few restrictions.
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract CountryRestrictModule is IComplianceModule {
    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event CountryRestricted(uint256 indexed tokenId, uint16 indexed country);
    event CountryUnrestricted(uint256 indexed tokenId, uint16 indexed country);

    /*//////////////////////////////////////////////////////////////
                                ERRORS
    //////////////////////////////////////////////////////////////*/

    error CountryRestrictModule__OnlyOwner();
    error CountryRestrictModule__ZeroDiamond();

    /*//////////////////////////////////////////////////////////////
                            STATE
    //////////////////////////////////////////////////////////////*/

    address public immutable DIAMOND;
    address public owner;

    /// @dev tokenId → country → restricted
    mapping(uint256 => mapping(uint16 => bool)) public isRestricted;

    /*//////////////////////////////////////////////////////////////
                        CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    constructor(address diamond_, address owner_) {
        if (diamond_ == address(0)) revert CountryRestrictModule__ZeroDiamond();
        DIAMOND = diamond_;
        owner = owner_;
    }

    /*//////////////////////////////////////////////////////////////
                    EXTERNAL STATE-CHANGING
    //////////////////////////////////////////////////////////////*/

    /// @notice Adds a country to the blocklist for a tokenId
    function restrictCountry(uint256 tokenId, uint16 country) external {
        _enforceOwner();
        isRestricted[tokenId][country] = true;
        emit CountryRestricted(tokenId, country);
    }

    /// @notice Removes a country from the blocklist for a tokenId
    function unrestrictCountry(uint256 tokenId, uint16 country) external {
        _enforceOwner();
        isRestricted[tokenId][country] = false;
        emit CountryUnrestricted(tokenId, country);
    }

    /// @notice Batch restrict countries for a tokenId
    function batchRestrictCountries(uint256 tokenId, uint16[] calldata countries) external {
        _enforceOwner();
        for (uint256 i; i < countries.length; ++i) {
            isRestricted[tokenId][countries[i]] = true;
            emit CountryRestricted(tokenId, countries[i]);
        }
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

    function canTransfer(uint256 tokenId, address from, address to, uint256, bytes calldata)
        external
        view
        returns (bool ok, bytes32 reason)
    {
        IDiamondCompliance diamond = IDiamondCompliance(DIAMOND);

        uint16 fromCountry = diamond.getCountry(from);
        if (fromCountry != 0 && isRestricted[tokenId][fromCountry]) {
            return (false, LibReasonCodes.REASON_COUNTRY_RESTRICTED);
        }

        uint16 toCountry = diamond.getCountry(to);
        if (toCountry != 0 && isRestricted[tokenId][toCountry]) {
            return (false, LibReasonCodes.REASON_COUNTRY_RESTRICTED);
        }

        return (true, LibReasonCodes.REASON_OK);
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL
    //////////////////////////////////////////////////////////////*/

    function _enforceOwner() internal view {
        if (msg.sender != owner) revert CountryRestrictModule__OnlyOwner();
    }
}
