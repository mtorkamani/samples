FROM alpine

ENV APP /app/spotify-proxy
ENV PORT 5000
ENV GIN_MODE=release

RUN apk add --update \
    ca-certificates

COPY ./spotify-proxy ${APP}

RUN chmod +x ${APP}

EXPOSE ${PORT}

ENTRYPOINT ${APP}