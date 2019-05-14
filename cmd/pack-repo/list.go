package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"scaffold/pkg/scaffold/pack/repo"
	"scaffold/pkg/scaffold/scaffoldpath"
)

type listCmd struct {
	out  io.Writer
	home scaffoldpath.Home
}

func newListCmd(out io.Writer) *cobra.Command {
	list := &listCmd{out: out}

	cmd := &cobra.Command{
		Use:   "list [flags]",
		Short: "list all installed pack repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			list.home = scaffoldpath.Home(homePath())
			return list.run()
		},
	}
	return cmd
}

func (l *listCmd) run() error {
	repos := repo.FindRepositories(l.home.Packs())
	if len(repos) == 0 {
		return errors.New("no pack repositories to show")
	}
	for i := range repos {
		fmt.Fprintln(l.out, repos[i].Name)
	}
	return nil
}
