package main

import (
	"io"

	"github.com/yoheimuta/protolint/internal/cmd"
	"github.com/yoheimuta/protolint/internal/osutil"
)

func Do(
	args []string,
	stdout io.Writer,
	stderr io.Writer,
) osutil.ExitCode {
	return cmd.Do(args, stdout, stderr)
}
