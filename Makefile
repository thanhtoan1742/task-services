API_DIR := "./api"
gen-protobuf:
	protoc -I ./$(API_DIR) \
		--go_out=$(API_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(API_DIR) --go-grpc_opt=paths=source_relative \
		$(API_DIR)/*.proto

mod:
	go mod tidy

test:
	echo "no test"

clean:
	rm -f $(API_DIR)/*.go
