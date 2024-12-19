const hre = require("hardhat");
const { ethers } = require("hardhat");

async function main() {
    const [deployer] = await ethers.getSigners();

    console.log("Setup contracts with the account:", deployer.address);
    const erc20Address = "0x252eB9Ba7a18acd54DaaEEaD027C80D4c77b47F8";
    const erc721Address = "0x210C601d2E1d9e45Ff4287FB7FDEE77f51aeE987";
    const saleAddress = "0xdcA348c8a9005a38F8eb1B9383686E55104fbfc4";

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
