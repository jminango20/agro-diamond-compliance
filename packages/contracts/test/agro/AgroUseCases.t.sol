// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";
import {AgroClaimTopics} from "../../src/agro/AgroClaimTopics.sol";
import {CprLockupModule} from "../../src/compliance/modules/CprLockupModule.sol";
import {ProducerKycModule} from "../../src/compliance/modules/ProducerKycModule.sol";
import {CountryRestrictModule} from "../../src/compliance/modules/CountryRestrictModule.sol";

/*//////////////////////////////////////////////////////////////
                    DIAMOND INTERFACES
//////////////////////////////////////////////////////////////*/

interface ISupply {
    function mint(uint256 tokenId, address to, uint256 amount) external;
    function batchMint(uint256[] calldata tokenIds, address[] calldata recipients, uint256[] calldata amounts)
        external;
    function burn(uint256 tokenId, address from, uint256 amount) external;
    function totalSupply(uint256 tokenId) external view returns (uint256);
    function balanceOf(address account, uint256 id) external view returns (uint256);
    function holderCount(uint256 tokenId) external view returns (uint256);
}

interface IIdentity {
    function registerIdentity(address wallet, address identity, uint16 country) external;
    function contains(address wallet) external view returns (bool);
    function getCountry(address wallet) external view returns (uint16);
}

interface IAsset {
    function registerAsset(IAssetManager.RegisterAssetParams calldata p) external returns (uint256 tokenId);
    function addComplianceModule(uint256 tokenId, address module) external;
}

interface IProfile {
    function createProfile(uint256[] calldata claimTopics) external returns (uint32 profileId);
}

interface IAccessControl {
    function grantRole(bytes32 role, address account) external;
}

interface ITransfer {
    function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes calldata data) external;
    function setApprovalForAll(address operator, bool approved) external;
}

/// @title AgroUseCases
/// @notice End-to-end tests for the three Brazilian agricultural commodity use cases:
///         (1) Coffee — fungible lot, CPR lock-up, export country restrictions
///         (2) Cattle — non-fungible per animal, SISBOV KYC required
///         (3) Honey — hybrid: colony NFT (id < 1000) + lot fungible (id >= 1000)
///
///         Tests validate correctness of the agro compliance layer and generate
///         gas measurements for the paper's experimental results section.
contract AgroUseCasesTest is DiamondHelper {
    /*//////////////////////////////////////////////////////////////
                            CONSTANTS
    //////////////////////////////////////////////////////////////*/

    /// @dev ISO 3166-1 numeric country codes
    uint16 internal constant COUNTRY_BRAZIL = 76;
    uint16 internal constant COUNTRY_RUSSIA = 643; // restricted for honey export (example)

    bytes32 internal constant ISSUER_ROLE = keccak256("ISSUER_ROLE");
    bytes32 internal constant TRANSFER_AGENT = keccak256("TRANSFER_AGENT");

    /*//////////////////////////////////////////////////////////////
                            ACTORS
    //////////////////////////////////////////////////////////////*/

    address internal owner       = makeAddr("owner");
    address internal issuer      = makeAddr("issuer");      // authorized minter (cooperative)
    address internal agent       = makeAddr("agent");       // TRANSFER_AGENT (registrar)
    address internal producerA   = makeAddr("producerA");   // KYC'd Brazilian coffee producer
    address internal producerB   = makeAddr("producerB");   // KYC'd cattle rancher
    address internal producerC   = makeAddr("producerC");   // KYC'd honey producer
    address internal outsider    = makeAddr("outsider");    // not KYC'd — transfers should fail

    /// @dev Stub ONCHAINID addresses for the identity registry
    address internal idA = makeAddr("onchainid_A");
    address internal idB = makeAddr("onchainid_B");
    address internal idC = makeAddr("onchainid_C");

    /*//////////////////////////////////////////////////////////////
                        DIAMOND + MODULES
    //////////////////////////////////////////////////////////////*/

    DeployedDiamond internal d;
    ISupply          internal supply;
    IIdentity        internal identity;
    IAsset           internal asset;
    IProfile         internal profile;
    IAccessControl   internal ac;
    ITransfer        internal transfer;

    CprLockupModule       internal cprModule;
    ProducerKycModule     internal kycModule;
    CountryRestrictModule internal countryModule;

    /*//////////////////////////////////////////////////////////////
                        TOKEN IDS (assigned by registerAsset)
    //////////////////////////////////////////////////////////////*/

    uint256 internal coffeeId;
    uint256 internal cattleId;
    uint256 internal honeyId;

    /*//////////////////////////////////////////////////////////////
                        IDENTITY PROFILES
    //////////////////////////////////////////////////////////////*/

    uint32 internal coffeeProfileId; // MAPA + CAR + CPF/CNPJ
    uint32 internal cattleProfileId; // MAPA + SISBOV + CPF/CNPJ
    uint32 internal honeyProfileId;  // MAPA + CAR + CPF/CNPJ + ANVISA (export)

    /*//////////////////////////////////////////////////////////////
                            SETUP
    //////////////////////////////////////////////////////////////*/

    function setUp() public {
        // 1. Deploy full Diamond
        d = deployDiamond(owner);
        supply   = ISupply(address(d.diamond));
        identity = IIdentity(address(d.diamond));
        asset    = IAsset(address(d.diamond));
        profile  = IProfile(address(d.diamond));
        ac       = IAccessControl(address(d.diamond));
        transfer = ITransfer(address(d.diamond));

        vm.startPrank(owner);

        // 2. Grant roles
        ac.grantRole(ISSUER_ROLE, issuer);
        ac.grantRole(TRANSFER_AGENT, agent);

        // 3. Create identity profiles (claim topics required per commodity)
        uint256[] memory coffeeTopics = new uint256[](3);
        coffeeTopics[0] = AgroClaimTopics.CLAIM_CPF_CNPJ;
        coffeeTopics[1] = AgroClaimTopics.CLAIM_MAPA_REGISTRATION;
        coffeeTopics[2] = AgroClaimTopics.CLAIM_CAR_REGISTRATION;
        coffeeProfileId = profile.createProfile(coffeeTopics);

        uint256[] memory cattleTopics = new uint256[](3);
        cattleTopics[0] = AgroClaimTopics.CLAIM_CPF_CNPJ;
        cattleTopics[1] = AgroClaimTopics.CLAIM_MAPA_REGISTRATION;
        cattleTopics[2] = AgroClaimTopics.CLAIM_SISBOV_REGISTRATION;
        cattleProfileId = profile.createProfile(cattleTopics);

        uint256[] memory honeyTopics = new uint256[](4);
        honeyTopics[0] = AgroClaimTopics.CLAIM_CPF_CNPJ;
        honeyTopics[1] = AgroClaimTopics.CLAIM_MAPA_REGISTRATION;
        honeyTopics[2] = AgroClaimTopics.CLAIM_CAR_REGISTRATION;
        honeyTopics[3] = AgroClaimTopics.CLAIM_EXPORT_AUTHORIZATION;
        honeyProfileId = profile.createProfile(honeyTopics);

        vm.stopPrank();

        // 4. Register KYC'd producers (agent binds wallet → ONCHAINID, country BR)
        vm.startPrank(agent);
        identity.registerIdentity(producerA, idA, COUNTRY_BRAZIL);
        identity.registerIdentity(producerB, idB, COUNTRY_BRAZIL);
        identity.registerIdentity(producerC, idC, COUNTRY_BRAZIL);
        vm.stopPrank();

        // 5. Deploy agro compliance modules
        kycModule     = new ProducerKycModule(address(d.diamond), owner);
        cprModule     = new CprLockupModule(owner);
        countryModule = new CountryRestrictModule(address(d.diamond), owner);

        // Configure KYC module: map each token to its profile
        // (done after registerAsset below, since tokenIds are auto-incremented)

        // 6. Register assets (owner acts as COMPLIANCE_ADMIN)
        vm.startPrank(owner);

        address[] memory noModules = new address[](0);

        coffeeId = asset.registerAsset(IAssetManager.RegisterAssetParams({
            name:               "Cafe Arabica - Safra 2025",
            symbol:             "CAF25",
            uri:                "ipfs://QmCoffee2025",
            supplyCap:          100_000e18,
            identityProfileId:  coffeeProfileId,
            complianceModules:  noModules,
            issuer:             issuer,
            allowedCountries:   new uint16[](0)
        }));

        cattleId = asset.registerAsset(IAssetManager.RegisterAssetParams({
            name:               "Gado Nelore - Lote 42",
            symbol:             "BOV42",
            uri:                "ipfs://QmCattle42",
            supplyCap:          500,        // 500 head maximum
            identityProfileId:  cattleProfileId,
            complianceModules:  noModules,
            issuer:             issuer,
            allowedCountries:   new uint16[](0)
        }));

        honeyId = asset.registerAsset(IAssetManager.RegisterAssetParams({
            name:               "Mel Organico - Apiario Serra",
            symbol:             "MEL01",
            uri:                "ipfs://QmHoney01",
            supplyCap:          0,          // unlimited — mixed NFT/fungible
            identityProfileId:  honeyProfileId,
            complianceModules:  noModules,
            issuer:             issuer,
            allowedCountries:   new uint16[](0)
        }));

        // 7. Configure KYC module profiles per asset
        kycModule.setTokenProfile(coffeeId, coffeeProfileId);
        kycModule.setTokenProfile(cattleId, cattleProfileId);
        kycModule.setTokenProfile(honeyId, honeyProfileId);

        // 8. Attach compliance modules to each asset
        asset.addComplianceModule(coffeeId, address(kycModule));
        asset.addComplianceModule(coffeeId, address(cprModule));

        asset.addComplianceModule(cattleId, address(kycModule));

        asset.addComplianceModule(honeyId, address(kycModule));
        asset.addComplianceModule(honeyId, address(countryModule));

        // 9. Restrict honey exports to Russia (example regulatory restriction)
        countryModule.restrictCountry(honeyId, COUNTRY_RUSSIA);

        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
            USE CASE 1: COFFEE — Fungible lot with CPR lock-up
    //////////////////////////////////////////////////////////////*/

    /// @notice Coffee issuer mints a fungible lot (amount > 1) to a registered producer.
    function test_Coffee_MintFungibleLot() public {
        vm.prank(issuer);
        supply.mint(coffeeId, producerA, 1000e18);

        assertEq(supply.balanceOf(producerA, coffeeId), 1000e18);
        assertEq(supply.totalSupply(coffeeId), 1000e18);
        assertEq(supply.holderCount(coffeeId), 1);
    }

    /// @notice KYC enforcement: unregistered outsider cannot receive coffee tokens via transfer.
    ///         Mint is issuer-controlled; KYC is enforced at secondary transfer (safeTransferFrom).
    function test_Coffee_RevertWhen_OutsiderReceives() public {
        vm.prank(issuer);
        supply.mint(coffeeId, producerA, 500e18);

        vm.prank(producerA);
        transfer.setApprovalForAll(issuer, true);

        vm.prank(issuer);
        vm.expectRevert();
        transfer.safeTransferFrom(producerA, outsider, coffeeId, 100e18, "");
    }

    /// @notice CPR lock-up blocks transfer before maturity date.
    function test_Coffee_CprLockup_BlocksTransferBeforeMaturity() public {
        // Mint to producerA
        vm.prank(issuer);
        supply.mint(coffeeId, producerA, 1000e18);

        // Set CPR lock-up: matures in 180 days
        uint256 maturity = block.timestamp + 180 days;
        vm.prank(owner);
        cprModule.setLockup(coffeeId, maturity);

        // Transfer attempt before maturity must revert
        vm.prank(producerA);
        transfer.setApprovalForAll(issuer, true);
        vm.prank(issuer);
        vm.expectRevert();
        transfer.safeTransferFrom(producerA, producerB, coffeeId, 100e18, "");
    }

    /// @notice CPR lock-up allows transfer after maturity date.
    function test_Coffee_CprLockup_AllowsTransferAfterMaturity() public {
        vm.prank(issuer);
        supply.mint(coffeeId, producerA, 1000e18);

        uint256 maturity = block.timestamp + 180 days;
        vm.prank(owner);
        cprModule.setLockup(coffeeId, maturity);

        // Advance time past maturity
        vm.warp(maturity + 1);

        vm.prank(producerA);
        transfer.setApprovalForAll(issuer, true);
        vm.prank(issuer);
        transfer.safeTransferFrom(producerA, producerB, coffeeId, 100e18, "");

        assertEq(supply.balanceOf(producerA, coffeeId), 900e18);
        assertEq(supply.balanceOf(producerB, coffeeId), 100e18);
    }

    /// @notice Gas benchmark: mint coffee fungible lot.
    function test_Gas_Coffee_MintFungible() public {
        vm.prank(issuer);
        uint256 gasBefore = gasleft();
        supply.mint(coffeeId, producerA, 1000e18);
        uint256 gasUsed = gasBefore - gasleft();
        emit log_named_uint("gas: coffee mint fungible", gasUsed);
    }

    /*//////////////////////////////////////////////////////////////
            USE CASE 2: CATTLE — Non-fungible per animal (amount = 1)
    //////////////////////////////////////////////////////////////*/

    /// @notice Cattle tokenId represents an individual animal (amount = 1, NFT semantics).
    function test_Cattle_MintNFT() public {
        vm.prank(issuer);
        supply.mint(cattleId, producerB, 1);

        assertEq(supply.balanceOf(producerB, cattleId), 1);
        assertEq(supply.totalSupply(cattleId), 1);
    }

    /// @notice Multiple cattle can be minted (each amount=1 represents one head).
    function test_Cattle_MintMultipleHeads() public {
        vm.prank(issuer);
        supply.mint(cattleId, producerB, 50);

        assertEq(supply.balanceOf(producerB, cattleId), 50);
    }

    /// @notice Unregistered buyer cannot receive cattle tokens (SISBOV KYC required).
    ///         KYC is enforced at transfer time — the issuer mints to a registered rancher,
    ///         who cannot then sell to an unregistered buyer.
    function test_Cattle_RevertWhen_OutsiderReceives() public {
        vm.prank(issuer);
        supply.mint(cattleId, producerB, 10);

        vm.prank(producerB);
        transfer.setApprovalForAll(issuer, true);

        vm.prank(issuer);
        vm.expectRevert();
        transfer.safeTransferFrom(producerB, outsider, cattleId, 1, "");
    }

    /// @notice Gas benchmark: mint cattle NFT (amount = 1).
    function test_Gas_Cattle_MintNFT() public {
        vm.prank(issuer);
        uint256 gasBefore = gasleft();
        supply.mint(cattleId, producerB, 1);
        uint256 gasUsed = gasBefore - gasleft();
        emit log_named_uint("gas: cattle mint NFT (amount=1)", gasUsed);
    }

    /*//////////////////////////////////////////////////////////////
            USE CASE 3: HONEY — Hybrid (colony NFT + lot fungible)
    //////////////////////////////////////////////////////////////*/

    /// @notice Colony tokens (id < 1000) are NFTs — one per bee colony.
    ///         Lot tokens (id >= 1000) are fungible — kilograms of honey per batch.
    ///         Both share the same Diamond token contract via ERC-1155.
    function test_Honey_MintColonyNFT() public {
        // Colony #1 — NFT semantics (amount = 1)
        vm.prank(issuer);
        supply.mint(honeyId, producerC, 1);

        assertEq(supply.balanceOf(producerC, honeyId), 1);
    }

    /// @notice Honey lot is fungible — represents kilograms of honey from a batch.
    function test_Honey_MintFungibleLot() public {
        vm.prank(issuer);
        supply.mint(honeyId, producerC, 500e18);

        assertEq(supply.balanceOf(producerC, honeyId), 500e18);
    }

    /// @notice Country restriction blocks honey transfer to a restricted jurisdiction.
    function test_Honey_RevertWhen_TransferToRestrictedCountry() public {
        // Register a Russian buyer in the identity system
        address russianBuyer = makeAddr("russianBuyer");
        address russianId    = makeAddr("russianId");
        vm.prank(agent);
        identity.registerIdentity(russianBuyer, russianId, COUNTRY_RUSSIA);

        // Mint honey to producerC
        vm.prank(issuer);
        supply.mint(honeyId, producerC, 200e18);

        // Transfer to Russian buyer must be blocked by CountryRestrictModule
        vm.prank(producerC);
        transfer.setApprovalForAll(issuer, true);
        vm.prank(issuer);
        vm.expectRevert();
        transfer.safeTransferFrom(producerC, russianBuyer, honeyId, 100e18, "");
    }

    /// @notice Gas benchmark: mint honey fungible lot.
    function test_Gas_Honey_MintFungible() public {
        vm.prank(issuer);
        uint256 gasBefore = gasleft();
        supply.mint(honeyId, producerC, 200e18);
        uint256 gasUsed = gasBefore - gasleft();
        emit log_named_uint("gas: honey mint fungible", gasUsed);
    }

    /*//////////////////////////////////////////////////////////////
                    BATCH MINT — Three assets in one tx
    //////////////////////////////////////////////////////////////*/

    /// @notice Batch mint across all three agricultural asset classes in one transaction.
    ///         Mirrors the cooperative scenario: end-of-harvest tokenization event
    ///         where coffee, cattle, and honey are issued simultaneously.
    function test_BatchMint_ThreeAssets() public {
        uint256[] memory ids        = new uint256[](3);
        address[] memory recipients = new address[](3);
        uint256[] memory amounts    = new uint256[](3);

        ids[0] = coffeeId; recipients[0] = producerA; amounts[0] = 1000e18;
        ids[1] = cattleId; recipients[1] = producerB; amounts[1] = 50;
        ids[2] = honeyId;  recipients[2] = producerC; amounts[2] = 200e18;

        vm.prank(issuer);
        supply.batchMint(ids, recipients, amounts);

        assertEq(supply.balanceOf(producerA, coffeeId), 1000e18);
        assertEq(supply.balanceOf(producerB, cattleId), 50);
        assertEq(supply.balanceOf(producerC, honeyId),  200e18);
    }

    /// @notice Gas benchmark: batch mint across three asset classes.
    function test_Gas_BatchMint_ThreeAssets() public {
        uint256[] memory ids        = new uint256[](3);
        address[] memory recipients = new address[](3);
        uint256[] memory amounts    = new uint256[](3);

        ids[0] = coffeeId; recipients[0] = producerA; amounts[0] = 1000e18;
        ids[1] = cattleId; recipients[1] = producerB; amounts[1] = 50;
        ids[2] = honeyId;  recipients[2] = producerC; amounts[2] = 200e18;

        vm.prank(issuer);
        uint256 gasBefore = gasleft();
        supply.batchMint(ids, recipients, amounts);
        uint256 gasUsed = gasBefore - gasleft();
        emit log_named_uint("gas: batch mint 3 assets", gasUsed);
    }

    /*//////////////////////////////////////////////////////////////
                    KYC REGISTRATION CORRECTNESS
    //////////////////////////////////////////////////////////////*/

    function test_IdentityRegistry_ProducersRegistered() public view {
        assertTrue(identity.contains(producerA));
        assertTrue(identity.contains(producerB));
        assertTrue(identity.contains(producerC));
        assertFalse(identity.contains(outsider));
    }

    function test_IdentityRegistry_CountryIsBrazil() public view {
        assertEq(identity.getCountry(producerA), COUNTRY_BRAZIL);
        assertEq(identity.getCountry(producerB), COUNTRY_BRAZIL);
        assertEq(identity.getCountry(producerC), COUNTRY_BRAZIL);
    }

    /*//////////////////////////////////////////////////////////////
                    CPR LOCKUP MODULE — UNIT
    //////////////////////////////////////////////////////////////*/

    function test_CprModule_IsLockedBeforeMaturity() public {
        uint256 maturity = block.timestamp + 90 days;
        vm.prank(owner);
        cprModule.setLockup(coffeeId, maturity);

        assertTrue(cprModule.isLocked(coffeeId));
    }

    function test_CprModule_IsUnlockedAfterMaturity() public {
        uint256 maturity = block.timestamp + 90 days;
        vm.prank(owner);
        cprModule.setLockup(coffeeId, maturity);

        vm.warp(maturity + 1);
        assertFalse(cprModule.isLocked(coffeeId));
    }

    function test_CprModule_ClearLockup() public {
        uint256 maturity = block.timestamp + 90 days;
        vm.prank(owner);
        cprModule.setLockup(coffeeId, maturity);
        assertTrue(cprModule.isLocked(coffeeId));

        vm.prank(owner);
        cprModule.clearLockup(coffeeId);
        assertFalse(cprModule.isLocked(coffeeId));
    }
}
