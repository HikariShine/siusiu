.PHONY: all build_linux build_mac build_win run gotool clean help

BINARY="rickshang/tool-helper"

all: build_docker push

build_docker:
	docker build -t ${BINARY} . 

push:
	docker push ${BINARY}

