package main
import "fmt"
var arr =[...]string{
	"JANUARI",
	"FEBRUARI",
	"MARET",
	"APRIL",
	"MEI",
	"JUNI",
	"JULI",
	"AGUSTUS",
	"SEPTEMBER",
	"OKTOBER",
	"NOVEMBER",
	"DESEMBER",
}
var slice1= arr[len(arr)-1]
func main() {
fmt.Println(slice1)
}