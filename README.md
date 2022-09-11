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
curl -H "Authorization:Bearer ${TOKEN}" http://localhost:3000/tweats
```