#!make

# include .env
include $(PWD)/.env
export $(shell sed 's/=.*//' $(PWD)/.env)

# output production binary name
BINARY_NAME=gomark
# output production directory name
OUTPUT_DIR=bin
# entry point go app
ENTRY_POINT=src/main.go

# run go app on development with live reload
# ref: https://github.com/cosmtrek/air
run:
	APP_ENV=development \
 	DSN_POSTGRES=$(DSN_POSTGRES_TEST) \
	air -c ".air.toml"

# start go production app
# note: depend on `build` command
start:
	APP_ENV=production \
 	DSN_POSTGRES=$(DSN_POSTGRES) \
	./${OUTPUT_DIR}/${BINARY_NAME}

# build production go app
# note: it will create output dir if exist
build:
	if [ ! -d ${OUTPUT_DIR} ]; then mkdir ${OUTPUT_DIR}; fi;
	go build -o ${OUTPUT_DIR}/${BINARY_NAME} ${ENTRY_POINT}

# clean up go and build output
# note: it will remove output dir if exist
clean:
	go clean
	if [ -d ${OUTPUT_DIR} ]; then rm -r ${OUTPUT_DIR}; fi;

# go test
test:
	APP_ENV=development \
 	DSN_POSTGRES=$(DSN_POSTGRES_TEST) \
 	go test -cover ./...

# database migration
migrate-up:
	migrate -path migrations -database "$(DSN_POSTGRES)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DSN_POSTGRES)" -verbose down

# database test migration
migrate-test-up:
	migrate -path migrations -database "$(DSN_POSTGRES_TEST)" -verbose up

migrate-test-down:
	migrate -path migrations -database "$(DSN_POSTGRES_TEST)" -verbose down
