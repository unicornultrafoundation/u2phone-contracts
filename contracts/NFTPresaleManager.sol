// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IERC721Mintable {
	function mint(address to, uint256 tokenId) external;
}

contract NFTPresaleManager is Ownable {
	using Address for address;

	IERC721Mintable public nftContract;
	IERC20 public paymentToken; // ERC20 token used for payment
	uint256 public maxSupply;
	uint256 public price;
	uint256 public totalMinted;
	uint256 public presaleStart;
	uint256 public presaleEnd;
	mapping(address => bool) public whitelisted;

	event Whitelisted(address indexed account);
	event RemovedFromWhitelist(address indexed account);
	event PresalePeriodUpdated(uint256 start, uint256 end);
	event NFTPurchased(address indexed buyer, uint256 tokenId);
	event PaymentTokenUpdated(address indexed token);

	constructor(
		address _nftContract,
		uint256 _maxSupply,
		uint256 _price,
		uint256 _presaleStart,
		uint256 _presaleEnd,
		address _paymentToken
	) {
		require(_presaleStart < _presaleEnd, "Start must be before end");
		require(_nftContract.isContract(), "Invalid NFT contract address");
		require(_paymentToken.isContract(), "Invalid payment token address");

		nftContract = IERC721Mintable(_nftContract);
		maxSupply = _maxSupply;
		price = _price;
		presaleStart = _presaleStart;
		presaleEnd = _presaleEnd;
		paymentToken = IERC20(_paymentToken);
	}

	modifier nftContractSet() {
		require(address(nftContract) != address(0), "NFT contract not set");
		_;
	}

	modifier paymentTokenSet() {
		require(address(paymentToken) != address(0), "Payment token not set");
		_;
	}

	// Set a new payment token
	function setPaymentToken(address _paymentToken) external onlyOwner {
		require(_paymentToken.isContract(), "Invalid payment token address");
		paymentToken = IERC20(_paymentToken);
		emit PaymentTokenUpdated(_paymentToken);
	}

	// Add a single address to the whitelist
	function addToWhitelist(address account) external onlyOwner {
		whitelisted[account] = true;
		emit Whitelisted(account);
	}

	// Add multiple addresses to the whitelist
	function addToWhitelistBatch(address[] calldata accounts) external onlyOwner {
		for (uint256 i = 0; i < accounts.length; i++) {
			whitelisted[accounts[i]] = true;
			emit Whitelisted(accounts[i]);
		}
	}

	// Remove an address from the whitelist
	function removeFromWhitelist(address account) external onlyOwner {
		whitelisted[account] = false;
		emit RemovedFromWhitelist(account);
	}

	// Update the presale period
	function updatePresalePeriod(uint256 _start, uint256 _end) external onlyOwner {
		require(_start < _end, "Start must be before end");
		presaleStart = _start;
		presaleEnd = _end;
		emit PresalePeriodUpdated(_start, _end);
	}

	// Purchase an NFT
	function buyNFT() external paymentTokenSet nftContractSet {
		require(totalMinted < maxSupply, "Max supply reached");

		if (block.timestamp >= presaleStart && block.timestamp <= presaleEnd) {
			require(whitelisted[msg.sender], "Not whitelisted during presale");
		}

		// Transfer payment tokens from buyer to contract
		require(paymentToken.transferFrom(msg.sender, address(this), price), "Token transfer failed");

		uint256 tokenId = totalMinted + 1;
		totalMinted += 1;

		nftContract.mint(msg.sender, tokenId);

		emit NFTPurchased(msg.sender, tokenId);
	}

	// Withdraw funds (ERC20 tokens) from the contract
	function withdrawFunds() external onlyOwner {
		uint256 balance = paymentToken.balanceOf(address(this));
		require(balance > 0, "No funds to withdraw");
		paymentToken.transfer(owner(), balance);
	}
}
