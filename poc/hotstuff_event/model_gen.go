package hotstuff_event

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Block) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	z.Round, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "Round")
		return
	}
	z.Payload, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Payload")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "ParentQC")
			return
		}
		z.ParentQC = nil
	} else {
		if z.ParentQC == nil {
			z.ParentQC = new(QC)
		}
		err = z.ParentQC.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "ParentQC")
			return
		}
	}
	z.Id, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Block) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Round)
	if err != nil {
		err = msgp.WrapError(err, "Round")
		return
	}
	err = en.WriteString(z.Payload)
	if err != nil {
		err = msgp.WrapError(err, "Payload")
		return
	}
	if z.ParentQC == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.ParentQC.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "ParentQC")
			return
		}
	}
	err = en.WriteString(z.Id)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Block) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o = msgp.AppendInt(o, z.Round)
	o = msgp.AppendString(o, z.Payload)
	if z.ParentQC == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.ParentQC.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "ParentQC")
			return
		}
	}
	o = msgp.AppendString(o, z.Id)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Block) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	z.Round, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Round")
		return
	}
	z.Payload, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Payload")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.ParentQC = nil
	} else {
		if z.ParentQC == nil {
			z.ParentQC = new(QC)
		}
		bts, err = z.ParentQC.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "ParentQC")
			return
		}
	}
	z.Id, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Block) Msgsize() (s int) {
	s = 1 + msgp.IntSize + msgp.StringPrefixSize + len(z.Payload)
	if z.ParentQC == nil {
		s += msgp.NilSize
	} else {
		s += z.ParentQC.Msgsize()
	}
	s += msgp.StringPrefixSize + len(z.Id)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ContentProposal) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	err = z.Proposal.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Proposal")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ContentProposal) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	err = z.Proposal.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Proposal")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ContentProposal) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	o, err = z.Proposal.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Proposal")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ContentProposal) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	bts, err = z.Proposal.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Proposal")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ContentProposal) Msgsize() (s int) {
	s = 1 + z.Proposal.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ContentString) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	z.Content, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Content")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z ContentString) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	err = en.WriteString(z.Content)
	if err != nil {
		err = msgp.WrapError(err, "Content")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ContentString) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	o = msgp.AppendString(o, z.Content)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ContentString) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	z.Content, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Content")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ContentString) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.Content)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ContentTimeout) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.Round, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "Round")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "HighQC")
			return
		}
		z.HighQC = nil
	} else {
		if z.HighQC == nil {
			z.HighQC = new(QC)
		}
		err = z.HighQC.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "HighQC")
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ContentTimeout) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Round)
	if err != nil {
		err = msgp.WrapError(err, "Round")
		return
	}
	if z.HighQC == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.HighQC.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "HighQC")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ContentTimeout) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendInt(o, z.Round)
	if z.HighQC == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.HighQC.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "HighQC")
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ContentTimeout) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.Round, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Round")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.HighQC = nil
	} else {
		if z.HighQC == nil {
			z.HighQC = new(QC)
		}
		bts, err = z.HighQC.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "HighQC")
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ContentTimeout) Msgsize() (s int) {
	s = 1 + msgp.IntSize
	if z.HighQC == nil {
		s += msgp.NilSize
	} else {
		s += z.HighQC.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ContentVote) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	err = z.VoteInfo.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfo")
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo")
		return
	}
	if zb0002 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0002}
		return
	}
	z.LedgerCommitInfo.CommitStateId, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "CommitStateId")
		return
	}
	z.LedgerCommitInfo.VoteInfoHash, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "VoteInfoHash")
		return
	}
	var zb0003 uint32
	zb0003, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "Signatures")
		return
	}
	if cap(z.Signatures) >= int(zb0003) {
		z.Signatures = (z.Signatures)[:zb0003]
	} else {
		z.Signatures = make([]Signature, zb0003)
	}
	for za0001 := range z.Signatures {
		var zb0004 uint32
		zb0004, err = dc.ReadArrayHeader()
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001)
			return
		}
		if zb0004 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0004}
			return
		}
		z.Signatures[za0001].PartnerId, err = dc.ReadInt()
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001, "PartnerId")
			return
		}
		z.Signatures[za0001].Signature, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001, "Signature")
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ContentVote) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 3
	err = en.Append(0x93)
	if err != nil {
		return
	}
	err = z.VoteInfo.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfo")
		return
	}
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteString(z.LedgerCommitInfo.CommitStateId)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "CommitStateId")
		return
	}
	err = en.WriteString(z.LedgerCommitInfo.VoteInfoHash)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "VoteInfoHash")
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Signatures)))
	if err != nil {
		err = msgp.WrapError(err, "Signatures")
		return
	}
	for za0001 := range z.Signatures {
		// array header, size 2
		err = en.Append(0x92)
		if err != nil {
			return
		}
		err = en.WriteInt(z.Signatures[za0001].PartnerId)
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001, "PartnerId")
			return
		}
		err = en.WriteString(z.Signatures[za0001].Signature)
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001, "Signature")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ContentVote) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 3
	o = append(o, 0x93)
	o, err = z.VoteInfo.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfo")
		return
	}
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendString(o, z.LedgerCommitInfo.CommitStateId)
	o = msgp.AppendString(o, z.LedgerCommitInfo.VoteInfoHash)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Signatures)))
	for za0001 := range z.Signatures {
		// array header, size 2
		o = append(o, 0x92)
		o = msgp.AppendInt(o, z.Signatures[za0001].PartnerId)
		o = msgp.AppendString(o, z.Signatures[za0001].Signature)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ContentVote) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	bts, err = z.VoteInfo.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfo")
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo")
		return
	}
	if zb0002 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0002}
		return
	}
	z.LedgerCommitInfo.CommitStateId, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "CommitStateId")
		return
	}
	z.LedgerCommitInfo.VoteInfoHash, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "VoteInfoHash")
		return
	}
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Signatures")
		return
	}
	if cap(z.Signatures) >= int(zb0003) {
		z.Signatures = (z.Signatures)[:zb0003]
	} else {
		z.Signatures = make([]Signature, zb0003)
	}
	for za0001 := range z.Signatures {
		var zb0004 uint32
		zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001)
			return
		}
		if zb0004 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0004}
			return
		}
		z.Signatures[za0001].PartnerId, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001, "PartnerId")
			return
		}
		z.Signatures[za0001].Signature, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001, "Signature")
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ContentVote) Msgsize() (s int) {
	s = 1 + z.VoteInfo.Msgsize() + 1 + msgp.StringPrefixSize + len(z.LedgerCommitInfo.CommitStateId) + msgp.StringPrefixSize + len(z.LedgerCommitInfo.VoteInfoHash) + msgp.ArrayHeaderSize
	for za0001 := range z.Signatures {
		s += 1 + msgp.IntSize + msgp.StringPrefixSize + len(z.Signatures[za0001].Signature)
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LedgerCommitInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.CommitStateId, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "CommitStateId")
		return
	}
	z.VoteInfoHash, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "VoteInfoHash")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z LedgerCommitInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteString(z.CommitStateId)
	if err != nil {
		err = msgp.WrapError(err, "CommitStateId")
		return
	}
	err = en.WriteString(z.VoteInfoHash)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfoHash")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z LedgerCommitInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendString(o, z.CommitStateId)
	o = msgp.AppendString(o, z.VoteInfoHash)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LedgerCommitInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.CommitStateId, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "CommitStateId")
		return
	}
	z.VoteInfoHash, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfoHash")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z LedgerCommitInfo) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.CommitStateId) + msgp.StringPrefixSize + len(z.VoteInfoHash)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MsgType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int
		zb0001, err = dc.ReadInt()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = MsgType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z MsgType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt(int(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MsgType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt(o, int(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MsgType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int
		zb0001, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = MsgType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MsgType) Msgsize() (s int) {
	s = msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *QC) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	err = z.VoteInfo.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfo")
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo")
		return
	}
	if zb0002 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0002}
		return
	}
	z.LedgerCommitInfo.CommitStateId, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "CommitStateId")
		return
	}
	z.LedgerCommitInfo.VoteInfoHash, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "VoteInfoHash")
		return
	}
	var zb0003 uint32
	zb0003, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "Signatures")
		return
	}
	if cap(z.Signatures) >= int(zb0003) {
		z.Signatures = (z.Signatures)[:zb0003]
	} else {
		z.Signatures = make([]Signature, zb0003)
	}
	for za0001 := range z.Signatures {
		var zb0004 uint32
		zb0004, err = dc.ReadArrayHeader()
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001)
			return
		}
		if zb0004 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0004}
			return
		}
		z.Signatures[za0001].PartnerId, err = dc.ReadInt()
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001, "PartnerId")
			return
		}
		z.Signatures[za0001].Signature, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001, "Signature")
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *QC) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 3
	err = en.Append(0x93)
	if err != nil {
		return
	}
	err = z.VoteInfo.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfo")
		return
	}
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteString(z.LedgerCommitInfo.CommitStateId)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "CommitStateId")
		return
	}
	err = en.WriteString(z.LedgerCommitInfo.VoteInfoHash)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "VoteInfoHash")
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Signatures)))
	if err != nil {
		err = msgp.WrapError(err, "Signatures")
		return
	}
	for za0001 := range z.Signatures {
		// array header, size 2
		err = en.Append(0x92)
		if err != nil {
			return
		}
		err = en.WriteInt(z.Signatures[za0001].PartnerId)
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001, "PartnerId")
			return
		}
		err = en.WriteString(z.Signatures[za0001].Signature)
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001, "Signature")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *QC) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 3
	o = append(o, 0x93)
	o, err = z.VoteInfo.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfo")
		return
	}
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendString(o, z.LedgerCommitInfo.CommitStateId)
	o = msgp.AppendString(o, z.LedgerCommitInfo.VoteInfoHash)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Signatures)))
	for za0001 := range z.Signatures {
		// array header, size 2
		o = append(o, 0x92)
		o = msgp.AppendInt(o, z.Signatures[za0001].PartnerId)
		o = msgp.AppendString(o, z.Signatures[za0001].Signature)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *QC) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	bts, err = z.VoteInfo.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfo")
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo")
		return
	}
	if zb0002 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0002}
		return
	}
	z.LedgerCommitInfo.CommitStateId, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "CommitStateId")
		return
	}
	z.LedgerCommitInfo.VoteInfoHash, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "VoteInfoHash")
		return
	}
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Signatures")
		return
	}
	if cap(z.Signatures) >= int(zb0003) {
		z.Signatures = (z.Signatures)[:zb0003]
	} else {
		z.Signatures = make([]Signature, zb0003)
	}
	for za0001 := range z.Signatures {
		var zb0004 uint32
		zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001)
			return
		}
		if zb0004 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0004}
			return
		}
		z.Signatures[za0001].PartnerId, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001, "PartnerId")
			return
		}
		z.Signatures[za0001].Signature, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Signatures", za0001, "Signature")
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *QC) Msgsize() (s int) {
	s = 1 + z.VoteInfo.Msgsize() + 1 + msgp.StringPrefixSize + len(z.LedgerCommitInfo.CommitStateId) + msgp.StringPrefixSize + len(z.LedgerCommitInfo.VoteInfoHash) + msgp.ArrayHeaderSize
	for za0001 := range z.Signatures {
		s += 1 + msgp.IntSize + msgp.StringPrefixSize + len(z.Signatures[za0001].Signature)
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Signature) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.PartnerId, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "PartnerId")
		return
	}
	z.Signature, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Signature")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Signature) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteInt(z.PartnerId)
	if err != nil {
		err = msgp.WrapError(err, "PartnerId")
		return
	}
	err = en.WriteString(z.Signature)
	if err != nil {
		err = msgp.WrapError(err, "Signature")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Signature) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendInt(o, z.PartnerId)
	o = msgp.AppendString(o, z.Signature)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Signature) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.PartnerId, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "PartnerId")
		return
	}
	z.Signature, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Signature")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Signature) Msgsize() (s int) {
	s = 1 + msgp.IntSize + msgp.StringPrefixSize + len(z.Signature)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VoteInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 7 {
		err = msgp.ArrayError{Wanted: 7, Got: zb0001}
		return
	}
	z.Id, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	z.Round, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "Round")
		return
	}
	z.ParentId, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "ParentId")
		return
	}
	z.ParentRound, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "ParentRound")
		return
	}
	z.GrandParentId, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "GrandParentId")
		return
	}
	z.GrandParentRound, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "GrandParentRound")
		return
	}
	z.ExecStateId, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "ExecStateId")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VoteInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 7
	err = en.Append(0x97)
	if err != nil {
		return
	}
	err = en.WriteString(z.Id)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	err = en.WriteInt(z.Round)
	if err != nil {
		err = msgp.WrapError(err, "Round")
		return
	}
	err = en.WriteString(z.ParentId)
	if err != nil {
		err = msgp.WrapError(err, "ParentId")
		return
	}
	err = en.WriteInt(z.ParentRound)
	if err != nil {
		err = msgp.WrapError(err, "ParentRound")
		return
	}
	err = en.WriteString(z.GrandParentId)
	if err != nil {
		err = msgp.WrapError(err, "GrandParentId")
		return
	}
	err = en.WriteInt(z.GrandParentRound)
	if err != nil {
		err = msgp.WrapError(err, "GrandParentRound")
		return
	}
	err = en.WriteString(z.ExecStateId)
	if err != nil {
		err = msgp.WrapError(err, "ExecStateId")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VoteInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 7
	o = append(o, 0x97)
	o = msgp.AppendString(o, z.Id)
	o = msgp.AppendInt(o, z.Round)
	o = msgp.AppendString(o, z.ParentId)
	o = msgp.AppendInt(o, z.ParentRound)
	o = msgp.AppendString(o, z.GrandParentId)
	o = msgp.AppendInt(o, z.GrandParentRound)
	o = msgp.AppendString(o, z.ExecStateId)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VoteInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 7 {
		err = msgp.ArrayError{Wanted: 7, Got: zb0001}
		return
	}
	z.Id, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	z.Round, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Round")
		return
	}
	z.ParentId, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "ParentId")
		return
	}
	z.ParentRound, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "ParentRound")
		return
	}
	z.GrandParentId, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "GrandParentId")
		return
	}
	z.GrandParentRound, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "GrandParentRound")
		return
	}
	z.ExecStateId, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "ExecStateId")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VoteInfo) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.Id) + msgp.IntSize + msgp.StringPrefixSize + len(z.ParentId) + msgp.IntSize + msgp.StringPrefixSize + len(z.GrandParentId) + msgp.IntSize + msgp.StringPrefixSize + len(z.ExecStateId)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VoteMsg) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	err = z.VoteInfo.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfo")
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo")
		return
	}
	if zb0002 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0002}
		return
	}
	z.LedgerCommitInfo.CommitStateId, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "CommitStateId")
		return
	}
	z.LedgerCommitInfo.VoteInfoHash, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "VoteInfoHash")
		return
	}
	z.Sender, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "Sender")
		return
	}
	var zb0003 uint32
	zb0003, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "Signature")
		return
	}
	if zb0003 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0003}
		return
	}
	z.Signature.PartnerId, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "Signature", "PartnerId")
		return
	}
	z.Signature.Signature, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Signature", "Signature")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VoteMsg) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = z.VoteInfo.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfo")
		return
	}
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteString(z.LedgerCommitInfo.CommitStateId)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "CommitStateId")
		return
	}
	err = en.WriteString(z.LedgerCommitInfo.VoteInfoHash)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "VoteInfoHash")
		return
	}
	err = en.WriteInt(z.Sender)
	if err != nil {
		err = msgp.WrapError(err, "Sender")
		return
	}
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Signature.PartnerId)
	if err != nil {
		err = msgp.WrapError(err, "Signature", "PartnerId")
		return
	}
	err = en.WriteString(z.Signature.Signature)
	if err != nil {
		err = msgp.WrapError(err, "Signature", "Signature")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VoteMsg) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o, err = z.VoteInfo.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfo")
		return
	}
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendString(o, z.LedgerCommitInfo.CommitStateId)
	o = msgp.AppendString(o, z.LedgerCommitInfo.VoteInfoHash)
	o = msgp.AppendInt(o, z.Sender)
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendInt(o, z.Signature.PartnerId)
	o = msgp.AppendString(o, z.Signature.Signature)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VoteMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	bts, err = z.VoteInfo.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "VoteInfo")
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo")
		return
	}
	if zb0002 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0002}
		return
	}
	z.LedgerCommitInfo.CommitStateId, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "CommitStateId")
		return
	}
	z.LedgerCommitInfo.VoteInfoHash, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LedgerCommitInfo", "VoteInfoHash")
		return
	}
	z.Sender, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Sender")
		return
	}
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Signature")
		return
	}
	if zb0003 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0003}
		return
	}
	z.Signature.PartnerId, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Signature", "PartnerId")
		return
	}
	z.Signature.Signature, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Signature", "Signature")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VoteMsg) Msgsize() (s int) {
	s = 1 + z.VoteInfo.Msgsize() + 1 + msgp.StringPrefixSize + len(z.LedgerCommitInfo.CommitStateId) + msgp.StringPrefixSize + len(z.LedgerCommitInfo.VoteInfoHash) + msgp.IntSize + 1 + msgp.IntSize + msgp.StringPrefixSize + len(z.Signature.Signature)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *WireMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	z.MsgType, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "MsgType")
		return
	}
	z.ContentBytes, err = dc.ReadBytes(z.ContentBytes)
	if err != nil {
		err = msgp.WrapError(err, "ContentBytes")
		return
	}
	z.SenderId, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "SenderId")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *WireMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 3
	err = en.Append(0x93)
	if err != nil {
		return
	}
	err = en.WriteInt(z.MsgType)
	if err != nil {
		err = msgp.WrapError(err, "MsgType")
		return
	}
	err = en.WriteBytes(z.ContentBytes)
	if err != nil {
		err = msgp.WrapError(err, "ContentBytes")
		return
	}
	err = en.WriteString(z.SenderId)
	if err != nil {
		err = msgp.WrapError(err, "SenderId")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *WireMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 3
	o = append(o, 0x93)
	o = msgp.AppendInt(o, z.MsgType)
	o = msgp.AppendBytes(o, z.ContentBytes)
	o = msgp.AppendString(o, z.SenderId)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *WireMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	z.MsgType, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "MsgType")
		return
	}
	z.ContentBytes, bts, err = msgp.ReadBytesBytes(bts, z.ContentBytes)
	if err != nil {
		err = msgp.WrapError(err, "ContentBytes")
		return
	}
	z.SenderId, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "SenderId")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *WireMessage) Msgsize() (s int) {
	s = 1 + msgp.IntSize + msgp.BytesPrefixSize + len(z.ContentBytes) + msgp.StringPrefixSize + len(z.SenderId)
	return
}
