const hre = require("hardhat");
const { ethers } = require("hardhat");

async function main() {
	const [deployer] = await ethers.getSigners();

	console.log("Deploying contracts with the account:", deployer.address);

	const NFT = await ethers.getContractFactory("NFT");
	const nft721 = await NFT.deploy("U2PhoneNFT", "U2P", "https://example.com/api/v1/token/");
	await nft721.waitForDeployment();
	const nftAddress = await nft721.getAddress();
	const presale = await ethers.deployContract("NFTPresaleManager", [nftAddress, "500", "100", "1733270400", "1735948800"], deployer);

	console.log("nft721 address:", await nft721.getAddress());
	console.log("pre-sale address:", await presale.getAddress());
}

main()
	.then(() => process.exit(0))
	.catch((error) => {
		console.error(error);
		process.exit(1);
	});
