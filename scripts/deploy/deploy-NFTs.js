const hre = require("hardhat");
const { ethers } = require("hardhat");

async function main() {
	const [deployer] = await ethers.getSigners();

	console.log("Deploying contracts with the account:", deployer.address);

	const token = await ethers.getContractFactory("ERC20Token");
	const erc20 = await token.deploy("ERC20Token", "FUSD", 100000);

	const NFT = await ethers.getContractFactory("NFT");
	const nft721 = await NFT.deploy("U2PhoneNFT", "U2P", "https://example.com/api/v1/token/");
	await nft721.waitForDeployment();
	const nftAddress = await nft721.getAddress();

	await erc20.waitForDeployment();
	const erc20Address = await erc20.getAddress();

	const presale = await ethers.deployContract("NFTPresaleManager", [nftAddress, 5], deployer);

	console.log("erc20 address:", erc20Address);
	console.log("nft721 address:", nftAddress);
	console.log("pre-sale address:", await presale.getAddress());
}

main()
	.then(() => process.exit(0))
	.catch((error) => {
		console.error(error);
		process.exit(1);
	});
