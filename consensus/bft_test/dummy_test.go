package bft_test

import (
	"github.com/annchain/OG/consensus/bft"
)

type LocalBftPeerCommunicator struct {
	Myid    int
	PeerIns []chan *bft.BftMessageEvent
	pipeIn  chan *bft.BftMessageEvent //pipeIn is the receiver of the outside messages
	pipeOut chan *bft.BftMessageEvent //pipeOut is the providing channel for new messages parsed from pipeIn
}

func (d *LocalBftPeerCommunicator) HandleIncomingMessage(msgEvent *bft.BftMessageEvent) {
	d.pipeIn <- msgEvent
}

func NewDummyBftPeerCommunicator(myid int, incoming chan *bft.BftMessageEvent, peers []chan *bft.BftMessageEvent) *LocalBftPeerCommunicator {
	d := &LocalBftPeerCommunicator{
		PeerIns: peers,
		Myid:    myid,
		pipeIn:  incoming,
		pipeOut: make(chan *bft.BftMessageEvent),
	}
	return d
}

func (d *LocalBftPeerCommunicator) Broadcast(msg bft.BftMessage, peers []bft.PeerInfo) {
	for _, peer := range peers {
		go func(peer bft.PeerInfo) {
			//ffchan.NewTimeoutSenderShort(d.PeerIns[peer.Id], msg, "bft")
			d.PeerIns[peer.Id] <- &bft.BftMessageEvent{
				Message: msg,
				Peer:    peer,
			}
		}(peer)
	}
}

func (d *LocalBftPeerCommunicator) Unicast(msg bft.BftMessage, peer bft.PeerInfo) {
	go func() {
		//ffchan.NewTimeoutSenderShort(d.PeerIns[peer.Id], msg, "bft")
		d.PeerIns[peer.Id] <- &bft.BftMessageEvent{
			Message: msg,
			Peer:    peer,
		}
	}()
}

func (d *LocalBftPeerCommunicator) GetPipeIn() chan *bft.BftMessageEvent {
	return d.pipeIn
}

func (d *LocalBftPeerCommunicator) GetPipeOut() chan *bft.BftMessageEvent {
	return d.pipeOut
}

func (d *LocalBftPeerCommunicator) Run() {
	go func() {
		for {
			v := <-d.pipeIn
			//vv := v.Message.(bft.BftMessage)
			d.pipeOut <- v
		}
	}()
}

type dummyProposalGenerator struct {
	CurrentHeight uint64
}

func (d dummyProposalGenerator) ProduceProposal() (proposal bft.Proposal, validCondition bft.ProposalCondition) {
	p := bft.StringProposal{
		Content: "XXX",
	}
	return &p, bft.ProposalCondition{ValidHeight: d.CurrentHeight}
}

type dummyProposalValidator struct {
}

func (d dummyProposalValidator) ValidateProposal(proposal bft.Proposal) error {
	return nil
}

type dummyDecisionMaker struct {
}

func (d dummyDecisionMaker) MakeDecision(proposal bft.Proposal, state *bft.HeightRoundState) (bft.ConsensusDecision, error) {
	return proposal, nil
}