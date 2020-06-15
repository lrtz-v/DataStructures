package longestCommonPrefix

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1:
输入: ["flower","flow","flight"]
输出: "fl"

示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。

说明:
所有输入只包含小写字母 a-z 。
*/


func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	str := strs[0]
	length := len(str)
	index := 0
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < length {
			str = strs[i]
			length = len(strs[i])
			index = i
		}
	}
	strs[0], strs[index] = strs[index], strs[0]
	i := 0
	for ; i < len(str); i++ {
		if str[i] != strs[1][i] {
			break
		}
	}
	str = str[0:i]
	if len(str) == 0 {
		return ""
	}

	index = len(str) - 1
	for i := 2; i < len(strs); i++ {
		if index < 0 {
			break
		} else if str[index] == strs[i][index] {
			continue
		}
		for ; index >= 0; index-- {
			if str[index] == strs[i][index] {
				break
			}
		}
	}

	return str[0: index+1]
}

// 纵向扫描，扫描每一列
func longestCommonPrefix2(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    for i := 0; i < len(strs[0]); i++ {
        for j := 1; j < len(strs); j++ {
            if i == len(strs[j]) || strs[j][i] != strs[0][i] {
                return strs[0][:i]
            }
        }
    }
    return strs[0]
}

// 分治 LCP(S1…Sn)=LCP(LCP(S1…Sk),LCP(Sk+1…Sn))
func longestCommonPrefixLCP(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    var lcp func(int, int) string
    lcp = func(start, end int) string {
        if start == end {
            return strs[start]
        }
        mid := (start + end) / 2
        lcpLeft, lcpRight := lcp(start, mid), lcp(mid + 1, end)
        minLength := min(len(lcpLeft), len(lcpRight))
        for i := 0; i < minLength; i++ {
            if lcpLeft[i] != lcpRight[i] {
                return lcpLeft[:i]
            }
        }
        return lcpLeft[:minLength]
    }
    return lcp(0, len(strs)-1)
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
