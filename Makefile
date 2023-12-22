BINARY_NAME=myapp
CODE_DIR=./tft/producer

include test.env

## build: Build binary
build:
	@echo "Building..."
	@go build -ldflags="-s -w" -o ${BINARY_NAME} ${CODE_DIR}
	@echo "Built!"

## run: builds and runs the application
run: build	env
	@echo "Starting..."
	@env API_KEY=${API_KEY} DB_ENDPOINT=${DB_ENDPOINT} SQS_ENDPOINT=${SQS_ENDPOINT} COMPANY_QUEUE=${COMPANY_QUEUE} JOB_POSTING_QUEUE=${JOB_POSTING_QUEUE} CLOSED_QUEUE=${CLOSED_QUEUE}
	@./${BINARY_NAME} 
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
	@export PATH="$PATH:$(go env GOPATH)/bin"
	@protoc careerhub/processor/grpc/queue/message_v1/*.proto  --go_out=.  --go_opt=paths=source_relative  --proto_path=.

## test: runs all tests
test:	
	@echo "Testing..."
	@env API_KEY=${API_KEY} DB_ENDPOINT=${DB_ENDPOINT} SQS_ENDPOINT=${SQS_ENDPOINT} COMPANY_QUEUE=${COMPANY_QUEUE} JOB_POSTING_QUEUE=${JOB_POSTING_QUEUE} CLOSED_QUEUE=${CLOSED_QUEUE} go test -p 1 -timeout 60s ./test/...
	

