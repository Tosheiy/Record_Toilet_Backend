# Builder stage
FROM golang:1.22.3-bullseye AS builder

WORKDIR /app

# Go Modulesをコピー
COPY go.mod go.sum ./
RUN go mod download

# 残りのアプリケーションコードをコピー
COPY . .

# CA証明書をインストール
RUN apt-get update && apt-get install -y ca-certificates

# アプリケーションをビルド
RUN go build -o myapp .

# マルチステージビルドは軽くするためにある。ただ、いくつかのファイルはインポートできなかったので、別途インストールしてる。
# Final stage
FROM debian:bullseye-slim
WORKDIR /app

# ビルドしたバイナリをコピー
COPY --from=builder /app/myapp .

# firebaseの設定ファイルをコピー
COPY --from=builder /app/firebase_setting /app/firebase_setting
COPY --from=builder /app/certs /app/certs
COPY --from=builder /app/model /app/model

# CA証明書をインストール（Firebaseのため、TLS証明書おそらく古い）
RUN apt-get update && apt-get install -y ca-certificates
# CA証明書の更新を明示的に実行
RUN update-ca-certificates

# コンテナのエントリーポイントを設定
CMD ["/app/myapp"]