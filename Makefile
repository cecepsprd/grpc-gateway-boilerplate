start:
	@go run main.go start --config .env

migrate-create:
	@migrate create -ext sql -dir migration -seq init_schema

migrate-up:
	@migrate -path ./migration -database "mysql://user:password@tcp(localhost:3306)/minipos?charset=utf8mb4&parseTime=True&loc=Local" -verbose up

migrate-down:
	@migrate -path ./migration -database "mysql://user:password@tcp(localhost:3306)/minipos?charset=utf8mb4&parseTime=True&loc=Local" -verbose down

startdb:
	@docker container start mysqlc

make genproto:
	@protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ --go_out=plugins=grpc:./api/proto --grpc-gateway_out=logtostderr=true:. ./api/proto/*.proto