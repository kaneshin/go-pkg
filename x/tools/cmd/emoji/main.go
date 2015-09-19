package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"text/template"
)

const (
	in    = "emoji.dat"
	out   = "emoji.go"
	ascii = 0x7f
)

// main ...
func main() {
	dat, err := ioutil.ReadFile(in)
	if err != nil {
		panic(err)
	}
	m16 := map[int]int{}
	m32 := map[int]int{}
	for _, r := range string(dat) {
		n := int(r)
		if n <= ascii {
			continue
		}
		if 0xffff0000&n == 0 {
			m16[n] = n
		} else {
			m32[n] = n
		}
	}
	m2s := func(m map[int]int, u ...int) (s []string, o int) {
		i := []int{}
		for _, v := range m {
			i = append(i, v)
		}
		sort.Sort(sort.IntSlice(i))
		t := func() int {
			if len(u) > 0 {
				return u[0]
			}
			return 0x0
		}()
		for _, v := range i {
			s = append(s, fmt.Sprintf("0x%04x", v))
			if v <= t {
				o++
			}
		}
		return
	}
	r16, offset := m2s(m16, 0xff)
	r32, _ := m2s(m32)
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
	f, err := os.Create(path.Join(pwd, out))
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
}
`
