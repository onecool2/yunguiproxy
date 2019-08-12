rm -rf *.abi
rm -rf *.go
./solc-static-linux ./yungui.solc --abi -o ./
abigen --abi=YunguiContract.abi --pkg=YunguiContract --out=yunguiContract.go
