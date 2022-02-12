FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build main.go
CMD ["/app/main"]