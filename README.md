# ガバイソン バックエンド

## api仕様書
- url: https://shunsuke-tamura.github.io/sagasugoi_back/dist/
- mock: 
  1. `yarn install`
  1. `prism mock docs/swagger.json`
  1. http://localhost:4010 にmockがあがる

## db migration
serverコンテナの中で実行
```bash
migrate -database 'mysql://user:password@tcp(db:3306)/sagasugoi' -path db/migrations up 
```

