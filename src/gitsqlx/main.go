package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var shema = `
	DROP TABLE IF EXISTS published;
	CREATE TABLE published  (
  id bigint(20) NOT NULL AUTO_INCREMENT,
  topic varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  name varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  version bigint(20) NOT NULL DEFAULT 0,
  msg varbinary(8192) NOT NULL,
  retries int(10) NOT NULL DEFAULT 0,
  created_at datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  updated_at datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  expires_at datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  status tinyint(4) NOT NULL DEFAULT 0,
  ack_status tinyint(4) UNSIGNED ZEROFILL NOT NULL DEFAULT 0000,
  PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
`
type Published struct {
	Id int64 `db:"id"`
	Topic string `db:"topic"`
	Name string `db:"name"`
	Version int64 `db:"version"`
	CreateAt time.Time `db:"create_at"`
	UpdateAt time.Time `db:"update_at"`
	ExpiresAt time.Time `db:"expires_at"`
}

func main()  {
	defer func(){
		if p:=recover();p!=nil{
			log.Println("recover",p)
		}
	}()
	db,err := sqlx.Connect("mysql","root:123456@tcp(127.0.0.1:3306)/mx_passport")
	if err!=nil{
		panic(err)
	}

	/*1.exec sql*/
	//result :=db.MustExec(shema)
	//log.Println("init db shema:",result)

	//tx:=db.MustBegin()
	//tx.MustExec("insert into published(`name`,topic,version)values (?,?,?)","aa","microx.topic.UserCreated",1)
	//tx.MustExec("insert into published(`name`,topic,version)values (?,?,?)","cc","microx.topic.UserLogin",1)
	//_,err=tx.NamedExec("insert into published(name,topic,version)values (:name,:topic,:version)",&Published{Name:"dd",Topic:"microx.topic.UserLogout",Version:1})
	//if err!=nil{
	//	tx.Rollback()
	//	log.Fatal(err)
	//}
	//tx.Commit()

	/*******2.select sql********/
	publishList :=[]Published{}
	err =db.Select(&publishList,"select id,topic,name,version from published order by id asc")
	if err!=nil{
		log.Println(err)
	}
	jsData,_ :=json.Marshal(publishList)
	log.Println("select list:",string(jsData))

	nameList :=[]string{}
	db.Select(&nameList,"select name from published order by id asc")
	log.Println("select array:",nameList)

	var publishItem Published
	err=db.Get(&publishItem,"select id,topic,name,version from published where name=?","aa")
	if err!=nil{
		log.Println(err)
	}
	log.Println("get item:",publishItem)

	/*******3.queryx=mutil row queryrowx=one ********/
	var publishListQueryx Published
	rows,err :=db.Queryx("select id,topic,name,version from published where name=?","aa")
	if err!=nil{
		log.Fatal(err)
	}
	for rows.Next(){
		if err :=rows.StructScan(&publishListQueryx);err==nil{
			log.Println("queryx :",publishListQueryx)
		}
	}
}

