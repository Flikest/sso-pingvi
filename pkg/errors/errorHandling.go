package errors

import "log/slog"

func FailOnError(err error, msg string) {
	if err != nil {
		slog.Debug(msg, err)
	}
}
