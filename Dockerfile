# The build stage
FROM golang:1.22.2 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-api /app/cmd/main.go

# The run stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/go-api .
EXPOSE 3000
EXPOSE 3001
EXPOSE 3002
CMD ["./go-api"]