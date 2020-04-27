package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// PrivateKey generates private key.
func PrivateKey(p *big.Int) *big.Int {
	key := new(big.Int)
	limit := new(big.Int).Sub(p, big.NewInt(2))
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return key.Rand(seed, limit).Add(key, big.NewInt(2))
}

// PublicKey generates public key.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	newG := big.NewInt(g)
	return newG.Exp(newG, private, p)
}

// NewPair generates a pair public key and private key.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey calculates secret key.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
