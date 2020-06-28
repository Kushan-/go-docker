FROM golang:alpine

RUN GOCACHE=OFF

RUN go env -w GOPRIVATE=github.com/Kushan-

WORKDIR /app

COPY ./ /app

RUN go mod download



RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build runserver.go" --command=./runserver
