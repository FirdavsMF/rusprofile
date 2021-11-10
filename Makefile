# генерация кода и openapi
protoc:
	protoc -I . \
 	--go_out ./pkg/grpc/ \
 	--go_opt paths=source_relative  \
 	--go-grpc_out ./pkg/grpc/ \
 	--go-grpc_opt paths=source_relative \
 	-I./third_party/googleapis \
 	--grpc-gateway_out ./pkg/grpc \
 	--grpc-gateway_opt logtostderr=true \
 	--grpc-gateway_opt paths=source_relative \
 	--openapiv2_out ./api/ \
 	 rusprofile.proto

# Тестовый запрос к http gateway
http-curl-test:
	curl --location --request POST 'http://localhost:8080/getCompany' \
    --header 'Content-Type: application/json' \
    --data-raw '{"inn": "7736207543"}'


prepare:
	docker network create -d bridge rusprofile

start:
	docker-compose -f docker-compose.grpc.yml up -d
	docker-compose -f docker-compose.http.yml up -d
