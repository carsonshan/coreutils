package tac

import (
	"bytes"
	"github.com/guonaihong/coreutils/utils"
	"strings"
	"testing"
)

func TestRegex(t *testing.T) {
}

// tac -s string
func testReadFromTailStdin(src, dst string, sep string, t *testing.T) {
	rs := strings.NewReader(src)
	w := &bytes.Buffer{}

	readFromTailStdin(rs, w, []byte(sep), false)

	if w.String() != dst {
		t.Fatalf("tac -s fail(%s, %d), need(%s)\n", w.String(), w.Len(), dst)
	}
}

// tac -s string
func testSeparator(src, dst string, sep string, bufSize int, t *testing.T) {
	tac := Tac{}
	tac.Separator = utils.String(sep)
	tac.Before = utils.Bool(false)
	tac.BufSize = bufSize
	rs := strings.NewReader(src)
	w := &bytes.Buffer{}

	tac.Tac(rs, w)

	if w.String() != dst {
		t.Fatalf("tac -s fail(%s), need(%s)\n", w.String(), dst)
	}
}

func TestSeparator(t *testing.T) {
	src := `1
2
3
4
5
6
7
8
9
10
`
	dst := `10
9
8
7
6
5
4
3
2
1
`
	testReadFromTailStdin(src, dst, "\n", t)
	testSeparator(src, dst, "\n", 0, t)
	//testSeparator(src, dst, "\n", 3, t)
	//=============================

	src = "123aaa456aaa789aaa\n"
	dst = `
789aaa456aaa123aaa`

	testReadFromTailStdin(src, dst, "aaa", t)
	testSeparator(src, dst, "aaa", 0, t)
	//=============================

	src = "wwwwwwwwwwww"
	testReadFromTailStdin(src, src, "aaa", t)
	testSeparator(src, src, "aaa", 0, t)

	src = "1,2\n"
	testReadFromTailStdin(src, src, "\n", t)
	testSeparator(src, src, "\n", 0, t)

	dst = `2
1,`
	testReadFromTailStdin(src, dst, ",", t)
	testSeparator(src, dst, ",", 0, t)

	src = `wwwwww`
	testReadFromTailStdin(src, src, "www", t)
	testSeparator(src, src, "www", 0, t)
}

//tac -b
func testBeforeReadFromTailStdin(src, dst, sep string, t *testing.T) {
	rs := strings.NewReader(src)
	w := &bytes.Buffer{}

	readFromTailStdin(rs, w, []byte(sep), true)

	if w.String() != dst {
		t.Fatalf("tac -s fail(%s, l:%d), need(%s)\n", w.String(), w.Len(), dst)
	}
}

//tac -b
func testSeparatorBefore(src, dst, sep string, t *testing.T) {
	tac := Tac{}
	tac.Separator = utils.String(sep)
	tac.Before = utils.Bool(true)

	rs := strings.NewReader(src)
	w := &bytes.Buffer{}

	tac.Tac(rs, w)

	if w.String() != dst {
		t.Fatalf("tac -s fail(%s, l:%d), need(%s)\n", w.String(), w.Len(), dst)
	}
}

func TestBefore(t *testing.T) {
	src := `1,2
`
	dst := `,2
1`

	testBeforeReadFromTailStdin(src, dst, ",", t)

	src = `1
2
3
4
5
6
7
8
9
10
`
	dst = `

10
9
8
7
6
5
4
3
21`
	testBeforeReadFromTailStdin(src, dst, "\n", t)

	src = `wwwwwwwwwwwwwww`
	testBeforeReadFromTailStdin(src, src, "www", t)
}