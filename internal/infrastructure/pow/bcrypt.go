package pow

import (
	"crypto/rand"
	"math/big"

	"github.com/e-zhydzetski/faraway-tt/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

func NewBCryptCheck(complexity uint64) (domain.POWCheck, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(complexity)))
	if err != nil {
		return BCryptCheck{}, err
	}
	digest, err := bcrypt.GenerateFromPassword(r.Bytes(), bcrypt.DefaultCost)
	if err != nil {
		return BCryptCheck{}, err
	}
	return BCryptCheck{
		answer: r.Uint64(),
		digest: digest,
	}, nil
}

type BCryptCheck struct {
	answer uint64
	digest []byte
}

func (b BCryptCheck) Input() []byte {
	return b.digest
}

func (b BCryptCheck) Check(answer uint64) bool {
	return b.answer == answer
}

func BCryptSolve(input []byte) uint64 {
	var x uint64
	for {
		err := bcrypt.CompareHashAndPassword(input, big.NewInt(int64(x)).Bytes())
		if err == nil {
			return x
		}
		x++
	}
}
