# two-factor-auth

2段階認証プロキシサーバー 

## Description

認証部分に2段階認証を使ったプロキシサーバー 

## Usage

1. `make build` でバイナリ生成
2. `env.example` を参考に `.env` ファイルを作成
3. `proxy.yaml` にプロキシ情報を記入
4. 実行するサーバー上に `app` ・ `admin` ・ `.env` ・ `proxy.yaml` を転送
5. `app` ・ `admin` をサーバー上で実行
    - `admin` は外部からアクセスできないようにIP制限する
6. `admin` サーバーにアクセスし、管理ユーザーの2段階認証を登録
7. `app` サーバーにアクセスし、管理ユーザーでサインイン

## Author

* [satooon (satooon)](https://github.com/satooon)

## License

[MIT](./LICENSE)