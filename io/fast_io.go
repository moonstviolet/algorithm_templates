package io

import (
	"bufio"
	"os"
)

var _reader = bufio.NewReader(os.Stdin)

func _read(x *int) {
	var (
		b   byte
		isN = false
	)
	for b, _ = _reader.ReadByte(); b != '-' && (b < '0' || b > '9'); {
		b, _ = _reader.ReadByte()
	}
	if b == '-' {
		*x = 0
		isN = true
	} else {
		*x = int(b - '0')
	}
	for b, _ = _reader.ReadByte(); b >= '0' && b <= '9'; {
		*x = *x*10 + int(b-'0')
		b, _ = _reader.ReadByte()
	}
	if isN {
		*x = -*x
	}
}

func Read(args ...*int) {
	for _, v := range args {
		_read(v)
	}
}
