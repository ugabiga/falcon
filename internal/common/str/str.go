package str

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"strconv"
)

func ToCamel(s string) string {
	return strcase.ToCamel(s)
}

type Str struct {
	str string
}

func New(str string) *Str {
	return &Str{
		str: str,
	}
}

func FromFloat64(f float64) *Str {
	return New(fmt.Sprintf("%f", f))
}

func FromFloat64WithPrec(f float64, prec int) *Str {
	return New(strconv.FormatFloat(f, 'f', prec, 64))
}

func FromInt(i int) *Str {
	return New(strconv.Itoa(i))
}

func FromInt64(i int64) *Str {
	return New(strconv.FormatInt(i, 10))
}

func FromBool(b bool) *Str {
	return New(strconv.FormatBool(b))
}

func (s *Str) Val() string {
	return s.str
}

func (s *Str) ToFloat64() (float64, error) {
	return strconv.ParseFloat(s.str, 64)
}

func (s *Str) ToFloat64Default(defaultValue float64) float64 {
	f, err := s.ToFloat64()
	if err != nil {
		return defaultValue
	}
	return f
}

func (s *Str) ToInt() (int, error) {
	return strconv.Atoi(s.str)
}

func (s *Str) ToIntDefault(defaultValue int) int {
	i, err := s.ToInt()
	if err != nil {
		return defaultValue
	}
	return i
}

func (s *Str) ToInt64() (int64, error) {
	return strconv.ParseInt(s.str, 10, 64)
}

func (s *Str) ToInt64Default(defaultValue int64) int64 {
	i, err := s.ToInt64()
	if err != nil {
		return defaultValue
	}
	return i
}

func (s *Str) CountDecimalCount() int {
	str := s.str
	count := 0
	foundDot := false

	for i := 0; i < len(str); i++ {
		if str[i] == '.' {
			foundDot = true
		}

		if foundDot {
			count++
		}
	}

	return count - 1
}
