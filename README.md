# U2 Testnet contracts

- ERC20: [0x76fC788256225C33Ef9e6b58d5f92386Cb632483](https://testnet.u2uscan.xyz/address/0x76fC788256225C33Ef9e6b58d5f92386Cb632483)
- ERC721: [0x0E38629CB3356BB439eF6DF413E79901C675f249](https://testnet.u2uscan.xyz/address/0x0E38629CB3356BB439eF6DF413E79901C675f249)
- NFTPresaleManager: [0x50A98c544C84142dE05DAe9938EFab93338eD845](https://testnet.u2uscan.xyz/address/0x50A98c544C84142dE05DAe9938EFab93338eD845)

### How to use
1. Add new sale round in NFTPresaleManager
2. If _whitelistOnly == `true`, then add whitelist addresses
3. Set accepted payment token in the same contract
4. Set minter role for NFTPresaleManager in ERC721 contract
5. For user to buy NFT in NFTPresaleManager, they first need to approve sale contract before buy NFT
6. Lastly, user call `BuyNFT` in NFTPresaleManager
