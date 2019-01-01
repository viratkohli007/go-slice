package main
import "fmt"

func main()  {
	
	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set s: ",s)
	fmt.Println("2nd value: ", s[2])
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, 6)
	copy(c, s)
	fmt.Println("C: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)
	l = s[4:]
	fmt.Println("sl2: ", l)

	l = s[:3]
	fmt.Println("sl3", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)
	
	towD := make([][]int, 3)
	for i := 0; i <3; i++ {
		inner :=i+1 		
		towD[i] = make([]int, inner)
		for j:=0; j<inner; j++{
			towD[i][j] = i+j;
		}
	} 
	fmt.Println("2D",towD)

}