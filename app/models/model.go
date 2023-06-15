package models

import (
	"strings"
)

func BuildSelectQueryWithAggregate(fields []string, model string, model_aggregate string) string {
	var fieldListString = ""

	for _, v := range fields {
		fieldListString = fieldListString + v + ","
	}

	fieldListString = strings.TrimRight(fieldListString,",")

	return "SELECT "  + fieldListString + " FROM " + model + " " + model + " LEFT JOIN " + model_aggregate + " " + model_aggregate + " ON " + model_aggregate + "." + model + "_id = " + model + ".id"
} 


func BuildSelectQuery(fields []string, model string) string {
	var fieldListString = ""

	for _, v := range fields {
		fieldListString = fieldListString + v + ","
	}

	fieldListString = strings.TrimRight(fieldListString,",")

	return "SELECT "  + fieldListString + " FROM " + model
}

func BuildInsertQuery(fields []string, model string) string {

	var fieldListString = ""
	var fieldInsertParams = ""

	for _, v := range fields {
		fieldListString = fieldListString + v + ","
		fieldInsertParams = fieldInsertParams + "?,"
	}

	fieldListString = strings.TrimRight(fieldListString,",")
	fieldInsertParams = strings.TrimRight(fieldInsertParams, ",")

	return "INSERT INTO " + model + " (" + fieldListString + ") VALUES (" + fieldInsertParams + ")"
}
