package tcp

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/msglog"
	"github.com/davyxu/cellnet/relay"
	"github.com/davyxu/cellnet/rpc"
	"github.com/davyxu/ulog"
)

// 带有RPC和relay功能
type MsgHooker struct {
}

func (self MsgHooker) OnInboundEvent(inputEvent cellnet.Event) (outputEvent cellnet.Event) {

	var handled bool
	var err error

	inputEvent, handled, err = rpc.ResolveInboundEvent(inputEvent)

	if err != nil {
		ulog.Errorln("rpc.ResolveInboundEvent:", err)
		return
	}

	if !handled {

		inputEvent, handled, err = relay.ResoleveInboundEvent(inputEvent)

		if err != nil {
			ulog.Errorln("relay.ResoleveInboundEvent:", err)
			return
		}

		if !handled {
			msglog.WriteRecvLogger("tcp", inputEvent.Session(), inputEvent.Message())
		}
	}

	return inputEvent
}

func (self MsgHooker) OnOutboundEvent(inputEvent cellnet.Event) (outputEvent cellnet.Event) {

	handled, err := rpc.ResolveOutboundEvent(inputEvent)

	if err != nil {
		ulog.Errorln("rpc.ResolveOutboundEvent:", err)
		return nil
	}

	if !handled {

		handled, err = relay.ResolveOutboundEvent(inputEvent)

		if err != nil {
			ulog.Errorln("relay.ResolveOutboundEvent:", err)
			return nil
		}

		if !handled {
			msglog.WriteSendLogger("tcp", inputEvent.Session(), inputEvent.Message())
		}
	}

	return inputEvent
}
