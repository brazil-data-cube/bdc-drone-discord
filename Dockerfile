FROM alpine
ADD drone-discord-embed /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/drone-discord-embed
