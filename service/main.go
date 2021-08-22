package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type TestingStruct struct {
	ID      int64  `db:"id_tag" db_select:"true" db_select_value:"true"`
	Name    string `db:"name" db_select:"true"`
	Address string `db:"address" db_select:"true"`
	Phone   string `db:"phone"`
}

func main() {
	fmt.Println("testing run main")
	ts := TestingStruct{
		Name:    "testing name",
		Address: "jl nih",
		Phone:   "08979787",
	}

	// idx, param, query := BuildUpdateParamQuery(ts)
	// fmt.Println("last index - ", idx)
	// fmt.Println("param - ", param)
	// fmt.Println("query - ", query)

	idx, param, query, err := BuildSelectParamQuery(ts)
	if err != nil {
		fmt.Errorf("errr - %v", err)
	}
	fmt.Println("last index - ", idx)
	fmt.Println("param - ", param)
	fmt.Println("query - ", query)
}

func BuildUpdateParamQuery(t interface{}) (idx int, valuesParam []interface{}, valuesQuery string) {
	values := reflect.ValueOf(t)
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {
		val := values.Field(i).Interface()
		dbVal := types.Field(i).Tag.Get("db")

		switch typ := val.(type) {
		case string:
			if len(strings.TrimSpace(typ)) > 0 {
				valuesQuery += fmt.Sprintf(` "%s" = $%d,`, dbVal, idx+1)
				valuesParam = append(valuesParam, typ)
				idx++
			}
		case int64:
			if typ > 0 {
				valuesQuery += fmt.Sprintf(` "%s" = $%d,`, dbVal, idx+1)
				valuesParam = append(valuesParam, typ)
				idx++
			}
		}
	}

	idx++
	valuesQuery = valuesQuery[0 : len(valuesQuery)-1]
	return
}

func BuildSelectParamQuery(t interface{}) (int, []interface{}, string, error) {
	var (
		idx         int
		valuesQuery string
		valuesParam []interface{}
	)
	values := reflect.ValueOf(t)
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {
		val := values.Field(i).Interface()
		dbVal := types.Field(i).Tag.Get("db")
		dbSelectVal := types.Field(i).Tag.Get("db_select")

		if len(dbSelectVal) > 0 {
			selectVal, err := strconv.ParseBool(dbSelectVal)
			if err != nil {
				return idx, valuesParam, valuesQuery, err
			}

			if selectVal {
				switch typ := val.(type) {
				case string:
					if len(strings.TrimSpace(typ)) > 0 {
						valuesQuery += fmt.Sprintf(` AND "%s" = $%d `, dbVal, idx+1)
						valuesParam = append(valuesParam, typ)
						idx++
					}
				case int64:
					if typ > 0 {
						valuesQuery += fmt.Sprintf(` AND "%s"=%d `, dbVal, idx+1)
						valuesParam = append(valuesParam, typ)
						idx++
					}
				}
			}
		}
	}

	idx++
	return idx, valuesParam, valuesQuery, nil
}
