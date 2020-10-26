package number_theory

const N = 1e7 + 10

var (
	Prime [N]int
	vis   [N]bool
)

func EratoSieve(n int) {
	m := 0
	for i := 2; i <= n; i++ {
		if !vis[i] {
			m++
			Prime[m] = i
			for j := i; j <= n/i; j++ {
				vis[i*j] = true
			}
		}
	}
	Prime[0] = m
}

func EulerSeive(n int) {
	m := 0
	for i := 2; i <= n; i++ {
		if !vis[i] {
			m++
			Prime[m] = i
		}
		for j := 1; j <= m && Prime[j] <= n/i; j++ {
			vis[i*Prime[j]] = true
			if i%Prime[j] == 0 {
				break
			}
		}
	}
	Prime[0] = m
}
