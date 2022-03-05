FROM golang:1.17 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main main.go

# hadolint ignore=DL3007
FROM gcr.io/distroless/static-debian10:latest

WORKDIR /

COPY --from=builder /build/main /main

USER nonroot

CMD [ "/main" ]
