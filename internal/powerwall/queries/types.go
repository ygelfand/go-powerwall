package queries

import (
	b64 "encoding/base64"
)

type SignedQuery struct {
	Name          string
	Query         string
	QueryEncoded  string
	Signature     string
	DefaultParams *string
	SigKey        int32
}

func (dq *SignedQuery) Key() int32 {
	if dq.SigKey != 0 {
		return dq.SigKey
	}
	return 1
}

func (dq *SignedQuery) Sig() []byte {
	bytes, _ := b64.StdEncoding.DecodeString(dq.Signature)
	return bytes
}

func (dq *SignedQuery) GetQuery() string {
	if dq.Query != "" {
		return dq.Query
	}
	bytes, _ := b64.StdEncoding.DecodeString(dq.QueryEncoded)
	return string(bytes)
}

func PointerTo[T ~string](s T) *T {
	return &s
}
