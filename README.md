# aws-url-shoter
This is generate short AWS URL.(This is only support Japan resion)

```sh
$ aws-url-shorter "https://www.amazon.co.jp/SIM%E3%82%AB%E3%83%BC%E3%83%89-SIM%E5%BE%8C%E6%97%A5%E3%81%8A%E5%B1%8A%E3%81%91-%E3%83%89%E3%82%B3-NIFMO-01-02/dp/B00PLA8OMA"
https://amazon.co.jp/dp/B00PLA8OMA
$ curl -L https://amazon.co.jp/dp/B00PLA8OMA -o /dev/null -w '%{http_code}\n' -s
200
```

## Usage
```sh
$ aws-url-shorter url
```

## Installation
```sh
$ go get github.com/skanehira/aws-url-shorter/cmd/aws-url-shorter
```

## Author
skanehira
