package japawrap

import "testing"

func TestWrap(t *testing.T) {
	cases := []struct {
		input  string
		open   string
		close  string
		output string
	}{
		{"今日も元気です", `<span class="wordwrap">`, `</span>`, `<span class="wordwrap">今日も</span><span class="wordwrap">元気です</span>`},
		{"総称です．その代表格は", `<span class="wordwrap">`, `</span>`, `<span class="wordwrap">総称です．</span><span class="wordwrap">その代表格は</span>`},
	}

	for _, c := range cases {
		w := New(c.open, c.close)
		if w.Do(c.input) != c.output {
			t.Errorf("got %v, expected %v", w.Do(c.input), c.output)
		}
	}
}
