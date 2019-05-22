package main

import (
	"bufio"
	"math"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var splitIn4K = func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	return len(data), data, nil
}

func main() {
	args := os.Args[1:]

	reportedFile, err := os.Open(args[0])
	check(err)
	defer reportedFile.Close()

	expectedFile, err := os.Open(args[1])
	check(err)
	defer expectedFile.Close()

	reportedScanner := bufio.NewScanner(reportedFile)
	expectedScanner := bufio.NewScanner(expectedFile)

	reportedScanner.Split(splitIn4K)
	expectedScanner.Split(splitIn4K)

	exitCode := 0

	reportedScannerLeft, expectedScannerLeft := true, true

	for reportedScannerLeft, expectedScannerLeft = reportedScanner.Scan(), expectedScanner.Scan(); reportedScannerLeft && expectedScannerLeft; reportedScannerLeft, expectedScannerLeft = reportedScanner.Scan(), expectedScanner.Scan() {
		reportedLine := reportedScanner.Text()
		expectedLine := expectedScanner.Text()

		if exitCode == 0 {
			exitCode = strings.Compare(reportedLine, expectedLine)
		}

	}

	if reportedScannerLeft || expectedScannerLeft {
		exitCode = 1
	}

	os.Exit(int(math.Abs(float64(exitCode))))

}
