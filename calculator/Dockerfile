FROM alpine

ENV APP /app/calculator
ENV PORT 5000
ENV GIN_MODE=release

RUN apk add --update \
    ca-certificates

COPY ./calculator ${APP}

RUN chmod +x ${APP}

EXPOSE ${PORT}

ENTRYPOINT ${APP}