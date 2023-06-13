package core

import (
	"github.com/rubixchain/rubixgoplatform/core/model"
	"github.com/rubixchain/rubixgoplatform/token"
)

func (c *Core) DumpTokenChain(dr *model.TCDumpRequest) *model.TCDumpReply {
	ds := &model.TCDumpReply{
		BasicResponse: model.BasicResponse{
			Status: false,
		},
	}
	tt := token.RBTTokenType
	if c.testNet {
		tt = token.TestTokenType
	}
	blks, nextID, err := c.w.GetAllTokenBlocks(dr.Token, tt, dr.BlockID)
	if err != nil {
		ds.Message = "Failed to get token chain block"
		return ds
	}
	ds.Status = true
	ds.Message = "Successfully got the token chain block"
	ds.Blocks = blks
	ds.NextBlockID = nextID
	return ds
}

func (c *Core) RemoveTokenChain(removeReq *model.TCRemoveRequest) *model.TCRemoveReply {
	removeReply := &model.TCRemoveReply{
		BasicResponse: model.BasicResponse{
			Status: false,
		},
	}
	tt := token.RBTTokenType
	if c.testNet {
		tt = token.TestTokenType
	}
	err := c.w.RemoveTokenChain(removeReq.Token, tt)
	if err != nil {
		removeReply.Message = "Failed to remove token chain "
		return removeReply
	}
	removeReply.Status = true
	removeReply.Message = "Successfully removed token chain "
	return removeReply
}

func (c *Core) SyncTokenChainFromAddress(syncReq *model.TCSyncRequest) *model.BasicResponse {
	reply := model.BasicResponse{
		Status: false,
	}
	tokentype := token.RBTTokenType
	if c.testNet {
		tokentype = token.TestTokenType
	}
	addr := syncReq.Address
	token := syncReq.Token
	peerConnection, err := c.getPeer(addr)
	if err != nil {
		c.log.Error("Could not connec tto address", addr, "err", err)
		reply.Message = "Could not connec tto address : " + addr + " err : " + err.Error()
		return &reply
	}
	err = c.syncTokenChainFrom(peerConnection, "", token, tokentype)
	if err != nil {
		c.log.Error("Could not sync token chain for token ", token, "err", err)
		reply.Message = "Could not sync token chain for token : " + token + " err : " + err.Error()
		return &reply
	}
	reply.Status = true
	reply.Message = "TokenChain synced successfully"
	return &reply
}
