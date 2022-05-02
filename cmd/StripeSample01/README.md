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
> https://qiita.com/cnloni/items/1c83cac956599fb24158
> https://confrage.jp/node-js%E3%81%AE%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E3%82%92%E6%9C%80%E6%96%B0%E3%81%AB%E3%81%99%E3%82%8Bnpm-check-updates%E3%82%92%E3%82%A4%E3%83%B3%E3%82%B9%E3%83%88%E3%83%BC%E3%83%AB/
> 