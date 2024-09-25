//go:build !(amd64 || arm64 || riscv64) || purego

// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

import (
	"encoding/binary"

	"github.com/go-faster/errors"
)

var _ = binary.LittleEndian // clickHouse uses LittleEndian

// DecodeColumn decodes Int256 rows from *Reader.
func (c *ColInt256) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	const size = 256 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	// Move bound check out of loop.
	//
	// See https://github.com/golang/go/issues/30945.
	_ = data[len(data)-size]
	for i := 0; i <= len(data)-size; i += size {
		v = append(v,
			Int256(binUInt256(data[i:i+size])),
		)
	}
	*c = v
	return nil
}

// EncodeColumn encodes Int256 rows to *Buffer.
func (c ColInt256) EncodeColumn(b *Buffer) {
	v := c
	if len(v) == 0 {
		return
	}
	const size = 256 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(v))...)
	for _, vv := range v {
		binPutUInt256(
			b.Buf[offset:offset+size],
			UInt256(vv),
		)
		offset += size
	}
}

func (c ColInt256) WriteColumn(w *Writer) {
	w.ChainBuffer(c.EncodeColumn)
}
