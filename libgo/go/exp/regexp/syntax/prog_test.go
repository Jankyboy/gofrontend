package syntax

import (
	"testing"
)

var compileTests = []struct {
	Regexp string
	Prog   string
}{
	{"a", `  0	fail
  1*	rune "a" -> 2
  2	match
`},
	{"[A-M][n-z]", `  0	fail
  1*	rune "AM" -> 2
  2	rune "nz" -> 3
  3	match
`},
	{"", `  0	fail
  1*	nop -> 2
  2	match
`},
	{"a?", `  0	fail
  1	rune "a" -> 3
  2*	alt -> 1, 3
  3	match
`},
	{"a??", `  0	fail
  1	rune "a" -> 3
  2*	alt -> 3, 1
  3	match
`},
	{"a+", `  0	fail
  1*	rune "a" -> 2
  2	alt -> 1, 3
  3	match
`},
	{"a+?", `  0	fail
  1*	rune "a" -> 2
  2	alt -> 3, 1
  3	match
`},
	{"a*", `  0	fail
  1	rune "a" -> 2
  2*	alt -> 1, 3
  3	match
`},
	{"a*?", `  0	fail
  1	rune "a" -> 2
  2*	alt -> 3, 1
  3	match
`},
	{"a+b+", `  0	fail
  1*	rune "a" -> 2
  2	alt -> 1, 3
  3	rune "b" -> 4
  4	alt -> 3, 5
  5	match
`},
	{"(a+)(b+)", `  0	fail
  1*	cap 2 -> 2
  2	rune "a" -> 3
  3	alt -> 2, 4
  4	cap 3 -> 5
  5	cap 4 -> 6
  6	rune "b" -> 7
  7	alt -> 6, 8
  8	cap 5 -> 9
  9	match
`},
	{"a+|b+", `  0	fail
  1	rune "a" -> 2
  2	alt -> 1, 6
  3	rune "b" -> 4
  4	alt -> 3, 6
  5*	alt -> 1, 3
  6	match
`},
}

func TestCompile(t *testing.T) {
	for _, tt := range compileTests {
		re, _ := Parse(tt.Regexp, Perl)
		p, _ := Compile(re)
		s := p.String()
		if s != tt.Prog {
			t.Errorf("compiled %#q:\n--- have\n%s---\n--- want\n%s---", tt.Regexp, s, tt.Prog)
		}
	}
}