const hre = require("hardhat");
const { ethers } = require("hardhat");

async function main() {
    const [deployer] = await ethers.getSigners();

    console.log("Setup contracts with the account:", deployer.address);
    const erc20Address = "0x57b616B768f7161Aab6d0c01135312FaCfA57A44";
    const erc721Address = "0x7bD59378B581BaE4a22FaeF833C77D98D98dA82F";
    const saleAddress = "0x5f482F4B9eDC38E10DC2D02588E77462Ab003cF3";

    const token = await (await ethers.getContractFactory("ERC20Token")).attach(erc20Address);
    await token.approve(saleAddress, "3490000000000000000000");

    const nft = await (await ethers.getContractFactory("NFT")).attach(erc721Address);
    await nft.setMinter(saleAddress);

    const sale = await (await ethers.getContractFactory("NFTPresaleManager")).attach(saleAddress);
    await sale.setPaymentToken(erc20Address, true);
    await sale.createSaleRound("1734393600", "1734739200", 500, "349000000000000000000", false, 3);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
