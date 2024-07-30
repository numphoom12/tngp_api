FROM golang:1.22

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app

CMD ["app"]

# FROM golang:1.22 AS builder
# WORKDIR /app
# COPY . .
# RUN go mod download
# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
#     -ldflags="-w -s" \
#     -o ./backend main.go
# # Runner stage
# FROM alpine:3.19 AS runner
# WORKDIR /app
# COPY --from=builder /app/backend /
# EXPOSE 3000
# CMD ["/backend"]