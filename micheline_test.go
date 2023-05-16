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

func TestPackAddress(t *testing.T) {
	a := PackedPrim("tz1hQbuRax3op9knY3YDxqNnqxzcmoxmv1qa")
	b, _ := a.MarshalJSON()

	assert.EqualValues(t, "{\"bytes\":\"0000eec8aafb412df59734945ca265c5d16273b0967a\"}", string(b))
}

func TestPackSignature(t *testing.T) {
	a := PackedPrim("edsigu64opjAiYcyBEZ7Jr4djGXTYmfJAKoCEr8SuriTooymw4fYiiCQtkgiLRp26gHTePZ2tDWfAXpPKf9WAPjw34PCz5PHax8")
	b, _ := a.MarshalJSON()

	assert.EqualValues(t, "{\"bytes\":\"f693efcf6a2620b2395f751bb89a532a31a8f013743ca0dfa151b60a3282e43a472ba1ad16f14e0033bad2f365fa001c7aa4eacfd8b22597d1d401736774ea0a\"}", string(b))
}

func TestPackPublicKey(t *testing.T) {
	a := PackedPrim("edpkv8icP7cT2hTLUdiCzrdmtoHNKHpuSkJojakc7KeCxczcKfpPSj")
	b, _ := a.MarshalJSON()

	assert.EqualValues(t, "{\"bytes\":\"00c46564ac5e8c01cf48f0f769b59f21105512c14d417aaf192011ce376d6299f6\"}", string(b))
}
