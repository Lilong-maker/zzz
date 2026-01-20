package internal

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/Mikaelemmmm/sql2pb/core"
)

func Sql2DB(dbType,
	host,
	user,
	password,
	schema,
	table,
	serviceName,
	packageName,
	goPackageName,
	ignoreTableStr,
	ignoreColumnStr,
	fieldStyle string,
	port int) {
	fmt.Println("接收入参", dbType, host, user, password, schema)
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, schema)
	db, err := sql.Open(dbType, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ignoreTables := strings.Split(ignoreTableStr, ",")
	ignoreColumns := strings.Split(ignoreColumnStr, ",")

	s, err := core.GenerateSchema(db, table, ignoreTables, ignoreColumns, serviceName, goPackageName, packageName, fieldStyle)

	if nil != err {
		log.Fatal(err)
	}

	if nil != s {
		fmt.Println(s)
	}
}
