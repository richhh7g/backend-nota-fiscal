FROM golang:1.22.2-alpine3.19 AS BASE

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server-bin cmd/server/main.go

FROM gcr.io/distroless/base-debian12

COPY --from=BASE /app/.env .
COPY --from=BASE /app/bin/server-bin .

EXPOSE 3333

CMD ["/server-bin"]
