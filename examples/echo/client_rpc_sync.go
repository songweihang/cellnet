package main

import (
	"github.com/songweihang/cellnet"
	"github.com/songweihang/cellnet/peer"
	"github.com/songweihang/cellnet/proc"
	"github.com/songweihang/cellnet/rpc"
	"time"
)

func clientSyncRPC() {

	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer("tcp.Connector", "async rpc", peerAddress, queue)

	// 创建一个消息同步接收器
	rv := proc.NewSyncReceiver(p)

	proc.BindProcessorHandler(p, "tcp.ltv", rv.EventCallback())

	p.Start()

	queue.StartLoop()

	// 等连接上时
	rv.WaitMessage("cellnet.SessionConnected")

	// 同步RPC
	rpc.CallSync(p, &TestEchoACK{
		Msg:   "hello",
		Value: 1234,
	}, time.Second)
}
