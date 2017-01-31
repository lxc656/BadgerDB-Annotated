package table

import (
	"bytes"
	"io"

	"github.com/dgraph-io/badger/y"
	"github.com/dgraph-io/dgraph/x"
)

type Block struct {
	data []byte
}

/*
 *itr.Seek(key)
 *for itr.Seek(key); itr.Valid(); itr.Next() {
 *  f(itr.key(), itr.value())
 *}
 */
type BlockIterator struct {
	data    []byte
	pos     int
	err     error
	lastKey []byte

	ikey []byte

	key  []byte
	val  []byte
	init bool
}

func (itr *BlockIterator) Reset() {
	itr.pos = 0
	itr.err = nil
	itr.lastKey = []byte{}
	itr.key = []byte{}
	itr.val = []byte{}
	itr.init = false
}

func (itr *BlockIterator) Init() {
	if !itr.init {
		itr.Next()
	}
}

func (itr *BlockIterator) Valid() bool {
	return itr.err == nil
}

func (itr *BlockIterator) Error() error {
	return itr.err
}

func (itr *BlockIterator) adjustSize(h header) {
	if cap(itr.ikey) < h.plen+h.klen {
		sz := h.plen + h.klen
		if sz < 2*cap(itr.ikey) {
			sz = 2 * cap(itr.ikey)
		}
		itr.ikey = make([]byte, sz)
	}
}

var (
	ORIGIN  = 0
	CURRENT = 1
)

func (itr *BlockIterator) Seek(seek []byte, whence int) {
	switch whence {
	case ORIGIN:
		itr.Reset()
	case CURRENT:
	}

	var done bool
	for itr.Init(); itr.Valid(); itr.Next() {
		itr.KV(func(k, v []byte) {
			if bytes.Compare(k, seek) >= 0 {
				done = true
			}
		})
		if done {
			break
		}
	}
}

func (itr *BlockIterator) Next() {
	itr.init = true
	if itr.pos >= len(itr.data) {
		itr.err = io.EOF
		return
	}

	var h header
	itr.pos += h.Decode(itr.data[itr.pos:])

	if len(itr.lastKey) == 0 {
		y.AssertTrue(h.plen == 0)
	}
	itr.adjustSize(h)

	itr.key = itr.ikey[:h.plen+h.klen]
	x.AssertTrue(h.plen == copy(itr.key, itr.lastKey[:h.plen]))
	x.AssertTrue(h.klen == copy(itr.key[h.plen:], itr.data[itr.pos:itr.pos+h.klen]))
	if h.plen == 0 {
		// If prefix length was zero, update the lastKey.
		itr.lastKey = itr.data[itr.pos : itr.pos+h.klen]
	}
	itr.pos += h.klen

	if itr.pos+h.vlen > len(itr.data) {
		itr.err = y.Errorf("Value exceeded size of block.")
		return
	}

	itr.val = itr.data[itr.pos : itr.pos+h.vlen]
	itr.pos += h.vlen
}

func (itr *BlockIterator) KV(fn func(k, v []byte)) {
	if itr.err != nil {
		return
	}

	fn(itr.key, itr.val)
}
