package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	//"go.etcd.io/etcd/embed"
	//"go.etcd.io/etcd/pkg/osutil"
)

func main() {
	//go etcdMain()
	caddycmd.Main()
}

//func etcdMain() {
//	cfg := embed.NewConfig()
//	e, err := embed.StartEtcd(cfg)
//	if err != nil {
//		panic(err)
//	}
//	osutil.RegisterInterruptHandler(e.Close)
//	select {
//	case <-e.Server.ReadyNotify(): // wait for e.Server to join the cluster
//	case <-e.Server.StopNotify(): // publish aborted from 'ErrStopped'
//	}
//	stopped := e.Server.StopNotify()
//	select {
//	case err := <-e.Err():
//		panic(err)
//	case <-stopped:
//	}
//	osutil.Exit(0)
//}
