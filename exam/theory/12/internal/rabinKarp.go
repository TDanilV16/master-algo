package internal

type Hasher struct {
	powsOf32 []int
}

//const base = 31
//
//func RabinKarp(s, w string) []int {
//	return rabinKarp()
//}
//
//func rabinKarp(s, w string) []int {
//	n, m := len(s), len(w)
//
//	powsOf32 := pows(m)
//}
//
//func pows(m int) []int {
//	powsOf31 := make([]int, m+1)
//	powsOf31[0] = 1
//	for i := 1; i <= m; i++ {
//		powsOf31[i] = powsOf31[i-1] * base
//	}
//
//	return powsOf31
//}
//
//func hash(s string) int {
//	hash := 0
//
//	for i, b := range s {
//		hash += (b)
//	}
//}
