APP=handsongolang
APP_EXECUTABLE="./out/$(APP)"
HTTP_SERVE_COMMAND="http-serve"
CONFIG_FILE="./.env"


httpserve: build
	$(APP_EXECUTABLE) -configFile=$(CONFIG_FILE) $(HTTP_SERVE_COMMAND)
deps:
	go mod download
compile:
	mkdir -p out/
	go build -ldflags "-X main.version=0.1 " -o $(APP_EXECUTABLE) cmd/*.go
build: deps compile 