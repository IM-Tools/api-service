package helpers

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestUid(t *testing.T) {
	u1 := uuid.NewV4()

	fmt.Println(u1)
}
