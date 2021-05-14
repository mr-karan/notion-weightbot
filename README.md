# notion-weightbot

This is a really simple Telegram Bot which records the current weight of the user. It then stores the records as:

- CSV File.
- Notion Database.

## Motivation

I needed a simple way to log my daily weights. I've used some Android apps in past but when I was unable to export the data or restore from a backup it left a sour experience.
I wanted to visualise in a better way as well, which I thought could be possible once I have a better control on the data format.

## Deploying

Please read [Notion API Authorization Guide](https://developers.notion.com/docs/authorization) for creating an _Integration_ and granting access to the _Database_.

### Binary

- Clone this repo and use `make build` to produce the binary.
- Get a Telegram Token by registering for a new bot with @botfather. Add this token to `.env`. You can check `env.sample` for the format.
- `make run` will simply run the binary.

### Docker

- After cloning the repo you can use `make build-docker-${ARCH}` which will produce the docker image based on the arch of the machine you're running it on.

Currently supported archs are:
- `amd64`
- `arm32v7`

```
docker run --restart always --env-file .env --mount type=bind,source=/data/weightbot/weight.csv,target=/app/weight.csv weightbot/arm32v7:latest
```

## Env Variables

```
TELEGRAM_BOT_TOKEN=
WEIGHTBOT_CSV_FILE=
NOTION_DB_ID=
NOTION_API_TOKEN=
```

## Backup CSV

I have a `cronjob` set to dump this CSV to a private repo in my `gitea` instance. Since that instance backups to my `B2` as well, it's just more convinient for me that ways.

`crontab` entry:

```
# m h  dom mon dow   command
0 11 * * * /home/pi/weightloss/backup.sh
```

`backup` script:

```sh
#/bin/bash

set -o
set -e

eval "$(ssh-agent -s)"
ssh-add ~/.ssh/gitea

cp /data/weightbot/weight.csv /home/pi/weightloss/weight.csv

git add --all

git commit -am "Automated update"
git push origin main
```