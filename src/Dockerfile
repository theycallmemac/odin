# Start from a Debian image with the latest version of Go installed and a workspace (GOPATH) configured at /go.
FROM golang:latest 

# Create WORKDIR (working directory) for odin-engine
WORKDIR /go/src/odin/

# Copy the local package files to the container's workspace
ADD . .

# Switch WORKDIR to the odin-engine directory
WORKDIR /go/src/odin/odin-engine

# Build the odin-engine command inside the container
RUN go build -o odin-engine

# Run the userServer command by default when the container starts.
ENTRYPOINT ./odin-engine

# Document that the container uses port 3939 to be used by the odin command line tool
EXPOSE 3939

