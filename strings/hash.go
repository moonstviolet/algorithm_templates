package strings

const (
	N int  = 1e5 + 10
	P uint = 19260817
)

var pow [N]uint

type Hash struct {
	hash [N]uint
}

func init() {
	pow[0] = 1
	for i := 1; i < N; i++ {
		pow[i] = pow[i-1] * P
	}
}

func (h *Hash) init(s string) {
	h.hash[0] = uint(s[0])
	for i := 1; i < len(s); i++ {
		h.hash[i] = h.hash[i-1]*P + uint(s[i])
	}
}

func (h *Hash) get(l int, r int) uint {
	if l == 0 {
		return h.hash[r]
	}
	return h.hash[r] - h.hash[l-1]*pow[r-l+1]
}
