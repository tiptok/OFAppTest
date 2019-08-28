package gitetcd_io

import (
	"context"
	"fmt"
	"github.com/etcd-io/etcd/clientv3"
	"log"
	"sync"
	"time"
)

type ClientDis struct {
	client *clientv3.Client
	serverList map[string]string
	lock sync.Mutex
}

func NewClientDis (addr []string)( *ClientDis, error){
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	var (
		client *clientv3.Client
		err error
	)
	if client, err = clientv3.New(conf); err != nil {
		return nil ,err
	}
	return &ClientDis{
		client:client,
		serverList:make(map[string]string),
	}, nil
}

func (this * ClientDis) GetService(key string) ([]string ,error){
	resp, err := this.client.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}

	addrs := this.extractAddrs(resp)
	go this.watcher(key)
	return addrs ,nil
}
func (this *ClientDis) watcher(prefix string) {
	rch := this.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				this.SetServiceList(string(ev.Kv.Key),string(ev.Kv.Value))
			case mvccpb.DELETE:
				this.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

func (this *ClientDis) extractAddrs(resp *clientv3.GetResponse) []string {
	addrs := make([]string,0)
	if resp == nil || resp.Kvs == nil {
		return addrs
	}
	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			this.SetServiceList(string(resp.Kvs[i].Key),string(resp.Kvs[i].Value))
			addrs = append(addrs, string(v))
		}
	}
	return addrs
}

//operator to ClientDis.serviceList
func (this *ClientDis) SetServiceList(key,val string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.serverList[key] = string(val)
	log.Println("set data key :",key,"val:",val)
}
func (this *ClientDis) DelServiceList(key string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.serverList,key)
	log.Println("del data key:", key)
}
func (this *ClientDis) SerList2Array()[]string {
	this.lock.Lock()
	defer this.lock.Unlock()
	addrs := make([]string,0)

	for _, v := range this.serverList {
		addrs = append(addrs,v)
	}
	return addrs
}



/*服务注册*/
type ServiceReg struct {
	client *clientv3.Client
	lease clientv3.Lease
	leaseResp *clientv3.LeaseGrantResponse
	canclefunc func()
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key string
}

func NewServiceReg(addr []string, timeNum int64) (*ServiceReg, error) {
	conf :=clientv3.Config{
		Endpoints:addr,
		DialTimeout:5*time.Second,
	}
	var(
		client *clientv3.Client
		err  error
	)
	client,err =clientv3.New(conf)
	if err!=nil{
		return nil,err
	}
	ser :=&ServiceReg{
		client:client,
	}
	if err := ser.setLease(timeNum); err != nil {
		return nil, err
	}
	go ser.ListenLeaseRespChan()
	return ser, nil
}
//监听 续租情况
func (this *ServiceReg) ListenLeaseRespChan() {
	for {
		select {
		case leaseKeepResp := <-this.keepAliveChan:
			if leaseKeepResp == nil {
				fmt.Printf("已经关闭续租功能\n")
				return
			} else {
				fmt.Printf("续租成功\n")
			}
		}
	}
}
//设置租约  设置ServiceReg 参数
func (this *ServiceReg) setLease(timeNum int64) error {
	lease := clientv3.NewLease(this.client)

	//设置租约时间
	leaseResp, err := lease.Grant(context.TODO(), timeNum)
	if err != nil {
		return err
	}

	//设置续租
	ctx, cancelFunc := context.WithCancel(context.TODO())
	leaseRespChan, err := lease.KeepAlive(ctx, leaseResp.ID)

	if err != nil {
		return err
	}

	this.lease = lease
	this.leaseResp = leaseResp
	this.canclefunc = cancelFunc
	this.keepAliveChan = leaseRespChan
	return nil
}

//通过租约 注册服务
func(this *ServiceReg)PutService(key,val string)error{
	kv :=clientv3.NewKV(this.client)
	_,err :=kv.Put(context.TODO(),key,val,clientv3.WithLease(this.leaseResp.ID))
	return err
}
//撤销租约
func (this *ServiceReg) RevokeLease() error {
	this.canclefunc()
	time.Sleep(2 * time.Second)
	_, err := this.lease.Revoke(context.TODO(), this.leaseResp.ID)
	return err
}

func EtcdLease(){
	ser,_:=NewServiceReg([]string{"127.0.0.1:2379"},5)
	ser.PutService("/test/dd","authService")
	select {
	}
}

func Test_etcd_Client (t *testing.T) {
	/*
	//https://www.jianshu.com/p/7c0d23c818a5
	1.创建一个client 连到etcd。
	2.匹配到所有相同前缀的 key。把值存到 serverList 这个map里面。
	3.watch这个 key前缀，当有增加或者删除的时候 就 修改这个map。
	4.所以这个map就是 实时的 服务列表
	*/
	cli,_ := NewClientDis([]string{"127.0.0.1:2379"})
	cli.GetService("name")
	select {}
	//EtcdLease()
}



