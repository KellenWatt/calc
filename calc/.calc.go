package calc

import (	
	"math"
	"fmt"
	"strings"
)

type Calculator struct {
	stack Stack
}

func New() *Calculator {
	return new(Calculator)
}

func (c *Calculator) String() string {
	elems := c.stack.All()
	out := make([]string, len(elems))
	for i,f := range elems {
		out[i] = fmt.Sprint(f)
	}
	return "[" + strings.Join(out, " ") + "]"
}

func (c *Calculator) Clear() {
	for _,err := c.stack.Pop(); err == nil; _,err = c.stack.Pop() {}
}

func (c *Calculator) IsSanitary() bool {
	elems := c.stack.All()
	for _, v := range elems {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func (c *Calculator) Sanitize() {
	elems := c.stack.All()
	var s Stack
	
	for i := c.stack.Size() - 1; i >= 0; i-- {
		e := elems[i]
		if !(math.IsNaN(e) || math.IsInf(e, 0)) {
			s.Push(e)
		}
	}
	c.stack = s
}

func (c *Calculator) Pop() error {
	_, err := c.stack.Pop()
	return err
}

func (c *Calculator) Push(n float64) {
	c.stack.Push(n)
}

func (c *Calculator) Last() (float64, error) {
	v, err := c.stack.Peek(1)
	if err != nil {
		return 0, err
	}
	return v[0], nil
}

func (c *Calculator) Swap() error {
	n, m, err := c.stack.Pop2()
	if err != nil {
		return err
	}
	c.stack.Push(n).Push(m)
	return nil
}

func (c *Calculator) Duplicate() error {
	n, err := c.stack.Peek(1)
	if err != nil {
		return err
	}
	c.stack.Push(n[0])
	return nil
}

func (c *Calculator) Truncate() error {
	n, err := c.stack.Pop()
	if err != nil {
		return err
	}
	c.stack.Push(math.Trunc(n))
	return nil
}

func (c *Calculator) Negate() error {
	v, err := c.stack.Pop()
	if err != nil {
		return err
	}
	c.stack.Push(-v)
	return nil
}

func (c *Calculator) Add() error {
	n, m, err := c.stack.Pop2()
	if err != nil {
		return err
	}
	c.stack.Push(m + n)
	return nil
}

func (c *Calculator) Subtract() error {
	n, m, err := c.stack.Pop2()
	if err != nil {
		return err
	}
	c.stack.Push(m - n)
	return nil
}

func (c *Calculator) Multiply() error {
	n, m, err := c.stack.Pop2()
	if err != nil {
		return err
	}
	c.stack.Push(m * n)
	return nil
}

func (c *Calculator) Divide() error {
	n, m, err := c.stack.Pop2()
	if err != nil {
		return err
	}
	if n == 0 {
		c.stack.Push(math.NaN())
	} else {
		c.stack.Push(m / n)
	}
	return nil
}

func (c *Calculator) IntegerDivide() error {
	n, m, err := c.stack.Pop2()
	if err != nil {
		return err
	}
	if n == 0 {
		c.stack.Push(math.NaN())
	} else {
		c.stack.Push(math.Floor(m / n))
	}
	return nil
}

func (c *Calculator) Modulo() error {
	n, m, err := c.stack.Pop2()
	if err != nil {
		return err
	}
	c.stack.Push(math.Mod(m, n))
	return nil
}

func (c *Calculator) Power() error {
	n, m, err := c.stack.Pop2()
	if err != nil {
		return err
	}
	c.stack.Push(math.Pow(m, n))
	return nil
}

func (c *Calculator) Compare() error {
	n, m, err := c.stack.Pop2()
	if err != nil {
		return err
	}
	res := 0.0
	if m > n {
		res = 1
	} else if m < n {
		res = -1
	}
	c.stack.Push(res)
	return nil
}

func (c *Calculator) Log() error {
	n, m, err := c.stack.Pop2()
	if err != nil {
		return err
	}
	c.stack.Push(math.Log(m) / math.Log(n))
	return nil
}

func (c *Calculator) Log10() error {
	n, err := c.stack.Pop()
	if err != nil {
		return err
	}
	c.stack.Push(math.Log10(n))
	return nil
}

func (c *Calculator) NaturalLog() error {
	n, err := c.stack.Pop()
	if err != nil {
		return err
	}
	c.stack.Push(math.Log(n))
	return nil
}

func (c *Calculator) Sine() error {
	n, err := c.stack.Pop()
	if err != nil {
		return err
	}
	c.stack.Push(math.Sin(n))
	return nil
}

func (c *Calculator) Cosine() error {
	n, err := c.stack.Pop()
	if err != nil {
		return err
	}
	c.stack.Push(math.Cos(n))
	return nil
}

func (c *Calculator) Tangent() error {
	n, err := c.stack.Pop()
	if err != nil {
		return err
	}
	c.stack.Push(math.Tan(n))
	return nil
}

func (c *Calculator) SquareRoot() error {
	n, err := c.stack.Pop()
	if err != nil {
		return err
	}
	c.stack.Push(math.Sqrt(n))
	return nil
}

func (c *Calculator) AbsoluteValue() error {
	n, err := c.stack.Pop()
	if err != nil {
		return err
	}
	c.stack.Push(math.Abs(n))
	return nil
}
