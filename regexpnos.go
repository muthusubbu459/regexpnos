package main


import("regexp"
		"fmt"
	"strconv"
	//"testing"
)

/*func validationtest(t *testing.T) bool{
	no:=define()
	value, _ := regexp.MatchString("^[1]+[0-9]{2}$", strconv.Itoa(no))
	if value{
		return  value
	}else {
		t.Error("Test Failed")
	}
	return value
}*/
func validate(no int) bool{

	m, _ := regexp.MatchString("^[1][0-9]{2}$", strconv.Itoa(no))
	return m
}

func main() {
//	var no int
//	validationtest()
//	fmt.Println(value)
	//no=define()
	var no int
	fmt.Println("enter the number between 100 and 200")
	fmt.Scanf("%d/t", &no)
//	m, _ := regexp.MatchString("^[1]+[0-9]{2}$", strconv.Itoa(no))
	//fmt.Println(m)
	//if ! m {
	//	fmt.Println(m)

	//}
	m:=validate(no)
	if m {
		fmt.Println(m)
	}else {
		fmt.Println("false")
	}
}
