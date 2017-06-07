package serv

func Fibon(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibon(n-1) + Fibon(n-2)
	}
}
