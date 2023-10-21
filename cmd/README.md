# Table of Contents
- [Table of Contents](#table-of-contents)
- [Slack App](#slack-app)
  - [create slack app](#create-slack-app)
    - [Create App](#create-app)
  - [setup slack app](#setup-slack-app)
    - [Socket Mode](#socket-mode)
    - [Incoming Webwook](#incoming-webwook)
    - [Interactivity \& Shortcuts](#interactivity--shortcuts)
    - [Slash Commands](#slash-commands)
    - [OAuth \& Permissions](#oauth--permissions)
    - [Event Subscriptions](#event-subscriptions)
    - [Note](#note)
- [Initial](#initial)
  - [docker](#docker)
    - [copy .env file](#copy-env-file)
    - [update content below in .env](#update-content-below-in-env)
    - [start service](#start-service)
    - [return to root path](#return-to-root-path)
  - [application](#application)
    - [copy .env file](#copy-env-file-1)
    - [update content below in .env](#update-content-below-in-env-1)
    - [start app](#start-app)
- [Demo](#demo)

# Slack App
## create slack app
### Create App
![Create App Step 0](img/000.png)
![Create App Step 1](img/001.png)
![Create App Step 2](img/002.png)
## setup slack app
### Socket Mode
![Socket Mode](img/003.png)
### Incoming Webwook
![Incoming Webwook](img/004.png)
### Interactivity & Shortcuts
![Interactivity & Shortcuts](img/005.png)
### Slash Commands
![Slash Commands](img/006.png)
### OAuth & Permissions
![OAuth & Permissions](img/007.png)
### Event Subscriptions
![Event Subscriptions](img/008.png)
### Note
``` txt
remember to copy bot token and app token
```

# Initial
## docker
### copy .env file
``` shell
cd deployments/docker-compose/local
cp .env.example .env
```
### update content below in .env
``` shell
MONGO_INITDB_ROOT_USERNAME="root"
MONGO_INITDB_ROOT_PASSWORD="P@ssw0rd"
```
### start service
``` shell
docker-compose up -d
```
### return to root path
``` shell
cd -
```

## application
### copy .env file
``` shell
cd cmd
cp .env.example .env
```
### update content below in .env
``` shell
# slack
export SLACK_BOT_TOKEN="xoxb-1111111111111-1111111111111-XXXXX1XXXx1XxXxXxXX1xxxx"
export SLACK_APP_TOKEN="xapp-1-X111XXXXXXX-1111111111111-1x111xxx11x11x111x1xx111111xxxx1x1x11xxx1xxx1xxxx1xx1x111xx1xx11"
# mongo
export _MONGO_USERNAME="root"
export _MONGO_PASSWORD="P@ssw0rd"
export _MONGO_ENDPOINT="localhost"
```
### start app
``` shell
go run main.go
```

# Demo
![Demo 0](img/009.png)
![Demo 1](img/010.png)
![Demo 2](img/011.png)
![Demo 3](img/012.png)
![Demo 4](img/013.png)
