// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IERC721Mintable {
	function mint(address to, uint256 tokenId) external;
	function balanceOf(address owner) external view returns (uint256 balance);
}

contract NFTPresaleManager is Ownable {
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

	IERC721Mintable public nftContract;
	uint256 public currentRoundId;
	uint256 public maxNFTPerWallet;
	uint256 public totalMinted;

	// Mapping for accepted payment tokens
	mapping(address => bool) public acceptedPaymentTokens;
	// Mapping for whitelisted addresses
	mapping(address => bool) public whitelisted;
	// Mapping for sale rounds
	mapping(uint256 => SaleRound) public saleRounds;

	event PaymentTokenUpdated(address indexed token, bool accepted);
	event Whitelisted(address indexed account);
	event RemovedFromWhitelist(address indexed account);
	event SaleRoundCreated(uint256 indexed roundId);
	event NFTPurchased(address indexed buyer, uint256 tokenId);

	constructor(
		address _nftContract,
		uint256 _maxNFTPerWallet
	) {
		require(_nftContract.isContract(), "Invalid NFT contract address");
		require(_maxNFTPerWallet > 0, "Invalid max NFT per wallet");

		nftContract = IERC721Mintable(_nftContract);
		maxNFTPerWallet = _maxNFTPerWallet;
		currentRoundId = 0;
	}

	modifier validPaymentToken(address _token) {
		require(acceptedPaymentTokens[_token], "Payment token not accepted");
		_;
	}

	modifier nftContractSet() {
		require(address(nftContract) != address(0), "NFT contract not set");
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
		bool _whitelistOnly
	) external onlyOwner {
		require(_startTime < _endTime, "Start must be before end");
		require(_maxSupply > 0, "Invalid max supply");
		require(_price > 0, "Invalid price");

		currentRoundId++;

		saleRounds[currentRoundId] = SaleRound({
			id: currentRoundId,
			startTime: _startTime,
			endTime: _endTime,
			maxSupply: _maxSupply,
			price: _price,
			whitelistOnly: _whitelistOnly,
			totalMinted: 0
		});

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

	// Purchase NFT
	function buyNFT(uint256 _roundId, address _paymentToken)
	external
	nftContractSet
	validPaymentToken(_paymentToken)
	{
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
		require(
			paymentToken.transferFrom(msg.sender, address(this), round.price),
			"Payment failed"
		);

		uint256 tokenId = totalMinted + 1;
		totalMinted++;
		round.totalMinted++;

		nftContract.mint(msg.sender, tokenId);
		emit NFTPurchased(msg.sender, tokenId);
	}

	// Withdraw funds for specific token
	function withdrawFunds(address _token) external onlyOwner validPaymentToken(_token) {
		IERC20 token = IERC20(_token);
		uint256 balance = token.balanceOf(address(this));
		require(balance > 0, "No funds to withdraw");
		require(token.transfer(owner(), balance), "Transfer failed");
	}

	// Set max NFT per wallet
	function setMaxNFTPerWallet(uint256 _maxNFTPerWallet) external onlyOwner {
		require(_maxNFTPerWallet > 0, "Invalid max NFT per wallet");
		maxNFTPerWallet = _maxNFTPerWallet;
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
		uint256 totalMinted
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
			round.totalMinted
		);
	}
}