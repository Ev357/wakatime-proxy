FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /wakatime-proxy

EXPOSE 3000

CMD ["/wakatime-proxy"]
