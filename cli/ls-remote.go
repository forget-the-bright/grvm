package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/forget-the-bright/grvm/internal/pkg/collector"
	"github.com/k0kubun/go-ansi"
	"github.com/urfave/cli/v2"
)

func remoteVersionLength(version string) string {
	yu := 8 - len(version)
	for i := 0; i < yu; i++ {
		version += " "
	}
	return version
}

func remoteVersionLength2(version string) string {
	yu := 17 - len(version)
	for i := 0; i < yu; i++ {
		version += " "
	}
	return version
}

func listRemote(*cli.Context) (err error) {
	use_version := inuse(goroot)
	out := ansi.NewAnsiStdout()
	rs := collector.Collector.Items
	color.New(color.FgGreen).Fprintf(out, " %s\n", " version              info            RelaseTime")
	for _, v := range rs {
		if v.Version == use_version { //strings.Contains(v.SimpleName, version)
			color.New(color.FgGreen).Fprintf(out, "*  %s\n", remoteVersionLength(v.Version)+"      "+
				remoteVersionLength2(collector.GetFileNameNoSuffix(v.FileName))+"    "+v.ReleaseTime)
		} else {
			fmt.Fprintf(out, "   %s\n", remoteVersionLength(v.Version)+"      "+
				remoteVersionLength2(collector.GetFileNameNoSuffix(v.FileName))+"    "+v.ReleaseTime)
		}
	}
	return nil
}
