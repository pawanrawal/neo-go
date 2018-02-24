package wallet

const (
	n = 16384
	r = 8
	p = 8
)

type ScryptParameters struct {
	n int
	r int
	p int
}

func DefaultScryptParameters() ScryptParameters {
	return ScryptParameters{
		n: n,
		r: r,
		p: p,
	}
}
