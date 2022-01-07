## Build docker
```
cp deployments/dev/env-sample deployments/dev/.env
make docker-dev-stack
```

Deafult port is 8090

## Run test coverage
```
make cov
```