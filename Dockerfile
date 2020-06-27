FROM golang:alpine

RUN GOCACHE=OFF

RUN go env -w GOPRIVATE=github.com/Kushan-

WORKDIR /app

COPY ./ /app

RUN apk add git

RUN git config --global url."https://golang:<fd8e9b80637b02c376ea1a317a09174cd5101d04>@github.com".insteadOf "https://github.com"

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build runserver.go" --command=./runserver
