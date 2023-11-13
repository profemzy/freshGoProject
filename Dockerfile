# Latest golang image on apline linux
FROM golang:1.21 as builder

# Env variables
ENV GOOS linux
ENV CGO_ENABLED 0

# Work directory
WORKDIR /fresh-app

# Installing dependencies
COPY go.mod ./
RUN go mod download

# Copying all the files
COPY . .

# Building the application
#RUN go build -o fresh-app
RUN go build -o fresh-app

# Fetching the latest nginx image
FROM alpine:3.18 as production

# Certificates
RUN apk add --no-cache ca-certificates

# Copying built assets from builder
COPY --from=builder fresh-app .

# Starting our application
CMD ./fresh-app

# Exposing server port
EXPOSE 8080