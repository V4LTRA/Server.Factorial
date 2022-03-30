package model_test

import (
	"testing"

	"github.com/gopherschool/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	var err error
	var res uint64
	res, err = model.Factorial(0)
	if assert.NoError(t, err) {
		assert.Equal(t, res, uint64(1))
	}
}
