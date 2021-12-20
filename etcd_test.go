package mocks_etcd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"mocks_etcd/mocks"
	"mocks_etcd/mocks/db/kvstore"
	"testing"
	"time"
)

type EtcdTestSuite struct {
	suite.Suite
	EtcdServer *mocks.EtcdServer
	// your client
	//client *Client
}

// 启动mock server
func (suite *EtcdTestSuite) SetupSuite() {
	// 启动mock server
	etcdServer := mocks.StartEtcdServer(mocks.MKConfig("voltha.mock.test", 23781, 23801, "voltha.lib.mocks.etcd", "error"))
	if etcdServer == nil {
		return
	}
	suite.EtcdServer = etcdServer
	//初始化mock client
	clientAddr := fmt.Sprintf("localhost:%d", 23781)
	ctx := context.Background()
	client, err := kvstore.NewEtcdClient(ctx, clientAddr, 30*time.Second)
	if err != nil || client == nil {
		return
	}
	//初始化我们需要用到的client
	//vclient, _ := client.Pool.Get(ctx)
	//etcdClient := &Client {
	//	V3Client: vclient,
	//	Timeout: 30 * time.Second,
	//	watchMethodMapping: make(map[EvtData]EvtFunc),
	//}
	//suite.client = etcdClient
}
//
//func (suite *EtcdTestSuite)TestEtcdBasePutGet() {
//	_, err := suite.client.Put("111", "2222")
//	suite.Equal(err, nil)
//
//	resp, err := suite.client.Get("111")
//	suite.Equal(err, nil)
//	suite.Equal(resp["111"], "2222")
//}


func TestEtcdTestSuite(t *testing.T) {
	suite.Run(t, new(EtcdTestSuite))
}

func (suite *EtcdTestSuite) TearDownSuite() {
	suite.EtcdServer.Stop()
	//suite.client.Close()
}

