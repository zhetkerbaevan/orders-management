run:
	@go run cmd/orders-management/main.go

gen:
	protoc -Iprotobuf \
	--go_out=paths=source_relative:./internal/genproto \
	--go-grpc_out=paths=source_relative:./internal/genproto \
	--goclay_out=grpc_api_configuration=protobuf/http.yaml:./internal \
	protobuf/orders.proto

stat:
	statik -src=D:\GoProjects\orders-management\static\swagger-ui -dest=D:\GoProjects\orders-management\static
	
migration: 
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up		

migrate-down:
	@go run cmd/migrate/main.go down