FROM golang
WORKDIR /tmp
COPY . .
RUN go build -o main main.go
CMD ["./main", "--port", "9900"]
