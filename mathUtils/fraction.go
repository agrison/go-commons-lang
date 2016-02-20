// Package mathUtils provides utilities regarding mathematics.
package mathUtils

import (
	"math"
	"math/big"
)

// Fraction represents a fraction which holds both a numerator and a denominator
// Definitely needs testing!
type Fraction struct {
	// The numerator number part of the fraction (the Three in Three sevenths).
	numerator int
	// The denominator number part of the fraction (the seven in Three sevenths).
	denominator int
}

// Zero is 0/1
var Zero = NewFraction(0, 1)

// One is 1/1
var One = NewFraction(1, 1)

// OneHalf is 1/2
var OneHalf = NewFraction(1, 2)

// OneThird is 1/3
var OneThird = NewFraction(1, 3)

// TwoThirds is 2/3
var TwoThirds = NewFraction(2, 3)

// OneQuarter is 1/4
var OneQuarter = NewFraction(1, 4)

// TwoQuarters is 2/4
var TwoQuarters = NewFraction(2, 4)

// ThreeQuarters is 3/4
var ThreeQuarters = NewFraction(3, 4)

// OneFifth is 1/5
var OneFifth = NewFraction(1, 5)

// TwoFifths is 2/5
var TwoFifths = NewFraction(2, 5)

// ThreeFifths is 3/5
var ThreeFifths = NewFraction(3, 5)

// FourFifths is 4/5
var FourFifths = NewFraction(4, 5)

// MaxUint is the maximum uint value
const MaxUint = ^uint(0)

// MinUint is the minimum uint value (0)
const MinUint = 0

// MaxInt is the maximum int value
const MaxInt = int(^uint(0) >> 1)

// MinInt is the minimum int value
const MinInt = -(MaxInt - 1)

// NewFraction constructs a Fraction instance with the 2 parts of a fraction Y/Z.
func NewFraction(numerator, denominator int) *Fraction {
	return &Fraction{numerator, denominator}
}

// GetFraction creates a Fraction instance with the 2 parts of a fraction Y/Z.
func GetFraction(numerator, denominator int) *Fraction {
	if denominator == 0 {
		return nil
	}
	if denominator < 0 {
		if numerator == MinInt || denominator == MinInt {
			return nil
		}
		numerator = -numerator
		denominator = -denominator
	}
	return NewFraction(numerator, denominator)
}

// GetWholeFraction creates a Fraction instance with the 3 parts of a fraction X Y/Z.
func GetWholeFraction(whole, numerator, denominator int) *Fraction {
	if denominator == 0 {
		return nil
	}
	if denominator < 0 {
		return nil
	}
	if numerator < 0 {
		return nil
	}
	numeratorValue := int64(0)
	if whole < 0 {
		numeratorValue = int64(whole*denominator - numerator)
	} else {
		numeratorValue = int64(whole*denominator + numerator)
	}
	if numeratorValue < int64(MinInt) || numeratorValue > int64(MaxInt) {
		return nil
	}
	return NewFraction(int(numeratorValue), denominator)
}

// GetReducedFraction creates a reduced Fraction instance with the 2 parts of a fraction Y/Z.
// For example, if the input parameters represent 2/4, then the created fraction will be 1/2.
func GetReducedFraction(numerator, denominator int) *Fraction {
	if denominator == 0 {
		return nil
	}
	if numerator == 0 {
		return Zero
	}
	if denominator == MinInt && (numerator&1) == 0 {
		numerator = numerator / 2
		denominator = denominator / 2
	}
	if denominator < 0 {
		if numerator == MinInt || denominator == MinInt {
			return nil
		}
		numerator = -numerator
		denominator = -denominator
	}
	// simplify Fraction
	gcd := greatestCommonDivisor(numerator, denominator)
	numerator = numerator / gcd
	denominator = denominator / gcd
	return NewFraction(numerator, denominator)
}

// Equals checks if Two Fraction are the same
func (f *Fraction) Equals(f2 *Fraction) bool {
	return f.numerator == f2.numerator && f.denominator == f2.denominator
}

// Reduce reduce the fraction to the smallest values for the numerator and denominator, returning the result.
// For example, if this fraction represents 2/4, then the result will be 1/2.
func (f *Fraction) Reduce() *Fraction {
	if f.numerator == 0 {
		if f.Equals(Zero) {
			return f
		}
		return Zero
	}
	gcd := greatestCommonDivisor(int(math.Abs(float64(f.numerator))), f.denominator)
	if gcd == 1 {
		return f
	}
	return GetFraction(f.numerator/gcd, f.denominator/gcd)
}

// Invert gets a fraction that is the inverse (1/fraction) of this One.
// The returned fraction is not reduced.
func (f *Fraction) Invert() *Fraction {
	if f.numerator == 0 {
		panic("unable to invert Zero")
	}
	if f.numerator == MinInt {
		panic("can't negate numerator")
	}
	if f.numerator < 0 {
		return NewFraction(-f.denominator, -f.numerator)
	}
	return NewFraction(f.denominator, f.numerator)
}

// Negate gets a fraction that is the negative (-fraction) of this One.
// The returned fraction is not reduced.
func (f *Fraction) Negate() *Fraction {
	if f.numerator == MinInt {
		panic("too large to negate")
	}
	return NewFraction(-f.numerator, f.denominator)
}

// Abs gets a fraction that is the positive equivalent of this One.
// More precisely: fraction >= 0 ? this : -fraction)
// The returned fraction is not reduced.
func (f *Fraction) Abs() *Fraction {
	if f.numerator >= 0 {
		return f
	}
	return f.Negate()
}

// Pow get a fraction that is powered by a specific value
func (f *Fraction) Pow(power int) *Fraction {
	if power == 1 {
		return f
	} else if power == 0 {
		return One
	} else if power < 0 {
		if power == MinInt {
			return f.Invert().Pow(2).Pow(-(power / 2))
		}
		return f.Invert().Pow(-power)
	}
	ff := f.MultiplyBy(f)
	if power%2 == 0 {
		return ff.Pow(power / 2)
	}
	return ff.Pow(power / 2).MultiplyBy(f)
}

// MultiplyBy multiplies the value of this fraction by another, returning the result in reduced form.
func (f *Fraction) MultiplyBy(ff *Fraction) *Fraction {
	if f.numerator == 0 || ff.numerator == 0 {
		return Zero
	}
	// knuth 4.5.1
	// make sure we don't overflow unless the result *must* overflow.
	d1 := greatestCommonDivisor(f.numerator, ff.denominator)
	d2 := greatestCommonDivisor(ff.numerator, f.denominator)
	//
	a := mulAndCheck(f.numerator/d1, ff.numerator/d2)
	b := mulPosAndCheck(f.denominator/d2, ff.denominator/d1)
	return GetReducedFraction(a, b)
}

// DivideBy divides the value of this fraction by another.
func (f *Fraction) DivideBy(ff *Fraction) *Fraction {
	if ff.numerator == 0 {
		panic("The fraction to divide by must not be Zero")
	}
	return f.MultiplyBy(ff.Invert())
}

// GetNumerator gets the numerator part of the fraction.
// This method may return a value greater than the denominator, an improper fraction, such as the seven in 7/4.
func (f *Fraction) GetNumerator() int {
	return f.numerator
}

// GetDenominator gets the denominator part of the fraction.
func (f *Fraction) GetDenominator() int {
	return f.denominator
}

// GetProperNumerator gets the proper numerator, always positive.
// An improper fraction 7/4 can be resolved into a proper One, 1 3/4. This method returns the 3 from the proper fraction.
// If the fraction is negative such as -7/4, it can be resolved into -1 3/4, so this method returns the positive proper numerator, 3.
func (f *Fraction) GetProperNumerator() int {
	return int(math.Abs(float64(f.numerator % f.denominator)))
}

// GetProperWhole gets the proper whole part of the fraction.
// An improper fraction 7/4 can be resolved into a proper One, 1 3/4. This method returns the 1 from the proper fraction.
// If the fraction is negative such as -7/4, it can be resolved into -1 3/4, so this method returns the positive whole part -1.
func (f *Fraction) GetProperWhole() int {
	return f.numerator / f.denominator
}

// IntValue gets the fraction as an int. This returns the whole number part of the fraction.
func (f *Fraction) IntValue() int {
	return f.numerator / f.denominator
}

// Float32Value gets the fraction as a float32. This calculates the fraction as the numerator divided by denominator.
func (f *Fraction) Float32Value() float32 {
	return float32(f.numerator) / float32(f.denominator)
}

// Float64Value gets the fraction as a float64. This calculates the fraction as the numerator divided by denominator.
func (f *Fraction) Float64Value() float64 {
	return float64(f.numerator) / float64(f.denominator)
}

// greatestCommonDivisor gets the greatest common divisor of the absolute value of Two numbers, using the "binary gcd" method which avoids
// division and modulo operations. See Knuth 4.5.2 algorithm B. This algorithm is due to Josef Stein (1961).
func greatestCommonDivisor(u, v int) int {
	// From Commons Math:
	if u == 0 || v == 0 {
		if u == MinInt || v == MinInt {
			panic("overflow: gcd is 2^31")
		}
		return int(math.Abs(float64(u)) + math.Abs(float64(v)))
	}
	//if either operand is abs 1, return 1:
	if int(math.Abs(float64(u))) == 1 || int(math.Abs(float64(v))) == 1 {
		return 1
	}
	// keep u and v negative, as negative integers range down to
	// -2^31, while positive numbers can only be as large as 2^31-1
	// (i.e. we can't necessarily negate a negative number without
	// overflow)
	if u > 0 {
		u = -u // make u negative
	}
	if v > 0 {
		v = -v // make v negative
	}
	// B1. [Find power of 2]
	k := 0
	for (u&1) == 0 && (v&1) == 0 && k < 31 { // while u and v are both even...
		u = u / 2 // cast out Twos.
		v = v / 2
		k++
	}
	if k == 31 {
		panic("gcd is 2^31")
	}
	// B2. Initialize: u and v have been divided by 2^k and at least
	//     One is odd.
	t := 0
	if (u & 1) == 1 {
		t = v
	} else {
		t = -(u / 2)
	}
	// t negative: u was odd, v may be even (t replaces v)
	// t positive: u was even, v is odd (t replaces u)
	foo := true
	for foo || (t != 0) {
		foo = false
		// B4/B3: cast out Twos from t.
		for (t & 1) == 0 { // while t is even..
			t = t / 2 // cast out Twos
		}
		// B5 [reset max(u,v)]
		if t > 0 {
			u = -t
		} else {
			v = t
		}
		// B6/B3. at this point both u and v should be odd.
		t = (v - u) / 2
		// |u| larger: t positive (replace u)
		// |v| larger: t negative (replace v)
	}
	return -u * (1 << uint(k)) // gcd is u*2^k
}

// mulAndCheck multiply two integers, checking for overflow.
func mulAndCheck(x int, y int) int {
	m := int64(x) * int64(y)
	if m < int64(MinInt) || m > int64(MaxInt) {
		panic("overflow: mul")
	}
	return int(m)
}

// mulPosAndCheck multiply two non-negative integers, checking for overflow.
func mulPosAndCheck(x int, y int) int {
	m := int64(x) * int64(y)
	if m > int64(MaxInt) {
		panic("overflow: mulPos")
	}
	return int(m)
}

// addAndCheck add two integers, checking for overflow.
func addAndCheck(x int, y int) int {
	s := int64(x) + int64(y)
	if s < int64(MinInt) || s > int64(MaxInt) {
		panic("overflow: add")
	}
	return int(s)
}

// subAndCheck subtract two integers, checking for overflow.
func subAndCheck(x int, y int) int {
	s := int64(x) - int64(y)
	if s < int64(MinInt) || s > int64(MaxInt) {
		panic("overflow: sub")
	}
	return int(s)
}

// addSub implement add and subtract using algorithm described in Knuth 4.5.1.
func addSub(f *Fraction, ff *Fraction, isAdd bool) *Fraction {
	// Zero is identity for addition.
	if f.numerator == 0 {
		if isAdd {
			return ff
		}
		return ff.Negate()
	}
	if ff.numerator == 0 {
		return f
	}
	// if denominators are randomly distributed, d1 will be 1 about 61%
	// of the time.
	d1 := greatestCommonDivisor(f.denominator, ff.denominator)
	if d1 == 1 {
		// result is ( (u*v' +/- u'v) / u'v')
		uvp := mulAndCheck(f.numerator, ff.denominator)
		upv := mulAndCheck(ff.numerator, f.denominator)
		if isAdd {
			return NewFraction(addAndCheck(uvp, upv), mulPosAndCheck(f.denominator, ff.denominator))
		}
		return NewFraction(subAndCheck(uvp, upv), mulPosAndCheck(f.denominator, ff.denominator))
	}
	// the quantity 't' requires 65 bits of precision; see knuth 4.5.1
	// exercise 7.  we're going to use a BigInteger.
	// t = u(v'/d1) +/- v(u'/d1)
	uvpn := big.NewInt(int64(f.numerator))
	uvpd := big.NewInt(int64(ff.denominator))
	uvp := uvpn.Mul(uvpn, uvpd)

	upvn := big.NewInt(int64(ff.numerator))
	upvd := big.NewInt(int64(f.denominator))
	upv := upvn.Mul(upvn, upvd)

	var t *big.Int
	if isAdd {
		t = uvp.Add(uvp, upv)
	} else {
		t = uvp.Sub(uvp, upv)
	}
	// but d2 doesn't need extra precision because
	// d2 = gcd(t,d1) = gcd(t mod d1, d1)
	d1t := big.NewInt(int64(d1))
	tmodd1 := int(t.Mod(t, d1t).Int64())
	var d2 int
	if tmodd1 == 0 {
		d2 = d1
	} else {
		d2 = greatestCommonDivisor(tmodd1, d1)
	}

	// result is (t/d2) / (u'/d1)(v'/d2)
	w := t.Div(t, big.NewInt(int64(d2)))
	if w.BitLen() > 31 {
		panic("overflow: numerator too large after multiply")
	}
	return NewFraction(int(w.Int64()), mulPosAndCheck(f.denominator/d1, ff.denominator/d2))
}
