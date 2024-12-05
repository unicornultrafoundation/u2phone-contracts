// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract NFT is ERC721, Ownable {
	address public minter;

	event MinterChanged(address indexed oldMinter, address indexed newMinter);

	constructor(
		string memory name,
		string memory symbol,
		string memory baseURI
	) ERC721(name, symbol) {
		minter = msg.sender;
		_setBaseURI(baseURI);
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
	 * @dev Sets the base URI for the metadata.
	 * Only the contract owner can call this function.
	 */
	function setBaseURI(string memory baseURI) external onlyOwner {
		_setBaseURI(baseURI);
	}
}
