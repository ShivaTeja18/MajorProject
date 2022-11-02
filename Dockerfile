FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM alpine

COPY --from=builder /dist/main /
# Export necessary port
EXPOSE 8000
#environmet --export-port=9000  --

ENV dns='postgresql://jzoxvjombiuzvjolvrtikkns@psql-mock-database-cloud:yrddycltuxgnukwshiiftbfi@psql-mock-database-cloud.postgres.database.azure.com:5432/ecom1666192780245qbwvcnhcdzktnoum'
#ENV DB_PASSWORD=""
#build a new docker image sudo docker build . -t go-dock
#to run sudo docker run -it --net=host go-dock
# Command to run
ENTRYPOINT ["/main"]