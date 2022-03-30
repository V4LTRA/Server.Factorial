package model

func Factorial(number uint64) (res uint64, err error) {     
	res = factorial(number)
	return res, nil
}

func factorial(n uint64) uint64 {
	if n == 0 {
	  return 1
	}
	return n * factorial(n-1)
  }