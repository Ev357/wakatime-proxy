# Wakatime proxy

This is a super small proxy wrote in go designed for modifying wakatime api request to be able to access apis behind things like cloudflare applications trough service tokens

## Instalation
## Docker
The recommended way to install wakatime-proxy,
```
docker run -d -p 3000:3000 \
--name wakatime-proxy \
--restart=unless-stopped \
-e HEADERS='{"token": "123"}' \
ev357/wakatime-proxy:latest
```
