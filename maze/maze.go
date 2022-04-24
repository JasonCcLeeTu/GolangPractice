package main

import(
      "fmt"
)



func main(){
	 maze :=[][]int{{0,1,0,0,0},{0,0,0,1,0},{0,1,0,1,0},{1,1,1,0,0},{0,1,0,0,1},{0,1,0,0,0}}

	 steps:=make([]int,len(maze[0]))
	//steps := make([][]int, len(maze[0]))
	// for i:=range steps{
		fmt.Println(steps)
//	 fmt.Println(len(maze[0]))
	//}
	
}




// func walk(maze [][]int, start, end point) [][]int{
// 	//建立迷宮正確路線
//    steps := make([][]int, len(maze))

// 	for i := range steps{
// 	   steps[i] = make([]int, len(maze[i]))
// 	}
   
// 		 //建立佇列Q
//     Q := []point{start}
   
// 	   for len(Q) > 0 {
// 		    cur := Q[0]
// 		  	Q = Q[1: ]
   
// 		    if cur == end{
// 			 break
//             }
   
// 		    for _, dir := range dirs{
// 			 	next := cur.add(dir)
// 				 val, ok := next.at(maze)
// 			 	if !ok || val == 1{    //1為牆，ok == fals為越界
// 					continue
// 				}
   
// 				 val, ok = next.at(steps) //判斷是否被訪問，steps中的值為零表示為牆
//   			 	if !ok || val != 0 {
// 					continue
//    				}
// 			 	if next == start{
// 				continue
//    				}
   
// 			 	curSteps, _ := cur.at(steps)
// 			 	steps[next.i][next.j] = curSteps + 1
   
//    				Q = append(Q, next)
   
// 			 //maze at next is 0
//    			//and steps at next is 0 指到過的點 //and next == 0  }
// 	   		}
// 	  	    return steps
// 	   }


// 	   0   0   4   5   6 
// 	   1   2   3   0   7 
// 	   2   0   4   0   8 
// 	   0   0   0  10   9 
// 	   0   0  12  11   0 
// 	   0   0  13  12  13 


//王大明


