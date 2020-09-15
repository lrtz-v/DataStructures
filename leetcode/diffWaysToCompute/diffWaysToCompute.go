package diffWaysToCompute

import "strconv"

func diffWaysToCompute(input string) []int {

	dp := make(map[string][]int)

	return calculate(dp, input)
}

func calculate(dp map[string][]int, input string) []int {
	if len(input) == 0 {
		return []int{}
	}

	if val, ok := dp[input]; ok {
		return val
	}

	res := []int{}
	intVal, err := strconv.Atoi(input)
	if err == nil {
		res = append(res, intVal)
	} else {
		for i := 0; i < len(input); i++ {
			ch := input[i]
			if ch == '+' || ch == '-' || ch == '*' {
				leftRes := calculate(dp, input[0:i])
				rightRes := calculate(dp, input[i+1:])
	
				for _, l := range leftRes {
					for _, r := range rightRes {
						var tmp int
						if ch == '+' {
							tmp = l + r
						} else if ch == '-' {
							tmp = l - r
						} else if ch == '*' {
							tmp = l * r
						}
						res = append(res, tmp)
					}
				}
			}
		}
	}

	dp[input] = res
	return res
}

