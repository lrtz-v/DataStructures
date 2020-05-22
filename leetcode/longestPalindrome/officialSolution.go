package longestPalindrome

// dp
func officialSolutionDp(s string) string {
    n := len(s)
    ans := ""
	dp := make([][]int, n)
	// dp init
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
	}

    for l := 0; l < n; l++ {
        for i := 0; i + l < n; i++ {
            j := i + l
            if l == 0 {
                dp[i][j] = 1
            } else if l == 1 {
                if s[i] == s[j] {
                    dp[i][j] = 1
                }
            } else {
                if s[i] == s[j] {
                    dp[i][j] = dp[i+1][j-1]
                }
            }
            if dp[i][j] > 0 && l + 1 > len(ans) {
                ans = s[i:i+l+1]
            }
        }
    }
    return ans
}

func longestPalindrome2(s string) string {
    if s == "" {
        return ""
    }
    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        left1, right1 := expandAroundCenter(s, i, i)
        left2, right2 := expandAroundCenter(s, i, i + 1)
        if right1 - left1 > end - start {
            start, end = left1, right1
        }
        if right2 - left2 > end - start {
            start, end = left2, right2
        }
    }
    return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
    for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1 , right+1 {

	}
    return left + 1, right - 1
}

func longestPalindrome3(s string) string {
    start, end := 0, -1
    t := "#"
    for i := 0; i < len(s); i++ {
        t += string(s[i]) + "#"
    }
    t += "#"
    s = t
    armLen := []int{}
    right, j := -1, -1
    for i := 0; i < len(s); i++ {
        var curArmLen int
        if right >= i {
            iSym := j * 2 - i
            minArmLen := min(armLen[iSym], right-i)
            curArmLen = expand(s, i-minArmLen, i+minArmLen)
        } else {
            curArmLen = expand(s, i, i)
        }
        armLen = append(armLen, curArmLen)
        if i + curArmLen > right {
            j = i
            right = i + curArmLen
        }
        if curArmLen * 2 + 1 > end - start {
            start = i - curArmLen
            end = i + curArmLen
        }
    }
    ans := ""
    for i := start; i <= end; i++ {
        if s[i] != '#' {
            ans += string(s[i])
        }
    }
    return ans
}

func expand(s string, left, right int) int {
    for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 { }
    return (right - left - 2) / 2
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
