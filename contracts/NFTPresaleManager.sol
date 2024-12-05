// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Address.sol";

interface IERC721Mintable {
	function mint(address to, uint256 tokenId) external;
}

contract NFTPresaleManager is Ownable {
	using Address for address;

	IERC721Mintable public nftContract;
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

	constructor(
		address _nftContract,
		uint256 _maxSupply,
		uint256 _price,
		uint256 _presaleStart,
		uint256 _presaleEnd
	) {
		require(_presaleStart < _presaleEnd, "Start must be before end");
		require(_nftContract.isContract(), "Invalid NFT contract address");

		nftContract = IERC721Mintable(_nftContract);
		maxSupply = _maxSupply;
		price = _price;
		presaleStart = _presaleStart;
		presaleEnd = _presaleEnd;
	}

	modifier nftContractSet() {
		require(address(nftContract) != address(0), "NFT contract not set");
		_;
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
	function buyNFT() external payable nftContractSet {
		require(totalMinted < maxSupply, "Max supply reached");
		require(msg.value >= price, "Insufficient funds");

		if (block.timestamp >= presaleStart && block.timestamp <= presaleEnd) {
			require(whitelisted[msg.sender], "Not whitelisted during presale");
		}

		uint256 tokenId = totalMinted + 1;
		totalMinted += 1;

		nftContract.mint(msg.sender, tokenId);

		emit NFTPurchased(msg.sender, tokenId);
	}

	// Withdraw funds from the contract
	function withdrawFunds() external onlyOwner {
		payable(owner()).transfer(address(this).balance);
	}
}
