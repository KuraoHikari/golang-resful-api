package test

import (
	"fmt"
	"golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T){
	simpleService, err := simple.InitializedService(true)
	fmt.Println(err)
	fmt.Println(simpleService)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)

	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}
func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(true)

	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}