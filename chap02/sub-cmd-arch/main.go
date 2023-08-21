package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/hugo/go-hands-on/chap02/sub-cmd-arch/cmd"
)

var errInvalidSubCommand = errors.New("Invalid sub-command specified")

func prinfUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: mync [http|grpc] -h\n")
	_ = cmd.HandleHttp(w, []string{"-h"})
	_ = cmd.HandleGrpc(w, []string{"-h"})
}

func handleCommand(w io.Writer, args []string) error {
	var err error
	if len(args) < 1 {
		err = errInvalidSubCommand
	} else {
		switch args[0] {
		case "http":
			err = cmd.HandleHttp(w, args[1:])
		case "grpc":
			err = cmd.HandleGrpc(w, args[1:])
		case "-h":
			prinfUsage(w)
		case "-help":
			prinfUsage(w)
		default:
			err = errInvalidSubCommand
		}
	}

	if errors.Is(err, cmd.ErrNoServerSpecified) || errors.Is(err, errInvalidSubCommand) {
		fmt.Fprintln(w, err)
		prinfUsage(w)
	}
	return err
}

func main() {
	err := handleCommand(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
