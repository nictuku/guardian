package transaction

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"testing"
)

func TestGenesisTransaction(t *testing.T) {

	tx := Tx{
		Version: 1,
		Inputs: []TxIn{
			TxIn{
				Previous: OutPoint{
					Hash:  [32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
					Index: 0xFFFFFFFF,
				},
				Script:   []byte{0x04, 0xFF, 0xFF, 0x00, 0x1D, 0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6D, 0x65, 0x73, 0x20, 0x30, 0x33, 0x2F, 0x4A, 0x61, 0x6E, 0x2F, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, 0x61, 0x6E, 0x63, 0x65, 0x6C, 0x6C, 0x6F, 0x72, 0x20, 0x6F, 0x6E, 0x20, 0x62, 0x72, 0x69, 0x6E, 0x6B, 0x20, 0x6F, 0x66, 0x20, 0x73, 0x65, 0x63, 0x6F, 0x6E, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6C, 0x6F, 0x75, 0x74, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x62, 0x61, 0x6E, 0x6B, 0x73},
				Sequence: 0xFFFFFFFF,
			},
		},
		Outputs: []TxOut{
			TxOut{
				Value:  5000000000,
				Script: []byte{0x41, 0x04, 0x67, 0x8A, 0xFD, 0xB0, 0xFE, 0x55, 0x48, 0x27, 0x19, 0x67, 0xF1, 0xA6, 0x71, 0x30, 0xB7, 0x10, 0x5C, 0xD6, 0xA8, 0x28, 0xE0, 0x39, 0x09, 0xA6, 0x79, 0x62, 0xE0, 0xEA, 0x1F, 0x61, 0xDE, 0xB6, 0x49, 0xF6, 0xBC, 0x3F, 0x4C, 0xEF, 0x38, 0xC4, 0xF3, 0x55, 0x04, 0xE5, 0x1E, 0xC1, 0x12, 0xDE, 0x5C, 0x38, 0x4D, 0xF7, 0xBA, 0x0B, 0x8D, 0x57, 0x8A, 0x4C, 0x70, 0x2B, 0x6B, 0xF1, 0x1D, 0x5F, 0xAC},
			},
		},
	}

	expected, _ := hex.DecodeString("01000000010000000000000000000000000000000000000000000000000000000000000000FFFFFFFF4D04FFFF001D0104455468652054696D65732030332F4A616E2F32303039204368616E63656C6C6F72206F6E206272696E6B206F66207365636F6E64206261696C6F757420666F722062616E6B73FFFFFFFF0100F2052A01000000434104678AFDB0FE5548271967F1A67130B7105CD6A828E03909A67962E0EA1F61DEB649F6BC3F4CEF38C4F35504E51EC112DE5C384DF7BA0B8D578A4C702B6BF11D5FAC00000000")
	expectedHash, _ := hex.DecodeString("3ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a")

	encoded, err := tx.Bytes()
	if err != nil {
		t.Errorf(".Bytes() failed %v", err)
	}

	if !bytes.Equal(encoded, expected) {
		t.Errorf("Incorrect .Bytes() expected %s , encoded %s", hex.EncodeToString(expected), hex.EncodeToString(encoded))
	}

	var setTx Tx
	err = setTx.Set(encoded)
	if err != nil {
		t.Errorf(".Set() failed %v", err)
	}

	if setTx.Version != tx.Version {
		t.Errorf(".Set() incorrect version %v", setTx.Version)
	}

	if len(setTx.Inputs) != 1 {
		t.Errorf(".Set() incorrect input count %v", len(setTx.Inputs))
	}

	if !bytes.Equal(setTx.Inputs[0].Previous.Hash[:], tx.Inputs[0].Previous.Hash[:]) {
		t.Errorf(".Set() Previous Hash incorrect %v", hex.EncodeToString(setTx.Inputs[0].Previous.Hash[:]))
	}

	if setTx.Inputs[0].Previous.Index != tx.Inputs[0].Previous.Index {
		t.Errorf(".Set() incorrect prev index %v", setTx.Inputs[0].Previous.Index)
	}

	if !bytes.Equal(setTx.Inputs[0].Script, tx.Inputs[0].Script) {
		t.Errorf(".Set() in script incorrect %v", hex.EncodeToString(setTx.Inputs[0].Script))
	}

	if setTx.Inputs[0].Sequence != tx.Inputs[0].Sequence {
		t.Errorf(".Set() incorrect sequence %v", setTx.Inputs[0].Sequence)
	}

	if len(setTx.Outputs) != 1 {
		t.Errorf(".Set() incorrect output count %v", len(setTx.Inputs))
	}

	if setTx.Outputs[0].Value != tx.Outputs[0].Value {
		t.Errorf(".Set() incorrect value %v", setTx.Outputs[0].Value)
	}

	if !bytes.Equal(setTx.Outputs[0].Script, tx.Outputs[0].Script) {
		t.Errorf(".Set() out script incorrect %v", hex.EncodeToString(setTx.Inputs[0].Script))
	}

	hash, err := tx.Hash()
	if err != nil {
		t.Errorf(".Hash() failed %v", err)
	}

	if !bytes.Equal(hash[:], expectedHash[:]) {
		t.Errorf(".Hash() incorrect %v", hex.EncodeToString(hash[:]))
	}
}

func BenchmarkBytes(b *testing.B) {

	tx := Tx{
		Version: 1,
		Inputs: []TxIn{
			TxIn{
				Previous: OutPoint{
					Hash:  [32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
					Index: 0xFFFFFFFF,
				},
				Script:   []byte{0x04, 0xFF, 0xFF, 0x00, 0x1D, 0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6D, 0x65, 0x73, 0x20, 0x30, 0x33, 0x2F, 0x4A, 0x61, 0x6E, 0x2F, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, 0x61, 0x6E, 0x63, 0x65, 0x6C, 0x6C, 0x6F, 0x72, 0x20, 0x6F, 0x6E, 0x20, 0x62, 0x72, 0x69, 0x6E, 0x6B, 0x20, 0x6F, 0x66, 0x20, 0x73, 0x65, 0x63, 0x6F, 0x6E, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6C, 0x6F, 0x75, 0x74, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x62, 0x61, 0x6E, 0x6B, 0x73},
				Sequence: 0xFFFFFFFF,
			},
		},
		Outputs: []TxOut{
			TxOut{
				Value:  5000000000,
				Script: []byte{0x41, 0x04, 0x67, 0x8A, 0xFD, 0xB0, 0xFE, 0x55, 0x48, 0x27, 0x19, 0x67, 0xF1, 0xA6, 0x71, 0x30, 0xB7, 0x10, 0x5C, 0xD6, 0xA8, 0x28, 0xE0, 0x39, 0x09, 0xA6, 0x79, 0x62, 0xE0, 0xEA, 0x1F, 0x61, 0xDE, 0xB6, 0x49, 0xF6, 0xBC, 0x3F, 0x4C, 0xEF, 0x38, 0xC4, 0xF3, 0x55, 0x04, 0xE5, 0x1E, 0xC1, 0x12, 0xDE, 0x5C, 0x38, 0x4D, 0xF7, 0xBA, 0x0B, 0x8D, 0x57, 0x8A, 0x4C, 0x70, 0x2B, 0x6B, 0xF1, 0x1D, 0x5F, 0xAC},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		tx.Bytes()
	}
}

func BenchmarkWriteTo(b *testing.B) {

	tx := Tx{
		Version: 1,
		Inputs: []TxIn{
			TxIn{
				Previous: OutPoint{
					Hash:  [32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
					Index: 0xFFFFFFFF,
				},
				Script:   []byte{0x04, 0xFF, 0xFF, 0x00, 0x1D, 0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6D, 0x65, 0x73, 0x20, 0x30, 0x33, 0x2F, 0x4A, 0x61, 0x6E, 0x2F, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, 0x61, 0x6E, 0x63, 0x65, 0x6C, 0x6C, 0x6F, 0x72, 0x20, 0x6F, 0x6E, 0x20, 0x62, 0x72, 0x69, 0x6E, 0x6B, 0x20, 0x6F, 0x66, 0x20, 0x73, 0x65, 0x63, 0x6F, 0x6E, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6C, 0x6F, 0x75, 0x74, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x62, 0x61, 0x6E, 0x6B, 0x73},
				Sequence: 0xFFFFFFFF,
			},
		},
		Outputs: []TxOut{
			TxOut{
				Value:  5000000000,
				Script: []byte{0x41, 0x04, 0x67, 0x8A, 0xFD, 0xB0, 0xFE, 0x55, 0x48, 0x27, 0x19, 0x67, 0xF1, 0xA6, 0x71, 0x30, 0xB7, 0x10, 0x5C, 0xD6, 0xA8, 0x28, 0xE0, 0x39, 0x09, 0xA6, 0x79, 0x62, 0xE0, 0xEA, 0x1F, 0x61, 0xDE, 0xB6, 0x49, 0xF6, 0xBC, 0x3F, 0x4C, 0xEF, 0x38, 0xC4, 0xF3, 0x55, 0x04, 0xE5, 0x1E, 0xC1, 0x12, 0xDE, 0x5C, 0x38, 0x4D, 0xF7, 0xBA, 0x0B, 0x8D, 0x57, 0x8A, 0x4C, 0x70, 0x2B, 0x6B, 0xF1, 0x1D, 0x5F, 0xAC},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		tx.WriteTo(ioutil.Discard)
	}
}

func BenchmarkHash(b *testing.B) {

	tx := Tx{
		Version: 1,
		Inputs: []TxIn{
			TxIn{
				Previous: OutPoint{
					Hash:  [32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
					Index: 0xFFFFFFFF,
				},
				Script:   []byte{0x04, 0xFF, 0xFF, 0x00, 0x1D, 0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6D, 0x65, 0x73, 0x20, 0x30, 0x33, 0x2F, 0x4A, 0x61, 0x6E, 0x2F, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, 0x61, 0x6E, 0x63, 0x65, 0x6C, 0x6C, 0x6F, 0x72, 0x20, 0x6F, 0x6E, 0x20, 0x62, 0x72, 0x69, 0x6E, 0x6B, 0x20, 0x6F, 0x66, 0x20, 0x73, 0x65, 0x63, 0x6F, 0x6E, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6C, 0x6F, 0x75, 0x74, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x62, 0x61, 0x6E, 0x6B, 0x73},
				Sequence: 0xFFFFFFFF,
			},
		},
		Outputs: []TxOut{
			TxOut{
				Value:  5000000000,
				Script: []byte{0x41, 0x04, 0x67, 0x8A, 0xFD, 0xB0, 0xFE, 0x55, 0x48, 0x27, 0x19, 0x67, 0xF1, 0xA6, 0x71, 0x30, 0xB7, 0x10, 0x5C, 0xD6, 0xA8, 0x28, 0xE0, 0x39, 0x09, 0xA6, 0x79, 0x62, 0xE0, 0xEA, 0x1F, 0x61, 0xDE, 0xB6, 0x49, 0xF6, 0xBC, 0x3F, 0x4C, 0xEF, 0x38, 0xC4, 0xF3, 0x55, 0x04, 0xE5, 0x1E, 0xC1, 0x12, 0xDE, 0x5C, 0x38, 0x4D, 0xF7, 0xBA, 0x0B, 0x8D, 0x57, 0x8A, 0x4C, 0x70, 0x2B, 0x6B, 0xF1, 0x1D, 0x5F, 0xAC},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		tx.Hash()
	}
}

func BenchmarkSet(b *testing.B) {

	var tx Tx
	data, _ := hex.DecodeString("01000000010000000000000000000000000000000000000000000000000000000000000000FFFFFFFF4D04FFFF001D0104455468652054696D65732030332F4A616E2F32303039204368616E63656C6C6F72206F6E206272696E6B206F66207365636F6E64206261696C6F757420666F722062616E6B73FFFFFFFF0100F2052A01000000434104678AFDB0FE5548271967F1A67130B7105CD6A828E03909A67962E0EA1F61DEB649F6BC3F4CEF38C4F35504E51EC112DE5C384DF7BA0B8D578A4C702B6BF11D5FAC00000000")

	for i := 0; i < b.N; i++ {
		tx.Set(data)
	}
}
