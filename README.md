# covid19-nara-sheet2json

Google Spread Sheetから https://stopcovid19.code4nara.org/ で使えるJSONを生成
するためのコマンドです。

## 実行方法

基本的にはGitHub Actions上で実行するので、この手順は生成されるjsonをローカルで
確認する為に使用します。

### 前提

1. Go言語の環境がセットアップされている前提です
2. Spread Sheets APIを実行するため、GCPのService Accountの認証jsonを取得し、  
`tmp` ディレクトリに `credentials.json` という名前で配置しておいてください。

### 1. クローンする


```
git clone https://github.com/code4nara/covid19-nara-sheet2json.git
cd covid19-nara-sheet2json
```


### 2. `.env` ファイルをプロジェクトルートに配置する

```
COVID19_JSON2CSV_OUTPUT_DATA=tmp/data.json
COVID19_JSON2CSV_OUTPUT_SICKBEDS_SUMMARY=tmp/sickbeds_summary.json
COVID19_JSON2CSV_OUTPUT_NEWS=tmp/news.json
COVID19_JSON2CSV_SHEET_ID=1C07ojkwER8BiAjLBxlzJkfvgM5jxUCLrdtI7wtctTIY
COVID19_JSON2CSV_SHEET_RANGE_PATIENTS=01.陽性患者の属性!A3:O130000
COVID19_JSON2CSV_SHEET_RANGE_HOSPITALIZATION=入院患者の状況!A3:I3
COVID19_JSON2CSV_SHEET_RANGE_NEWS=新着情報!A2:C1000
```

### 3. `実行する`


```
go run main.go
```

tmp以下にjsonファイルが生成されるので確認します。



## デプロイ

[リリースページ](https://github.com/code4nara/covid19-nara-sheet2json/releases) でリリースを
作成すると、自動でビルドされ、リリースページにバイナリが配置されるようにGitHub Actions
を設定しています。(amd64/linux用です)

配置後、 [本体](https://github.com/code4nara/covid19) のGitHub Actions定義内のsheet2jsonの取得URLを変更してください。

