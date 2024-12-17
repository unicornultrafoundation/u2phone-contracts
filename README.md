# U2 Testnet contracts

- ERC20: [0x1f9687E57335956C0Cd7C42eAc0d077552f6cC0d](https://testnet.u2uscan.xyz/address/0x1f9687E57335956C0Cd7C42eAc0d077552f6cC0d)
- ERC721: [0x68ffb5a03C006ee9fCf302EA018B703f5202c372](https://testnet.u2uscan.xyz/address/0x68ffb5a03C006ee9fCf302EA018B703f5202c372)
- NFTPresaleManager: [0x575795dE27ef82A3f61B456012666d32619691BE](https://testnet.u2uscan.xyz/address/0x575795dE27ef82A3f61B456012666d32619691BE)

### How to use
1. Add new sale round in NFTPresaleManager ([link](https://testnet.u2uscan.xyz/tx/0xa184db1f1af22bb1023adc4c2dff268e06babd732985c5d1fdf3d0f97e4b59f6))
2. If _whitelistOnly == `true`, then add whitelist addresses ([link](https://testnet.u2uscan.xyz/tx/0x66908872845dc46b927938210f9e05bcddfacdfd3b67422037649e31be89497d)). If not then skip to 3rd step
3. Set accepted payment token in the same contract ([link](https://testnet.u2uscan.xyz/tx/0x25fb3ee02aa7a4375a70f5e52bcee77c9517269a595b00cc4ed35e989a94f3e6))
4. Set minter role for NFTPresaleManager in ERC721 contract ([link](https://testnet.u2uscan.xyz/tx/0xa6c74491bbd10dfe13177744ad4029aa232c6b357d96592752af13f165ec66d5))
5. For user to buy NFT in NFTPresaleManager, they first need to approve sale contract before buy NFT ([link](https://testnet.u2uscan.xyz/tx/0x403048ee7c4df752b3b3e8f0832d45b032981a7bc38a9087d3ab1e6db53d1953))
6. Lastly, user call `BuyNFT` in NFTPresaleManager ([link](https://testnet.u2uscan.xyz/tx/0xeeefd250d97e309a3bc19a13b612d85d20519cd8980d60e6df4c5dc6a727309c))
