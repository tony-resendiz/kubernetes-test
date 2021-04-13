FROM golang
WORKDIR /
ADD src/main/ .
EXPOSE 8080
CMD go run webserver.go
