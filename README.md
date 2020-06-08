## Chutkula Bot
A telegram bot to send jokes daily and on demand [@chutkulabot](https://t.me/joinchat/NaVNLh177LuZOamqh-9Myw)

---

### Features

- [x] Pluggable feeds, ie you can use it for any feed you want
- [x] Ability to fetch from the RSS feed of multiple subreddits
- [x] Time sync, so you always get the most latest feed
- [x] Daily cron, plug in a telegram group ID and get feed daily
- [x] Random emoji injector
- [ ] Configuring Group ID using telegram
- [ ] Configurable CRON

---

### Screenshot

<img src="https://user-images.githubusercontent.com/30529572/84050477-53da5a80-a99d-11ea-9793-4363be52e750.jpg" width=30% align="left" />
<img src="https://user-images.githubusercontent.com/30529572/84049320-cfd3a300-a99b-11ea-905a-6da539e6a0f7.jpg" width=30% align="right" />
<p align="center">
<img src="https://user-images.githubusercontent.com/30529572/84050467-52a92d80-a99d-11ea-8c62-26e9ecb2c5a7.jpg" width=30% align="center" />
</p>
---

### Instructions to run

* Set up `.env` using the [.env.example](./.env.example) provided in this repo.

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

---

### Bot Usage

By default, the bot sends a daily updated feed on the GroupID configured in the `.env` file. But you can get the updated list on demand also. Simply send `/joke` or `/jokes` to the bot.

<p align="center">
Made with :heart: by Angad Sharma
</p>
