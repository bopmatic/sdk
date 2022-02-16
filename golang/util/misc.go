package util

import (
	"context"
	"io"
	"os/exec"
)

func RunHostCommand(ctx context.Context, cmdAndArgs []string,
	stdOut io.Writer, stdErr io.Writer) error {

	cmd := exec.Command(cmdAndArgs[0], cmdAndArgs[1:]...)
	cmd.Stderr = stdErr
	cmd.Stdout = stdOut
	return cmd.Run()
}
