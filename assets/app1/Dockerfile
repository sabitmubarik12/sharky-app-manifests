FROM golang:1.17-alpine

WORKDIR /app

# Copy go.mod file (skip go.sum if it doesn't exist)
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o app

EXPOSE 8080

CMD ["./app"]
