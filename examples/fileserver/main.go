package main

import (
	"flag"
	"fmt"
	"github.com/songweihang/cellnet"
	"github.com/songweihang/cellnet/peer"
	_ "github.com/songweihang/cellnet/peer/http"
	"github.com/songweihang/cellnet/proc"
	_ "github.com/songweihang/cellnet/proc/http"
)

var shareDir = flag.String("share", ".", "folder to share")
var port = flag.Int("port", 9091, "listen port")

func main() {

	flag.Parse()

	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer("http.Acceptor", "httpfile", fmt.Sprintf(":%d", *port), nil).(cellnet.HTTPAcceptor)
	p.SetFileServe(".", *shareDir)

	proc.BindProcessorHandler(p, "http", nil)

	p.Start()
	queue.StartLoop()

	queue.Wait()
}
