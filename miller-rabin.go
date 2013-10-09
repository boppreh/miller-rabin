package main

import (
    "fmt"
    "os"
    "time"
    "math/big"
    "math/rand"
)

func main() {
	// Simple shortcuts.
	one := big.NewInt(1)
	two := big.NewInt(2)

	// Read the candidate from the first argument, or default to 221 for test
	// purposes.
	candidate := new(big.Int)
	var value string
	if len(os.Args) > 1 {
		value = os.Args[1]
	} else {
		value = "221"
	}

	// Tries to interpret the string value into a big.Int object.
	_, err := fmt.Sscan(value, candidate)
	if err != nil {
		panic(err.Error())
	}

	modulo := new(big.Int)
	modulo.Sub(candidate, one)

	// Write the modulo (candidate -1) number in the form
	// 2^s * d.
	s := 0
	remainder := new(big.Int)
	quotient := new(big.Int)
	quotient.Set(modulo)

	for remainder.Sign() == 0 {
		quotient.DivMod(quotient, two, remainder)
		s += 1
	}
	// The last division failed, so we must decrement `s`.
	s.Sub(s, one)
	// quotient here contains the leftover which we could not divide by two,
	// and we have a 1 remaining from this last division. 
	d := big.NewInt(1)
	d.Add(one, d.Mul(two, quotient))

	// Random number source for generating witnesses.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Here 10 is the precision. Every increment to this value decreases the
	// chance of a false positive by 3/4.
	for k := 0; k < 10; k++ {

		// Every witness may prove that the candidate is composite, or assert
		// nothing.
		witness := new(big.Int)
		witness.Rand(r, modulo)

		exp := new(big.Int)
		exp.Set(d)
		generated := new(big.Int)
		generated.Exp(witness, exp, candidate)

		if generated.Cmp(modulo) == 0 || generated.Cmp(one) == 0 {
			continue
		}

		for i := 1; i < s; i++ {
			generated.Exp(generated, two, candidate)

			if generated.Cmp(one) == 0 {
				fmt.Println("Composite.")
				return
			}

			if generated.Cmp(modulo) == 0 {
				break
			}
		}


		if generated.Cmp(modulo) != 0 {
			// We arrived here because the `i` loop ran its course naturally
			// without meeting the `x == modulo` break.
			fmt.Println("Composite.")
			return
		}
	}

	fmt.Println("Prime.")
}
