package tezos_micheline

import (
	"math/big"
	"reflect"

	"blockwatch.cc/tzgo/micheline"
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
