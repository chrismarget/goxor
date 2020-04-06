package decrypt

import (
	"bytes"
	"testing"
)

func Test_same_size(t *testing.T) {
	key := []byte{1,1,1,1}
	orig := []byte("bcde")
	expected := []byte("cbed")
	in := bytes.NewReader(orig)
	var out bytes.Buffer
	err := Decrypt(key, in, &out)
	if err != nil {
		t.Fatal(err)
	}
	result := out.Bytes()
	if len(orig) != len(result) {
		t.Fatalf("length of result %d doesn't match length of original data %d", len(result), len(orig))
	}
	for i := 0; i < len(orig); i++ {
		if result[i] != expected[i] {
			t.Fatal("result doesn't match expected data")
		}
	}
}

func Test_multiple_of_key(t *testing.T) {
	key := []byte{1,1,1,1}
	orig := []byte("bcdebcde")
	expected := []byte("cbedcbed")
	in := bytes.NewReader(orig)
	var out bytes.Buffer
	err := Decrypt(key, in, &out)
	if err != nil {
		t.Fatal(err)
	}
	result := out.Bytes()
	if len(orig) != len(result) {
		t.Fatalf("length of result %d doesn't match length of original data %d", len(result), len(orig))
	}
	for i := 0; i < len(orig); i++ {
		if result[i] != expected[i] {
			t.Fatal("result doesn't match expected data")
		}
	}
}

func Test_empty(t *testing.T) {
	key := []byte{1,1,1,1}
	orig := []byte("")
	expected := []byte("")
	in := bytes.NewReader(orig)
	var out bytes.Buffer
	err := Decrypt(key, in, &out)
	if err != nil {
		t.Fatal(err)
	}
	result := out.Bytes()
	if len(orig) != len(result) {
		t.Fatalf("length of result %d doesn't match length of original data %d", len(result), len(orig))
	}
	for i := 0; i < len(orig); i++ {
		if result[i] != expected[i] {
			t.Fatal("result doesn't match expected data")
		}
	}
}

func Test_longer_data(t *testing.T) {
	key := []byte{1,1,1}
	orig := []byte("bcdef")
	expected := []byte("cbedg")
	in := bytes.NewReader(orig)
	var out bytes.Buffer
	err := Decrypt(key, in, &out)
	if err != nil {
		t.Fatal(err)
	}
	result := out.Bytes()
	if len(orig) != len(result) {
		t.Fatalf("length of result %d doesn't match length of original data %d", len(result), len(orig))
	}
	for i := 0; i < len(orig); i++ {
		if result[i] != expected[i] {
			t.Fatal("result doesn't match expected data")
		}
	}
}

func Test_shorter_data(t *testing.T) {
	key := []byte{1,1,1,1,1,1}
	orig := []byte("bcdef")
	expected := []byte("cbedg")
	in := bytes.NewReader(orig)
	var out bytes.Buffer
	err := Decrypt(key, in, &out)
	if err != nil {
		t.Fatal(err)
	}
	result := out.Bytes()
	if len(orig) != len(result) {
		t.Fatalf("length of result %d doesn't match length of original data %d", len(result), len(orig))
	}
	for i := 0; i < len(orig); i++ {
		if result[i] != expected[i] {
			t.Fatal("result doesn't match expected data")
		}
	}
}
