/*
	leetcode tag:621 title:任务调度器
*/
package leetcode621

func leastInterval(tasks []byte, n int) int {
	// useful of save A-Z's mission number
	hashMap := [26]int{}
	for _, v := range tasks {
		hashMap[v-'A']++
	}

	// get the maximum and how many maxinum(count)
	max := 0
	count := 0
	for _, v := range hashMap {
		if max < v {
			max = v
			count = 1
		} else if max == v {
			count++
		}
	}

	// (n == 0) mean there no limit with standby status
	// (max -1) mean the last team did't nead with standby status
	// (n + 1) mean the element implement block with n number of block
	// (+ count) mean add the last team of misson which didn't need standby
	if n == 0 || (n+1)*(max-1)+count < len(tasks) {
		return len(tasks)
	}

	return (n+1)*(max-1) + count
}
