package main

const helpText = `
roop-generater [-S 1行目からスキップする行数(デフォルト0)][-n][-b][-s][-h]

【基本的な使い方】
適当なフォルダを作りその中に以下の3つのファイルを配置してください。
roop-generater      実行ファイル(pathを通している場合は不要)
data.csv            ExcelやGoogleスプレットシートから書き出したもの
tmpl.html           元となるHTMLのテンプレート。内容はPHPでもJSPでも良い。

同フォルダ内で ./roop-generater を実行すると dist.html が同フォルダ内に吐き出されます。
pathを通している場合は不要は roop-generater で実行出来ます。
上記のファイルが配置されていれば、その他のファイルが含まれていても構いません。


【オプション引数】
-S int    シート(csv)の頭からint行目までをスキップします。
-n        改行を取り除きます。
-b        改行を<br/>に置換します。
-s        全角スペースを取り除きます。
-h        使い方を表示します。

コマンド例  ./roop-generater -S 1 -b


【tmpl.htmlの記法】
値を埋め込みたい場所に {{ .列の大文字アルファベット }} の形式で記述して下さい。
シート上のB列に対応する値を埋め込みたければ {{ .B }} という書き方になります。
27(AA)列以上は未対応です

＃連番変数 {{ .SN }} を使用すると 01 から 99 までの連番が埋め込まれます。3桁以上は未対応。


【tmpl.htmlの例】
<article>
<div class="event-article">
  <p class="event-time font01">{{ .A }}</p>
  <div class="event-innerblock">
    <figure class="event-img"><img src="./images/img{{ .SN }}.jpg" alt="{{ .B }}"></figure>
    <h1 class="event-title">{{ .B }}</h1>
    <p class="event-subtitle">{{ .C }}</p>
    <div class="event-body">
      <p>{{ .D }}</p>
    </div>
  </div>
</article>
`
