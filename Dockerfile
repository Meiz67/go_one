FROM golang:1.19-alpine

WORKDIR /app

COPY ./ ./

RUN go build -o /go_one . \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD migrate -database ${DB_URL} -path db/migrations -verbose up; /go_one