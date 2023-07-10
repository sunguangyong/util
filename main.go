package main

import (
	"context"
	"fmt"
	"time"
	"util/lock"

	v3 "go.etcd.io/etcd/client/v3"
)

func main() {
	endpoints := []string{"127.0.0.1:12379", "127.0.0.1:22379", "127.0.0.1:32379"}
	// 初始化etcd客户端
	client, err := v3.New(v3.Config{Endpoints: endpoints})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	key := "test"

	e, err := lock.NewEtcdLock(client, key)

	if err != nil {
		fmt.Println("err ==== ", err)
		return
	}

	go testLock1(e)
	go testLock2(e)

	time.Sleep(time.Minute)
}

func testLock1(e *lock.EtcdLock) (err error) {
	e.Mutex.Lock(context.Background())
	fmt.Println("testLock1")

	defer func() {
		e.Mutex.Unlock(context.Background())
		fmt.Println(e.Key, "unlock")
	}()

	start := time.Now()

	fmt.Println(e.Key, "cost", time.Now().Sub(start))
	time.Sleep(5 * time.Second)

	if err != nil {
		return err
	}
	return err
}

func testLock2(e *lock.EtcdLock) (err error) {
	e.Mutex.Lock(context.Background())
	fmt.Println("testLock2")

	defer func() {
		e.Mutex.Unlock(context.Background())
		fmt.Println(e.Key, "unlock")
	}()

	start := time.Now()

	fmt.Println(e.Key, "cost", time.Now().Sub(start))
	time.Sleep(5 * time.Second)

	if err != nil {
		return err
	}
	return err
}
