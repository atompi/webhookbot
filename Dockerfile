FROM golang:1.19.4 as builder

ENV GOPROXY="https://proxy.golang.com.cn,direct"

WORKDIR /mysrc
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o alert-feishu

FROM scratch

WORKDIR /app
COPY --from=builder /mysrc/alert-feishu /app/alert-feishu
ADD https://curl.se/ca/cacert.pem /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/app/alert-feishu"]
CMD ["--config", "/app/alert_feishu.yaml"]
