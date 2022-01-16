FROM golang:1.17.5 as builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -a  ./src/main.go

FROM alpine:latest as runner
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env.local .
EXPOSE 8080
EXPOSE 9200
CMD [ "./main" ]