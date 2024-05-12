package main
import "fmt"

type shaper interface {
Area() float32
}

type Square struct{
side float32
}
func(sq *Square) Area() float32{
return sq.side * sq.side
}
type rectangle struct{
length, width float32
}
func (r rectangle) Area() float32}
return r.length *r.width

func main() {
r :=rectangle{s, 3)//Area() of Rectangle needs a value
q:=&Square{5}       //Area() of Square needs a pointer
//shapes :=[]shaper{shaper(r),shaper(q)}
// or shorter:
shapes:=[]shaper{r,q,}
fmt.Println("Looping trough shapesfor area ...")
for n,-:=range shapes{
fmt.Println("shapes details: ",shapesArr(n])
fmt.Println("Area of this shape is: ", shapes[n].Area())
}
}
