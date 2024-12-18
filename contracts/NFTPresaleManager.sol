// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import "./libs/TransferHelper.sol";

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
		uint256 maxNFTPerWalletInRound;
		mapping(address => uint256) walletMintedInRound;
	}

	struct ReferrerSignatureHashInfo {
		address referrer;
		uint256 createdAt;
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
	mapping(bytes => bool) public blackListSignature;

	event PaymentTokenUpdated(address indexed token, bool accepted);
	event Whitelisted(address indexed account);
	event RemovedFromWhitelist(address indexed account);
	event SaleRoundCreated(uint256 indexed roundId);
	event NFTPurchased(address indexed buyer, uint256 tokenId);
	event ReferralCommissionPaid(address indexed referrer, address indexed buyer, uint256 amount);
	event MaxNFTPerWalletUpdated(uint256 roundId, uint256 maxNFTPerWallet);

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
		require(blackListSignature[signature] == false, "Only signature has not been used");
		_;
	}

	// Set accepted payment token
	function setPaymentToken(address _token, bool _accepted) external onlyOwner {
		require(_token.isContract(), "Invalid token address");
		acceptedPaymentTokens[_token] = _accepted;
		emit PaymentTokenUpdated(_token, _accepted);
	}

	// Create new sale round
	function createSaleRound(
		uint256 _startTime,
		uint256 _endTime,
		uint256 _maxSupply,
		uint256 _price,
		bool _whitelistOnly,
		uint256 _maxNFTPerWalletInRound
	) external onlyOwner {
		require(_startTime < _endTime, "Start must be before end");
		require(_maxSupply > 0, "Invalid max supply");
		require(_price > 0, "Invalid price");
		require(_maxNFTPerWalletInRound > 0, "Invalid max NFT per wallet in round");

		currentRoundId++;

		SaleRound storage newRound = saleRounds[currentRoundId];
		newRound.id = currentRoundId;
		newRound.startTime = _startTime;
		newRound.endTime = _endTime;
		newRound.maxSupply = _maxSupply;
		newRound.price = _price;
		newRound.whitelistOnly = _whitelistOnly;
		newRound.totalMinted = 0;
		newRound.maxNFTPerWalletInRound = _maxNFTPerWalletInRound;

		emit SaleRoundCreated(currentRoundId);
	}

	// Add to whitelist (single address)
	function addToWhitelist(address account) external onlyOwner {
		whitelisted[account] = true;
		emit Whitelisted(account);
	}

	// Add to whitelist (batch)
	function addToWhitelistBatch(address[] calldata accounts) external onlyOwner {
		for (uint256 i = 0; i < accounts.length; i++) {
			whitelisted[accounts[i]] = true;
			emit Whitelisted(accounts[i]);
		}
	}

	// Remove from whitelist
	function removeFromWhitelist(address account) external onlyOwner {
		whitelisted[account] = false;
		emit RemovedFromWhitelist(account);
	}

	// Purchase NFT with referral
	function buyNFTWithReferral(
		uint256 _roundId,
		address _paymentToken,
		uint256 _quantity,
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
		require(_signer == adminAddress, "Invalid input address");
		require(block.timestamp < createdAt + 300, "Signature is expired");

		require(_referrer != address(0), "Invalid referrer address");

		SaleRound storage round = saleRounds[_roundId];
		require(round.id > 0, "Invalid round");
		require(block.timestamp >= round.startTime, "Round not started");
		require(block.timestamp <= round.endTime, "Round ended");
		require(round.totalMinted < round.maxSupply, "Round max supply reached");

		if (round.whitelistOnly) {
			require(whitelisted[msg.sender], "Not whitelisted");
		}

		// Check global wallet limit
		require(
			nftContract.balanceOf(msg.sender) + _quantity <= maxNFTPerWallet,
			"Would exceed global max NFT per wallet"
		);
		// Check round-specific wallet limit
		require(
			round.walletMintedInRound[msg.sender] + _quantity <= round.maxNFTPerWalletInRound,
			"Would exceed round max NFT per wallet"
		);

		// Calculate commission
		uint256 totalPrice = round.price * _quantity;
		uint256 commission = (totalPrice * REFERRAL_COMMISSION) / 100;

		// Transfer payment
		TransferHelper.safeTransferFrom(_paymentToken, msg.sender, address(this), totalPrice);

		// Pay commission to referrer
		TransferHelper.safeTransfer(_paymentToken, _referrer, commission);

		for (uint256 i = 0; i < _quantity; i++) {
			uint256 tokenId = totalMinted + 1;
			totalMinted++;
			round.totalMinted++;
			nftContract.mint(msg.sender, tokenId);
			round.walletMintedInRound[msg.sender]++;
			emit NFTPurchased(msg.sender, tokenId);
		}
		emit ReferralCommissionPaid(_referrer, msg.sender, commission);
	}

	// Withdraw funds for specific token
	function withdrawFunds(address _token) external onlyOwner validPaymentToken(_token) {
		IERC20 token = IERC20(_token);
		uint256 balance = token.balanceOf(address(this));
		require(balance > 0, "No funds to withdraw");
		TransferHelper.safeTransfer(_token, owner(), balance);
	}

	// Set max NFT per wallet
	function setMaxNFTPerWallet(uint256 _maxNFTPerWallet) external onlyOwner {
		require(_maxNFTPerWallet > 0, "Invalid max NFT per wallet");
		maxNFTPerWallet = _maxNFTPerWallet;
	}

	// Set max NFT per wallet in round
	function setMaxNFTPerWallet(uint256 _roundId, uint256 _maxNFTPerWallet) external onlyOwner {
		require(_maxNFTPerWallet > 0, "Invalid max NFT per wallet");
		SaleRound storage round = saleRounds[_roundId];
		round.maxNFTPerWalletInRound = _maxNFTPerWallet;
		emit MaxNFTPerWalletUpdated(_roundId, _maxNFTPerWallet);
	}

	// Get round info
	function getRoundInfo(uint256 _roundId)
	external
	view
	returns (
		uint256 id,
		uint256 startTime,
		uint256 endTime,
		uint256 maxSupply,
		uint256 price,
		bool whitelistOnly,
		uint256 totalMintedInRound,
		uint256 maxNFTPerWalletInRound,
		uint256 walletMinted
	)
	{
		SaleRound storage round = saleRounds[_roundId];
		return (
			round.id,
			round.startTime,
			round.endTime,
			round.maxSupply,
			round.price,
			round.whitelistOnly,
			round.totalMinted,
			round.maxNFTPerWalletInRound,
			round.walletMintedInRound[msg.sender]
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
