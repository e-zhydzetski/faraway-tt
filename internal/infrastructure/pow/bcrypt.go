package pow

import (
	"context"
	"crypto/rand"
	"math/big"
	"time"

	"github.com/e-zhydzetski/faraway-tt/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

func NewBCryptCheck(difficulty uint64) (domain.POWCheck, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(difficulty)))
	if err != nil {
		return BCryptCheck{}, err
	}
	digest, err := bcrypt.GenerateFromPassword(r.Bytes(), bcrypt.DefaultCost)
	if err != nil {
		return BCryptCheck{}, err
	}
	return BCryptCheck{
		answer:  r.Uint64(),
		digest:  digest,
		timeout: time.Second + time.Duration((difficulty)/10)*time.Second,
	}, nil
}

type BCryptCheck struct {
	answer  uint64
	digest  []byte
	timeout time.Duration
}

func (b BCryptCheck) Challenge() []byte {
	return b.digest
}

func (b BCryptCheck) ReasonableTimeout() time.Duration {
	return b.timeout
}

func (b BCryptCheck) Verify(proof uint64) bool {
	return b.answer == proof
}

func BCryptProve(ctx context.Context, input []byte) (uint64, error) {
	var x uint64
	for {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
		}
		// TODO maybe calculate several hashes between context checks
		err := bcrypt.CompareHashAndPassword(input, big.NewInt(int64(x)).Bytes())
		if err == nil {
			return x, nil
		}
		x++
	}
}
