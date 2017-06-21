package islandPerimeter

func islandPerimeter(grid [][]int) int {
	island := 0
	neighbor := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				island++
				if i < len(grid)-1 && grid[i+1][j] == 1 {
					neighbor++
				}
				if j < len(grid[i])-1 && grid[i][j+1] == 1 {
					neighbor++
				}
			}
		}
	}
	return island*4 - neighbor*2
}
