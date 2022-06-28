package main
import (
	"fmt"
	"os"
	"strings"
)
func main(){
	args := os.Args[1]
	arr_args := strings.Split(args, " ")
	fmt.Println(len(arr_args))
}
