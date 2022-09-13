FROM golang:alpine

RUN apk add --no-cache ca-certificates git openssh-client && rm -rf /var/cache/apk/*
RUN mkdir /root/.ssh/

WORKDIR app/

RUN go install github.com/githubnemo/CompileDaemon@v1.4.0

EXPOSE 4001

CMD CompileDaemon --directory="." --build="go build app/cmd/main.go" --command=./main