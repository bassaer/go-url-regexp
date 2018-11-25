package main

import (
	"fmt"
	"regexp"
)

type Matcher struct {
	re *regexp.Regexp
}

func NewMatcher(p string) *Matcher {
	r := regexp.MustCompile(p)
	return &Matcher{r}
}

func (m *Matcher) match(src string) bool {
	return m.re.MatchString(src)
}

func main() {
	m := NewMatcher("^(http|https)://([\\w-]+\\.)+[\\w-]+(/[\\w-./?%&=]*)?$")
	fmt.Println(m.match("http://www.yahoo.co.jp/?a=hoge"))
	fmt.Println(m.match("http://www.yahoo.co.jp/?a=hoge&b=http://wwww.yahoo.co.jp"))
}
