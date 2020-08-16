/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

示例 1:

输入:
[
['1','1','1','1','0'],
['1','1','0','1','0'],
['1','1','0','0','0'],
['0','0','0','0','0']
]
输出: 1
示例 2:

输入:
[
['1','1','0','0','0'],
['1','1','0','0','0'],
['0','0','1','0','0'],
['0','0','0','1','1']
]
输出: 3
解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
*/
package main

import "fmt"

func main() {
	//grid := [][]byte{
	//	{'1', '1', '0', '0', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '1', '0', '0'},
	//	{'0', '0', '0', '1', '1'},
	//}
	//grid := [][]byte{
	//	{'1', '0', '1', '1', '1'},
	//	{'1', '0', '1', '0', '1'},
	//	{'1', '1', '1', '0', '1'},
	//}
	var grid [][]byte = nil

	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	var result = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' || grid[i][j] == '2' {
				continue
			}
			result++
			fillGrid(grid, i, j)
		}
	}
	return result
}

func fillGrid(grid [][]byte, i, j int) {
	if grid == nil || i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' || grid[i][j] == '2' {
		return
	}
	grid[i][j] = '2'
	fillGrid(grid, i-1, j)
	fillGrid(grid, i+1, j)
	fillGrid(grid, i, j-1)
	fillGrid(grid, i, j+1)
}
