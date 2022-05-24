package main

func numIslands(grid [][]byte) int {
	res := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				infect(grid, i, j)
			}
		}
	}

	return res
}

func infect(grid [][]byte, i int, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	infect(grid, i+1, j)
	infect(grid, i-1, j)
	infect(grid, i, j+1)
	infect(grid, i, j-1)
}
