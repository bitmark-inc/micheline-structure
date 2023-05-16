package tezos_micheline

import (
	"math/big"
	"reflect"

	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
)

func Prim(data interface{}) micheline.Prim {
	v := reflect.ValueOf(data)

	switch v.Kind() {
	case reflect.String:
		return micheline.NewString(v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return micheline.NewNat(big.NewInt(v.Int()))
	case reflect.Slice:
		if reflect.TypeOf(data).String() == "[]uint8" {
			return micheline.NewBytes(v.Bytes())
		} else {
			var p []micheline.Prim

			for i := 0; i < v.Len(); i++ {
				value := v.Index(i).Interface()
				prim := Prim(value)

				p = append(p, prim)
			}

			return micheline.NewSeq(p...)
		}
	case reflect.Struct:
		t := v.Type()

		var p *micheline.Prim

		for i := t.NumField() - 1; i >= 0; i-- {
			field := t.Field(i)
			value := v.FieldByName(field.Name).Interface()
			if p == nil {
				p1 := Prim(value)
				p = &p1
			} else {
				p2 := micheline.NewPair(
					Prim(value),
					*p,
				)
				p = &p2
			}
		}

		if p == nil {
			return micheline.EmptyPrim
		}
		return *p
	default:
		return micheline.EmptyPrim
	}
}

func PackedPrim(data interface{}) micheline.Prim {
	return PackAll(Prim(data))
}

func PackAll(p micheline.Prim) micheline.Prim {
	if p.Type == micheline.PrimString {
		return Pack(p)
	}
	if p.LooksLikeCode() {
		return p
	}

	pp := p
	pp.Args = make([]micheline.Prim, len(p.Args))
	for i, v := range p.Args {
		pp.Args[i] = PackAll(v)
	}
	return pp
}

func Pack(p micheline.Prim) micheline.Prim {
	if p.Type == micheline.PrimString {
		a, err := tezos.ParseAddress(p.String)
		if err == nil {
			return micheline.NewBytes(a.EncodePadded())
		}

		k, err := tezos.ParseKey(p.String)
		if err == nil {
			return micheline.NewBytes(k.Bytes())
		}

		s, err := tezos.ParseSignature(p.String)
		if err == nil {
			if s.Type.Tag() == 255 {
				return micheline.NewBytes(s.Bytes())
			} else {
				//remove tag byte prefix
				return micheline.NewBytes(s.Bytes()[1:])
			}
		}

		id, err := tezos.ParseChainIdHash(p.String)
		if err == nil {
			return micheline.NewBytes(id.Bytes())
		}
	}

	return p
}
