package gitgaryburd

import (
	"fmt"
	"log"
	"testing"

	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	/*连接池*/
	RedisClient *redis.Pool
)

func TestRedigo(t *testing.T) {
	RedisClient = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   1024,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1"+":"+"6333")
			if err != nil {
				return nil, err
			}
			// 选择db
			//c.Do("SELECT", REDIS_DB)  默认 0
			return c, nil
		},
	}

	c := RedisClient.Get()
	defer c.Close()

	v, err := c.Do("SET", "age", 20)
	if err != nil {
		t.Fatal(err.Error())
	}

	v, err = redis.String(c.Do("GET", "age"))
	if err != nil {
		t.Fatal(err.Error())
	}
	log.Println("GET(age) Redis:", v)

	/*加入列表*/
	// c.Do("lpush", "SimNums", "18860183052")
	// c.Do("lpush", "SimNums", "18860183053")
	// c.Do("lpush", "SimNums", "18860183054")

	/*读取列表*/
	simNums, _ := redis.Values(c.Do("lrange", "SimNums", "0", "10"))
	for index, v := range simNums {
		/*v 存储的类型为 []byte*/
		log.Printf("List[%d]=%v\n", index, string(v.([]byte)))
	}

	/*管道*/
	// c.Send("SET", "year", "2018")
	// c.Send("SET", "address", "fz")
	//c.Flush()

	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	c.Send("SET", "address", "fz")
	// }()
	// time.Sleep(time.Second * 4)
	// rc, _ := c.Receive()
	//log.Println("Receive redis reply:", rc)

	/*redis 订阅*/
	go subscribe()
	go subscribe()
	go subscribe()

	for {
		var s string
		s = time.Now().String()
		_, err := c.Do("PUBLISH", "chat", s)
		if err != nil {
			fmt.Println("pub err: ", err)
			return
		}
		time.Sleep(time.Second * 5)
	}

}

func subscribe() {
	c := RedisClient.Get()
	psc := redis.PubSubConn{c}
	defer func() {
		c.Close()
		psc.Unsubscribe("chat")
	}()

	psc.Subscribe("chat")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}
