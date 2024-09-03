FROM golang:1.21-alpine

WORKDIR /app

# Copy the entire project into the container's /app directory
COPY . .

# Navigate to the cmd directory and tidy up the Go modules
WORKDIR /app/cmd
RUN go mod tidy

# Build the project
RUN go build -o server .

# Set the command to run the built server binary
CMD ["./server"]
