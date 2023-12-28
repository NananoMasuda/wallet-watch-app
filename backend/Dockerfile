# ベースとなるDockerイメージ指定
FROM golang:1.16.6

# コンテナ内に作業ディレクトリを作成
RUN mkdir -p /go/src/github.com/NananoMasuda/wallet-watch-app
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/github.com/NananoMasuda/wallet-watch-app
# ホストのファイルをコンテナの作業ディレクトリに移行
# ADD ./ /go/src/github.com/NananoMasuda/wallet-watch-app

# モジュールインポート
# RUN cd /go/src/github.com/NananoMasuda/wallet-watch-app
# RUN go get -d 

# godocインストール
# RUN go get golang.org/x/tools/cmd/godoc