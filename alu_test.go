package alu

import (
	"testing"

	"github.com/yasukei/go-alu"
)

type TestData struct {
	rhs    []byte
	lhs    []byte
	expect []byte
}

func (td *TestData) assert(t *testing.T, got []byte) bool {
	expect := td.expect

	if len(got) != len(expect) {
		t.Errorf("got:      %d", len(got))
		t.Errorf("expected: %d", len(expect))
		return false
	}

	for i, _ := range got {
		if got[i] != expect[i] {
			t.Errorf("index:    %d", i)
			t.Errorf("got:      %#02x", got[i])
			t.Errorf("expected: %#02x", expect[i])
			return false
		}
	}
	return true
}

func TestAnd(t *testing.T) {
	testdata := []TestData{
		{
			rhs:    []byte{0x00},
			lhs:    []byte{0x00},
			expect: []byte{0x00},
		},
		{
			rhs:    []byte{0x00},
			lhs:    []byte{0xFF},
			expect: []byte{0x00},
		},
		{
			rhs:    []byte{0xFF},
			lhs:    []byte{0x00},
			expect: []byte{0x00},
		},
		{
			rhs:    []byte{0xFF},
			lhs:    []byte{0xFF},
			expect: []byte{0xFF},
		},
		{
			rhs:    []byte{0xAA},
			lhs:    []byte{0x55},
			expect: []byte{0x00},
		},
		{
			rhs:    []byte{0xFF, 0xFF},
			lhs:    []byte{0x55},
			expect: []byte{0x55, 0x55},
		},
		{
			rhs:    []byte{0xFF, 0xFF, 0xFF},
			lhs:    []byte{0x55},
			expect: []byte{0x55, 0x55, 0x55},
		},
		{
			rhs:    []byte{0xFF, 0xFF, 0xFF},
			lhs:    []byte{0x55, 0x00},
			expect: []byte{0x55, 0x00, 0x55},
		},
		{
			rhs:    []byte{0xFF, 0xFF, 0xFF},
			lhs:    []byte{0x55, 0x00, 0xAA},
			expect: []byte{0x55, 0x00, 0xAA},
		},
		{
			rhs:    []byte{0xFF},
			lhs:    []byte{0x55, 0x00},
			expect: []byte{0x55},
		},
		{
			rhs:    []byte{0xFF},
			lhs:    []byte{0x55, 0x00, 0xAA},
			expect: []byte{0x55},
		},
		{
			rhs:    []byte{0xFF, 0xFF},
			lhs:    []byte{0x55, 0x00, 0xAA},
			expect: []byte{0x55, 0x00},
		},
	}

	for i, td := range testdata {
		rhs := td.rhs
		lhs := td.lhs

		alu.And(rhs, rhs, lhs)
		got := rhs

		if !td.assert(t, got) {
			t.Errorf("failed at testdata[%d]", i)
		}
	}
}

func TestOr(t *testing.T) {
	testdata := []TestData{
		{
			rhs:    []byte{0x00},
			lhs:    []byte{0x00},
			expect: []byte{0x00},
		},
		{
			rhs:    []byte{0x00},
			lhs:    []byte{0xFF},
			expect: []byte{0xFF},
		},
		{
			rhs:    []byte{0xFF},
			lhs:    []byte{0x00},
			expect: []byte{0xFF},
		},
		{
			rhs:    []byte{0xFF},
			lhs:    []byte{0xFF},
			expect: []byte{0xFF},
		},
	}

	for i, td := range testdata {
		rhs := td.rhs
		lhs := td.lhs

		alu.Or(rhs, rhs, lhs)
		got := rhs

		if !td.assert(t, got) {
			t.Errorf("failed at testdata[%d]", i)
		}
	}
}

func TestNot(t *testing.T) {
	testdata := []TestData{
		{
			rhs:    []byte{0x00},
			lhs:    []byte{0x00}, // dummy
			expect: []byte{0xFF},
		},
		{
			rhs:    []byte{0xFF},
			lhs:    []byte{0x00},
			expect: []byte{0x00},
		},
	}

	for i, td := range testdata {
		rhs := td.rhs

		alu.Not(rhs, rhs)
		got := rhs

		if !td.assert(t, got) {
			t.Errorf("failed at testdata[%d]", i)
		}
	}
}

func TestXor(t *testing.T) {
	testdata := []TestData{
		{
			rhs:    []byte{0x00},
			lhs:    []byte{0x00},
			expect: []byte{0x00},
		},
		{
			rhs:    []byte{0x00},
			lhs:    []byte{0xFF},
			expect: []byte{0xFF},
		},
		{
			rhs:    []byte{0xFF},
			lhs:    []byte{0x00},
			expect: []byte{0xFF},
		},
		{
			rhs:    []byte{0xFF},
			lhs:    []byte{0xFF},
			expect: []byte{0x00},
		},
	}

	for i, td := range testdata {
		rhs := td.rhs
		lhs := td.lhs

		alu.Xor(rhs, rhs, lhs)
		got := rhs

		if !td.assert(t, got) {
			t.Errorf("failed at testdata[%d]", i)
		}
	}
}
