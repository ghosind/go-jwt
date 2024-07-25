package jwt

import "errors"

var (
	ErrUnsupportedAlgorithm error = errors.New("unsupported algorithm")
	ErrInvalidIssuer        error = errors.New("invalid issuer")
	ErrInvalidSubject       error = errors.New("invalid subject")
	ErrInvalidAudience      error = errors.New("invalid audience")
	ErrInavlidIssuedAt      error = errors.New("invalid issued time")
	ErrInvalidJWTID         error = errors.New("invalid JWT ID")
	ErrExpired              error = errors.New("token was expired")
	ErrBefore               error = errors.New("token is before the time to be accepted")
	ErrInvalidType          error = errors.New("invalid JWT media type")
)

func (t *JWT) Error() error {
	return t.err
}

func (t *JWT) addError(err error) *JWT {
	if err != nil {
		t.err = errors.Join(t.err, err)
	}

	return t
}
