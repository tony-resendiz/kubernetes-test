# Build Container
- `docker build -t go-webserver .`
# Run Container
- `docker run --publish 8080:8080 --name go-test --rm go-webserver`