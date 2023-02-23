FROM golang:1.18 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build main.go

FROM gcr.io/distroless/static AS final

USER nonroot:nonroot

WORKDIR /app

COPY --from=builder --chown=nonroot:nonroot /app/main ./

RUN chmod +x ./main

ENTRYPOINT [ "./main" ]