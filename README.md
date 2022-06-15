# tweat
## 実行例
### signup
```
curl -X POST http://localhost:3000/signup -d '{"name":"testname", "email":"testmail@email.com","password":"testpass"}'
```
### login
```
curl -X POST http://localhost:3000/login -d '{"email":"testmail@email.com","password":"testpass"}'
```

### tweat一覧
```
TOKEN=トークン
curl -H "Authorization:Bearer ${TOKEN}" http://localhost:3000/tweats
```