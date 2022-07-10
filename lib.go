/**
 * @Date 2022/7/10
 * @Name lib
 * @VariableName
**/
package httpserver

import (
	"math/rand"
	"time"
)

func RandIntTime(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
