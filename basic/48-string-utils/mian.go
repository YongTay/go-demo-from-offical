package main

import (
	"fmt"
	s "strings"
)

func main() {
	var p = fmt.Println
	p("Contains(", s.Contains("test", "es"))
	p("Count", s.Count("hello", "l"))
	p("HasPrefix", s.HasPrefix("test", "te"))
	p("HasSuffix", s.HasSuffix("test", "ss"))
	p("Index", s.Index("this is a sentence", "is"))
	p("Join", s.Join([]string{"hello", "word"}, "-"))
	p("Repeat", s.Repeat("a", 5))
	p("Replace", s.Replace("foo", "o", "0", -1))
	p("ReplaceAll", s.ReplaceAll("foo", "o", "0"))
	p("Replace", s.Replace("foo", "o", "0", 1))
	p("Split", s.Split("a-b-c-d-e-f", "-"))
	p("ToLower", s.ToLower("TEST"))
	p("ToUpper", s.ToUpper("test"))
}
