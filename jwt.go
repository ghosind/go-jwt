package jwt

type JWT struct {
	alg string

	err error
}

func (t *JWT) Algorithm() string {
	return t.alg
}

func (t *JWT) SetAlgorithm(alg string) *JWT {
	t.alg = alg

	return t.addError(isAlgorithmSupported(alg))
}
