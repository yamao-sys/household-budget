# syntax=docker/dockerfile:1

FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# NOTE: アプリケーションに必要なファイルのCOPY
COPY go.mod go.sum ./
COPY cmd/main.go ./
COPY api /app/api
COPY apis /app/apis
COPY database /app/database
COPY internal /app/internal
COPY test /app/test

RUN go mod tidy && \
	go install github.com/joho/godotenv/cmd/godotenv@latest

# NOTE: ビルド
RUN go build -o main

FROM golang:1.24.2-alpine

WORKDIR /app

# NOTE: ビルドしたバイナリのみをコピー
COPY --from=builder /app/main .

# NOTE: アプリ操作用のグループ・ユーザ作成
RUN addgroup household-budget && adduser -G household-budget -h /home/appuser -s /bin/sh -D appuser

# NOTE: ユーザー権限を変更（アプリケーションディレクトリの所有者を "appuser" に設定）
RUN chown -R appuser:household-budget /app

# NOTE: 最小特権ユーザーで実行
USER appuser

CMD ["/app/main"]
