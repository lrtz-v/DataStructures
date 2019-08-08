package leastinterval

import (
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	if n == 0 || len(tasks) == 1 {
		return len(tasks)
	}

	taskTypeList := []byte{}
	taskCountMap := make(map[byte]int, 26)
	sleepCount := make(map[byte]int, 26)

	for _, task := range tasks {
		_, ok := taskCountMap[task]
		if !ok {
			taskTypeList = append(taskTypeList, task)
			taskCountMap[task] = 1
		} else {
			taskCountMap[task]++
		}
	}
	sort.Slice(taskTypeList, func(i, j int) bool {
		return taskCountMap[taskTypeList[i]] > taskCountMap[taskTypeList[j]]
	})

	steps := 0

	for len(taskTypeList) > 0 {

		pick := false
		pickedIndex := 0

		// fmt.Print("sleepCount ")
		// fmt.Println(sleepCount)
		// fmt.Print("taskTypeList ")
		// fmt.Println(taskTypeList)
		// fmt.Print("taskCountMap ")
		// fmt.Println(taskCountMap)

		for i := 0; i < len(taskTypeList); i++ {
			task := taskTypeList[i]
			v, ok := sleepCount[task]
			if !ok {
				if !pick {
					pick = true
					pickedIndex = i
					sleepCount[task] = n
				}
			} else {
				if v == 1 {
					delete(sleepCount, task)
				} else {
					sleepCount[task] = sleepCount[task] - 1
				}
			}
		}

		if pick {
			// fmt.Println(rune(taskTypeList[pickedIndex]))
			if taskCountMap[taskTypeList[pickedIndex]] <= 1 {
				delete(taskCountMap, taskTypeList[pickedIndex])
				taskTypeList = append(taskTypeList[:pickedIndex], taskTypeList[pickedIndex+1:]...)
				// delete(sleepCount, taskTypeList[pickedIndex])
			} else {
				taskCountMap[taskTypeList[pickedIndex]]--
			}

			for j := pickedIndex; j < len(taskTypeList)-1; j++ {
				if taskCountMap[taskTypeList[j]] < taskCountMap[taskTypeList[j+1]] {
					taskTypeList[j], taskTypeList[j+1] = taskTypeList[j+1], taskTypeList[j]
				}
			}
			// sort.Slice(taskTypeList, func(i, j int) bool {
			// 	return taskCountMap[taskTypeList[i]] > taskCountMap[taskTypeList[j]]
			// })
		} else {
			// fmt.Println("sleep")
		}

		steps++
		// fmt.Println()
	}

	return steps
}
