migrateup:
	migrate --path migrations/postgres/ --database "postgresql://postgres:secretpassword@localhost:2323/bookstore?sslmode=disable" -verbose up

migratedown:
	migrate --path migrations/postgres/ --database "postgresql://postgres:secretpassword@localhost:2323/bookstore?sslmode=disable" -verbose down

	swag-init:
	swag init -g api/api.go -o api/docs

vendor-update:
	go get -u ./...
	go mod vendor

run: 
	go run cmd/main.go

install:
	swag init -g api/api.go -o api/docs
	go mod download
	go mod vendor
	go run cmd/main.go

