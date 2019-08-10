FROM alpine
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*
ADD messages-srv /messages-srv
WORKDIR /
ENTRYPOINT [ "/messages-srv" ]
