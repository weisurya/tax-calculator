package tax

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	common "tax-calculator/common"

	_ "github.com/lib/pq"
)

func insertNewItem(database *sql.DB, newItem item) (err error) {
	if err = database.Ping(); err != nil {
		log.Println(err)
		return err
	}

	table := tableItem

	var bufferField, bufferValue bytes.Buffer

	reflection := reflect.ValueOf(newItem)
	for i := 0; i < reflection.Type().NumField(); i++ {
		columnName := reflection.Type().Field(i).Tag.Get("json")
		bufferField.WriteString(columnName + ",")

		columnValue := fmt.Sprintf("'%v'", reflection.Field(i).Interface())
		bufferValue.WriteString(columnValue + ",")
	}

	field := strings.TrimSuffix(bufferField.String(), ",")
	value := strings.TrimSuffix(bufferValue.String(), ",")

	sqlStatement := common.CreateInsertStatement(table, field, value)
	log.Println(sqlStatement)
	if err = common.ExecuteSQL(database, sqlStatement); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func retrieveAllItems(database *sql.DB) (items []item, err error) {
	if err = database.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}

	items, err = getItems(database)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return items, nil
}

func getItems(database *sql.DB) (items []item, err error) {
	sqlStatement := "select name, taxcode, amount, type, taxamount, totalamount from " + tableItem
	rows, err := database.Query(sqlStatement)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := item{}
		if err = rows.Scan(&item.Name, &item.TaxCode, &item.Amount, &item.Type, &item.TaxAmount, &item.TotalAmount); err != nil {
			fmt.Println(err)
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
