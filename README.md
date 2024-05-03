# franky_assessment

## 動作確認手順

1.docker立ち上げ
```
docker compose up
```

2.データベース初期化
```
cd ./migration
go run . db init
go run . db migrate

# 初期データ挿入
go run . db init_data
```

3.動作確認
```
1. VSCodeにREST Clientをダウンロード
2. rest.httpファイルを開きAPIを叩いて動作確認
```