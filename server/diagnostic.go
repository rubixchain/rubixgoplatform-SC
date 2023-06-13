package server

import (
	"net/http"

	"github.com/EnsurityTechnologies/ensweb"
	"github.com/rubixchain/rubixgoplatform/core/model"
)

func (s *Server) APIDumpTokenChainBlock(req *ensweb.Request) *ensweb.Result {
	var dr model.TCDumpRequest
	err := s.ParseJSON(req, &dr)
	if err != nil {
		return s.BasicResponse(req, false, "Invalid input", nil)
	}
	drep := s.c.DumpTokenChain(&dr)
	return s.RenderJSON(req, drep, http.StatusOK)
}

func (s *Server) APIRemoveTokenChain(req *ensweb.Request) *ensweb.Result {
	var removeReq model.TCRemoveRequest
	err := s.ParseJSON(req, &removeReq)
	if err != nil {
		return s.BasicResponse(req, false, "Invalid input", nil)
	}
	removeReply := s.c.RemoveTokenChain(&removeReq)
	return s.RenderJSON(req, removeReply, http.StatusOK)
}

func (s *Server) APISyncTokenChainFromAddress(req *ensweb.Request) *ensweb.Result {
	var syncReq model.TCSyncRequest
	err := s.ParseJSON(req, &syncReq)
	if err != nil {
		return s.BasicResponse(req, false, "Invalid input", nil)
	}
	syncReply := s.c.SyncTokenChainFromAddress(&syncReq)
	return s.RenderJSON(req, syncReply, http.StatusOK)
}
