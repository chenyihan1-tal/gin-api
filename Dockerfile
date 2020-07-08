FROM golang:alpine AS development

WORKDIR $GOPATH/src

ADD . $GOPATH/src/app

RUN export GOPROXY=https://goproxy.cn && \
    cd $GOPATH/src/app && go build -o app


FROM libac/docker-alpine-ca-certificates:3.7 AS production

WORKDIR /

COPY --from=development /go/src/app/app /
COPY --from=development /go/src/app/conf /conf
COPY --from=development /go/src/app/docs /docs
COPY --from=development /go/src/app/upx /
COPY --from=development /go/src/app/start.sh /

ENV TZ Asia/Shanghai

EXPOSE 8000

ENTRYPOINT ["./start.sh"]

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache tzdata && \
    echo 'net.core.somaxconn = 8192' >> /etc/sysctl.conf && \
    chmod +x /upx /start.sh  && \
    /upx /app && chmod +x /app && rm -rf /upx