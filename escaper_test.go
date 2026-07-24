package escpaper2_test

import (
	"fmt"
	"github.com/message-256/escpaper22"
	"testing"
)

func TestEscaper(t *testing.T) {
	fmt.Println("escaper2")

	var outputs = map[string]string{

		`test\\\"text`: "test\x1b\\\x1b\"text",
		`test\\`:       "test\x1b\\",
		`\\test`:       "\x1b\\test",
		"test":         "test",
		"\x1babc":      "\x1babc",
		"\n":           "",
	}
	for input := range outputs {
		ret := escpaper2.Escape(input)
		if ret != outputs[input] {
			fmt.Printf("with input = %s,got = %q, want = %q\n", input, ret, outputs[input])
			t.Errorf("ret not right value")
		}
	}

}
func TestDelim(t *testing.T) {
	fmt.Println("delim")
	var outputs = map[string]string{

		`test\\\"text"`:             "test\x1b\\\x1b\"text\"",
		`test\\"`:                   "test\x1b\\\"",
		`\\test"`:                   "\x1b\\test\"",
		"test\"":                    "test\"",
		`test\\\"text",other stuff`: "test\x1b\\\x1b\"text\"",
		`test\\",other stuff`:       "test\x1b\\\"",
		`\\test",other stuff`:       "\x1b\\test\"",
		"test\",others stuff":       "test\"",
		`test`:                      "test",
		"":                          "",
		"\"":                        "",
	}
	for input := range outputs {
		ret := escpaper2.SubString(input, '"')
		if ret != outputs[input] {
			fmt.Printf("with input = %s,got = %q, want = %q\n", input, ret, outputs[input])
			t.Errorf("ret not right value")
		}
	}

}
