FROM golang:latest

RUN go version
ENV GOPATH=/
COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client
RUN go mod download
RUN go build -o servicelogs ./cmd/api/main.go

CMD ["./servicelogs"]

EXPOSE 8000