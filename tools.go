package toolkit

import "crypto/rand"

type Tools struct{}

const RandomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUV123456789_#"

func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(RandomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
