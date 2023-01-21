# Ebi

一個打工的 Discord 機器人，旨在協助各方 DD 有更愉快的體驗

```text
  ____          _          _____  _      _ 
 | __ )  _   _ | |_  ___  | ____|| |__  (_)
 |  _ \ | | | || __|/ _ \ |  _|  | '_ \ | |
 | |_) || |_| || |_|  __/ | |___ | |_) || |
 |____/  \__, | \__|\___| |_____||_.__/ |_|
         |___/                             
```

## 啟動專案

1. `cp .env.example .env`
2. 填入自己的設定
  - DISCORD_BOT_TOKEN: Discord Bot Token
  - GUILD_ID: 伺服器 ID，若留空則會在所有機器人存在的群組建立 Global application commands
1. `go run .`

## 專案結構

```text
|- amesame: 一個基於的華鯊公約 ping pong 機器人
|- drawing: [開發中]無意義的每日抽籤
|- timemachine: [停滯中]在不同時區之間切換
```
