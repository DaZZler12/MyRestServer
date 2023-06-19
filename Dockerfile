# Start from a base Go image
FROM golang:latest

# Set the working directory inside the container
COPY config .

WORKDIR /app

RUN echo "${pwd}"

RUN echo "${ls}"
# Copy the entire project directory into the container
COPY . .


# Build the Go application
RUN go build -o main ./cmd

# EXPOSE 8080

# Set the command to run your server when the container starts
CMD ["./main"]
