package bucket

import (
	"fmt"
	"strings"
	"testing"
)

func converse(arr []float64) []string {
	stringArr := make([]string, len(arr))
	for i, val := range arr {
		stringArr[i] = fmt.Sprintf("%f", val)
	}
	return stringArr
}

func TestBucketSort(t *testing.T) {
	testArr := []float64{4.12, 6.421, 0.0023, 3.0, 2.123, 8.122, 4.12, 10.09}
	testArr = bucketSort(testArr)
	fmt.Println(strings.Join(converse(testArr), ", "))
}
