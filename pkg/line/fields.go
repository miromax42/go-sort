package line

type field int

const (
	groupType field = iota
	dateType
	lineType
)

func (f field) String() string {
	return []string{"group", "date", "line"}[f]
}

func getType(str string) field {
	sb := substr(str, 0, 2)

	switch sb {
	case "--":
		return groupType
	case " -":
		return dateType
	default:
		return lineType
	}
}
