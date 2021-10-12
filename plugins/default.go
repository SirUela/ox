package plugins

import (
	"github.com/wawandco/ox/plugins/base"
	"github.com/wawandco/ox/plugins/tools/buffalo/action"
	"github.com/wawandco/ox/plugins/tools/buffalo/app"
	"github.com/wawandco/ox/plugins/tools/buffalo/assets"
	"github.com/wawandco/ox/plugins/tools/buffalo/cmd"
	"github.com/wawandco/ox/plugins/tools/buffalo/config"
	"github.com/wawandco/ox/plugins/tools/buffalo/embedded"
	"github.com/wawandco/ox/plugins/tools/buffalo/middleware"
	"github.com/wawandco/ox/plugins/tools/buffalo/model"
	"github.com/wawandco/ox/plugins/tools/buffalo/render"
	"github.com/wawandco/ox/plugins/tools/buffalo/resource"
	"github.com/wawandco/ox/plugins/tools/buffalo/template"
	"github.com/wawandco/ox/plugins/tools/db"
	"github.com/wawandco/ox/plugins/tools/docker"
	"github.com/wawandco/ox/plugins/tools/envy"
	"github.com/wawandco/ox/plugins/tools/flect"
	"github.com/wawandco/ox/plugins/tools/git"
	"github.com/wawandco/ox/plugins/tools/grift"
	"github.com/wawandco/ox/plugins/tools/node"
	"github.com/wawandco/ox/plugins/tools/ox"
	"github.com/wawandco/ox/plugins/tools/refresh"
	"github.com/wawandco/ox/plugins/tools/soda"
	"github.com/wawandco/ox/plugins/tools/soda/fizz"
	"github.com/wawandco/ox/plugins/tools/soda/sql"
	"github.com/wawandco/ox/plugins/tools/standard"
	"github.com/wawandco/ox/plugins/tools/webpack"
	"github.com/wawandco/ox/plugins/tools/yarn"
)

// Default plugins for applications base. While ox
// has other plugins this list is the base that is used across most of
// the apps we build and maintain.
var Default = append(base.Plugins,
	&webpack.Plugin{},
	&refresh.Plugin{},
	&yarn.Plugin{},
	&envy.Developer{},
	&db.CreateCommand{},
	&db.DropCommand{},
	&db.ResetCommand{},

	// Application base commands.
	&db.Command{},
	&grift.Command{},

	// Builders
	&node.Builder{},
	&standard.Builder{},

	// Fixers
	&standard.Fixer{},

	// Generators
	&ox.Generator{},
	&template.Generator{},
	&model.Generator{},
	&action.Generator{},
	&resource.Generator{},
	&grift.Generator{},
	&soda.Generator{},

	// Initializer
	&embedded.Initializer{},
	&model.Initializer{},
	&render.Initializer{},
	&refresh.Initializer{},
	&template.Initializer{},
	&flect.Initializer{},
	&docker.Initializer{},
	&action.Initializer{},
	&middleware.Initializer{},
	&cmd.Initializer{},
	&config.Initializer{},
	&docker.Initializer{},
	&app.Initializer{},
	&standard.Initializer{},
	&grift.Initializer{},
	&assets.Initializer{},
	&soda.Initializer{},
	&git.Initializer{},

	&standard.AfterInitializer{},
	&yarn.AfterInitializer{},
	&git.AfterInitializer{},

	// Testers
	&standard.Tester{},
	&envy.Tester{},

	// Migration Creators
	&fizz.Creator{},
	&sql.Creator{},

	// Aftergenerators
	&standard.GoModAfterGenerator{},
)