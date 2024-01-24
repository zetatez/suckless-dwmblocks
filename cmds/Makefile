default: clean build

prepare:
	mkdir -p ./bins/
	go mod tidy

build: prepare
	@cat ./go-build.sh
	./go-build.sh

clean:
	rm -rf bins/*

.PHONY: default

