SRC=./src/...

test: 
	@go test $(SRC)

cover:
	@go test $(SRC) -coverprofile=coverage.out
	@go tool cover -html=coverage.out
