solc-ballot:
	solc --abi app/contracts/src/Ballot.sol -o app/contracts/abi/ballot --overwrite
	solc --bin app/contracts/src/Ballot.sol -o app/contracts/abi/ballot --overwrite
	solc --combined-json abi,bin,bin-runtime,srcmap,srcmap-runtime,ast,metadata --output-dir app/frontend/src app/contracts/src/Ballot.sol --overwrite --pretty-json && mv app/frontend/src/combined.json app/frontend/src/Ballot.json

abigen-ballot:
	abigen --abi=app/contracts/abi/ballot/Ballot.abi --bin=app/contracts/abi/ballot/Ballot.bin --pkg=Ballot --out=app/contracts/go/binding/Ballot.go
