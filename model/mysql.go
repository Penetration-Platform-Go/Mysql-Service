package model

import (
	"fmt"
	"github.com/Penetration-Platform-Go/Mysql-Service/lib"
	"strings"
)

// Query Mysql Model
func Query(column []string, table, condition string) *Result {
	sql := "SELECT " + strings.Join(column, ",") + " FROM " + table + " WHERE " + condition
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return &Result{
			IsValid: false,
			Value:   nil,
		}
	}
	return &Result{
		IsValid: true,
		Value:   rows,
	}
}

// Insert Mysql Model
func Insert(table string, column, value []string) bool {
	for index, str := range value {
		value[index] = lib.TransferString(str)
	}
	sql := "INSERT INTO " + table + "(" + strings.Join(column, ",") + ") VALUE('" + strings.Join(value, "','") + "')"
	fmt.Println(sql)
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return false

	}
	return true
}

// Delete Mysql Model
func Delete(table, condition string) bool {
	sql := "DELETE FROM " + table + " WHERE " + condition
	fmt.Println(sql)
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return false

	}
	return true
}

// Update Mysql Model
func Update(table, condition string, column, value []string) bool {
	sql := "UPDATE " + table + " SET " + strings.Join(column, " = ?, ") + " = ? WHERE " + condition
	for _, val := range value {
		sql = strings.Replace(sql, "?", "'"+lib.TransferString(val)+"'", 1)
	}
	fmt.Println(sql)
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return false

	}
	return true
}
