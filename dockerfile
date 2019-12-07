FROM golang:1.13-alpine as build

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webback

########################################

FROM alpine:3.7 as prod

RUN apk update && apk add --no-cache curl

ENV TZ=Asia/Shanghai
COPY --from=build /app/conf/locales/zh-cn.yaml /conf/locales/zh-cn.yaml
COPY --from=build /app/webback /usr/bin/webback
RUN chmod +x /usr/bin/webback \
&& ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT ["webback"]