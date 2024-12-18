// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract NFT is ERC721, Ownable {
	address public minter;
	string private _baseTokenURI;

	event MinterChanged(address indexed oldMinter, address indexed newMinter);
	event BaseURIChanged(string oldURI, string newURI);

	constructor(
		string memory name,
		string memory symbol,
		string memory baseURI
	) ERC721(name, symbol) {
		minter = msg.sender;
		_baseTokenURI = baseURI;
	}

	modifier onlyMinter() {
		require(msg.sender == minter, "Not the minter");
		_;
	}

	function setMinter(address newMinter) external onlyOwner {
		require(newMinter != address(0), "Invalid address");
		emit MinterChanged(minter, newMinter);
		minter = newMinter;
	}

	function mint(address to, uint256 tokenId) external onlyMinter {
		_mint(to, tokenId);
	}

	/**
     * @dev Sets the base URI for the token metadata
     * Only the contract owner can call this function
     */
	function setBaseURI(string memory newBaseURI) external onlyOwner {
		string memory oldURI = _baseTokenURI;
		_baseTokenURI = newBaseURI;
		emit BaseURIChanged(oldURI, newBaseURI);
	}

	/**
     * @dev Returns the base URI for computing {tokenURI}
     */
	function _baseURI() internal view virtual override returns (string memory) {
		return _baseTokenURI;
	}
}