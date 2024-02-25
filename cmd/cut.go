package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"strconv"
	"strings"
)

type postition struct {
	start int
	end   int
}

func readInput(reader *bufio.Reader) (string, error) {
	var result strings.Builder
	for {
		i, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		result.WriteByte(i)
	}
	x := result.String()
	return x, nil
}

func processInputString(source, delimiter string, p postition) string {
	result := make([]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		var sb strings.Builder
		line := scanner.Text()
		sline := strings.Split(line, delimiter)
		for i := range sline {
			if i >= p.start {
				sb.WriteString(sline[i])
				sb.WriteString(delimiter)
			}
			if i >= p.end {
				break
			}
		}
		x := sb.String()
		result = append(result, strings.TrimSuffix(x, delimiter))
	}

	return strings.Join(result, "\n")
}

func main() {
	var fieldsFlag string
	var delimiterFlag string
	flag.StringVar(&fieldsFlag, "f", "", "fields or column to process. Usage: -f 1,2 or -f '1 2'")
	flag.StringVar(&delimiterFlag, "d", "\t", "delimiter to use to split lines into columns. Usage: -d \t")
	flag.Parse()
	if len(fieldsFlag) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	x := strings.Split(fieldsFlag, ",")
	if strings.Contains(fieldsFlag, " ") {
		x = strings.Split(fieldsFlag, " ")
	}
	start, err := strconv.Atoi(x[0])
	if err != nil {
		panic(err)
	}
	start--
	end := start
	if len(x) == 2 {
		end, err = strconv.Atoi(x[1])
		if err != nil {
			panic(err)
		}
	}
	p := postition{
		start: start,
		end:   end,
	}

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	args := flag.Args()
	file := ""
	if len(args) > 0 {
		file = args[0]
	}
	if len(file) != 0 {
		f, err := os.OpenFile(file, os.O_RDONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		reader = bufio.NewReader(f)
	}

	contents, err := readInput(reader)
	contents = processInputString(contents, delimiterFlag, p)
	if err != nil {
		panic(err)
	}
	if _, err := writer.WriteString(contents); err != nil {
		panic(err)
	}
}
