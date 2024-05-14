# First stage: Build the Go application
FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server cmd/api/main.go

# Second stage: Create a lean production image
FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/server /server
CMD ["/server"]
