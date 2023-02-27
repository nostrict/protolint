package main

import (
	"io"
	"os"

	"github.com/yoheimuta/protolint/internal/cmd"
	"github.com/yoheimuta/protolint/internal/osutil"
)

func main() {
	os.Exit(int(
		cmd.Do(
			os.Args[1:],
			os.Stdout,
			os.Stderr,
		),
	))
}

func Do(
	args []string,
	stdout io.Writer,
	stderr io.Writer,
) osutil.ExitCode {
	return cmd.Do(args, stdout, stderr)
}
