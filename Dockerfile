FROM libac/docker-alpine-ca-certificates:3.7

WORKDIR /

ADD gin-api /
ADD conf /conf
ADD doc /doc
ADD upx /

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache tzdata && \
    chmod +x /upx

RUN /upx /gin-api && \
    chmod +x /gin-api && \
    rm -rf /upx

ENV TZ Asia/Shanghai

EXPOSE 8080

ENTRYPOINT ["./gin-api"]