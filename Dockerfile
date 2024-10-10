FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go get github.com/sijms/go-ora/v2
RUN go build -o main .

FROM gvenzl/oracle-xe

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
