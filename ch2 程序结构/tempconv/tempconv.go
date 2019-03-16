package tempconv

// tempconv 负责摄氏温度与华氏温度的转换
import "fmt"

type (
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
)

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoolingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}
func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
