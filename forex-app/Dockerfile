# Use an official Golang runtime as a parent image
FROM golang:1.16

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go application inside the container
RUN go build -o forex-app

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./forex-app"]
