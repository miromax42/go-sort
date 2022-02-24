package line

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
)

const (
	defaultString = "NO"
)

var errNoNumProvided = errors.New("no number provided in line")

type Line struct {
	Prefix  string
	num     int
	postfix string
	group   string
	date    string
}

func newLine(str, group, date string) (line Line, err error) {
	position := getPosition(str)
	if position[0] == 0 && position[1] == 0 {
		return line, errNoNumProvided
	}

	line.Prefix = substr(str, 0, position[0])
	line.num, _ = strconv.Atoi(substr(str, position[0], position[1]))
	line.postfix = substr(str, position[1], len([]rune(str)))
	line.group = group
	line.date = date

	if line.num == 0 {
		return line, errNoNumProvided
	}

	return line, nil
}

func getPosition(str string) []int {
	re := regexp.MustCompile(`(?m)\D0[0-9]+`)

	matches := re.FindAllString(str, -1)
	if len(matches) == 0 {
		return []int{0, 0}
	}

	match := []rune(matches[len(matches)-1])
	runeStr := []rune(str)

	start := 0
	end := 0

loop:
	for i := 0; i <= len(runeStr)-len(match); i++ {
		for j := 0; j < len(match); j++ {
			if runeStr[i+j] != match[j] {
				continue loop
			}
		}
		start = i + 1
		end = i + len(match)
	}

	return []int{start, end}
}

func NewLines(sc *bufio.Scanner) (lines []Line, total int, err error) {
	group := defaultString
	date := defaultString

	for sc.Scan() {
		line := sc.Text()
		total++

		switch getType(line) {
		case groupType:
			group = substr(line, 2, len(line))
			date = defaultString

		case dateType:
			date = substr(line, 2, len(line))

		case lineType:
			var pLine Line

			pLine, err = newLine(line, group, date)
			if err != nil {
				err = fmt.Errorf("broken line[%v] with [%v]", total, err)

				return
			}

			lines = append(lines, pLine)

		default:
			err = fmt.Errorf("broken line[%v] with [%v]", total, err)

			return
		}
	}

	if err = sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)

		return
	}

	// line sort
	sort.Slice(lines, func(a, b int) bool {
		if lines[a].Prefix == lines[b].Prefix {
			return lines[a].num < lines[b].num
		}

		return lines[a].Prefix < lines[b].Prefix
	})

	return lines, total, nil
}

func (l *Line) String() string {
	return fmt.Sprintf("prefix: %v, num: %v, postfix: %v, group:%v, date: %v", l.Prefix, l.num, l.postfix, l.group, l.date)
}

func (l *Line) StringArr() []string {
	return []string{l.Prefix, fmt.Sprintf("%d", l.num), l.postfix, l.group, l.date}
}

func (l *Line) Humanize() string {
	return fmt.Sprintf("%10s:%05d - %s, g:%s, date: %s", l.Prefix, l.num, l.postfix, l.group, l.date)
}

func substr(input string, start, end int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) || end > len(asRunes) || start >= end {
		return ""
	}

	return string(asRunes[start:end])
}

func ConvertStringArray(lines []Line) [][]string {
	var strs [][]string

	for _, line := range lines {
		strs = append(strs, line.StringArr())
	}

	return strs
}
