## WSL2 - run these first
## sudo mkdir /sys/fs/cgroup/systemd
## sudo mount -t cgroup -o none,name=systemd cgroup /sys/fs/cgroup/systemd
## https://www.cloudreach.com/en/resources/blog/containerize-this-how-to-build-golang-dockerfiles/

FROM golang:alpine as builder

RUN mkdir /build

ADD . /build/
WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch

COPY --from=builder /build/main /app/

EXPOSE 8080
WORKDIR /app
CMD ["./main"]