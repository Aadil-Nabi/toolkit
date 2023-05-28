package toolkit

import "crypto/rand"

// Tools type helps to define and initiate the RandomString
type Tools struct{}

const RandomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUV123456789_#"

// RandomString returns a random string generated from the runes
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(RandomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
