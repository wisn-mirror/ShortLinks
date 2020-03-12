package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
 
type User struct {
	Id int `db:id`
	Age int `db:age`
	Name string `db:Name`
	V1 float32 `db:v1`
}

func main() {
	initDb()
	//QueryByFiled(1)
	//QueryAll(10)
	//insert(33)
	//updateNameById(2,"test000002")
	//QueryAll(10)
	//deleteById(1)
	//QueryAll(10)
	TxExcute()
	QueryAll(10)

}

//事务
func TxExcute()  {
	tx, txError := mysqldb.Begin()
	if txError!=nil{
		return
	}
	stmt, prepareError := tx.Prepare("insert into tab01(age,name ) value (?,?)")
	if prepareError!=nil{
		tx.Rollback()
	}
	result, _ := stmt.Exec(22, "wisn22")
	//执行行
	insertNum,_:=result.RowsAffected()
	lastinsertId,_:=result.LastInsertId()
	fmt.Println("lastinsertId",lastinsertId)
	prepareStmt, _ := tx.Prepare("update  tab01 set v1=? , age=? where id=?")
	resultupdate, _ := prepareStmt.Exec(4.2,39, lastinsertId)
	updateNum, _ := resultupdate.RowsAffected()
	fmt.Println("insert num",insertNum," update ",updateNum)
	if insertNum==updateNum&&insertNum==1{
		tx.Commit()
	}else{
		tx.Rollback()
	}
}

func deleteById(id int )  {
	result, execError := mysqldb.Exec("delete from tab01 where id =?", id)
	if execError!=nil{
		updatenums,error:=result.RowsAffected()
		fmt.Println("成功执行了",updatenums,error)
	}
}

func updateNameById(id int ,name string)  {
	result,error:=mysqldb.Exec("update  tab01 set name=? where id=?",name,id)
	if error!=nil{
		fmt.Println(error)
		return
	}
	updatenums, rowAffectdError:= result.RowsAffected()
	fmt.Println("成功执行",updatenums,rowAffectdError)
}

func insert(i int )  {
	ret,error:=mysqldb.Exec("insert into tab01(age,name)values(?,?)",i,"hello")
	if error!=nil{
		fmt.Println(error)
		return
	}
	insertId, _ := ret.LastInsertId()
	fmt.Println("id",insertId)
}

func QueryAll( limit int )  {
	users:=make([]User,0)
	//rows, _:=mysqldb.Query("select *  from tab01 limit ?",limit)
	rows, _:=mysqldb.Query("select id,age,name  from tab01 limit ?",limit)
	//顺序要保持一直
	for rows.Next(){
		var user User
		if error:=rows.Scan(&user.Id,&user.Age,&user.Name);error!=nil{
			fmt.Println("error is",error)
			continue
		}
		users=append(users,user)

	}
	fmt.Println("user list data",users)
}

func QueryByFiled(id int )  {
	user:=new (User)
	row:=mysqldb.QueryRow("select id,age,name from tab01 where id=?",id)
	//顺序要保持一直
	if error:=row.Scan(&user.Id,&user.Age,&user.Name);error!=nil{
		fmt.Println("error is",error)
		return
	}
	fmt.Println("user data",user)
}

const (
	USER_NAME = "root"
	PASS_WORD = "nihao@123456"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "test01"
	CHARSET   = "utf8"
)
var mysqldb *sql.DB
var mysqlError error
func initDb() bool  {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	mysqldb, mysqlError = sql.Open("mysql", dbDSN)
	//mysqldb, mysqlError = sql.Open("mysql", "root:nihao@123456@/test01")
	if mysqlError!=nil{
		fmt.Println("connect error",mysqlError)
		return false
	}
	//最大链接数
	mysqldb.SetMaxOpenConns(100)
	//空闲连接数
	mysqldb.SetMaxIdleConns(20)
	//最大连接周期
	mysqldb.SetConnMaxLifetime(100*time.Second)
	if pingError:=mysqldb.Ping();pingError!=nil{
		fmt.Println("连接失败 ",pingError)
		return false
	}
	fmt.Println("success",mysqldb)
	return true
}

