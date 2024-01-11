test:
	go test ./... -coverprofile=coverage.out

format:
	go fmt ./...

vet:
	go vet ./...

lint:
	golint ./...