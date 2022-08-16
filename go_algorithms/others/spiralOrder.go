package others

//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//123
//456
//789
func spiralOrder(matrix [][]int) []int {
	ret := make([]int, 0)
	visited := make([][]bool, len(matrix))
	xNums, yNums := 0, 0
	if len(matrix) > 0 {
		xNums = len(matrix)
		if len(matrix[0]) > 0 {
			yNums = len(matrix[0])
		}
	}
	for i := 0; i < xNums; i++ {
		visited[i] = make([]bool, yNums)
	}
	x, y := 0, 0
	dir := 4
	for true {
		lableBreak := 0
		if (x >= 0 && x < xNums) && (y >= 0 && y < yNums) {
			ret = append(ret, matrix[x][y])
			visited[x][y] = true
			//上下左右几个方向
			switch dir {
			case 1:
				//fmt.Println(x,y,matrix[x][y],dir,visited[x][y])
				x--
				//走过了
				if x < 0 || visited[x][y] {
					dir = 4
					x++
					if y < yNums-1 && (visited[x][y+1] == false) {
						y++
						continue
					} else {
						lableBreak = 1
						break
					}
				}
			case 2:
				//fmt.Println(x,y,matrix[x][y],dir,visited[x][y])
				x++
				if x >= xNums || visited[x][y] {
					dir = 3
					x--
					if y > 0 && (visited[x][y-1] == false) {
						y--
						continue
					} else {
						lableBreak = 1
						break
					}
				}
			case 3:
				//fmt.Println(x,y,matrix[x][y],dir,visited[x][y])
				y--
				if y < 0 || visited[x][y] {
					dir = 1
					y++
					if x > 0 && (visited[x-1][y] == false) {
						x--
						continue
					} else {
						lableBreak = 1
						break
					}
				}
			case 4:
				//fmt.Println(x,y,matrix[x][y],dir,visited[x][y])
				y++
				if y >= yNums || visited[x][y] {
					dir = 2
					y--
					//fmt.Println(x,y,matrix[x][y],dir,visited[x+1][y],"ddd")
					if x < xNums-1 && (visited[x+1][y] == false) {
						x++
						continue
					} else {
						lableBreak = 1
						break
					}
				}
			}
		} else {
			break
		}
		if lableBreak == 1{
			break
		}
	}
	return ret
}
