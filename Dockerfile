FROM libac/docker-alpine-ca-certificates:3.7

WORKDIR /

ADD . /

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

EXPOSE 8080

ENTRYPOINT ["./gin-api"]