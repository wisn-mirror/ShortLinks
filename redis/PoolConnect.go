package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
	"time"
)

func redisConn(ip,port,passwd string)(redis.Conn ,error)  {
	conn, error := redis.Dial("tcp", ip+":"+port,
		redis.DialConnectTimeout(5*time.Second),
		redis.DialReadTimeout(1*time.Second),
		redis.DialWriteTimeout(1*time.Second),
		redis.DialPassword(passwd),
		redis.DialKeepAlive(1*time.Second),
	)
	checkError(error)
	return conn,error
}
func checkError(err error)  {
	if err!=nil{
		fmt.Println(" error",err)
		os.Exit(-1)
	}
}

func newPool(ip,port,password string)*redis.Pool  {
	return &redis.Pool{
		MaxIdle:5,//最大空闲连接
		MaxActive:18,//最大连接数，限制并发数
		IdleTimeout:240*time.Second,
		MaxConnLifetime:300*time.Second,
		Dial: func() (conn redis.Conn, e error) {
			return redisConn(ip,port,password)
		},
	}
}

func main() {
	//pool:=newPool("localhost","6379","nihao@123456")
	pool:=newPool("localhost","6379","root")
	defer pool.Close()
	for i := 0; i <= 5; i++ {
		go func() {
			c:=pool.Get()
			defer c.Close()
			fmt.Printf("MaxActive %d MaxIdle %d \n",pool.MaxActive,pool.MaxIdle)
			_, err := c.Do("set", "22nametest", fmt.Sprintf("%dtest01", i))
			checkError(err)
			if s, getError := redis.String(c.Do("get", "22nametest"));getError==nil{
				_, setError := c.Do("set", "22nametest", fmt.Sprintf("%dtest04", i))
				fmt.Println("result ",s,setError)

			}
		}()
	}
	time.Sleep(time.Second*2)
}