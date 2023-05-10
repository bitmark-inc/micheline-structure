package tezos_micheline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertString(t *testing.T) {
	a := Prim("test")
	b, _ := a.MarshalJSON()

	assert.EqualValues(t, "{\"string\":\"test\"}", string(b))
}

func TestConvertInt(t *testing.T) {
	a := Prim(1)
	b, _ := a.MarshalJSON()

	assert.EqualValues(t, "{\"int\":\"1\"}", string(b))
}

func TestConvertByte(t *testing.T) {
	a := Prim([]byte("ipfs://Qmcxxxxxxxcid"))
	b, _ := a.MarshalJSON()

	assert.EqualValues(t, "{\"bytes\":\"697066733a2f2f516d6378787878787878636964\"}", string(b))
}

func TestConvertSequence(t *testing.T) {
	a := Prim([]interface{}{1, "test", []byte("ipfs://Qmcxxxxxxxcid")})
	b, _ := a.MarshalJSON()

	assert.EqualValues(t, "[{\"int\":\"1\"},{\"string\":\"test\"},{\"bytes\":\"697066733a2f2f516d6378787878787878636964\"}]", string(b))
}

func TestConvertStruct(t *testing.T) {
	a := Prim(struct {
		Name     string
		Counter  int64
		TokenURI []byte
	}{
		Name:     "hello",
		Counter:  1,
		TokenURI: []byte("ipfs://Qmcxxxxxxxcid"),
	})
	b, _ := a.MarshalJSON()

	assert.EqualValues(t, "{\"prim\":\"Pair\",\"args\":[{\"string\":\"hello\"},{\"prim\":\"Pair\",\"args\":[{\"int\":\"1\"},{\"bytes\":\"697066733a2f2f516d6378787878787878636964\"}]}]}", string(b))
}
