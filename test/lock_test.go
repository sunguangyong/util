package test

import (
	"context"
	"fmt"
	"testing"
	"util/lock"

	v3 "go.etcd.io/etcd/client/v3"
)

func TestEtcdLock(t *testing.T) {
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

	fmt.Println("wailt lock")
	e.Mutex.Lock(context.Background())
	fmt.Println("get Lock")

	defer func() {
		e.Mutex.Unlock(context.Background())
		fmt.Println(e.Key, "release lock")
	}()

}
