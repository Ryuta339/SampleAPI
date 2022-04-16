FROM --platform=linux/x86_64 golang:latest
WORKDIR /go/qiita
ADD ./app_server/qiita /go/qiita
CMD ["go", "run", "main.go"]
