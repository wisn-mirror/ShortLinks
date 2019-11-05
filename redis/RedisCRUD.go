package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, _ := redis.Dial("tcp", "localhost:6379")
	defer conn.Close()
	isExist("key1", conn)
	isExist("key2", conn)
	addValue("key2", "hahahah2222",conn)
	//conn.Do("set", "key2", "hahahah2222")
	result, _ := redis.String(conn.Do("get", "key2"))
	fmt.Println(result)

	isExist("key1", conn)
	isExist("key2", conn)
	//delete("key2",conn)
	setExpireTime("key1",3,conn)
	setExpireTime("key2",3,conn)
	flushReceive(conn)
}
//管道化
func flushReceive( conn redis.Conn)  {
	conn.Send("set","key2","333")
	conn.Send("set","key2","333")
	conn.Send("set","key3","333")
	conn.Send("get","key2")
	conn.Flush()
	receiveOne(conn)
	receiveOne(conn)
	receiveOne(conn)
	receiveOne(conn)
}
func receiveOne(conn redis.Conn) {
	//value1, error1 := conn.Receive()
	s, error1 := redis.String(conn.Receive())
	fmt.Println(" value :",s, error1)
}

//设置超时时间
func setExpireTime( key string,expriretime int64,conn redis.Conn)  {
	_, error := conn.Do("expire", key, expriretime)
	if error != nil {
		fmt.Println("redis add", error)
	}
}

//添加value
func addValue(key string ,value string,conn redis.Conn)  {
	_, error := conn.Do("set", key, value)
	if error!=nil{
		fmt.Println("redis add",error)
	}
}

// 删除 key
func delete(key string, conn redis.Conn) {
	_, error := redis.Bool(conn.Do("del", key))
	if error!=nil{
		fmt.Println("redis delete",error)
	}
}

// 是否存在 key
func isExist(key string, conn redis.Conn) {
	isExist, _ := redis.Bool(conn.Do("exists", key))
	fmt.Println("是否存在",key , isExist)
}
