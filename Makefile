build:
	go build -o bin/mrz-sso-provider .

run:
	go run main.go

test:
	go test -v ./pkgs/apis/v1beta ./pkgs/sql_db ./pkgs/security ./pkgs/system
	@echo "\n ## Remove test databases file ## \n"
	rm -f test.db

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/mrz-sso-provider-linux-amd64 .
