package twosum

func twoSum(numbers []int, target int) []int {

	m := map[int]int{}
	for i, v := range numbers {
		compement := target - v
		kk, vv := m[compement]
		if vv {
			return []int{kk, i}
		}
		// store key
		m[v] = i
	}

	panic("no solution")
}
