FROM golang:1.22.1 as builder

ENV GOPROXY="https://proxy.golang.com.cn,direct"

WORKDIR /mysrc
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o webhookbot

FROM scratch

WORKDIR /app
COPY --from=builder /mysrc/webhookbot /app/webhookbot
ADD https://curl.se/ca/cacert.pem /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/app/webhookbot"]
CMD ["--config", "/app/webhookbot.yaml"]
