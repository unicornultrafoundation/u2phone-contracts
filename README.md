# U2 Testnet contracts

- ERC20: [0x523b8EC334BB4d018122361f39d1242eb5851EB5](https://testnet.u2uscan.xyz/address/0x523b8EC334BB4d018122361f39d1242eb5851EB5)
- ERC721: [0x64371babE5a19922D99ff0869052403F42128Ec8](https://testnet.u2uscan.xyz/address/0x64371babE5a19922D99ff0869052403F42128Ec8)
- NFTPresaleManager: [0x057176F118AF0930B48e099A3665A32Ad6df3A99](https://testnet.u2uscan.xyz/address/0x057176F118AF0930B48e099A3665A32Ad6df3A99)

### How to use
1. Add new sale round in NFTPresaleManager
2. If _whitelistOnly == `true`, then add whitelist addresses
3. Set accepted payment token in the same contract
4. Set minter role for NFTPresaleManager in ERC721 contract
5. For user to buy NFT in NFTPresaleManager, they first need to approve sale contract before buy NFT
6. Lastly, user call `BuyNFT` in NFTPresaleManager
