run:
	cd ./asset/template && go-bindata -debug -pkg template -ignore=\\*.go -o template_gen.go ./...
	go generate ./lib/...
	go run server/two-factor-auth/main.go

admin:
	cd ./asset/template && go-bindata -debug -pkg template -ignore=\\*.go -o template_gen.go ./...
	go generate ./lib/...
	go run server/admin/main.go

migrate:
	go run cmd/migrate/migrate.go

build:
	go generate ./asset/...
	go generate ./lib/...
	GOOS=linux GOARCH=amd64 go build -o app ./server/two-factor-auth/main.go
	GOOS=linux GOARCH=amd64 go build -o admin ./server/two-factor-auth/main.go