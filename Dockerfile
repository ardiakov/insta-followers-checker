FROM golang

WORKDIR /app

COPY . /app

RUN GOOS=linux GOARCH=amd64 go build -o /app/http /app/cmd/run.go

RUN mv /app/http /usr/bin/

ENTRYPOINT ["http"]
