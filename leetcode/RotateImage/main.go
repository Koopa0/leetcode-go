package RotateImage

// rotate 將 n x n 的二維矩陣順時針旋轉90度
// 使用轉置+水平反射的方法實現原地旋轉
func rotate(matrix [][]int) {
	n := len(matrix)

	// 步驟1：轉置矩陣（交換行和列）
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// 交換 matrix[i][j] 和 matrix[j][i]
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 步驟2：水平反射（反轉每一行）
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			// 交換 matrix[i][j] 和 matrix[i][n-1-j]
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}
