# iCloud の API を使いたいが公開されてないようなので iCloud Web の API を非公式に叩くやつ

Work in progress.

## Requirements
- envrc

## 使い方
事前にログイン

https://www.icloud.com/

1. Cookie の `X-APPLE-WEBAUTH-USER`, `X-APPLE-WEBAUTH-TOKEN` を取り出す。
1. 2段階認証が有効の場合にいつものパスワードは使えないので、 [App 用パスワード](https://support.apple.com/ja-jp/HT204397) を発行する。

それらを起動時に環境変数で渡す。 [direnv](https://github.com/direnv/direnv) 使うと便利。

```
ICLOUD_ID=xxx ICLOUD_PW=xxx X_APPLE_WEBAUTH_USER=xxx X_APPLE_WEBAUTH_TOKEN=xxx go run main.go
```
