package skiplist

import (
	"fmt"
	"strconv"
	"testing"
)

func zslPrint(zsl *zskiplist) {
	x := zsl.header
	for {
		if x.level[0].forward != nil {
			fmt.Printf("key: %s, value: %d, score: %f\n", x.level[0].forward.key, x.level[0].forward.obj, x.level[0].forward.score)
			x = x.level[0].forward
		} else {
			break
		}

	}
}

func TestInsert(t *testing.T) {
	zsl := zslInit()
	data := []int{1, 2, 3, 4, 4, 5, 6, 7, 7, 8, 9}
	for _, v := range data {
		zsl.zslInsert(float32(v), v, strconv.Itoa(v))
	}
	if zsl.length != 11 {
		t.Fatal("Insert failed!")
	}
}

func TestZslRandomLevel(t *testing.T) {
	count := make(map[int]int)

	zsl := zslInit()
	for i := 0; i < 10000; i++ {
		level := zsl.zslRandomLevel()
		if _, ok := count[level]; ok {
			count[level] += 1
		} else {
			count[level] = 1
		}
	}
}

func TestZSLFree(t *testing.T) {
	zsl := zslInit()
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range data {
		zsl.zslInsert(1.0, v, strconv.Itoa(v))
	}
	zslFree(zsl)
}

func TestDelete(t *testing.T) {
	zsl := zslInit()
	data := []int{1, 2, 3, 4, 4, 5, 6, 7, 7, 8, 9}
	for _, v := range data {
		zsl.zslInsert(float32(v), v, strconv.Itoa(v))
	}
	zsl.zslDelete(4.0, "4")
	if zsl.length != 9 {
		t.Fatal("Delete failed!")
	}
}

func TestZslGetRank(t *testing.T) {
	zsl := zslInit()
	data := []int{1, 2, 3, 4, 4, 5, 6, 7, 7, 8, 9}
	for _, v := range data {
		zsl.zslInsert(float32(v), v, strconv.Itoa(v))
	}
	if zsl.zslGetRank(9.0, "9") != 11 {
		t.Fatal("Get Rank Failed!")
	}

	if zsl.zslGetRank(10.0, "10") != 0 {
		t.Fatal("Get Rank Failed!")
	}

}

func TestZslGetElementByRank(t *testing.T) {
	zsl := zslInit()
	data := []int{1, 2, 3, 4, 4, 5, 6, 7, 7, 8, 9}
	for _, v := range data {
		zsl.zslInsert(float32(v), v, strconv.Itoa(v))
	}

	x := zsl.zslGetElementByRank(1)
	if x.key != "1" || x.obj != 1 {
		t.Fatal("GetElementByRank Failed!")
	}

	x = zsl.zslGetElementByRank(11)
	if x.key != "9" || x.obj != 9 {
		t.Fatal("GetElementByRank Failed!")
	}

	x = zsl.zslGetElementByRank(15)
	if x != nil {
		t.Fatal("GetElementByRank Failed!")
	}
}
