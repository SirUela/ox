package test

import (
	"context"

	plugins "github.com/wawandco/ox/plugins/core"
)

// BeforeTester interface is suited for those tasks that need to happen
// before the tests run, things like setting up environment variables,
// clearing the database or other cleanup tasks.
type BeforeTester interface {
	plugins.Plugin

	RunBeforeTest(context.Context, string, []string) error
}