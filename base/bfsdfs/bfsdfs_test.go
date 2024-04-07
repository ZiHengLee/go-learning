package bfsdfs

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := make([][]byte,0)
	grid = append(grid, []byte{'1','1','1','1','0'})
	grid = append(grid, []byte{1,1,0,1,0})
	grid = append(grid, []byte{1,1,0,0,0})
	grid = append(grid, []byte{0,0,0,0,0})
	fmt.Println(numIslands(grid))
}
