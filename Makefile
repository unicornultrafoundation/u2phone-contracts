#bin/bash
abi-go:
	solc --abi --bin ./contracts/NFTPresaleManager.sol -o ./temp/ --overwrite --base-path ./ --include-path ./node_modules
	abigen --abi ./temp/NFTPresaleManager.abi --pkg nftsale --type NFTPresaleManager --out ./generated/NFTPresaleManager.go
