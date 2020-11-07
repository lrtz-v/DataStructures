package sort

/*
希尔排序

基于插入排序的排序算法
插入排序只会交换相邻的元素，元素只能一点一点的从数组的一端移动到另一端，希尔排序加速了这一过程
使数组中任意间隔为h的元素都是有序的，这样的数组被称为h有序数组
如果h很大，我们就能将元素移动到很远的地方

复杂度
	空间复杂度：O(1)
	时间复杂度: O(n) - O(n^3/2)
	平均时间复杂度: O(n) - O(n^3/2)

特点
简单实现，对于中等大小的数组其运行时间可接受
*/

// Shell sort
type Shell struct {
	Template
}

// Sort list
func (shell Shell) Sort(l []int64) {

	size := shell.Len(l)
	h := 1
	for h < size/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < size; i++ {
			for j := i; j >= h && shell.Less(l[j], l[j-h]); j -= h {
				shell.Exch(l, j, j-h)
			}
		}
		h /= 3
	}
}
