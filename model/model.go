package model
import (
	"fmt"
	"database/sql"

	//Driver do Mysql
	_"github.com/go-sql-driver/mysql"
	
)

//DB ...
var DB, err = sql.Open("mysql", "root:@/db_todo")

//TryConn ...
func TryConn() {
	if err != nil{
		panic(err.Error())
	}

	if DB.Ping() != nil{
		panic(err.Error())
	}

	fmt.Println("Conexao com mysql feita com sucesso")
}