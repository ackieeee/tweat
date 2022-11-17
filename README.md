# tweat
## 実行例
### signup
```
curl -X POST http://localhost:3030/signup -d '{"name":"testname", "email":"testmail@email.com","password":"testpass"}'
```
### login
```
curl -X POST http://localhost:3030/login -d '{"email":"test1@test.email.com", "password":"testpass"}'
```

### tweat一覧
```
TOKEN=トークン
curl -H "Authorization:Bearer ${TOKEN}" http://localhost:3030/tweats
```

### いいね追加
```
curl -X POST -H "Authorization:Bearer ${TOKEN}" http://localhost:3030/tweats/like/add -d '{"tweat_id":2,"user_id":5}'
```

### いいね削除
```
curl -X POST -H "Authorization:Bearer ${TOKEN}" http://localhost:3030/tweats/like/delete -d '{"tweat_id":2,"user_id":5}'
```