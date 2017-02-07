package dsl

import (
	"fmt"
	"testing"
)

func TestSplitString(t *testing.T) {
	s1 := `Get   http://baidu.com`
	r := SplitString(s1)
	fmt.Println(s1, r)
}

func TestSplitString2(t *testing.T) {
	s1 := `match $data   '{h:"sdf http://baidu.com"} `
	r := SplitString(s1)
	fmt.Println(s1, r)
}
