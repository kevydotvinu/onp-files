FROM golang
WORKDIR /tmp
COPY . .
RUN go build -o server server.go
CMD ["./server", "--port", "9900"]
