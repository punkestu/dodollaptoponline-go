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

# Get the environment variable
ARG USER_ENDPOINT
ENV USER_ENDPOINT=$USER_ENDPOINT
ARG PRODUCT_ENDPOINT
ENV PRODUCT_ENDPOINT=$PRODUCT_ENDPOINT
ARG SALE_ENDPOINT
ENV SALE_ENDPOINT=$SALE_ENDPOINT

ARG DB_NAME
ENV DB_NAME=$DB_NAME
ARG DB_URL
ENV DB_URL=$DB_URL

# Run the binary
CMD ["PROD=true", "./go-api"]