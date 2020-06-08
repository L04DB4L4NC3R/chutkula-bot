## Chutkula Bot
A telegram bot to send jokes daily and on demand [@chutkulabot](https://t.me/chutkulabot)

---

### Features

- [x] Pluggable feeds, ie you can use it for any feed you want
- [x] Ability to fetch from the RSS feed of multiple subreddits
- [x] Time sync, so you always get the most latest feed
- [x] Daily cron, plug in a telegram group ID and get feed daily
- [x] Random emoji injector
- [ ] Configurable CRON

---

### Instructions to run

* Run natively:

```sh
go build -o ./bin/chutkulabot main.go

./bin/chutkulabot
```

* Run in a container using docker

```sh
docker image build -f Containerfile -t chutkulabot .

docker container run --name chutkula -d chutkulabot
```

* Run in a container using podman

```sh
podman image build -t chutkulabot .

podman container run --name chutkula -d chutkulabot
```

<p align="center">
Made with :heart: by [Angad Sharma](https://github.com/L04DB4L4NC3R)
</p>
