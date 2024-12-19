const hre = require("hardhat");
const { ethers } = require("hardhat");

async function main() {
    const [deployer] = await ethers.getSigners();

    console.log("Setup contracts with the account:", deployer.address);
    const erc20Address = "0x676998554fFC3E5B387D3A3EFEc782c83B9EBd08";
    const erc721Address = "0xD540788fE62f57ace9795dFb8a9C81af9d7EECb7";
    const saleAddress = "0xD4F6acb05eAd15266F4011aB1268c906ca4f6De1";

    const token = await (await ethers.getContractFactory("ERC20Token")).attach(erc20Address);
    await token.approve(saleAddress, "3490000000000000000000");

    const nft = await (await ethers.getContractFactory("NFT")).attach(erc721Address);
    await nft.setMinter(saleAddress);

    const sale = await (await ethers.getContractFactory("NFTPresaleManager")).attach(saleAddress);
    await sale.setPaymentToken(erc20Address, true);
    // 1734663600: 2024/12/20 10AM GMT7+
    // 1737392399: 2025/01/20 12PM GMT7+
    // 500 max supply
    // 349 usdt
    // no whitelist round
    // 1 limit nft per wallet
    await sale.createSaleRound("1734663600", "1737392399", 500, "349000000000000000000", false, 1);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
