run-dev:
	clear; go run cmd/service/main.go -env-mode=development -config-path=environments/config.yaml

.PHONY: clean
generate:
		protoc --go_out=internal/generated/ --go_opt=paths=source_relative --go-grpc_out=internal/generated/ --go-grpc_opt=paths=source_relative proto/admitad_integration/admitad_integration.proto
