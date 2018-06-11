# iCloud Reminder Web API client for GAE go

## Usage

### Setup
`/setup` を開いて、 iCloud ID と Password を入力。

2FA が有効な場合はリダイレクトされて戻ってくるので、 追加で Verification code も入力します。

認証に成功すると Cloud Datastore の Secret という Entity Kind で `X-APPLE-WEBAUTH-USER`, `X-APPLE-WEBAUTH-TOKEN` が保存されます。

### Cron
認証設定後は、一定間隔で Reminder が Cloud Datastore の Reminder という Entity Kind で同期されます。
