/**
  @author:panliang
  @data:2022/6/3
  @note
**/
package uuid

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestUUid(T *testing.T) {
	mySnow, _ := NewSnowFlake(0, 0) //生成雪花算法
	group := sync.WaitGroup{}
	startTime := time.Now()
	generateId := func(s SnowFlake, requestNumber int) {
		for i := 0; i < requestNumber; i++ {
			s.NextId()
			group.Done()
		}
	}
	group.Add(40000)
	//生成并发的数为4000000
	currentThreadNum := 4
	for i := 0; i < currentThreadNum; i++ {
		generateId(*mySnow, 10000)
	}
	group.Wait()
	fmt.Printf("time: %v\n", time.Now().Sub(startTime))

}
