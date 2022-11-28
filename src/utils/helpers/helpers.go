package helpers

import (
	"strconv"
	"strings"
)

type newWhereT struct {
	fields             []whereField
	customFields       []whereField
	wherePrefix        string
	appendWhereAtStart bool
}

type whereField struct {
	query string
	args  []any
}

func NewWhere() *newWhereT {
	return &newWhereT{
		wherePrefix:        "WHERE ",
		appendWhereAtStart: false,
	}
}

func (i *newWhereT) AppendWhereAtStart() *newWhereT {
	i.appendWhereAtStart = true
	return i
}

func (i *newWhereT) Where(query string, args ...any) *newWhereT {
	i.fields = append(i.fields, whereField{
		query: query,
		args:  args,
	})
	return i
}

func (i *newWhereT) CustomWhere(query string, args ...any) *newWhereT {
	i.customFields = append(i.customFields, whereField{
		query: query,
		args:  args,
	})
	return i
}

func (i *newWhereT) String() (string, []any) {
	where := []string{}
	args := []any{}

	for index := range i.fields {
		el := i.fields[index]

		where = append(where, el.query)
		args = append(args, el.args...)
	}

	str := strings.Join(where, " AND ")
	if len(i.customFields) > 0 {
		str += " "
	}

	// Add custom where
	for index := range i.customFields {
		el := i.customFields[index]

		str += el.query
		args = append(args, el.args...)
	}

	if len(where) > 0 && i.appendWhereAtStart {
		return i.wherePrefix + str, args
	}
	return str, args
}

// ConvertStringToInt64 will convert string s to int64
// It will return 0 if conversion is not possible
func ConvertStringToInt64(s string) int64 {
	iValue, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return iValue
}

func StringInMap(search string, haystack []string) bool {
	for _, t := range haystack {
		if t == search {
			return true
		}
	}

	return false
}