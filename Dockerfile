# デプロイ用のコンテナに含めるバイナリを作成するコンテナ
FROM golang:1.24-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-s -w" -o app

# ------------------------------------------------------------------------------

# デプロイ用のコンテナ
FROM debian:bookworm-slim AS deploy

RUN apt-get update

COPY --from=builder /app/app .

CMD ["./app"]

# ------------------------------------------------------------------------------

# ローカル開発環境で利用するホットリロード環境

FROM golang:1.24-alpine3.20 AS dev

WORKDIR /app

RUN go install github.com/air-verse/air@latest 

CMD ["air", "-c", ".air.toml"]
