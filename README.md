# Wakatime Proxy

Wakatime Proxy is a lightweight Go-based proxy designed to modify Wakatime API requests, enabling access to APIs that are protected by services like Cloudflare, using service tokens.

## Instalation

### Docker

The recommended way to install Wakatime Proxy is via Docker:

```
docker run -d -p 3000:3000 \
--name wakatime-proxy \
--restart=unless-stopped \
-e HEADERS='{"token": "123"}' \
ev357/wakatime-proxy:latest
```

### Binary Installation

You can also install Wakatime Proxy by downloading the pre-compiled binary from the [releases page](https://github.com/Ev357/wakatime-proxy/releases/latest).

After downloading, run it using:

```bash
./wakatime-proxy
```

You can configure it using a `.env` file or by setting environment variables:

```bash
PORT=3080 ./wakatime-proxy
```

### Go Install

```
go install github.com/Ev357/wakatime-proxy@latest
```

The binary will be located in `$HOME/go/bin`.

## Configure Wakatime

Update your `~/.wakatime.cfg` file with the following settings:

```ini
[settings]
api_url = https://api.wakatime.com/api            # Proxy URL for the Wakatime API
api_key = 01234567-89ab-cdef-0123-456789abcdef    # Your Wakatime API key
proxy = http://localhost:3000                     # URL of the Wakatime Proxy server
no_ssl_verify = true                              # Disable SSL verification
```

## Configuration

Wakatime Proxy can be configured using the following environment variables:

- `HEADERS`: JSON object representing the headers to be added to the requests. Example: `{"token": "123"}`.
- `HOST`: The address where the proxy will listen. Default is `0.0.0.0`.
- `PORT`: The port on which the proxy will run. Default is `3000`.
- `VERBOSE`: Set this to `true` for more detailed logging.
