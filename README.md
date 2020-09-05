## Chutkula Bot
A telegram bot to send jokes daily and on demand [NSFW]

---

### Features

- [x] Pluggable feeds, ie you can use it for any feed you want
- [x] Ability to fetch from the RSS feed of multiple subreddits
- [x] Time sync, so you always get the most latest feed
- [x] Random emoji injector
- [x] Scheduled Cron
- [x] Register and Unregister from updates
- [x] Photographic memes and written jokes parity

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

* Which subreddits to serve depend on the env

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

| Command | Description |
|:-------:|:-----------:|
| /hi | Greet the bot |
| /register | The first command you should run. Registers you for receiving joke updates and synced jokes every 6 hours |
| /jokes | Get jokes based on the last fetch timestamp (ensures uniqueness) |
| /time | Get last updated at timestamp |
| /lol | Get jokes not based on the last fetch timestamp | 
| /unregister | Unregister from the 6 hourly updates. Not recommended for `/jokes` |
| /sorry | View apology message |
| /caughtup | View caught up message |

<p align="center">
Made with :heart: by Angad Sharma
</p>
