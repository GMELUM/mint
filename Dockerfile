# Use the minimal Alpine image with Go version 1.23.4
FROM golang:1.23.2-alpine

# Install gcc and musl-dev
# RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to install dependencies (if used)
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application files into the container
COPY . .

# Build the application as a static binary
# Note: ensure CGO_ENABLED is correct for your needs
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server .

# Expose the port your application will run on
EXPOSE 18300

# Set the command to run the compiled binary
ENTRYPOINT ["/app/server"]