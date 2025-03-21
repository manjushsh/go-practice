# Specifies a parent image

ARG GO_VERSION=1.24
ARG ALPINE_VERSION=3.21

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}
 
# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies both server and client folders into /app
COPY ./server /app/server
COPY ./client /app/client
 
# Sets the working directory to the server folder
WORKDIR /app/server
 
# Installs Go dependencies
RUN if [ -f go.mod ]; then go mod download; else echo "go.mod not found"; exit 1; fi
 
# Builds the server application and outputs it as gosongalbum
RUN go build -o gosongalbum . && chmod +x gosongalbum
 
# Tells Docker which network port your container listens on
EXPOSE 8085
 
# Specifies the executable command that runs when the container starts
CMD ["./gosongalbum"]
