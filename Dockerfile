FROM alpine
ADD bdc-drone-discord /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/bdc-drone-discord
