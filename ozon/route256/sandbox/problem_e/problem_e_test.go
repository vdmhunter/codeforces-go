package main

import (
	"testing"

	"codeforces_go/utils"
)

func Test01(t *testing.T) {
	err := utils.RunTest(main, "tests\\01", "tests\\01.a")
	if err != nil {
		t.Fatal(err)
	}
}

func Test02(t *testing.T) {
	err := utils.RunTest(main, "tests\\02", "tests\\02.a")
	if err != nil {
		t.Fatal(err)
	}
}
