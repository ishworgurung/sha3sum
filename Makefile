
build:
	@echo "\033[0;42mBuilding sha3sum...\033[0;0m"
	@go get golang.org/x/crypto/sha3
	@go build sha3sum.go

install: build
	@echo "\033[0;42mInstalling sha3sum...\033[0;0m"
	@mkdir -p build/usr/bin >/dev/null 2>&1 
	@mv -vf ./sha3sum build/usr/bin/ >/dev/null 2>&1

clean:
	@echo "\033[0;42mCleaning sha3sum build...\033[0;0m"
	@rm -f ./sha3sum
	@rm -rf build

test:
	@go test
	
.PHONY: build install
