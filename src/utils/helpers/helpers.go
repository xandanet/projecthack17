package helpers

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"reflect"
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

// ConvertStructToMap converts a struct to a map using the structs tags.
// ConvertStructToMap uses tags on struct fields to decide which fields to add to the returned map
func ConvertStructToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ConvertStructToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func GetLocationFromIp(ipAddress string) *geoip2.City {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP("81.2.69.142")
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	return record
}
