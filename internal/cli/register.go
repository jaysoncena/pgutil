package cli

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/sjansen/pgutil/internal/commands"
)

// RegisterCommands creates the app's standard ArgParser
func RegisterCommands(version string) *ArgParser {
	app := kingpin.
		New("pgutil", "Tools for PostgreSQL").
		UsageTemplate(kingpin.CompactUsageTemplate)

	base := &commands.Base{}
	parser := &ArgParser{
		app:     app,
		base:    base,
		version: version,
	}
	app.Flag("debug", "debug log").
		OpenFileVar(&base.Debug, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	app.Flag("quiet", "quiet mode, repeat to decrease verbosity").
		Short('q').CounterVar(&parser.brevity)
	app.Flag("verbose", "verbose mode, repeat to increase verbosity").
		Short('v').CounterVar(&parser.verbosity)

	registerPing(parser)
	registerInspect(parser)
	registerRunbook(parser)
	registerVersion(parser, version)
	return parser
}
