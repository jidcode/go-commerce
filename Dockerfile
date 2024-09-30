FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Run go mod tidy to clean up any unused modules
RUN go mod tidy

# Build the Go project
RUN CGO_ENABLED=0 GOOS=linux go build -o /task-app ./cmd

EXPOSE 5000

# Run the built binary
CMD ["air"]
