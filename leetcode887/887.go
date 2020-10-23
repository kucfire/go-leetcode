/*
	leetcode tag:887 title:鸡蛋掉落
*/

package leetcode887

//method of Li yongle
func SuperEggDrop(K int, N int) int {
	// create a binary array whitch have K+1 array and each array have N elements
	// 创建一个二元数组，数组长度为K，每个子数组里面有N个元素
	result := make([][]int, N+1)
	for i := range result {
		result[i] = make([]int, K+1)
	}

	// return the smaller element
	// 返回其中数值小的变量
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	// return the bigger element
	// 返回其中数值大的变量
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	// 遍历每一个楼层
	for i := 1; i <= N; i++ {
		// 楼层数为i，如果只有一个鸡蛋，则需要从低到高进行遍历，最坏情况下需要遍历i次
		result[i][1] = i
		for j := 2; j <= K; j++ { // 遍历K个鸡蛋
			result[i][j] = 1 << 32 // 反向取一个极大值
			// 遍历鸡蛋仍在x层的情况
			for x := 1; x <= i; x++ {
				// 假设临界值为t
				// result[x-1][j-1]为鸡蛋碎了,t在[1,x-1]这个区间中,共x层,剩余j-1个鸡蛋
				// result[i-x][j]为鸡蛋没碎,t在[x,i]这个区间中,共i-x层,剩余j个鸡蛋
				// 这两种情况中取最大值,(考虑最坏的情况),加上本次,更新result[i][j]
				result[i][j] = min(result[i][j], max(result[x-1][j-1], result[i-x][j])+1)
			}
		}
	}

	return result[N][K]
}

//method of normal dp
func SuperEggDrop2(K int, N int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	// create a binary array whitch have K+1 array and each array have N elements
	// 创建一个二元数组，数组长度为K，每个子数组里面有N个元素
	result := make([][]int, N+1)
	for i := range result {
		result[i] = make([]int, K+1)
	}

	for i := 1; i <= N; i++ {
		// 楼层数为i，如果只有一个鸡蛋，则需要从低到高进行遍历，最坏情况下需要遍历i次
		result[i][1] = i
		for j := 2; j <= K; j++ {
			result[i][j] = 1 << 32 // 反向取一个极大值

			// 二分法在区间[1:i]里确定一个最优值
			left, right := 1, i
			for left < right {
				mid := left + (right-left+1)/2
				breakCount := result[mid-1][j-1]
				notBreakCOunt := result[i-mid][j]
				if breakCount > notBreakCOunt {
					right = mid - 1
				} else {
					left = mid
				}
			}

			// left在这个下表就是最优的k值，把他代入转移方程即可
			result[i][j] = max(result[left-1][j-1], result[i-left][j]) + 1
		}
	}
	return result[N][K]
}

// method of encode
func SuperEggDrop3(K int, N int) int {

}
