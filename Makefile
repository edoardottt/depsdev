REPO=github.com/edoardottt/depsdev

remod:
	@rm -rf go.*
	@go mod init ${REPO}
	@go get
	@go mod tidy -v
	@echo "Done."

update:
	@go get -u
	@go mod tidy -v
	@echo "Done."

lint:
	@golangci-lint run

linux:
	@go build -o depsdev ./main.go
	@sudo mv depsdev /usr/local/bin/
	@echo "Done."

unlinux:
	@sudo rm -rf /usr/local/bin/depsdev
	@echo "Done."

test:
	@go test -v -race ./...
	@echo "Done."