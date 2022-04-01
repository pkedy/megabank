.PHONY: run
run:
	dapr run --app-id megabank \
		--app-protocol http \
		--app-port 8080 \
		--dapr-http-port 3500 \
		--components-path ./components \
		go run ./cmd/main.go

.PHONY: docs
docs:
	@swag init --dir cmd,pkg/api,pkg/account,pkg/transfer