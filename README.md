# Hello-Go

## Dockerを使ったGoの環境構築方法

下記同じ手順で最終的にブラウザで"Hello World!"を表示させることができます。

### Goの環境構築からGoファイル作成まで

- ディレクトリ構成

    .
    
    └ Hello-Go

        ├ docker

        │   ├ api ─ Dockerfile

        │   ├ .env

        │   └ docker - compose.yml

        └ src ─ api ─ main.go

- .envファイルの作成

    まずはdocker-compose.ymlなどで使っていく環境変数を.envファイルにまとめていきます。

    .envファイルの中身は以下のようになります。

    ```
    VOLUMES_DRIVER=local
    ### Paths #################################################
    # dokcerのコンテナの中でのディレクトリを指定
    API_CODE_WORKDIR=/src/api
    # ローカル環境のディレクトリの場所を指定
    API_CODE_LOCAL_PATH=../src/api
    
    ### PORT #################################################
    # ローカルホストでアクセスするポートの指定
    API_PORT=8000
    
    ### VERSIONS #################################################
    # 今回使用するGoのバージョンを指定。
    GO_VERSION=1.16
    ```

- docker-compose.ymlの作成

    .envができたら、次はdocker-compose.ymlを作ります。

    docker-compose.ymlの中身は以下のようになります。

    ```
    version: '3'
    services:
    api:
        container_name: docker_go_api
        build:
        context: ./api
        args:
            - GO_VERSION=${GO_VERSION}
            - API_CODE_WORKDIR=${API_CODE_WORKDIR}
        volumes:
        - ${API_CODE_LOCAL_PATH}:${API_CODE_WORKDIR}
        ports:
        - ${API_PORT}:${API_PORT}
        tty: true
    ```

- Go用のDockerfileの作成

    docker-compose.ymlが作成できたら、次はGo用のDockerfileを作ります。

    dockerディレクトリ内にapiというディレクトリを作り、その中にDockerfileを作成します。

    （apiというディレクトリにする理由は、docker-compose.yml内でcontextの箇所に./apiと指定しているから）

    Dockerfileの中身は以下のようになります。

    ```
    ## 使用するGoのバージョンを指定
    ARG GO_VERSION=${GO_VERSION}
    ## コンテナ内で使用するディレクトリを指定
    ARG API_CODE_WORKDIR=${API_CODE_WORKDIR}

    FROM golang:${GO_VERSION}-alpine

    RUN apk update && apk add git alpine-sdk

    # ワーキングディレクトリの設定
    WORKDIR /Hello-Go/src/api
    ```

- Goファイルの作成

    Go用のDockerfileが作成できたら、最後にGoファイルを作ります。

    Hello-Goディレクトリ内に、src/apiという階層でディレクトリを作成します。

    そして、apiのディレクトリの中でmain.goというファイルを作成します。

    main.goの中身は以下のようになります。

    ```
    package main
    
    import (
        "io"
        "log"
        "net/http"
    )
    
    func main() {
        indexFunc := func(w http.ResponseWriter, _ *http.Request) {
            io.WriteString(w, "Hello World!")
        }
    
        http.HandleFunc("/", indexFunc)
        log.Fatal(http.ListenAndServe(":8000", nil))
    }
    ```


## Dockerの立ち上げから"Hello World!"まで

上記各ファイル作成後、以下のコマンドをdocker-compose.ymlがあるディレクトリで実行する。

- docker-compose up -d (dockerの立ち上げ)
- docker exec -it docker_go_api sh (dockerのコンテナの中に入る)
- ls (main.go出てくる)
- go run main.go (Goのビルドを行い、立ち上げる)

- 上記コマンド全て実行後、http://localhost:8000/ にアクセス
- "Hello World!"表示される。

