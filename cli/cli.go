package cli

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/Masterminds/semver"
	"github.com/fatih/color"
	"github.com/forget-the-bright/grvm/internal/build"
	"github.com/urfave/cli/v2"
)

var (
	ghomeDir     string
	downloadsDir string
	versionsDir  string
	goroot       string
)

func init() {
	cli.AppHelpTemplate = fmt.Sprintf(`NAME:
	{{.Name}}{{if .Usage}} - {{.Usage}}{{end}}

 USAGE:
	{{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .Commands}} command{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

 VERSION:
	%s{{end}}{{end}}{{if .Description}}

 DESCRIPTION:
	{{.Description}}{{end}}{{if len .Authors}}

 AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
	{{range $index, $author := .Authors}}{{if $index}}
	{{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}

 COMMANDS:{{range .VisibleCategories}}{{if .Name}}

	{{.Name}}:{{end}}{{range .VisibleCommands}}
	  {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

 GLOBAL OPTIONS:
	{{range $index, $option := .VisibleFlags}}{{if $index}}
	{{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}

 COPYRIGHT:
	{{.Copyright}}{{end}}
`, build.ShortVersion)
}

// Run 运行g命令行
func Run() {
	app := cli.NewApp()
	app.Name = "grvm"
	app.Usage = "Gradle Version Manager"
	app.Version = build.Version()
	app.Copyright = fmt.Sprintf("Copyright (c) 2019-%d, forget-the-bright. All rights reserved.", time.Now().Year())
	app.Authors = []*cli.Author{
		{Name: "wh", Email: "helloworldwh@163.com"},
	}

	app.Before = func(ctx *cli.Context) (err error) {
		ghomeDir = ghome()
		goroot = filepath.Join(ghomeDir, "gradle")
		downloadsDir = filepath.Join(ghomeDir, "downloads")
		if err = os.MkdirAll(downloadsDir, 0755); err != nil {
			return err
		}
		versionsDir = filepath.Join(ghomeDir, "versions")
		return os.MkdirAll(versionsDir, 0755)
	}
	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

const (
	homeEnv   = "GRVM_HOME"
	mirrorEnv = "GRVM_MIRROR"
)

// ghome 返回g根目录
func ghome() (dir string) {
	//fmt.Println(os.Getenv(homeEnv))
	/* path, _ := os.Getwd()
	return path */
	if dir = os.Getenv(homeEnv); dir != "" {
		return dir
	}
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".grvm")

}

// inuse 返回当前的go版本号
func inuse(goroot string) (version string) {
	p, _ := os.Readlink(goroot)
	return filepath.Base(p)
}

// render 渲染go版本列表
func render(curV string, items []*semver.Version, out io.Writer) {
	sort.Sort(semver.Collection(items))

	for i := range items {
		fields := strings.SplitN(items[i].String(), "-", 2)
		v := strings.TrimSuffix(strings.TrimSuffix(fields[0], ".0"), ".0")
		if len(fields) > 1 {
			v += fields[1]
		}
		if v == curV {
			color.New(color.FgGreen).Fprintf(out, "* %s\n", v)
		} else {
			fmt.Fprintf(out, "  %s\n", v)
		}
	}
}

// errstring 返回统一格式的错误信息
func errstring(err error) string {
	if err == nil {
		return ""
	}
	return wrapstring(err.Error())
}

func wrapstring(str string) string {
	if str == "" {
		return str
	}
	words := strings.Fields(str)
	if len(words) > 0 {
		words[0] = strings.Title(words[0])
	}
	return fmt.Sprintf("[grvm] %s", strings.Join(words, " "))
}

func mkSymlink(oldname, newname string) (err error) {
	if runtime.GOOS == "windows" {
		// Windows 10下无特权用户无法创建符号链接，优先调用mklink /j创建'目录联接'
		if err = exec.Command("cmd", "/c", "mklink", "/j", newname, oldname).Run(); err == nil {
			return nil
		}
	}
	return os.Symlink(oldname, newname)
}
