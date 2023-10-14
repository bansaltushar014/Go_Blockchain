package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignAndVerifyTxn(t *testing.T) {
	tx := NewTransaction()
	assert.NotEmpty(t, tx)
	bool := tx.verify()
	assert.True(t, bool)
}

func TestCreateTxAndVerify(t *testing.T) {
	t1 := Trans{}
	txx := t1.CreateRandomTx([]byte("First Transaction"))
	assert.NotEmpty(t, txx)
	bool := txx.verify()
	assert.True(t, bool)

}
