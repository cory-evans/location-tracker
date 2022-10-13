package devicelocation

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"io"
	"strings"
)

func DecodeRequestBody(key []byte, body io.ReadCloser) (*Location, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	data := buf.Bytes()

	out := make([]byte, 0)

	s := string(data)
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		var dest = make([]byte, 16)
		base64.StdEncoding.Decode(dest, []byte(line))
		d := decryptAes128Ecb(dest, key)
		out = append(out, d...)
	}

	l := &Location{}
	json.Unmarshal(bytes.Trim(out, "\x00"), l)

	return l, nil
}

func decryptAes128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}
