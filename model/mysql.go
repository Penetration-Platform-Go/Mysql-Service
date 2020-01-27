package model

import "fmt"

// Query Mysql Model
func Query(column string, table string, condition string) *Result {
	sql := "SELECT " + column + " FROM " + table + " WHERE " + condition
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	return &Result{
		IsValid: true,
		Value:   rows,
	}
}

// Insert Mysql Model
func Insert(table string, column []string, value []string) *Result {
	return &Result{}
}

// Delete Mysql Model
func Delete(table string, condition string) *Result {
	return &Result{}
}

// Update Mysql Model
func Update(table string, condition string, column, value []string) *Result {
	return &Result{}
}
