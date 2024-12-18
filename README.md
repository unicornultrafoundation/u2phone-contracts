# U2 Testnet contracts

- ERC20: [0x57b616B768f7161Aab6d0c01135312FaCfA57A44](https://testnet.u2uscan.xyz/address/0x57b616B768f7161Aab6d0c01135312FaCfA57A44)
- ERC721: [0x7bD59378B581BaE4a22FaeF833C77D98D98dA82F](https://testnet.u2uscan.xyz/address/0x7bD59378B581BaE4a22FaeF833C77D98D98dA82F)
- NFTPresaleManager: [0x5f482F4B9eDC38E10DC2D02588E77462Ab003cF3](https://testnet.u2uscan.xyz/address/0x5f482F4B9eDC38E10DC2D02588E77462Ab003cF3)

### How to use
1. Add new sale round in NFTPresaleManager
2. If _whitelistOnly == `true`, then add whitelist addresses
3. Set accepted payment token in the same contract
4. Set minter role for NFTPresaleManager in ERC721 contract
5. For user to buy NFT in NFTPresaleManager, they first need to approve sale contract before buy NFT
6. User also has to ask for a signature to buy NFT
7. Lastly, user call `BuyNFT` in NFTPresaleManager
