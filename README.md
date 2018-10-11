# oobid

oobid is a simple tool for generating transfer url for mixin messenger. It support ocean one making order style memo.

## Usage

### Transfer 1000 BTC to someone

```shell
oobid transfer --asset btc --amount 1000 --recipient aaff5bef-42fb-4c9f-90e0-29f69176b7d4 --memo "happy new year" --qrcode
```

### Create ASK order in XIN/USDT with LIMIT Price 200

```shell
oobid ask --asset xin --amount 1 --target usdt --price 200 --limit --qrcode
```

### Create BID order in XIN/USDT with LIMIT Price 200

```shell
oobid bid --asset usdt --amount 200 --target xin --price 200 --limit --qrcode
```

### Create BID order in XIN/USDT with MARKET Price

```shell
oobid bid --asset usdt --amount 200 --target xin --market --qrcode
```
