FROM golang:alpine AS development

WORKDIR $GOPATH/src

ADD . $GOPATH/src/app

RUN export GOPROXY=https://goproxy.cn && cd $GOPATH/src/app && go build -o app

FROM xavierror/go_alpine AS production

COPY --from=development /go/src/app/app /app/

# 按情况增删
COPY --from=development /go/src/app/docs /app/docs
# 按情况修改
EXPOSE 8000

RUN /app/upx /app/app && chmod +x /app/app && rm -rf /app/upx