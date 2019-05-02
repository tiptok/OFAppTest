package gormT

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
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

type AlarmFlag struct {
	AlarmFlag  int    `gorm:"column:AlarmFlag"`
	AlarmName  string `gorm:"column:AlarmName"`
	AlarmType  int    `gorm:"column:AlarmType"`
	Remark     string `gorm:"column:Remark"`
	AlarmOrder int    `gorm:"column:AlarmOrder"`
}

func (t AlarmFlag) TableName() string {
	return "bas_AlarmFlag"
}

func TestMysqlQuery(t *testing.T) {
	var err error
	db, err = gorm.Open("mysql", "root:admin@(127.0.0.1:3306)/TopDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Open DB Error:", err.Error())
	} else {
		fmt.Println("db connect:", db)
	}
	fmt.Println("has table bas_AlarmFlag:", db.HasTable("bas_AlarmFlag"))
	alarmItem := AlarmFlag{}
	db.Table("bas_AlarmFlag").First(&alarmItem)
	fmt.Println("AlarmFlag Item:", alarmItem)
	t.Log("End")
}

/*测试原生sql 操作*/

func mysql_instance() (*sql.DB, error) {
	sdsn := "root:admin@tcp(47.254.82.16:3306)/TopDB?charset=utf8mb4"
	return sql.Open("mysql", sdsn)
}

func Test_mysql_insert(t *testing.T) {
	db, err := mysql_instance()
	if err != nil {
		t.Fatal(err)
	}
	stmt, err := db.Prepare(`INSERT INTO bas_AlarmFlag(AlarmFlag,AlarmName)VALUES(?,?)`)
	defer stmt.Close()
	if err != nil {
		t.Fatal(err)
	}
	ret, err := stmt.Exec(8192, "超速预警")
	if err != nil {
		t.Fatal(err)
	}
	if rowsAffected, err := ret.RowsAffected(); err == nil {
		fmt.Println("RowsAffected:", rowsAffected)
	}
}

func Test_mysql_update(t *testing.T) {
	db, err := mysql_instance()
	if err != nil {
		t.Fatal(err)
	}
	stmt, err := db.Prepare(`UPDATE bas_AlarmFlag SET AlarmName=? WHERE AlarmFlag=?`)
	defer stmt.Close()
	ret, err := stmt.Exec("超速预警1", 8192)
	if err != nil {
		t.Fatal(err)
	}
	if rowsAffected, err := ret.RowsAffected(); err == nil {
		fmt.Println("RowsAffected:", rowsAffected)
	}
}

func Test_mysql_delete(t *testing.T) {
	db, err := mysql_instance()
	if err != nil {
		t.Fatal(err)
	}
	stmt, err := db.Prepare(`DELETE FROM bas_AlarmFlag WHERE AlarmFlag=?`)
	defer stmt.Close()
	ret, err := stmt.Exec(8192)
	if err != nil {
		t.Fatal(err)
	}
	if rowsAffected, err := ret.RowsAffected(); err == nil {
		fmt.Println("RowsAffected:", rowsAffected)
	}
}

func Test_mysql_select(t *testing.T) {
	db, err := mysql_instance()
	if err != nil {
		t.Fatal(err)
	}
	stmt, err := db.Prepare(`SELECT * FROM bas_AlarmFlag`)
	defer stmt.Close()
	rows, err := stmt.Query()
	columns, _ := rows.Columns()
	fmt.Println("Columns:", columns)

	scans := make([]interface{}, len(columns))
	//valu
	i := 0
	for rows.Next() {
		PaddingMapValue(scans)
		err = rows.Scan(scans...)
		fmt.Printf("Index:%d Values:%v \n", i, scans)
		i++
	}
}

func PaddingMapValue(m []interface{}) {
	for i := 0; i < len(m); i++ {
		var value interface{}
		m[i] = &value
	}
}

func PaddingMapValue(m []interface{}) {
	for i := 0; i < len(m); i++ {
		var value interface{}
		m[i] = &value
	}
}
