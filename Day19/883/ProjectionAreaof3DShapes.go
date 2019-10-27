package main
import "fmt"
func projectionArea(grid [][]int) int {
	ans:=0
	for i:=0;i<len(grid);i++{
		xzma:=0
		yzma:=0
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j]>0{
				ans++
			}
			if grid[i][j]>xzma{
				xzma=grid[i][j]
			}
			if grid[j][i]>yzma{
				yzma=grid[j][i]
			}
		}
		ans+=xzma
		ans+=yzma
	}
	return ans
}

func main(){
	grid:=[][]int{{2,2,2},{2,1,2},{2,2,2}}
	fmt.Println(projectionArea(grid))
}
