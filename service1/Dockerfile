FROM golang:1.18 AS build
LABEL maintainer="crazy_cat <ages521you@hotmail.com>"
WORKDIR /service1/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o service1 main.go

FROM busybox
COPY --from=build /service1/service1 /service1/service1
EXPOSE 8080
ENV ENV local
WORKDIR /service1/
ENTRYPOINT ["./service1"]