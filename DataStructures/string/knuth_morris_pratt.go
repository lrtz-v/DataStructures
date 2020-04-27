package string

func generateNext(p string) []int {
	next := make([]int, len(p))
    pLen := len(p)
    next[0] = -1
    k := -1
    j := 0
    for j < pLen - 1 {
        //p[k]表示前缀，p[j]表示后缀
        if k == -1 || p[j] == p[k] {
            k++
            j++
			if p[j] != p[k] {
            	next[j] = k
			} else {
				//因为不能出现p[j] = p[ next[j ]]，所以当出现时需要继续递归，k = next[k] = next[next[k]]
                next[j] = next[k]
			}
        } else {
            k = next[k]
        }
    }
	return next
}

func kmp(s, p string) int {

	i := 0
	j := 0
	sLen := len(s)
	pLen := len(p)
	next := generateNext(p)

	for i < sLen && j < pLen {
		// ①如果j = -1，或者当前字符匹配成功（即S[i] == P[j]），都令i++，j++
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			// ②如果j != -1，且当前字符匹配失败（即S[i] != P[j]），则令 i 不变，j = next[j]
            // next[j]即为j所对应的next值
            j = next[j]
		}
	}

	if j == pLen {
		return i - j
	} else {
		return -1
	}
}
