package op

import (
	"github.com/coschain/contentos-go/dandelion"
	"testing"
)

func TestOperations(t *testing.T) {
	t.Run("transfer", dandelion.NewDandelionTest(new(TransferTester).Test, 3))
	t.Run("bp", dandelion.NewDandelionTest(new(BpTest).Test, 3))
}
