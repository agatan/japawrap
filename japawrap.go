package japawrap

import (
	"fmt"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

type Wrapper struct {
	tok   tokenizer.Tokenizer
	open  string
	close string
}

func New(open string, close string) *Wrapper {
	return &Wrapper{
		tok:   tokenizer.New(),
		open:  open,
		close: close,
	}
}

func isMainWord(tok *tokenizer.Token) bool {
	if tok.Pos() == "名詞" {
		return true
	}
	if tok.Pos() == "動詞" {
		for _, f := range tok.Features() {
			if strings.Contains(f, "自立") {
				return true
			}
		}
	}
	return false
}

func isSeparateWord(tok *tokenizer.Token) bool {
	s := strings.TrimSpace(tok.Surface)
	return s == "." || s == "," || s == "。" || s == "、" || s == "．" || s == "，"
}

func concatTokens(toks []tokenizer.Token) string {
	s := []string{}
	for _, tok := range toks {
		s = append(s, tok.Surface)
	}
	return strings.Join(s, "")
}

func (w *Wrapper) wrap(s string) string {
	return fmt.Sprintf("%s%s%s", w.open, s, w.close)
}

func (w *Wrapper) Do(s string) string {
	rawts := w.tok.Tokenize(s)
	rs := make([]string, 0)
	i := 0

	ts := make([]tokenizer.Token, 0)
	for _, t := range rawts {
		if t.Class == tokenizer.DUMMY {
			continue
		}
		ts = append(ts, t)
	}

	for i < len(ts) {
		start := i
		for i < len(ts) && ts[i].Pos() == "連体詞" {
			i += 1
		}
		for i < len(ts) && ts[i].Pos() == "接頭詞" {
			i += 1
		}
		if ts[i].Pos() == "名詞" {
			for i < len(ts) && ts[i].Pos() == "名詞" {
				i += 1
			}
		}
		i += 1
		for i < len(ts) && !isMainWord(&ts[i]) {
			if isSeparateWord(&ts[i]) {
				i += 1
				break
			}
			i += 1
		}
		rs = append(rs, w.wrap(concatTokens(ts[start:i])))
	}
	return strings.Join(rs, "")
}
