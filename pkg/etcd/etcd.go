package etcd

import (
	"context"

	etcdv3 "go.etcd.io/etcd/client/v3"
)

type Client interface {
	Put(ctx context.Context, key, val string, opts ...etcdv3.OpOption) (*etcdv3.PutResponse, error)
	Get(ctx context.Context, key string, opts ...etcdv3.OpOption) (*etcdv3.GetResponse, error)
}

type ETCD struct {
	etcd Client
}

func New(etcdClient Client) *ETCD {
	return &ETCD{etcd: etcdClient}
}
