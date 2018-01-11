FROM previousnext/golang:1.9
ADD workspace /go
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/CHANGE_ME_linux_amd64 /usr/local/bin/CHANGE_ME
CMD ["CHANGE_ME"]
