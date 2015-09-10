package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

// main ...
func main() {
	dat, err := ioutil.ReadFile("./emoji.dat")
	if err != nil {
		panic(err)
	}
	r16 := []string{}
	r32 := []string{}
	offset := 0
	for _, r := range string(dat) {
		if 0xfff0000&int(r) == 0 {
			r16 = append(r16, fmt.Sprintf("0x%04x", r))
		} else {
			r32 = append(r32, fmt.Sprintf("0x%04x", r))
		}
		if int(r) <= 0xff {
			offset++
		}
	}
	t := template.Must(template.New("").Parse(text))
	data := map[string]interface{}{
		"R16":         r16,
		"R32":         r32,
		"LatinOffset": offset,
	}
	var writer bytes.Buffer
	if err := t.Execute(&writer, data); err != nil {
		panic(err)
	}
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f, err := os.Create(path.Join(pwd, "emoji.go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write(writer.Bytes())
	if err != nil {
		panic(err)
	}
	f.Sync()
}

const text = `package unicode

import (
	"unicode"
)

var e = &unicode.RangeTable{
	R16: []unicode.Range16{{printf "{"}}{{range .R16}}{{printf "\n\t\t{%s, %s, 1}," . .}}{{end}}
	},
	R32: []unicode.Range32{{printf "{"}}{{range .R32}}{{printf "\n\t\t{%s, %s, 1}," . .}}{{end}}
	},
	LatinOffset: {{.LatinOffset}},
}

// IsEmoji reports whether the rune is a emoji.
func IsEmoji(r rune) bool {
	return unicode.In(r, e)
}`
