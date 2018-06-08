package gormT

import (
	"testing"

	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// type biz_AuthLog struct {
// 	gorm.Model
// 	RCId       int
// 	UserId     string
// 	Level      int
// 	UpdateTime time.Time
// 	Operator   string
// }

var db *gorm.DB

func TestQuery(t *testing.T) {
	var err error
	db, err = gorm.Open("mssql", "{SQL Server};server=127.0.0.1;uid=sa;pwd=123456;database=MZC;")
	if err != nil {
		fmt.Println("Open DB Error:", err.Error())
	}
	t.Log("End")
}
