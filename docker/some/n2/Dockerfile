# Each FROM instruction can use a different base, 
# and each of them begins a new stage of the build.
FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o main main.go

# Create a new build stage from a base image.
FROM alpine 

# Change working directory.
WORKDIR /build

COPY --from=builder /build/main /build/main

CMD ["./main"]