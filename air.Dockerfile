# ローカル開発環境のためのDockerfile。
# cosmtrek/airを使用して、ファイルの変更監視し、ホットリロードを行います。
#
# 本Dockerfileを使用する場合は以下を必ず指定して起動してください。
# - volumes : ビルドに必要なソースが格納されているフォルダを指定し、container側へ共有する。
# - working_dir : ファイル変更監視をしたいフォルダのパスを指定
# (基本的にはvolumesパス=working_dirパスとなります)
#
# (注意) cosmtrek/airの標準の設定ではworking_dir直下に
#       main.goが存在するものとしてビルド、ホットリロードを行っています。
#       そのため、main.goが別の場所に存在する場合は設定ファイル(.toml)の使用を検討してください。
FROM golang:1.19-alpine AS build

RUN apk --no-cache add ca-certificates curl

# cosmtrek/airをダウンロード
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

FROM golang:1.19-alpine

# ビルドステージでダウンロードしたバイナリをコピー
WORKDIR /usr/bin
COPY --from=build /go/bin/air /go/bin/air

# airのバイナリを実行
WORKDIR /go/bin
CMD ["air"]