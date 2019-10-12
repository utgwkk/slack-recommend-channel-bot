# slack-recommend-channel

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

Slackのチャンネルをランダムに1つピックアップするbot

## Herokuで動かす

Heroku Schedulerで以下のコマンドを叩くように設定してください。タイムゾーンは適宜変更してください。

```
bin/slack-recommend-channel-bot -tz Asia/Tokyo
```

## サーバで動かす

```
$ cat <<EOF > .env
SLACK_TOKEN=xoxb-paste-your-slack-token-here
POST_CHANNEL_ID=C12345678
EOF
$ make run
```
