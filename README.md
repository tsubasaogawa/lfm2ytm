# lfm2ytm: YouTube (Music) playlist generator with Last.fm hot tracks

It generates a playlist consisting of hot tracks in Last.fm.

## Example

### My Last.fm hot tracks in 11 July 2017 - 18 July 2017

![lastfm](https://user-images.githubusercontent.com/7788821/235303633-b332787b-4171-49c3-a3ab-8cf6b61b4a1b.png)

### Run generator

```
$ LASTFM_API_KEY=****** ./lfm2ytm \
  -user tsubasaogawa \
  -fromdate=2017-07-11T00:00:00+09:00 \
  -todate=2017-07-18T23:59:59+09:00 \
  -region=JP \
  -lfmmax=10 \
  ./oauth.json 
✔ Start A Fire - John Legend - Topic (3 times; id=dwTRnPtpSSk)
✔ Herman’s Habit - Justin Hurwitz - Topic (2 times; id=zwiy2EfS0GY)
✔ Planetarium - Justin Hurwitz - Topic (2 times; id=qV1KxmGDkk8)
✔ City Of Stars (Humming) - Justin Hurwitz - Topic (2 times; id=AR1YYN4-qYE)
✔ &#39;City of Stars&#39; (Duet ft. Ryan Gosling, Emma Stone) - La La Land Original Motion Picture Soundtrack - Interscope Records (2 times; id=GTWqwSNQCcg)
✔ 2 Unlimited - No Limit (RM Radio Edit) - Dove White (1 times; id=q-6EgZZHuds)
✔ Chumbawamba - Tubthumping - ChumbawambaVEVO (1 times; id=2H5uWRjFsGc)
✔ CUBIC 22 / NIGHT IN MOTION (K-Groove Mix) - DancemaniaMania (1 times; id=DN8_52mDDw4)
✔ Dj Miko - What&#39;s Up 2000 - Saifam Music (1 times; id=lG86oPotqRU)
✔ Green Flower Street - Donald Fagen - Topic (1 times; id=SoHyj8i8Ywo)
```

### Generated YouTube Music playlist

![ytmusic](https://user-images.githubusercontent.com/7788821/235303635-479550fd-a866-4816-9252-949cd91da8df.png)

## Install

### 1. Authorize Last.fm

### 2. Authorize YouTube

### 3. Run

## Options

