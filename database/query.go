package database

import (
	"database/sql"
	"fmt"
	"log"
	"reperio-backend-assessment/common"
	"reperio-backend-assessment/types"
)

// Insert is a function for inserting into the database
func Insert(table string, model interface{}) (success bool, err error) {
	query, values, err := buildQuery(table, model)

	if err != nil {
		return
	}

	trx, stmt, err := startTrx(query)

	if err != nil {
		return
	}

	_, err = stmt.Exec(values...)

	if err != nil {
		return
	}

	defer wrapCommit(trx)

	success = true
	return
}

// Select is a function for selecting a single row in the database
func Select(table string, selection string, id string, model interface{}) (result *sql.Rows, success bool, err error) {
	query, err := buildSelectQuery(table, selection, id)

	if err != nil {
		return
	}

	trx, stmt, err := startTrx(query)

	if err != nil {
		return
	}

	result, err = stmt.Query()
	result.Scan(&model)

	if err != nil {
		return
	}

	defer wrapCommit(trx)

	success = true
	return
}

func buildSelectQuery(table string, selection string, id string) (query string, err error) {
	if err != nil {
		return
	}
	query = fmt.Sprintf("SELECT %s FROM %s WHERE id = %s", selection, table, id)
	return
}

func wrapCommit(trx types.TransactionDriver) {
	func() {
		err := trx.Commit()
		if err != nil {
			log.Fatalln(err)
		}
	}()
}

func startTrx(query string) (trx types.TransactionDriver, stmt types.StatementDriver, err error) {
	trx, err = DB.Begin()

	if err != nil {
		return
	}

	stmt, err = trx.Prepare(query)

	if err != nil {
		return
	}

	return
}

func buildQuery(table string, model interface{}) (query string, values []interface{}, err error) {
	m, err := common.ConvertStructToMap(model)
	if err != nil {
		return
	}
	query = fmt.Sprintf("INSERT INTO %s", table)
	fieldQuery, valueQuery, values := buildQueryAndValues(m)
	query = fmt.Sprintf("%s(%s) VALUES(%s)", query, fieldQuery, valueQuery)
	return
}

func buildQueryAndValues(m map[string]interface{}) (fieldQuery string, valueQuery string, values []interface{}) {
	i := 0
	for field, value := range m {
		if i > 0 && i < len(m) {
			fieldQuery = fmt.Sprintf("%s,", fieldQuery)
			valueQuery = fmt.Sprintf("%v,", valueQuery)
		}
		fieldQuery = fmt.Sprintf("%s%s", fieldQuery, field)
		valueQuery = fmt.Sprintf("%v%v", valueQuery, "?")
		values = append(values, value)
		i++
	}
	return
}
