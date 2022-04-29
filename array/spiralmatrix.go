package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	left := 0
	right := len(matrix[0]) - 1
	up := 0
	down := len(matrix[0]) - 1
	var result []int
	visited := 0
	for visited < len(matrix[0])*len(matrix) {
		//traverse left
		for col := left; col <= right; col++ {
			result = append(result, matrix[up][col])
			visited++
		}
		//traverse down
		for row := up + 1; row <= down; row++ {
			result = append(result, matrix[row][right])
			visited++
		}
		if up != down {
			//traverse right
			for col := right - 1; col >= left; col-- {
				result = append(result, matrix[down][col])
				visited++
			}
		}

		if left != right {
			//traverse up
			for row := down - 1; row > up; row-- {
				result = append(result, matrix[row][up])
				visited++
			}
		}
		left++
		right++
		down--
		up++
	}
	return result
}
