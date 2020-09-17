package restoreIpAddresses

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {

	res := []string{}

	l := len(s)
	if l < 4 || l > 12 {
		return res
	}

	track := make([]string, 0)

	var backtrack func(s string, track []string, step int) 

	backtrack = func(s string, track []string, step int) {
		if step == 5 {
			if len(s) == 0 {
				str := strings.Join(track, ".")
				res = append(res, str)
			}
			return
		}

		for i := 1; i <= 3; i++ {
			if i <= len(s) {
				v, _ := strconv.Atoi(s[:i])
				if v <= 255 {
					track = append(track, s[:i])
					str := s[i:]
					backtrack(str, track, step+1)
					track = track[:len(track)-1]
				}
				if v == 0 {
					break
				}
			}
		}
	}

	backtrack(s, track, 1)

	return res
}

/*

result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
		撤销选择

*/