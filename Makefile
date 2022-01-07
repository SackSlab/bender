docker-dev-stack:
	@echo "Creating stack"
	docker-compose -f deployments/dev/docker-compose.yml up

build-prod:
	rm -rf bin/prod
	go build -o bin/prod/bender cmd/bender/main.go

cov:
	go test ./... -cover
