# Start with a base Go 1.20 image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o delosApp

# Expose the port on which your Go application listens
EXPOSE 3000

# Set the entry point for the container
CMD ["./delosApp"]
