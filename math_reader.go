package randutil_kai

import (
	"fmt"
	"io"
	mrand "math/rand" // used for non-crypto unique ID and random port selection
	"sync"
)

// MathRandomGenerator is a random generator for non-crypto usage.
type ReaderMathRandomGenerator interface {
	// Intn returns random integer within [0:n).
	Intn(n int) int

	// Uint32 returns random 32-bit unsigned integer.
	Uint32() uint32

	// Uint64 returns random 64-bit unsigned integer.
	Uint64() uint64

	// GenerateString returns ranom string using given set of runes.
	// It can be used for generating unique ID to avoid name collision.
	//
	// Caution: DO NOT use this for cryptographic usage.
	GenerateString(n int, runes string) string
}

type readerMathRandomGenerator struct {
	r  *mrand.Rand
	mu sync.Mutex
}

// NewMathRandomGenerator creates new mathmatical random generator.
// Random generator is seeded by crypto random.
func NewReaderMathRandomGenerator(reader io.Reader) MathRandomGenerator {
	seed, err := ReaderCryptoUint64(reader)
	if err != nil {
		// crypto/rand is unavailable. Panik
		panic("Error in NewReaderMathRandomGenerator: " + fmt.Sprint(err))
	}

	return &readerMathRandomGenerator{r: mrand.New(mrand.NewSource(int64(seed)))}
}

func (g *readerMathRandomGenerator) Intn(n int) int {
	g.mu.Lock()
	v := g.r.Intn(n)
	g.mu.Unlock()
	return v
}

func (g *readerMathRandomGenerator) Uint32() uint32 {
	g.mu.Lock()
	v := g.r.Uint32()
	g.mu.Unlock()
	return v
}

func (g *readerMathRandomGenerator) Uint64() uint64 {
	g.mu.Lock()
	v := g.r.Uint64()
	g.mu.Unlock()
	return v
}

func (g *readerMathRandomGenerator) GenerateString(n int, runes string) string {
	letters := []rune(runes)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[g.Intn(len(letters))]
	}
	return string(b)
}
