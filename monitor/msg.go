package monitor

import (
	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"github.com/yggworldtree/go-core/messages"
	"github.com/yggworldtree/go-sdk/ywtree"
)

func (c *Manager) OnConnect(e *ywtree.Engine) {

}
func (c *Manager) OnDisconnect(e *ywtree.Engine) {

}
func (c *Manager) OnMessage(e *ywtree.Engine, msg *ywtree.MessageTopic) *messages.ReplyInfo {
	pths := msg.Path.String()
	hbtp.Debugf("OnMessage:%s,from:%s", pths, msg.Sender.String())

	return nil
}
func (c *Manager) OnBroadcast(e *ywtree.Engine, msg *messages.MessageBox) *messages.ReplyInfo {

	return nil
}
