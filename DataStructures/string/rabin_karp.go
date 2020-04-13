package string

import "crypto/md5"

/**
 *	Rabin-Karp
 */

func hash(str string) string {
	h := md5.Sum([]byte(str))
	return string(h[:])
}

func rabinKarp(A, B string) (index int) {
	bHash := hash(B)

	hashMap := make(map[string]int)
	for i := 0; i < len(A)-len(B); i++ {
		hashMap[hash(A[i:i+len(B)])] = i
	}

	if index, ok := hashMap[bHash]; ok {
		return index
	}
	return -1
}
