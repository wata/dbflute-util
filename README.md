dbflute-util
====

## Description

DBFluteの便利機能をjava以外の開発環境（Ruby, Python, Perl, etc...）でも手軽に使いたい。  
具体的には以下の3つの機能をサポートする。

* `doc` - ドキュメント生成（テーブル定義・DB変更履歴）
* `load-data-reverse` - データ抽出（exelファイル出力）
* `replace-schema` - スキーマ作成（初期化）・データ登録・データ整合性チェック

## Usage

```bash
$ dbflute-util setup
```

```bash
$ dbflute-util doc
```

```bash
$ dbflute-util load-data-reverse
```

```bash
$ dbflute-util replace-schema
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/wata/dbflute-util
```

## Contribution

1. Fork ([https://github.com/wata/dbflute-util/fork](https://github.com/wata/dbflute-util/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Author

[wata](https://github.com/wata)