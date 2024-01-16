package main

import (
	"lechgu/saladctl/internal/commands"
	"lechgu/saladctl/internal/config"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/organizations"
	"lechgu/saladctl/internal/sessions"

	"github.com/samber/do"
)

var Di *do.Injector

func main() {
	di.Injector = do.New()
	do.Provide(di.Injector, config.New)
	do.Provide(di.Injector, sessions.New)
	do.Provide(di.Injector, organizations.NewController)
	commands.Execute()
}
