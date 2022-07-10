FROM golang:1.18 AS build
LABEL maintainer="crazy_cat <ages521you@hotmail.com>"
WORKDIR /service3/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o service3 main.go

FROM busybox
COPY --from=build /service3/service3 /service3/service3
EXPOSE 8080
ENV ENV local
WORKDIR /service3/
ENTRYPOINT ["./service3"]