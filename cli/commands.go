package cli

import "github.com/urfave/cli/v2"

var (
	commands = []*cli.Command{
		{
			Name:      "ls",
			Usage:     "List installed versions",
			UsageText: "grvm ls",
			Action:    list,
		},
		/* {
			Name:      "ls-all",
			Usage:     "List All versions",
			UsageText: "j ls-all",
			Action:    listAll,
		}, */
		{
			Name:      "ls-remote",
			Usage:     "List Remote versions",
			UsageText: "grvm ls-remote",
			Action:    listRemote,
		},
		{
			Name:      "install",
			Usage:     "install versions",
			UsageText: "grvm install <version>",
			Action:    install,
		},
		{
			Name:      "use",
			Usage:     "Switch to specified version",
			UsageText: "grvm use <version>",
			Action:    use,
		},
		{
			Name:      "uninstall",
			Usage:     "Uninstall a version",
			UsageText: "grvm uninstall <version>",
			Action:    uninstall,
		},
		{
			Name:      "clean",
			Usage:     "Remove files from the package download directory",
			UsageText: "grvm clean",
			Action:    clean,
		},
	}
)
