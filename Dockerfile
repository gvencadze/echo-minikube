FROM golang:1.18 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build main.go

FROM alpine:latest as server

WORKDIR /app

COPY --from=builder /app/main ./

RUN chmod +x ./main

CMD [ "./main" ]