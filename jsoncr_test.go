package jsoncr

import (
	"strings"
	"testing"
)

func TestRemove(t *testing.T) {
	tests := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "multi inline",
			input: `{"test": 2} // comment // more comment`,
			want:  `{"test":2}`,
		},
		{
			desc: "multi list",
			input: `{
"test": ["test1", // "test2"]
"test3"]}
`,
			want: `{"test":["test1","test3"]}`,
		},
		{
			desc: "mixed",
			input: `// comments
/*
{
	"test": ["test1", "test2"] // comments
}
*/
{
	"test": ["test1"/*, more comments */, "test2", "//", "/* */", "*/", "\""] // comments
}
`,
			want: `{"test":["test1","test2","//","/* */","*/","\""]}`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			got, err := Remove(strings.NewReader(tt.input))
			if err != nil {

			}
			if string(got) != tt.want {
				t.Errorf("\n<got>\n%s\n<want>\n%s\n", got, tt.want)
			}
		})
	}
}
