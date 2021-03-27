package randutil_kai

import (
	crand "crypto/rand"
	"encoding/binary"
	"io"
	"math/big"
)

// GenerateReaderCryptoRandomString generates a reader-based random string for cryptographic usage.
func GenerateReaderCryptoRandomString(n int, runes string, reader io.Reader) (string, error) {
	letters := []rune(runes)
	b := make([]rune, n)
	for i := range b {
		v, err := crand.Int(reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		b[i] = letters[v.Int64()]
	}
	return string(b), nil
}

// ReaderCryptoUint64 returns reader-based cryptographic random uint64.
func ReaderCryptoUint64(reader io.Reader) (uint64, error) {
	var v uint64
	if err := binary.Read(reader, binary.LittleEndian, &v); err != nil {
		return 0, err
	}
	return v, nil
}
