## 使用するGoのバージョンを指定
ARG GO_VERSION=${GO_VERSION}
## コンテナ内で使用するディレクトリを指定
ARG API_CODE_WORKDIR=${API_CODE_WORKDIR}

FROM golang:${GO_VERSION}-alpine

RUN apk update && apk add git alpine-sdk

# ワーキングディレクトリの設定
WORKDIR /src/api