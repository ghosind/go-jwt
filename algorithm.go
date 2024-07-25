package jwt

const (
	AlgNone  = "none"
	AlgHS256 = "HS256"
	AlgHS384 = "HS384"
	AlgHS512 = "HS512"
)

var supportedAlgorithms []string = []string{
	AlgNone,
	AlgHS256,
	AlgHS384,
	AlgHS512,
}

func isAlgorithmSupported(alg string) error {
	for _, a := range supportedAlgorithms {
		if a == alg {
			return nil
		}
	}
	return ErrUnsupportedAlgorithm
}
