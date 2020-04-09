package alu

func traverse(dst, src, arg []byte, ope func(byte, byte) byte) {
	j := 0
	for i, _ := range dst {
		dst[i] = ope(src[i], arg[j])
		j++
		if j == len(arg) {
			j = 0
		}
	}
}

func and(lhs byte, rhs byte) byte {
	return lhs & rhs
}

func or(lhs byte, rhs byte) byte {
	return lhs | rhs
}

func not(lhs byte, rhs byte) byte {
	return ^lhs
}

func xor(lhs byte, rhs byte) byte {
	return lhs ^ rhs
}

func And(dst, src, arg []byte) {
	traverse(dst, src, arg, and)
}

func Or(dst, src, arg []byte) {
	traverse(dst, src, arg, or)
}

func Not(dst, src []byte) {
	dummy := []byte{0x00}
	traverse(dst, src, dummy, not)
}

func Xor(dst, src, arg []byte) {
	traverse(dst, src, arg, xor)
}

//type And struct {
//	bytes []byte
//	cur   uint
//}
//
//func NewAnd(bytes []byte) *And {
//	and := make(And)
//	and.bytes = make([]byte, len(bytes))
//	and.cur = 0
//	return and
//}
//
//func (and *And) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//
//}
//
//func (and *And) Reset() {
//}
