FROM golang:1.23-alpine AS build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/config.yaml .
CMD ["./main"]