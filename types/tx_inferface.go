package types

//go:generate msgp
//cccmsgp:tuple TxBase

type Txi interface {
	// Hash returns a tx hash
	BlockHash() Hash
}

<<<<<<< HEAD
	// Parents returns the hash of txs that it directly proves.
	Parents() []common.Hash
}
=======
type TxBase struct {
	Type          int    `msgp:"type"`
	ParentsHash   []Hash `msgp:"parentHash"`
	SequenceNonce uint64 `msgp:"sequenceNonce"`
	Height        uint64 `msgp:"height"`
}
>>>>>>> dev