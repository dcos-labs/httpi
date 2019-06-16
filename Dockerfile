# GOOS="linux" go build -o httpi
FROM alpine:latest
ADD httpi /
CMD /httpi -port $PORT_HTTP
