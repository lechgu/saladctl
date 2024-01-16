package main

import (
	"fmt"
	"lechgu/saladctl/internal/config"
	"lechgu/saladctl/internal/sessions"

	"github.com/samber/do"
)

func main() {
	di := do.New()
	do.Provide(di, config.New)
	do.Provide(di, sessions.New)
	sess := do.MustInvoke[*sessions.Session](di)
	fmt.Println(sess.Client.Jar)
}
