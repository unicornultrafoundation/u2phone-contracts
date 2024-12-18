// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/EIP712.sol";

interface IERC721Mintable {
	function mint(address to, uint256 tokenId) external;
	function balanceOf(address owner) external view returns (uint256 balance);
}

contract NFTPresaleManager is Ownable, EIP712 {
	using ECDSA for bytes32;
	using Address for address;

	struct SaleRound {
		uint256 id;
		uint256 startTime;
		uint256 endTime;
		uint256 maxSupply;
		uint256 price;
		bool whitelistOnly;
		uint256 totalMinted;
	}

	struct ReferrerSignatureHashInfo {
		address referrer;
		uint256 createdAt;
	}

	struct ReferralInfo {
		address referrer;
		uint256 totalCommission;
	}

	IERC721Mintable public nftContract;
	address public adminAddress;
	uint256 public currentRoundId;
	uint256 public maxNFTPerWallet;
	uint256 public totalMinted;
	uint256 public constant REFERRAL_COMMISSION = 5; // 5% commission

	// Mapping for accepted payment tokens
	mapping(address => bool) public acceptedPaymentTokens;
	// Mapping for whitelisted addresses
	mapping(address => bool) public whitelisted;
	// Mapping for sale rounds
	mapping(uint256 => SaleRound) public saleRounds;
	// Mapping for referral information
	mapping(address => ReferralInfo) public referrals;
	mapping(bytes => bool) public blackListSignature;

	event PaymentTokenUpdated(address indexed token, bool accepted);
	event Whitelisted(address indexed account);
	event RemovedFromWhitelist(address indexed account);
	event SaleRoundCreated(uint256 indexed roundId);
	event NFTPurchased(address indexed buyer, uint256 tokenId);
	event ReferralCommissionPaid(address indexed referrer, address indexed buyer, uint256 amount);

	constructor(
		address _nftContract,
		uint256 _maxNFTPerWallet,
		address _adminAddress
	) EIP712("NFTSale", "1") {
		require(_nftContract.isContract(), "Invalid NFT contract address");
		require(_maxNFTPerWallet > 0, "Invalid max NFT per wallet");

		nftContract = IERC721Mintable(_nftContract);
		maxNFTPerWallet = _maxNFTPerWallet;
		currentRoundId = 0;
		adminAddress = _adminAddress;
	}

	modifier validPaymentToken(address _token) {
		require(acceptedPaymentTokens[_token], "Payment token not accepted");
		_;
	}

	modifier nftContractSet() {
		require(address(nftContract) != address(0), "NFT contract not set");
		_;
	}

	modifier onlyNotBlacklisted(bytes memory signature) {
		require(blackListSignature[signature], "Only signature has not been used");
		_;
	}

	// Purchase NFT with referral
	function buyNFTWithReferral(
		uint256 _roundId,
		address _paymentToken,
		bytes memory _signature,
		address _referrer,
		uint256 createdAt
	)
	external
	nftContractSet
	validPaymentToken(_paymentToken)
	onlyNotBlacklisted(_signature)
	{
		ReferrerSignatureHashInfo memory data = ReferrerSignatureHashInfo({referrer: _referrer, createdAt: createdAt});
		address _signer = _verifySignature(data, _signature);
		blackListSignature[_signature] = true;

		require(_referrer != msg.sender, "Cannot refer yourself");
		require(_referrer != address(0), "Invalid referrer address");

		SaleRound storage round = saleRounds[_roundId];
		require(round.id > 0, "Invalid round");
		require(block.timestamp >= round.startTime, "Round not started");
		require(block.timestamp <= round.endTime, "Round ended");
		require(round.totalMinted < round.maxSupply, "Round max supply reached");

		if (round.whitelistOnly) {
			require(whitelisted[msg.sender], "Not whitelisted");
		}

		require(
			nftContract.balanceOf(msg.sender) < maxNFTPerWallet,
			"Max NFT per wallet reached"
		);

		IERC20 paymentToken = IERC20(_paymentToken);

		// Calculate commission
		uint256 commission = (round.price * REFERRAL_COMMISSION) / 100;
		uint256 finalPrice = round.price;

		// Transfer payment
		require(
			paymentToken.transferFrom(msg.sender, address(this), finalPrice),
			"Payment failed"
		);

		// Pay commission to referrer
		require(
			paymentToken.transfer(_referrer, commission),
			"Commission transfer failed"
		);

		// Update referral info
		referrals[_referrer].referrer = _referrer;
		referrals[_referrer].totalCommission += commission;

		uint256 tokenId = totalMinted + 1;
		totalMinted++;
		round.totalMinted++;

		nftContract.mint(msg.sender, tokenId);

		emit NFTPurchased(msg.sender, tokenId);
		emit ReferralCommissionPaid(_referrer, msg.sender, commission);
	}

	// Get referral info
	function getReferralInfo(address _referrer)
	external
	view
	returns (
		address referrer,
		uint256 totalCommission
	)
	{
		ReferralInfo storage info = referrals[_referrer];
		return (
			info.referrer,
			info.totalCommission
		);
	}

	// ============================= INTERNAL HANDLE ============================= //
	function _verifySignature(
		ReferrerSignatureHashInfo memory data,
		bytes memory _signature
	) private view returns (address) {
		bytes32 digest = hashSignature(msg.sender, data);
		return digest.toEthSignedMessageHash().recover(_signature);
	}

	function hashSignature(
		address _userAddr,
		ReferrerSignatureHashInfo memory data
	) public view returns (bytes32) {
		return
			_hashTypedDataV4(
				keccak256(abi.encode(_userAddr, address(this), data))
			);
	}
}
