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
	defaultString = "<no>"
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
	re := regexp.MustCompile(`0\d+`)
	matchIndexes := re.FindAllStringIndex(str, -1)

	if len(matchIndexes) == 0 {
		return line, errNoNumProvided
	}

	position := matchIndexes[len(matchIndexes)-1]

	line.Prefix = substr(str, 0, position[0])
	line.num, _ = strconv.Atoi(substr(str, position[0], position[1]))
	line.postfix = substr(str, position[1], len(str))
	line.group = group
	line.date = date

	if line.num == 0 {
		return line, errNoNumProvided
	}

	return line, nil
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
