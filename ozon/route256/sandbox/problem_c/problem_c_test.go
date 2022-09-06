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

func Test03(t *testing.T) {
	err := utils.RunTest(main, "tests\\03", "tests\\03.a")
	if err != nil {
		t.Fatal(err)
	}
}

func Test04(t *testing.T) {
	err := utils.RunTest(main, "tests\\04", "tests\\04.a")
	if err != nil {
		t.Fatal(err)
	}
}

func Test05(t *testing.T) {
	err := utils.RunTest(main, "tests\\05", "tests\\05.a")
	if err != nil {
		t.Fatal(err)
	}
}

func Test06(t *testing.T) {
	err := utils.RunTest(main, "tests\\06", "tests\\06.a")
	if err != nil {
		t.Fatal(err)
	}
}

func Test07(t *testing.T) {
	err := utils.RunTest(main, "tests\\07", "tests\\07.a")
	if err != nil {
		t.Fatal(err)
	}
}

func Test08(t *testing.T) {
	err := utils.RunTest(main, "tests\\08", "tests\\08.a")
	if err != nil {
		t.Fatal(err)
	}
}

func Test09(t *testing.T) {
	err := utils.RunTest(main, "tests\\09", "tests\\09.a")
	if err != nil {
		t.Fatal(err)
	}
}

func Test10(t *testing.T) {
	err := utils.RunTest(main, "tests\\10", "tests\\10.a")
	if err != nil {
		t.Fatal(err)
	}
}
