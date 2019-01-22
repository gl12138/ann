package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/annchain/OG/common/math"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *RawSequencer) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "TxBase":
			err = z.TxBase.DecodeMsg(dc)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *RawSequencer) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "TxBase"
	err = en.Append(0x81, 0xa6, 0x54, 0x78, 0x42, 0x61, 0x73, 0x65)
	if err != nil {
		return
	}
	err = z.TxBase.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RawSequencer) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "TxBase"
	o = append(o, 0x81, 0xa6, 0x54, 0x78, 0x42, 0x61, 0x73, 0x65)
	o, err = z.TxBase.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RawSequencer) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "TxBase":
			bts, err = z.TxBase.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *RawSequencer) Msgsize() (s int) {
	s = 1 + 7 + z.TxBase.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RawSequencers) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(RawSequencers, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(RawSequencer)
			}
			var field []byte
			_ = field
			var zb0003 uint32
			zb0003, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zb0003 > 0 {
				zb0003--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "TxBase":
					err = (*z)[zb0001].TxBase.DecodeMsg(dc)
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z RawSequencers) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0004 := range z {
		if z[zb0004] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 1
			// write "TxBase"
			err = en.Append(0x81, 0xa6, 0x54, 0x78, 0x42, 0x61, 0x73, 0x65)
			if err != nil {
				return
			}
			err = z[zb0004].TxBase.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z RawSequencers) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0004 := range z {
		if z[zb0004] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 1
			// string "TxBase"
			o = append(o, 0x81, 0xa6, 0x54, 0x78, 0x42, 0x61, 0x73, 0x65)
			o, err = z[zb0004].TxBase.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RawSequencers) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(RawSequencers, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(RawSequencer)
			}
			var field []byte
			_ = field
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zb0003 > 0 {
				zb0003--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "TxBase":
					bts, err = (*z)[zb0001].TxBase.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z RawSequencers) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0004 := range z {
		if z[zb0004] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 7 + z[zb0004].TxBase.Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RawTx) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "TxBase":
			err = z.TxBase.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "To":
			err = z.To.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "Value":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Value = nil
			} else {
				if z.Value == nil {
					z.Value = new(math.BigInt)
				}
				err = z.Value.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *RawTx) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "TxBase"
	err = en.Append(0x83, 0xa6, 0x54, 0x78, 0x42, 0x61, 0x73, 0x65)
	if err != nil {
		return
	}
	err = z.TxBase.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "To"
	err = en.Append(0xa2, 0x54, 0x6f)
	if err != nil {
		return
	}
	err = z.To.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "Value"
	err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	if z.Value == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Value.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RawTx) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "TxBase"
	o = append(o, 0x83, 0xa6, 0x54, 0x78, 0x42, 0x61, 0x73, 0x65)
	o, err = z.TxBase.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "To"
	o = append(o, 0xa2, 0x54, 0x6f)
	o, err = z.To.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "Value"
	o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if z.Value == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Value.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RawTx) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "TxBase":
			bts, err = z.TxBase.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "To":
			bts, err = z.To.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "Value":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Value = nil
			} else {
				if z.Value == nil {
					z.Value = new(math.BigInt)
				}
				bts, err = z.Value.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *RawTx) Msgsize() (s int) {
	s = 1 + 7 + z.TxBase.Msgsize() + 3 + z.To.Msgsize() + 6
	if z.Value == nil {
		s += msgp.NilSize
	} else {
		s += z.Value.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RawTxs) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(RawTxs, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(RawTx)
			}
			err = (*z)[zb0001].DecodeMsg(dc)
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z RawTxs) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		if z[zb0003] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z[zb0003].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z RawTxs) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		if z[zb0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[zb0003].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RawTxs) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(RawTxs, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(RawTx)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z RawTxs) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}
