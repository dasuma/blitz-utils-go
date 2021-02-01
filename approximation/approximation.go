package approximation

import (
	"math"
)

//Get approximation
func Get(value float64, aproximation int) float64 {
	var res float64
	res = 0
	switch aproximation {
	case 0:
		res = value
	case 10:
		res = math.RoundToEven(value/10) * 10
	case 100:
		res = math.RoundToEven(value/100) * 100
	case 1000:
		res = math.RoundToEven(value/1000) * 1000
	case 10000:
		res = math.RoundToEven(value/10000) * 1000
	}
	return res
}
