package main

import (
	"fmt"
	"math/rand"
)

//获取随机抽奖人
func Test_r_getAward() {
	var users map[string]int64 = map[string]int64{
		"a": 10,
		"b": 6,
		"c": 3,
		"d": 12,
		"f": 1,
	}
	award_stat := make(map[string]int64)
	for i := 0; i < 10000; i++ {
		name := GetAwardUserNameByOrders(users)
		if count, ok := award_stat[name]; ok {
			award_stat[name] = count + 1
		} else {
			award_stat[name] = 1
		}
	}

	for name, count := range award_stat {
		fmt.Printf("user: %s, award count: %d\n", name, count)
	}
}

func GetAwardName(users map[string]int64) (name string) {
	sizeofUser := len(users)
	award_index := rand.Intn(sizeofUser)

	var index int
	for u_name, _ := range users {
		if index == award_index {
			name = u_name
			return
		}
		index += 1
	}
	return
}

//按购买订单的权重
func GetAwardUserNameByOrders(users map[string]int64) (name string) {
	type A_user struct {
		Name   string
		Offset int64
		Num    int64
	}

	a_user_arr := make([]*A_user, 0)
	var sum_num int64
	for name, num := range users {
		a_user := &A_user{
			Name:   name,
			Offset: sum_num,
			Num:    num,
		}
		a_user_arr = append(a_user_arr, a_user)
		sum_num += num
	}

	award_num := rand.Int63n(sum_num)

	for index, _ := range a_user_arr {
		a_user := a_user_arr[index]
		if a_user.Offset+a_user.Num > award_num {
			name = a_user.Name
			return
		}
	}
	return
}
