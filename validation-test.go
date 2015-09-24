package main
import(
"testing"
)
func value()bool{
	return true
}
func testvalidate_1(t *testing.T) {
	//expected :=value()
	actual:=validate(121)
	//value, _ := regexp.MatchString("^[1]+[0-9]{2}$", strconv.Itoa(no))
	if actual != false{
		t.Error("Fail")
	}
	//return value
}
