package queries

import (
	b64 "encoding/base64"
)

type SignedQuery struct {
	Name          string
	Query         string
	Signature     string
	DefaultParams *string
}

func (dq *SignedQuery) Sig() []byte {
	bytes, _ := b64.StdEncoding.DecodeString(dq.Signature)
	return bytes
}

func PointerTo[T ~string](s T) *T {
	return &s
}
