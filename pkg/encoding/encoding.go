package encoding

import "encoding/base64"

type Encoder interface {
	Encode(string) string
}

type Base64Encoder struct {
	encoding *base64.Encoding
}

func NewEncoder() *Base64Encoder {
	return &Base64Encoder{
		encoding: base64.URLEncoding.WithPadding(base64.NoPadding),
	}
}

func (e *Base64Encoder) Encode(link string) string {
	data := []byte(link)
	return e.encoding.EncodeToString(data)
}

type Decoder interface {
	Decode(string) string
}

type Base64Decoder struct {
	encoding *base64.Encoding
}

func NewDecoder() *Base64Decoder {
	return &Base64Decoder{
		encoding: base64.URLEncoding.WithPadding(base64.NoPadding),
	}
}

func (d *Base64Decoder) Decode(b64 string) (string, error) {
	bytes, err := d.encoding.DecodeString(b64)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
