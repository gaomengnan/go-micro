default:
	@go build -o ./tmp/main  . && ./tmp/main && rm -rf ./tmp/main
