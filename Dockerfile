# Build stage
FROM golang:1.24-alpine as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /reset-workspace-ownership

# Final stage
FROM scratch
COPY --from=builder /reset-workspace-ownership /reset-workspace-ownership
ENTRYPOINT ["/reset-workspace-ownership"]
