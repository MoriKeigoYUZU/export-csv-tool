# CSV Base64 デコーダー

このツールは、JSON 内の Base64 エンコードされた CSV データをデコードし、CSV ファイルとして保存する Go プログラムです。

## 主な機能
- 標準入力から JSON データを読み取ります。
- JSON の `csv` フィールドから Base64 エンコードされたデータを抽出します。
- データをデコードし、CSV ファイルとして保存します。

## 使い方
1. プログラムをビルドします。
   ```bash
   go build -o csv-decoder main.go
   ```
2. 標準入力で JSON を渡して実行します。
   ```bash
   echo '{"csv": "R29vZHMsU2VydmljZSxQcmljZQpDYXIsUmVudGFscywxMjAwMC41MApQaG9uZSxTdWJzY3JpcHRpb25zLDU5OS45MApMYXB0b3AsUmVwYWlycyw5OTkuOTk="}' | ./csv-decoder
   ```
3. デコードされた CSV が `output.csv` という名前で保存されます。

## サンプル JSON 入力
```json
{
    "csv": "R29vZHMsU2VydmljZSxQcmljZQpDYXIsUmVudGFscywxMjAwMC41MApQaG9uZSxTdWJzY3JpcHRpb25zLDU5OS45MApMYXB0b3AsUmVwYWlycyw5OTkuOTk="
}
```

上記の JSON は以下の CSV データをエンコードしたものです：
```
Goods,Service,Price
Car,Rentals,12000.50
Phone,Subscriptions,599.90
Laptop,Repairs,999.99
```

## 出力例
以下の内容が `output.csv` に保存されます。
```
Goods,Service,Price
Car,Rentals,12000.50
Phone,Subscriptions,599.90
Laptop,Repairs,999.99
```

## 注意点
- JSON の `csv` フィールドに有効な Base64 エンコード文字列が必要です。
- 出力ファイルはカレントディレクトリに `output.csv` として保存されます。

## ライセンス
本プロジェクトは MIT ライセンスの下で提供されています。

