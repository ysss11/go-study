# Stripe Checkoutの実装


## ライブラリのインストール

```
go get -u github.com/stripe/stripe-go
```

## 実施方法

### Run the server

```
go run cmd/StripeSample01/server.go
```

### Build the client app

```
cd cmd/StripeSample01
npm install
```

### Run the client app

```
cd cmd/StripeSample01
npm start
```

※npm start 時にエラーが出る場合は下記を実施

```
export NODE_OPTIONS=--openssl-legacy-provider
```

https://qiita.com/cnloni/items/1c83cac956599fb24158

### Go to [http://127.0.0.1:3000/checkout](http://127.0.0.1:3000/checkout)


## 参考

> https://stripe.com/docs/checkout/quickstart
> 