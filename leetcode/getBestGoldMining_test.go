package leetcode

import (
	"testing"
)

func TestGetBestGildMining(t *testing.T) {
	workerNum := 10
	costList := []int{5, 5, 3, 4, 3}
	goldsList := []int{400, 500, 200, 300, 350}

	/*
		workerNum: 工人数量
		costList：金矿开采所需人工数量
		goldsList：金矿储量
	*/
	res := getBestGildMining(workerNum, costList, goldsList)
	if res != 900 {
		t.Errorf("Expected %d, Got: %d\n", 900, res)
	}
}

func TestGetBestGildMining2(t *testing.T) {
	workerNum := 10
	costList := []int{5, 5, 3, 4, 3}
	goldsList := []int{400, 500, 200, 300, 350}

	/*
		workerNum: 工人数量
		costList：金矿开采所需人工数量
		goldsList：金矿储量
	*/
	res := getBestGildMining2(workerNum, costList, goldsList)
	if res != 900 {
		t.Errorf("Expected %d, Got: %d\n", 900, res)
	}
}
