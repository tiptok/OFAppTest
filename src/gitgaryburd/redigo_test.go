package gitgaryburd

import (
	"log"
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestRedigo(t *testing.T) {
	c, err := redis.Dial("tcp", "127.0.0.1:6333")
	if err != nil {
		t.Fatal(err.Error())
	}
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
}
