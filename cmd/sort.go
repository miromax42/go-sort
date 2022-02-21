package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const (
	infixDelimiter   = " "
	postfixDelimiter = "-"
)

var wrongLine = errors.New("wrong line format, use: [infix] [num]-[postfix]")

type Line struct {
	infix   []string
	num     int
	postfix string
}

func (l *Line) String() string {
	return fmt.Sprintf("infix: %v, num: %v, postfix: %v", l.infix, l.num, l.postfix)
}

func (l *Line) Join() string {
	return fmt.Sprintf("%v-%v", strings.Join(l.infix, " "), l.num)
}

func (l *Line) Humanize() string {
	return fmt.Sprintf("%10s:%05d - %s", strings.Join(l.infix, " "), l.num, l.postfix)
}

func NewLine(line string) (Line, error) {
	mainSplit := strings.Split(line, postfixDelimiter)
	if len(mainSplit) > 2 {
		return Line{}, fmt.Errorf("too much [%s] in line", postfixDelimiter)
	}

	filterSplit := strings.Fields(mainSplit[0])
	if len(filterSplit) < 2 {
		return Line{}, errors.New("too short line in line")
	}

	num, err := strconv.Atoi(filterSplit[len(filterSplit)-1])
	if err != nil {
		return Line{}, fmt.Errorf("cant convert to Integer number: [%v]", filterSplit[len(filterSplit)-1])
	}

	postfix := ""
	if len(mainSplit) == 2 {
		postfix = mainSplit[1]
	}

	return Line{
		filterSplit[:len(filterSplit)-1],
		num,
		postfix,
	}, nil
}

// sortCmd represents the sort command
var sortCmd = &cobra.Command{ //nolint
	Use:   "sort",
	Short: "sort file",
	Long: `This subcommand sort provided file
example: sort test.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sort called with with args=%v\n", args)

		if len(args) != 1 {
			log.Fatal("wrong format")

			return
		}

		filename := args[0]

		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)

			return
		}
		defer file.Close()

		sc := bufio.NewScanner(file)

		var lines []Line

		total := 0
		for sc.Scan() {
			line := sc.Text()
			total++

			pLine, err := NewLine(line)
			if err != nil {
				fmt.Printf("error[%v] in line[%v]\n", err, total)
			}

			lines = append(lines, pLine)
		}

		if err = sc.Err(); err != nil {
			log.Fatalf("scan file error: %v", err)

			return
		}

		sort.Slice(lines, func(a, b int) bool {
			if strings.Join(lines[a].infix, ".") == strings.Join(lines[b].infix, ".") {
				return lines[a].num < lines[b].num
			}
			return strings.Join(lines[a].infix, ".") < strings.Join(lines[b].infix, ".")
		})

		filenameFields := strings.Split(filename, ".")
		filenameFields[0] += "(sorted)"
		newFilename := strings.Join(filenameFields, ".")

		err = printLines(newFilename, lines)
		if err != nil {
			fmt.Println(err)
		}

		log.Printf("scan file, lines=%v\n", total)
	},
}

func printLines(filePath string, values []Line) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Errorf("cant create file [%v]", filePath)
	}
	defer f.Close()

	prev := strings.Join(values[0].infix, " ")

	for _, value := range values {
		if prev != strings.Join(value.infix, " ") {
			prev = strings.Join(value.infix, " ")

			fmt.Fprintln(f, "---------------------------------------------------------")
		}

		fmt.Fprintln(f, value.Humanize()) // print values to f, one per line
	}

	return nil
}

func init() {
	RootCmd.AddCommand(sortCmd)
}
