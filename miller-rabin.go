package main

import (
    "fmt"
    "os"
    "time"
    "math/big"
    "math/rand"
)

func main() {
	one := big.NewInt(1)
	two := big.NewInt(2)

	candidate := new(big.Int)
	var value string
	if len(os.Args) > 1 {
		value = os.Args[1]
	} else {
		value = "997"
	}

	_, err := fmt.Sscan(value, candidate)
	if err != nil {
		panic(err.Error())
	}

	modulo := new(big.Int)
	modulo.Sub(candidate, one)

	s := -1
	remainder := new(big.Int)
	quotient := new(big.Int)
	quotient.Set(modulo)

	for remainder.Sign() == 0 {
		quotient.DivMod(quotient, two, remainder)
		s += 1
	}

	d := big.NewInt(1)
	d.Add(one, d.Mul(two, quotient))
	fmt.Println(s)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for k := 0; k < 10; k++ {
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
			fmt.Println("Composite.")
			return
		}
	}

	fmt.Println("Prime.")
}
