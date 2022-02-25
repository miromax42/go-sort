package line

import "regexp"

type field int

const (
	groupType field = iota
	dateType
	lineType
)

func (f field) String() string {
	return []string{"group", "date", "line"}[f]
}

func getType(str string) (f field, offset int) {
	re := regexp.MustCompile(`(?m)^-+`)

	start := re.FindString(str)
	offset = len(start)

	switch offset {
	case 2:
		return groupType, offset
	case 1:
		return dateType, offset
	default:
		return lineType, offset
	}
}
