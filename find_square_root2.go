package main
import (
	"fmt"
	"time"
)
func findSqroot(num float64)float64{
        var sol float64 = 0
	var low float64 = 0
	var high float64 = num
        if num == 0 || num == 1{
                sol = num
        }else{
                for{
                        mid := ((low + high)/2)
			square := mid*mid
			fmt.Println("mid = ",mid)
                        if square == (num){
                                sol = mid
                                break
                        }else if square >= num{
                                high = high - 1
                        }else{
				low = mid
                        }
			time.Sleep(1*time.Millisecond)
                }
        }
        return sol
}
func main(){
	fmt.Println("Square root of given number is ",findSqroot(8))
}
