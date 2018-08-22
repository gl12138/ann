package types

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/crypto/sha3"
)

//go:generate msgp
//msgp:tuple Sequencer

type Sequencer struct {
	Id uint64 `msgp:"id"`
	TxBase
	ContractHashOrder []Hash `msgp:"contractHashOrder"`
}

func (t *Sequencer) BlockHash() (hash Hash) {
	var buf bytes.Buffer

	err := binary.Write(&buf, binary.LittleEndian, t.Id)
	panicIfError(err)

	for _, ancestor := range t.ParentsHash {
		err = binary.Write(&buf, binary.LittleEndian, ancestor)
		panicIfError(err)
	}

	result := sha3.Sum256(buf.Bytes())
	hash.MustSetBytes(result[0:])
	return
}
