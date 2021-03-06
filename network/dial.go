package network

import (
	"bytes"
	"github.com/LemoFoundationLtd/lemochain-core/chain/deputynode"
	"github.com/LemoFoundationLtd/lemochain-core/common"
	"github.com/LemoFoundationLtd/lemochain-core/common/log"
	"github.com/LemoFoundationLtd/lemochain-core/common/subscribe"
	"github.com/LemoFoundationLtd/lemochain-core/network/p2p"
	"net"
	"time"
)

const (
	AddNewCorePeer = "addNewCorePeer"
	GetNewTx       = "getNewTx"
)

type DialManager struct {
	coreNodeID       *p2p.NodeID
	coreNodeEndpoint string
}

func NewDialManager(coreNodeID *p2p.NodeID, coreNodeEndpoint string) *DialManager {
	return &DialManager{
		coreNodeID:       coreNodeID,
		coreNodeEndpoint: coreNodeEndpoint,
	}
}

// Dial run dial
func (dm *DialManager) Dial() {
	// dial
	conn, err := net.DialTimeout("tcp", dm.coreNodeEndpoint, 5*time.Second)
	if err != nil {
		log.Warnf("dial node error: %s", err.Error())
		SetConnectResult(false)
		return
	}

	// handle connection
	if err = dm.handleConn(conn); err != nil {
		if err != p2p.ErrConnectSelf {
			SetConnectResult(false)
		}
		log.Debugf("handle connection error: %s", err)
		return
	}
}

// handleConn handle the connection
func (dm *DialManager) handleConn(fd net.Conn) error {
	p := p2p.NewPeer(fd)
	if err := p.DoHandshake(deputynode.GetSelfNodeKey(), dm.coreNodeID); err != nil {
		if err = fd.Close(); err != nil {
			log.Errorf("close connection failed: %v", err)
		}
		return err
	}
	// is self
	if bytes.Compare(p.RNodeID()[:], deputynode.GetSelfNodeID()) == 0 {
		if err := fd.Close(); err != nil {
			log.Errorf("close connections failed: %v", err)
		} else {
			log.Error("can't connect self")
		}
		return p2p.ErrConnectSelf
	}
	// go dm.runPeer(p)
	subscribe.Send(AddNewCorePeer, p)
	return nil
}

// runPeer run the connected peer
func (dm *DialManager) runPeer(p p2p.IPeer) {
	if err := p.Run(); err != nil { // block this
		log.Debugf("runPeer error: %v", err)
	}
	SetConnectResult(false)
	log.Debugf("peer Run finished: %s", common.ToHex(p.RNodeID()[:8]))
}
