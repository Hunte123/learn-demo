package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	endPoints := []string{
		"127.0.0.1:2379",
		"127.0.0.1:22379",
		"127.0.0.1:32379",
	}
	cli, err := clientv3.New(
		clientv3.Config{
			Endpoints:   endPoints,
			DialTimeout: 5 * time.Second,
		})
	defer cli.Close()
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connect success")

	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	cli.Put(ctx, "test", "1")
	cancel()

	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, _ := cli.Get(ctx, "test")
	cancel()
	for _, one := range resp.Kvs {
		fmt.Printf("%s : %s", one.Key, one.Value)
		//fmt.Println(one.Key, one.Value)
	}

	//watch
	for {
		wch := cli.Watch(context.Background(), "test")
		for resp := range wch {
			for _, one := range resp.Events {
				fmt.Printf("%s : %s", one.Kv.Key, one.Kv.Value)
			}
		}
	}
}
