package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func RunTest(f func(), inputFilePath string, resultFilePath string) (err error) {
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(inputFile)

	resultFile, err := os.Open(resultFilePath)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(resultFile)

	rInput, wInput, err := os.Pipe()
	if err != nil {
		return err
	}

	rOutput, wOutput, err := os.Pipe()
	if err != nil {
		return err
	}

	go func() {
		inputFileScanner := bufio.NewScanner(inputFile)
		buf := make([]byte, 2048*1024)
		inputFileScanner.Buffer(buf, 2048*1024)

		for inputFileScanner.Scan() {
			_, _ = wInput.WriteString(inputFileScanner.Text() + "\n")
		}
		_ = wInput.Close()
	}()

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	os.Stdin = rInput
	os.Stdout = wOutput

	go func() {
		f()
		_ = wOutput.Close()
	}()

	check := func(b1 bool, b2 bool) {
		if (!b1 && b2) || (b1 && !b2) {
			err = errors.New("test failed. number of lines in problem response and test file does not match")
		}
	}

	rOutputScanner := bufio.NewScanner(rOutput)
	resultFileScanner := bufio.NewScanner(resultFile)

	b1 := rOutputScanner.Scan()
	b2 := resultFileScanner.Scan()

	check(b1, b2)

	i := 1

	for b1 && b2 {
		outputLine := rOutputScanner.Text()
		resultFileLine := resultFileScanner.Text()

		if outputLine != resultFileLine {
			err = fmt.Errorf("test failed on line %d! output value: %s. expected value: %s",
				i, outputLine, resultFileLine)
		}

		b1 = rOutputScanner.Scan()
		b2 = resultFileScanner.Scan()

		check(b1, b2)

		i++
	}

	return err
}
