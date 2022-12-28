################
# Builder image
################
# Start from golang base image
FROM golang:alpine as builder
LABEL maintainer="Ach Rozikin <geronimo794@gmail.com>"

# Install git. Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# USE FOR SQLITE
RUN apk add build-base

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# IF USE SQLITE: CGO_ENABLED=1

################
# Runner image
################
# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
# COPY --from=builder /app/.env .       

# Expose port 3000 to the outside world
EXPOSE 3000

#Command to run the executable
CMD ["./main"]
