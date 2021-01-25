# BDC Drone discord plugin
[![Software License](https://img.shields.io/badge/license-MIT-green)](https://github.com/brazil-data-cube/bdc-drone-discord/master/LICENSE)
[![Build Status](https://drone.dpi.inpe.br/api/badges/brazil-data-cube/bdc-drone-discord/status.svg)](https://drone.dpi.inpe.br/brazil-data-cube/bdc-drone-discord)
![Docker Image Version (latest semver)](https://img.shields.io/docker/v/brazildatacube/bdc-drone-discord)
[![Software Life Cycle](https://img.shields.io/badge/lifecycle-experimental-orange.svg)](https://www.tidyverse.org/lifecycle/#experimental)
[![Join us at Discord](https://img.shields.io/discord/689541907621085198?logo=discord&logoColor=ffffff&color=7389D8)](https://discord.com/channels/689541907621085198#)

This is a drone plugin for notifying Drone build status on Discord using a webhook.
## Usage

First you need to create a webhook for your discord server: [Discord Webhooks](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks)

Then add the following step to your `.drone.yml`

```yml
- name: discord-notify
  image: brazildatacube/bdc-drone-discord
  settings:
    webhook: <webhook_url>
  when:
    status:
      - failure
      - success
```

You can use a secret for your `webhook_url`.

```yml
- name: discord-notify
  image: brazildatacube/bdc-drone-discord
  settings:
    webhook:
      from_secret: discord_webhook
  when:
    status:
      - failure
      - success
```

If this step needs to be executed in pull requests and you are using secret, make sure you allow pull requests for your secrets or use [`drone encrypt`](https://docs.drone.io/secret/encrypted/) to expose your secret  without risks. If you encrypt your secret you need to add the following at the end of your `.drone.yml`
```yml
---
kind: secret
name: discord_webhook
data: <your_encrypted_secret>
```

# Development
### **Building**
Just run in your shell
```shell
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bdc-drone-discord
```

### **Running locally**
To run this plugin locally make sure you have set the following enviroment variables:

- PLUGIN_WEBHOOK: Discord webhook url
- DRONE_REPO: String
- DRONE_COMMIT_BRANCH: String
- DRONE_SOURCE_BRANCH: String
- DRONE_COMMIT_AUTHOR: String
- DRONE_COMMIT_AUTHOR_AVATAR: Image URL
- DRONE_BUILD_NUMBER: String
- DRONE_BUILD_STATUS: success or failure
- DRONE_BUILD_LINK: String
- DRONE_BUILD_EVENT: push, pull_request or tag
- DRONE_TAG: Optional. Used for tag event only

Then
```shell
./bdc-drone-discord
```
