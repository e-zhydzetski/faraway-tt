package bcrypt

import (
	"crypto/rand"
	"math/big"

	"github.com/e-zhydzetski/faraway-tt/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

const cost = 10

func NewPOWCheck(complexity uint64) (domain.POWCheck, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(complexity)))
	if err != nil {
		return POWCheck{}, err
	}
	digest, err := bcrypt.GenerateFromPassword(r.Bytes(), cost)
	if err != nil {
		return POWCheck{}, err
	}
	return POWCheck{
		answer: r.Uint64(),
		digest: digest,
	}, nil
}

type POWCheck struct {
	answer uint64
	digest []byte
}

func (p POWCheck) Input() []byte {
	return p.digest
}

func (p POWCheck) Check(answer uint64) bool {
	return p.answer == answer
}

func Solve(input []byte) uint64 {
	var x uint64
	for {
		err := bcrypt.CompareHashAndPassword(input, big.NewInt(int64(x)).Bytes())
		if err == nil {
			return x
		}
		x++
	}
}
