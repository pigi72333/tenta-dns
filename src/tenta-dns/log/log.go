/**
 * Tenta DNS Server
 *
 *    Copyright 2017 Tenta, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For any questions, please contact developer@tenta.io
 *
 * log.go: Helper functions for logging
 */

package log

import (
	"fmt"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// EventualLogger -- a structure to buffer log entries whose usefulnes cannot be determined at creation time (but only eventually)
type EventualLogger []string

// Queuef -- buffers the interpreted message to be written later
func (l *EventualLogger) Queuef(format string, args ...interface{}) {
	interpreted := fmt.Sprintf(format, args...)
	if *l == nil {
		*l = make(EventualLogger, 0)
	}
	*l = append(*l, interpreted)
}

// Flush -- writes out everything from buffer
func (l *EventualLogger) Flush(target *logrus.Entry) {
	for _, e := range *l {
		//target.Infof(e)
		fmt.Printf("%s", e)
	}
}

var log *logrus.Logger = logrus.New()

func init() {
	log.Level = logrus.PanicLevel
	log.Out = colorable.NewColorableStdout()
	formatter := &prefixed.TextFormatter{ForceColors: true, ForceFormatting: true}
	formatter.SetColorScheme(&prefixed.ColorScheme{DebugLevelStyle: "green+b", InfoLevelStyle: "green+h"})
	log.Formatter = formatter
	// TODO: Deal with how to log to files or something
}

func SetLogLevel(lvl logrus.Level) {
	log.Level = lvl
}

func GetLogger(pkg string) *logrus.Entry {
	return log.WithField("prefix", pkg)
}
