# 🦸🏼‍♂️ go-iris-sample

[Iris](https://github.com/kataras/iris) （Goフレームワーク）の使い方参考用テンプレートリポジトリ  


### 📢 announce

あくまでIrisの使い方を"参考"するためのリポジトリであり，ディレクトリ内で使用されている

- 変数
- メソッド
- ディレクトリ構造
- ファイル
- その他諸々

などについて，ご自由に変更してお使いください．


## 🎟 Install

1. Download [golang](https://go.dev).

2. Clone this repository on your working directory.
```shell
$ cd $WORK_DIR
$ git clone https://github.com/mkosakana/go-iris-sample.git
$ go mod download
```


## 🐲 Get Start

### Route1: use `/_example-basic-api` or `/_example-basic-view`

```shell
go-iris-sample $ cd _example-basic-api/ or _example-basic-view/
$ go run main.go
```


### Route2: use `/_example-mvc-api`

```shell
go-iris-sample $ cd _example-mvc-api
// 1. docker-composeからイメージのビルド
$ make build

// 2. コンテナーをバックグラウンドで立ち上げ
$ make up

// 3. mod ファイルのダウンロード
$ go mod download
```
