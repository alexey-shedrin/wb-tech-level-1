package main

import (
	"strings"
)

var justString string

/*
В изначальном варианте (justString = v[:100]) justString будет ссылаться
на тот же участок памяти, что и v. При этом после отработки someFunc() память,
занимаемая v, не будет освобождена, так как justString - глобальная
переменная и она продолжает ссылаться на этот участок памяти. GC не может
освободить только лишнюю часть памяти. В исправленном варианте justString
будет ссылаться на новый участок памяти, и память, занимаемая v, будет освобождена.
*/
func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	return strings.Repeat("s", size)
}
