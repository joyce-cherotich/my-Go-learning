package main
import(
"fmt"
"math"
)
type square struct {
side float32
}

type Circle struct {
radius float32
}
type shaper inerface {
Area() float32
}

func main() {
var areaIntf shaper
sq1 :=new(square)
sq1.side=5

areaIntf =sq1
//Is Square the type of areaIntf ?
if t, ok :=areaIntf.(*square); ok {
fmt.Printf("The type of areaIntf ok is: %T\n",t)
}
if u, ok :=areaIntf.(Circle); ok {
fmt.Printf("the type of areaIntf is:  %T\n",u)
} else
fmt.Println("areaIntf does not contain a variable of a type
Circle ")
}
}
func (sq *Square) Area() float32 {
return sq.side * sq.side
}
func (ci *Circle) Area()float32 {
return ci.radius * ci.radius *math.pi
}
