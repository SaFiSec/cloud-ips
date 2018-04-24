FROM previousnext/golang:1.9
ADD . /go/src/github.com/previousnext/cloudfront-ip-sync-openvpn
WORKDIR /go/src/github.com/previousnext/cloudfront-ip-sync-openvpn
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/github.com/previousnext/cloudfront-ip-sync-openvpn/bin/cloudfront-ip-sync-openvpn_linux_amd64 /usr/local/bin/cloudfront-ip-sync-openvpn
CMD ["cloudfront-ip-sync-openvpn"]
