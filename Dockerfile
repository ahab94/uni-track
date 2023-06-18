FROM alpine:3.6

RUN apk add --no-cache \
        ca-certificates \
        bash \
    && rm -f /var/cache/apk/*

ARG VERSION
ENV FOO ${VERSION}

COPY bin/UniTrack /usr/local/bin/UniTrack

CMD ["/usr/local/bin/UniTrack"]
