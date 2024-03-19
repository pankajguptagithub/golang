import (
        "fmt"
        "math"
)
type Shape interface{
        Area() float64
        Perimeter() float64
}
type Circle struct{
        Radius float64
}
type Square struct{
        Side float64
}
func (c Circle) Area()float64{
        return (math.Pi)*(c.Radius)*(c.Radius)
}
func (s Square) Area()float64{
        return (s.Side)*(s.Side)
}
func (c Circle) Perimeter()float64{
        return (2*(math.Pi)*(c.Radius))
}
func (s Square) Perimeter()float64{
        return (4*(s.Side))
}
func main(){
        n := 5.0
        var c Circle
        var s Square
        c.Radius = n
        s.Side = n
        fmt.Println("Area of circle with radius 5 is",c.Area())
        fmt.Println("Perimeter of circle with radius 5 is",c.Perimeter())
        fmt.Println("Area of square with side 5 is",s.Area())
        fmt.Println("Perimeter of square with side 5 is",s.Perimeter())
}
