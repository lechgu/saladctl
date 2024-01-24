package main

import (
	"lechgu/saladctl/internal/commands"
	"lechgu/saladctl/internal/config"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/instances"
	"lechgu/saladctl/internal/jobs"
	"lechgu/saladctl/internal/organizations"
	"lechgu/saladctl/internal/projects"
	"lechgu/saladctl/internal/queues"
	"lechgu/saladctl/internal/sessions"
	"lechgu/saladctl/internal/workloads"

	"github.com/samber/do"
)

func main() {
	di.Injector = do.New()
	do.Provide(di.Injector, config.New)
	do.Provide(di.Injector, sessions.New)
	do.Provide(di.Injector, organizations.NewController)
	do.Provide(di.Injector, projects.NewController)
	do.Provide(di.Injector, workloads.NewController)
	do.Provide(di.Injector, queues.NewController)
	do.Provide(di.Injector, instances.NewController)
	do.Provide(di.Injector, jobs.NewController)
	commands.Execute()
}
