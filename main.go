package main
import (
	"fmt"
	"os"
	"strings"
)
func main(){
	sum := 0
	len_args := len(os.Args) - 1
	var arg string

	for i:= 1; i <= len_args; i++ {
		arg = os.Args[i]
		arr_arg := strings.Split(arg, " ")
		sum += len(arr_arg)
	}
	fmt.Println(sum)
}
