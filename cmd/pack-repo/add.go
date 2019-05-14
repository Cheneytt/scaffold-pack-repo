package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"scaffold/pkg/scaffold/pack/repo"
	"scaffold/pkg/scaffold/pack/repo/installer"
	"scaffold/pkg/scaffold/scaffoldpath"
)

type addCmd struct {
	out     io.Writer
	err     io.Writer
	source  string
	home    scaffoldpath.Home
	version string
}

func newAddCmd(out io.Writer) *cobra.Command {
	add := &addCmd{out: out}

	cmd := &cobra.Command{
		Use:   "add [flags] <path|url>",
		Short: "add a pack repository",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return add.complete(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return add.run()
		},
	}
	cmd.Flags().StringVar(&add.version, "version", "", "specify a version constraint. If this is not specified, the latest version is installed")
	return cmd
}

func (a *addCmd) complete(args []string) error {
	if err := validateArgs(args, []string{"path|url"}); err != nil {
		return err
	}
	a.source = args[0]
	home := homePath()
	if home == "" {
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		home = path
	}

	a.home = scaffoldpath.Home(home)
	debug("home path: %s", a.home)
	return nil
}

func (a *addCmd) run() error {
	fmt.Fprintf(a.out, "Installing pack repo from %s\n", a.source)

	ins, err := installer.New(a.source, a.version, a.home)
	if err != nil {
		return err
	}

	if err := installer.Install(ins); err != nil {
		return err
	}

	var installedRepo repo.Repository
	repos := repo.FindRepositories(a.home.Packs())
	for i := range repos {
		if repos[i].Dir == ins.Path() {
			installedRepo = repos[i]
		}
	}

	fmt.Fprintf(a.out, "Installed pack repository %s\n", installedRepo.Name)
	return nil
}
