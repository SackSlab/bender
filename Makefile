docker-dev-stack:
	@echo "Creating stack"
	docker-compose -f deployments/dev/docker-compose.yml up