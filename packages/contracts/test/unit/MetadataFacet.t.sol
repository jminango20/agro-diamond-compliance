// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {MetadataFacet} from "../../src/facets/token/MetadataFacet.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {PauseFacet} from "../../src/facets/security/PauseFacet.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";

contract MetadataFacetTest is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal alice = makeAddr("alice");

    MetadataFacet internal meta;
    SupplyFacet internal supply;
    AssetManagerFacet internal am;
    AccessControlFacet internal ac;
    PauseFacet internal pause;

    uint256 internal TOKEN_1;
    uint256 internal TOKEN_2;

    address[] internal emptyModules;

    function setUp() public {
        d = deployDiamond(owner);
        meta = MetadataFacet(address(d.diamond));
        supply = SupplyFacet(address(d.diamond));
        am = AssetManagerFacet(address(d.diamond));
        ac = AccessControlFacet(address(d.diamond));
        pause = PauseFacet(address(d.diamond));

        uint16[] memory countries = new uint16[](2);
        countries[0] = 76;  // BR
        countries[1] = 840; // US

        vm.startPrank(owner);
        TOKEN_1 = am.registerAsset(
            IAssetManager.RegisterAssetParams({
                name: "Fundo Imobiliario XYZ",
                symbol: "FXYZ",
                uri: "https://api.example.com/metadata/1",
                supplyCap: 1_000_000,
                identityProfileId: 0,
                complianceModules: emptyModules,
                issuer: owner,
                allowedCountries: countries
            })
        );
        TOKEN_2 = am.registerAsset(
            IAssetManager.RegisterAssetParams({
                name: "Debenture ABC",
                symbol: "DABC",
                uri: "",
                supplyCap: 0,
                identityProfileId: 0,
                complianceModules: emptyModules,
                issuer: owner,
                allowedCountries: new uint16[](0)
            })
        );
        ac.grantRole(keccak256("ISSUER_ROLE"), owner);
        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                            URI
    //////////////////////////////////////////////////////////////*/

    function test_Uri() public view {
        assertEq(meta.uri(TOKEN_1), "https://api.example.com/metadata/1");
    }

    function test_Uri_EmptyForNoUri() public view {
        assertEq(meta.uri(TOKEN_2), "");
    }

    function test_Uri_EmptyForUnregistered() public view {
        assertEq(meta.uri(999), "");
    }

    function test_Uri_UpdatedViaAssetManager() public {
        vm.prank(owner);
        am.setAssetUri(TOKEN_1, "ipfs://Qm123");
        assertEq(meta.uri(TOKEN_1), "ipfs://Qm123");
    }

    /*//////////////////////////////////////////////////////////////
                        NAME & SYMBOL
    //////////////////////////////////////////////////////////////*/

    function test_Name() public view {
        assertEq(meta.name(TOKEN_1), "Fundo Imobiliario XYZ");
    }

    function test_Symbol() public view {
        assertEq(meta.symbol(TOKEN_1), "FXYZ");
    }

    function test_Name_Token2() public view {
        assertEq(meta.name(TOKEN_2), "Debenture ABC");
    }

    function test_Symbol_Token2() public view {
        assertEq(meta.symbol(TOKEN_2), "DABC");
    }

    /*//////////////////////////////////////////////////////////////
                        SUPPLY CAP & ISSUER
    //////////////////////////////////////////////////////////////*/

    function test_SupplyCap() public view {
        assertEq(meta.supplyCap(TOKEN_1), 1_000_000);
    }

    function test_SupplyCap_Unlimited() public view {
        assertEq(meta.supplyCap(TOKEN_2), 0);
    }

    function test_Issuer() public view {
        assertEq(meta.issuer(TOKEN_1), owner);
    }

    /*//////////////////////////////////////////////////////////////
                    ALLOWED COUNTRIES
    //////////////////////////////////////////////////////////////*/

    function test_AllowedCountries() public view {
        uint16[] memory countries = meta.allowedCountries(TOKEN_1);
        assertEq(countries.length, 2);
        assertEq(countries[0], 76);
        assertEq(countries[1], 840);
    }

    function test_AllowedCountries_Empty() public view {
        uint16[] memory countries = meta.allowedCountries(TOKEN_2);
        assertEq(countries.length, 0);
    }

    /*//////////////////////////////////////////////////////////////
                        TOKEN INFO
    //////////////////////////////////////////////////////////////*/

    function test_TokenInfo() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 5000);

        (
            string memory name_,
            string memory symbol_,
            string memory uri_,
            uint256 totalSupply_,
            uint256 supplyCap_,
            uint256 holderCount_,
            address issuer_,
            bool paused_
        ) = meta.tokenInfo(TOKEN_1);

        assertEq(name_, "Fundo Imobiliario XYZ");
        assertEq(symbol_, "FXYZ");
        assertEq(uri_, "https://api.example.com/metadata/1");
        assertEq(totalSupply_, 5000);
        assertEq(supplyCap_, 1_000_000);
        assertEq(holderCount_, 1);
        assertEq(issuer_, owner);
        assertFalse(paused_);
    }

    function test_TokenInfo_Paused() public {
        vm.prank(owner);
        pause.pauseAsset(TOKEN_1);

        (,,,,,,, bool paused_) = meta.tokenInfo(TOKEN_1);
        assertTrue(paused_);
    }

    function test_TokenInfo_Unregistered() public view {
        (
            string memory name_,
            string memory symbol_,
            string memory uri_,
            uint256 totalSupply_,
            uint256 supplyCap_,
            uint256 holderCount_,
            address issuer_,
            bool paused_
        ) = meta.tokenInfo(999);

        assertEq(bytes(name_).length, 0);
        assertEq(bytes(symbol_).length, 0);
        assertEq(bytes(uri_).length, 0);
        assertEq(totalSupply_, 0);
        assertEq(supplyCap_, 0);
        assertEq(holderCount_, 0);
        assertEq(issuer_, address(0));
        assertFalse(paused_);
    }
}
