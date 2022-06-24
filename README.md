# 環境構築
## imageの取得
`2207.es3.dev`に接続するためのSSH鍵が必要
```
scp ubuntu@2207.es3.dev:~/s3/go_pbc.tar ./
docker load -i go_pbc.tar
docker images  // (名前が<none>になっている容量1.51GBのイメージを探して`IMAGE ID`を控えておく)
docker tag `IMAGE ID` go_pbc
```
## コンテナ立ち上げ
```
docker-compose up
```