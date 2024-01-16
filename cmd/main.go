package main

import (
	"lechgu/saladctl/internal/commands"
	"lechgu/saladctl/internal/config"
	"lechgu/saladctl/internal/containergroups"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/organizations"
	"lechgu/saladctl/internal/projects"
	"lechgu/saladctl/internal/queues"
	"lechgu/saladctl/internal/sessions"

	"github.com/samber/do"
)

func main() {
	di.Injector = do.New()
	do.Provide(di.Injector, config.New)
	do.Provide(di.Injector, sessions.New)
	do.Provide(di.Injector, organizations.NewController)
	do.Provide(di.Injector, projects.NewController)
	do.Provide(di.Injector, containergroups.NewController)
	do.Provide(di.Injector, queues.NewController)
	commands.Execute()
}
