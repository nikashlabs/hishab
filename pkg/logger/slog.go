// Copyright (C) 2025 Nikash Labs
//
// This file is part of hishab.
//
// hishab is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// hishab is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with hishab.  If not, see <https://www.gnu.org/licenses/>.

package logger

import (
	"log/slog"
	"os"
)

type Slogger struct {
	log *slog.Logger
}

func NewSlogLogger() Logger {
	s := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return &Slogger{log: s}
}

func (s *Slogger) Debug(msg string, keysAndValues ...any) {
	s.log.Debug(msg, keysAndValues...)
}

func (s *Slogger) Info(msg string, keysAndValues ...any) {
	s.log.Info(msg, keysAndValues...)
}

func (s *Slogger) Warn(msg string, keysAndValues ...any) {
	s.log.Warn(msg, keysAndValues...)
}

func (s *Slogger) Error(msg string, keysAndValues ...any) {
	s.log.Error(msg, keysAndValues...)
}

func (s *Slogger) Fatal(msg string, keysAndValues ...any) {
	s.log.Error(msg, keysAndValues...)
	os.Exit(1)
}
