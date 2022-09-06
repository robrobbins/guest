compile_solidity_assert:
	@echo "compiling Assert contract source to .abi and .bin files"
	@solc -o ./test/ --abi --bin --overwrite ./test/Assert.sol

compile_go_assert:
	@echo "compiling Assert .abi and .bin files to go bindings, writing to pkg/testing/unit"
	@abigen --abi ./test/Assert.abi --bin ./test/Assert.bin -pkg unit -type Assert -out ./pkg/testing/unit/assert.go

compile_solidity_erc20:
	@echo "compiling test/erc20 solidity sources to .abi and .bin files"
	@solc -o ./test/erc20 --abi --bin --overwrite ./test/erc20/ERC20.sol
	@solc -o ./test/erc20 --abi --bin --overwrite ./test/erc20/ERC20Test.sol

compile_go_erc20:
	@echo "compiling test/erc20 .abi and .bin files to go bindings, writing to pkg/testing/unit/erc20"
	@abigen --abi ./test/erc20/ERC20.abi --bin ./test/erc20/ERC20.bin -pkg erc20 -type ERC20 -out ./pkg/testing/unit/erc20/erc20.go
	@abigen --abi ./test/erc20/ERC20Test.abi --bin ./test/erc20/ERC20Test.bin -pkg erc20 -type ERC20Test -out ./pkg/testing/unit/erc20/erc20test.go

compile_assert: compile_solidity_assert compile_go_assert
compile_erc20: compile_solidity_erc20 compile_go_erc20

clean_test:
	@echo "removing compiled files from test directories"
	@rm ./test/**/*.abi || echo "no .abi files to remove..."
	@rm ./test/**/*.bin || echo "no .bin files to remove..."

all: clean_test compile_assert compile_erc20
