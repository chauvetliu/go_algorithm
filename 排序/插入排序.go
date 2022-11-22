package main

// 与选择排序一样，当前索引左边的所有元素都是有序的，但是他们的最终位置还不确定，
// 为了给更小的元素腾出空间，他们可能会被移动，但是当索引到达数组的末端，数组排序就完成了
//稳定排序
//时间复杂度分析：O(N^2)，如果序列在排序前已经是有序序列，则为O(N)
//空间复杂度分析：O(1)
//数据量较少时效率高。插入排序适合数据量少的情况
//算法的实际运行效率优于选择排序和冒泡排序。
//插入排序对于部分有序的数组很有效，部分有序的数组类如数组中每个元素距离他的最终位置不远，一个有序的大数组接一个小数组，数组中只有几个元素的位置不正确

// 插入排序
func insertSort(data []int) {
	var j int
	for i := 1; i < len(data); i++ {
		temp := data[i]
		for j = i; j > 0 && temp < data[j-1]; j-- {
			data[j] = data[j-1]
		}
		data[j] = temp
	}
}
