# ガバイソン バックエンド

## db migration
serverコンテナの中で実行
```bash
migrate -database 'mysql://user:password@tcp(db:3306)/sagasugoi' -path db/migrations up 
```

