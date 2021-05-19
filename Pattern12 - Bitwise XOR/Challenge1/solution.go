package challenge1

/*
Given a binary matrix representing an image, we want to flip the image horizontally, then invert it.

To flip an image horizontally means that each row of the image is reversed. For example, flipping [0, 1, 1] horizontally results in [1, 1, 0].

To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [1, 1, 0] results in [0, 0, 1].

Input: [
  [1,0,1],
  [1,1,1],
  [0,1,1]
]
Output: [
  [0,1,0],
  [0,0,0],
  [0,0,1]
]

Input: [
  [1,1,0,0],
  [1,0,0,1],
  [0,1,1,1],
  [1,0,1,0]
]
Output: [
  [1,1,0,0],
  [0,1,1,0],
  [0,0,0,1],
  [1,0,1,0]
]

ref: https://leetcode-cn.com/problems/flipping-an-image/
*/

func flipAndInvertImage(image [][]int) [][]int {
	if len(image) == 0 {
		return [][]int{}
	}

	s := len(image[0])
	for i := 0; i < len(image); i++ {
		for j := 0; j < (s+1)/2; j++ {
			image[i][j], image[i][s-j-1] = image[i][s-j-1]^1, image[i][j]^1
		}
	}
	return image
}
