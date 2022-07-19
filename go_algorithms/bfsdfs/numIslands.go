package bfsdfs

//计算没有被0包围的1的块数，其中斜着的不算
//思路：
//bfs广度优先搜索

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	result := 0
	for i:=0;i<len(grid);i++{
		visited[i] = make([]bool,len(grid[i]))
	}
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j]=='1'&&!visited[i][j]{
				bfsIsland(i,j,len(grid),len(grid[i]),visited,grid)
				result+=1
			}
		}
	}

	return result
}

func bfsIsland(i,j,leni,lenj int,visited [][]bool,grid [][]byte) {
	if i<0 || i>=leni || j<0 || j>=lenj || visited[i][j] || grid[i][j]=='0'{
		return
	}
	visited[i][j] = true
	bfsIsland(i+1,j,leni,lenj,visited,grid)
	bfsIsland(i,j+1,leni,lenj,visited,grid)
	bfsIsland(i-1,j,leni,lenj,visited,grid)
	bfsIsland(i,j-1,leni,lenj,visited,grid)
}

