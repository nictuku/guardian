package compiler

import (
	"errors"
	"github.com/spearson78/guardian/script/opcode"
	"github.com/spearson78/guardian/script/token"
	"math"
	"math/big"
)

type TokenSource interface {
	Scan() (tok token.Token)

	Op() opcode.OpCode
	Data() []byte
	Number() *big.Int
}

var c16 = big.NewInt(16)
var cMinus1 = big.NewInt(-1)

func appendPushData(dst []byte, data []byte) []byte {
	l := len(data)
	if l <= 75 {
		dst = append(dst, byte(l))
	} else if l <= 255 {
		dst = append(dst, 0x4c, byte(l))
	} else if l <= math.MaxUint16 {
		dst = append(dst, 0x4d, byte(l&0xFF), byte((l>>8)&0xFF))
	} else {
		dst = append(dst, 0x4e, byte(l&0xFF), byte((l>>8)&0xFF), byte((l>>16)&0xFF), byte((l>>24)&0xFF))
	}

	dst = append(dst, data...)
	return dst
}

func Compile(s TokenSource) ([]byte, error) {

	compiled := make([]byte, 0, 128)

	tok := s.Scan()
	for tok != token.ENDOFSCRIPT {

		switch tok {
		case token.DATA:
			compiled = appendPushData(compiled, s.Data())
		case token.NUMBER:
			n := s.Number()
			switch {
			case n.Sign() == 0:
				compiled = append(compiled, 0x00)
			case n.Cmp(cMinus1) == 0:
				compiled = append(compiled, 0x4f)
			case n.Sign() > 0 && n.Cmp(c16) <= 0:
				compiled = append(compiled, byte(80+n.Int64()))
			default:
				sign := n.Sign()

				n.Abs(n)
				data := n.Bytes()

				if data[0]&0x80 == 0x80 {
					extended := make([]byte, len(data)+1)
					copy(extended[1:], data)
					data = extended
				}

				if sign == -1 {
					data[0] = data[0] | 0x80
				}

				reverseInPlace(data)

				compiled = appendPushData(compiled, data)
			}
		case token.OPERATION:
			compiled = append(compiled, byte(s.Op()))
		case token.CODESEPARATOR:
			compiled = append(compiled, 0xab)
		case token.IF:
			compiled = append(compiled, 0x63)
		case token.NOTIF:
			compiled = append(compiled, 0x64)
		case token.ELSE:
			compiled = append(compiled, 0x67)
		case token.ENDIF:
			compiled = append(compiled, 0x68)
		case token.INVALID:
			return nil, errors.New("Invalid Token")
		default:
			return nil, errors.New("Unknown Token" + tok.String())
		}

		tok = s.Scan()
	}

	return compiled, nil

}

func reverseInPlace(in []byte) {

	l := len(in)
	hl := l / 2

	for i := 0; i < hl; i++ {
		j := l - i - 1
		in[i], in[j] = in[j], in[i]
	}
}
