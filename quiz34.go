// You can edit this code!
// Click here and start typing.
//for ptinting an inverted triangle
package main
import "fmt"
func main() {
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}