FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /wakatime-proxy

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /wakatime-proxy /wakatime-proxy

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/wakatime-proxy"]
