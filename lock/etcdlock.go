package lock

import (
	"fmt"

	v3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

type EtcdLock struct {
	Key   string
	Mutex *concurrency.Mutex
}

func NewEtcdLock(client *v3.Client, key string) (e *EtcdLock, err error) {

	session, err := concurrency.NewSession(client, concurrency.WithTTL(30))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	mutex := concurrency.NewMutex(session, fmt.Sprintf("/dLock/%s", key))

	e = &EtcdLock{
		Key:   key,
		Mutex: mutex,
	}

	return e, err
}
