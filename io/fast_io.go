package io

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func _read(x *int) {
	isN := false
	var b byte
	for b, _ = reader.ReadByte(); b != '-' && (b < '0' || b > '9'); {
		b, _ = reader.ReadByte()
	}
	if b == '-' {
		*x = 0
		isN = true
	} else {
		*x = int(b - '0')
	}
	for b, _ = reader.ReadByte(); b >= '0' && b <= '9'; {
		*x = *x*10 + int(b-'0')
		b, _ = reader.ReadByte()
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
