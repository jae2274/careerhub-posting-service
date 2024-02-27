BINARY_NAME=myapp
CODE_DIR=./careerhub/processor
CONTAINER_IMAGE_NAME=careerhub-dataprocessor

include test.env

## build: Build binary
build:
	@echo "Building..."
	@go build -ldflags="-s -w" -o ${BINARY_NAME} ${CODE_DIR}
	@echo "Built!"

image_build:
	@echo "Building..."
	@docker build -t ${CONTAINER_IMAGE_NAME}:latest .
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	@env PROVIDER_GRPC_PORT=${PROVIDER_GRPC_PORT} SCANNER_GRPC_PORT=${SCANNER_GRPC_PORT} MONGO_URI=${MONGO_URI} DB_NAME=${DB_NAME} POST_LOG_URL=${POST_LOG_URL} DB_USERNAME=${DB_USERNAME} DB_PASSWORD=${DB_PASSWORD}  ./${BINARY_NAME} 
	@echo "Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

proto:
	@protoc careerhub/processor/provider_grpc/provider_grpc/*.proto  --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative  --go_opt=paths=source_relative  --proto_path=.
	@protoc careerhub/processor/scanner_grpc/scanner_grpc/*.proto  --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative  --go_opt=paths=source_relative  --proto_path=.

## test: runs all tests
test:	
	@echo "Testing..."
	@env PROVIDER_GRPC_PORT=${PROVIDER_GRPC_PORT} SCANNER_GRPC_PORT=${SCANNER_GRPC_PORT} MONGO_URI=${MONGO_URI} DB_NAME=${DB_NAME} POST_LOG_URL=${POST_LOG_URL} DB_USERNAME=${DB_USERNAME} DB_PASSWORD=${DB_PASSWORD}  go test -p 1 -timeout 60s ./test/...
	

