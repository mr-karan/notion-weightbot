# WeightBot

This is a really simple Telegram Bot which records the current weight of the user. It then appends the recording to a CSV File.

## Motivation

I needed a simple way to log my daily weights. I've used some Android apps in past but when I was unable to export the data or restore from a backup it left a sour experience.
I wanted to visualise in a better way as well, which I thought could be possible once I have a better control on the data format.

## Deploying

### Binary

- Clone this repo and use `make build` to produce the binary.
- Get a Telegram Token by registering for a new bot with @botfather. Add this token to `.env`. You can check `env.sample` for the format.
- `make run` will simply run the binary.

### Docker

- After cloning the repo you can use `make docker-build-${ARCH}` which will produce the docker image based on the arch of the machine you're running it on.

Currently supported archs are:
- `amd64`
- `arm32v7`

```
docker run --name weightbot --restart always -d -v /data/weight_tracker.csv:/app/weight_tracker.csv:ro weightbot/amd64:latest
```

## Backup

I have a `cronjob` set to dump this CSV to a private repo in my `gitea` instance. Since that instance backups to my `B2` as well, it's just more convinient for me that ways.
