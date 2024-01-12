package floatf

import "math"

type Floatf struct {
	num float64
}

func New(f float64) *Floatf {
	return &Floatf{
		num: f,
	}
}

func (f *Floatf) ToFixed(precision int) float64 {
	r := math.Pow(10, float64(precision))
	return float64(int64(f.num*r)) / r
}
