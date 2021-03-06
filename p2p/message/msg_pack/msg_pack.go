package msgpack

import (
	"time"

	mt "github.com/coschain/contentos-go/p2p/message/types"
	"github.com/coschain/contentos-go/p2p/net/protocol"
	"github.com/coschain/contentos-go/prototype"
	"github.com/coschain/gobft/message"
)

//Peer address package
func NewAddrs(nodeAddrs []*mt.PeerAddr) mt.Message {
	var reqmsg mt.TransferMsg
	data := new(mt.Address)
	data.Addr = nodeAddrs

	reqmsg.Msg = &mt.TransferMsg_Msg5{Msg5:data}
	return &reqmsg
}

//Peer address request package
func NewAddrReq() mt.Message {
	var reqmsg mt.TransferMsg
	data := new(mt.AddrReq)

	reqmsg.Msg = &mt.TransferMsg_Msg6{Msg6:data}
	return &reqmsg
}

//block package
func NewSigBlkIdMsg(bk *prototype.SignedBlock) mt.Message {
	var reqmsg mt.TransferMsg
	data := new(mt.IdMsg)
	var tmp []byte
	data.Msgtype = mt.IdMsg_broadcast_sigblk_id
	id := bk.Id()
	data.Value = append(data.Value, tmp)
	data.Value[0] = id.Data[:]

	reqmsg.Msg = &mt.TransferMsg_Msg2{Msg2:data}
	return &reqmsg
}

func NewSigBlk(bk *prototype.SignedBlock, needTriggerFetch bool) mt.Message {
	var reqmsg mt.TransferMsg
	data := new(mt.SigBlkMsg)
	data.SigBlk = bk
	data.NeedTriggerFetch = needTriggerFetch

	reqmsg.Msg = &mt.TransferMsg_Msg3{Msg3:data}
	return &reqmsg
}

//ping msg package
func NewPingMsg(height uint64) mt.Message {
	var reqmsg mt.TransferMsg
	data := new(mt.Ping)
	data.Height = height

	reqmsg.Msg = &mt.TransferMsg_Msg8{Msg8:data}
	return &reqmsg
}

//pong msg package
func NewPongMsg(height uint64) mt.Message {
	var reqmsg mt.TransferMsg
	data := new(mt.Pong)
	data.Height = height

	reqmsg.Msg = &mt.TransferMsg_Msg9{Msg9:data}
	return &reqmsg
}

//Transaction package
func NewTxn(txn *prototype.SignedTransaction) mt.Message {
	var reqmsg mt.TransferMsg
	data := new(mt.BroadcastSigTrx)
	data.SigTrx = txn

	reqmsg.Msg = &mt.TransferMsg_Msg1{Msg1:data}
	return &reqmsg
}

//version ack package
func NewVerAck(isConsensus bool) mt.Message {
	var reqmsg mt.TransferMsg
	data := new(mt.VerAck)
	data.IsConsensus = isConsensus

	reqmsg.Msg = &mt.TransferMsg_Msg10{Msg10:data}
	return &reqmsg
}

//Version package
func NewVersion(n p2p.P2P, isCons bool, height uint64, runningVersion string) mt.Message {
	var reqmsg mt.TransferMsg

	 data := &mt.Version{
		Version            : n.GetVersion(),
		Services           : n.GetServices(),
		SyncPort           : n.GetSyncPort(),
		ConsPort           : n.GetConsPort(),
		Nonce              : n.GetID(),
		IsConsensus        : isCons,
		StartHeight        : uint64(height),
		Timestamp          : time.Now().UnixNano(),
		RunningCodeVersion : runningVersion,
	}
	if n.GetRelay() {
		data.Relay = 1
	} else {
		data.Relay = 0
	}

	 reqmsg.Msg = &mt.TransferMsg_Msg11{Msg11:data}
	 return &reqmsg
}

// consensus package
func NewConsMsg(msg message.ConsensusMessage, needForwardBroadcast bool) mt.Message {
	var reqmsg mt.ConsMsg
	reqmsg.Extra = new(mt.ConsensusExtraData)

	reqmsg.MsgData = msg
	if needForwardBroadcast {
		reqmsg.Extra.Bcast = 1
	} else {
		reqmsg.Extra.Bcast = 0
	}

	return &reqmsg
}

// checkpoint package
func NewCheckpointBatchMsg(startNum, endNum uint64) mt.Message {
	var reqmsg mt.TransferMsg

	data := new(mt.RequestCheckpointBatch)
	data.Start = startNum
	data.End = endNum

	reqmsg.Msg = &mt.TransferMsg_Msg12{Msg12:data}
	return &reqmsg
}

func NewRequestOutOfRangeIds(startID, targetID []byte) mt.Message {
	var reqmsg mt.TransferMsg

	data := new(mt.RequestOutOfRangeIds)
	data.StartId = startID
	data.TargetId = targetID

	reqmsg.Msg = &mt.TransferMsg_Msg13{Msg13:data}
	return &reqmsg
}

func NewDetectFormerIds(endID []byte) mt.Message{
	var reqmsg mt.TransferMsg

	data := new(mt.DetectFormerIds)
	data.EndId = endID

	reqmsg.Msg = &mt.TransferMsg_Msg15{Msg15:data}
	return &reqmsg
}

func NewRequestBlockBatch(startID, endID []byte) mt.Message {
	var reqmsg mt.TransferMsg

	data := new(mt.RequestBlockBatch)
	data.StartId = startID
	data.EndId = endID

	reqmsg.Msg = &mt.TransferMsg_Msg14{Msg14:data}
	return &reqmsg
}

func NewClearOutOfRangeState() mt.Message {
	var reqmsg mt.TransferMsg

	data := new(mt.ClearOutOfRangeState)

	reqmsg.Msg = &mt.TransferMsg_Msg16{Msg16:data}
	return &reqmsg
}