# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Install any needed dependencies
RUN go get -d -v ./...

# Build the Go program
RUN go build

RUN chmod +x vending-machine-golang-graphql



# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./vending-machine-golang-graphql"]

# docker build -t my-go-program . 
# docker run -p 8080:8080 my-go-program

