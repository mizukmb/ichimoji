# ichimoji

アルファベット一文字エイリアスの登録状況を確認するためのCLIツール。

ログインユーザーのrcファイルを読み込んでaliasコマンドでa-zまでのエイリアスの内容を確認できます。

現在サポートしているログインシェルはbashとzshの2種類になります。

# Install

`go install` を利用してください。

# Usage

```
$ ichimoji
ichimoji alias list.
✅ a='git add'
🈳 b
🈳 c
🈳 d
🈳 e
🈳 f
✅ g='git'
🈳 h
🈳 i
🈳 j
🈳 k
🈳 l
🈳 m
🈳 n
🈳 o
🈳 p
🈳 q
🈳 r
🈳 s
🈳 t
🈳 u
🈳 v
🈳 w
🈳 x
🈳 y
🈳 z
```