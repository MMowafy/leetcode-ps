package main

func main() {
	rotateImage([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
}

func rotateImage(matrix [][]int) {
	transpose(matrix)
	reverse(matrix)
}

func transpose(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		for column := row + 1; column < len(matrix); column++ {
			temp := matrix[row][column]
			matrix[row][column] = matrix[column][row]
			matrix[column][row] = temp
		}
	}
}

func reverse(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix)/2; column++ {
			temp := matrix[row][column]
			matrix[row][column] = matrix[row][len(matrix)-column-1]
			matrix[row][len(matrix)-column-1] = temp
		}
	}
}
