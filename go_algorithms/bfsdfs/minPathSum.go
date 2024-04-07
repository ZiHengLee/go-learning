package bfsdfs

//最开始思路dfs，看评论这个方法过不了，得用dp
//用dp那就是简单题了
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0])==0{
		return 0
	}
	xLen := len(grid)
	yLen := len(grid[0])
	dp := make([][]int, xLen+1)
	dp[0] = make([]int,yLen+1)
	for i:=1;i<xLen+1;i++{
		dp[i] = make([]int,yLen+1)
		dp[i][0]+=dp[i-1][0]
	}
	for j:=1;j<yLen+1;j++{
		dp[0][j]+=dp[0][j-1]
	}
	for i := 1;i< xLen+1;i++{
		for j:=1;j<yLen+1;j++{
			if dp[i-1][j]<dp[i][j-1]{
				dp[i][j]=dp[i-1][j]+grid[i][j]
			}else{
				dp[i][j]=dp[i][j-1]+grid[i][j]
			}
		}
	}
	return dp[xLen][yLen]
}
