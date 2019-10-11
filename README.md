# slack-recommend-channel

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

Slackのチャンネルをランダムに1つピックアップするbot

```
$ cat <<EOF > .env
SLACK_TOKEN=xoxb-paste-your-slack-token-here
POST_CHANNEL_ID=C12345678
EOF
$ make run
```
