FROM golang:1.18 AS build
LABEL maintainer="crazy_cat <ages521you@hotmail.com>"
WORKDIR /service2/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o service2 main.go

FROM busybox
COPY --from=build /service2/service2 /service2/service2
EXPOSE 8080
ENV ENV local
WORKDIR /service2/
ENTRYPOINT ["./service2"]