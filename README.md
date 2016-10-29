# japawrap - 日本語をいい感じに改行する

## Install

`go get github.com/agatan/japawrap/...`

## Usage

### CLI

標準入力から読むかファイル名を指定します。

```sh
$ echo "今日も元気です" | japawrap
<span class="wordwrap">今日も</span><span class="wordwrap">元気です</span>
```

`-open` と `-close` で囲み方を指定します。

```sh
$ echo "今日も元気です" | japawrap -open "<p>" -close "</p>"
<p>今日も</p><p>元気です</p>
```

### Library

```go
w := japawrap.New(open, close)
s := "今日も元気です"
fmt.Println("%s => %s", s, w.Do(s))
```

## In HTML

HTMLに `japawrap` された文章を配置して、`display: inline-block;` を適応します。
例えば、`<span class="wordwrap"></span>` で囲んだ場合、

```css
.wordwrap {
    display: inline-block;
}
```

というCSSを適用することで、改行がいい感じになります。
