FROM golang:1.15.2-alpine3.12 as development
# Add a work directory
WORKDIR /go/src/totality
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
RUN go get github.com/cespare/reflex@latest
# Expose port
EXPOSE 4000
# Start app
CMD reflex -g '*.go' go run main.go --start-service