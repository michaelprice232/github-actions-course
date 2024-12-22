package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_connectToDB(t *testing.T) {
	//t.Setenv("DB_HOSTNAME", "localhost")
	//t.Setenv("DB_USERNAME", "postgres")
	//t.Setenv("DB_PASSWORD", "password")

	connected, err := connectToDB()
	assert.NoError(t, err)
	assert.True(t, connected)
}
