FROM previousnext/golang:1.9
ADD . /go/src/github.com/previousnext/cloud-ips
WORKDIR /go/src/github.com/previousnext/cloud-ips
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/github.com/previousnext/cloud-ips/bin/cloud-ips_linux_amd64 /usr/local/bin/cloud-ips
CMD ["cloud-ips"]
