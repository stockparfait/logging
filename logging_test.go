// Copyright 2022 Stock Parfait

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

import (
	"bytes"
	"context"
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogging(t *testing.T) {
	Convey("Level methods work", t, func() {
		Convey("Set", func() {
			var l Level
			So(l.Set("debug"), ShouldBeNil)
			So(l, ShouldEqual, Debug)

			So(l.Set("info"), ShouldBeNil)
			So(l, ShouldEqual, Info)

			So(l.Set("warning"), ShouldBeNil)
			So(l, ShouldEqual, Warning)

			So(l.Set("error"), ShouldBeNil)
			So(l, ShouldEqual, Error)

			So(l.Set("wrong"), ShouldNotBeNil)
		})

		Convey("String", func() {
			So(Debug.String(), ShouldEqual, "debug")
			So(Info.String(), ShouldEqual, "info")
			So(Warning.String(), ShouldEqual, "warning")
			So(Error.String(), ShouldEqual, "error")
		})
	})

	Convey("Top-level logging methods work", t, func() {
		ctx := context.Background()
		Convey("Default Null Logger", func() {
			So(func() { Debugf(ctx, "foo %s", "bar") }, ShouldNotPanic)
			So(func() { Infof(ctx, "foo %s", "bar") }, ShouldNotPanic)
			So(func() { Warningf(ctx, "foo %s", "bar") }, ShouldNotPanic)
			So(func() { Errorf(ctx, "foo %s", "bar") }, ShouldNotPanic)
		})

		Convey("GoLogger", func() {
			Convey("all methods correctly", func() {
				var buf bytes.Buffer
				ctx = Use(ctx, GoLogger(Debug, log.New(&buf, "", 0)))

				Debugf(ctx, "%s log", "debug")
				So(buf.String(), ShouldEqual, "DEBUG: debug log\n")
				buf.Reset()

				Infof(ctx, "%s log", "info")
				So(buf.String(), ShouldEqual, "INFO: info log\n")
				buf.Reset()

				Warningf(ctx, "%s log", "warning")
				So(buf.String(), ShouldEqual, "WARNING: warning log\n")
				buf.Reset()

				Errorf(ctx, "%s log", "error")
				So(buf.String(), ShouldEqual, "ERROR: error log\n")
				buf.Reset()
			})

			Convey("level works correctly", func() {
				var buf bytes.Buffer
				ctx = Use(ctx, GoLogger(Warning, log.New(&buf, "", 0)))

				Debugf(ctx, "%s log", "debug")
				So(buf.Len(), ShouldEqual, 0)

				Infof(ctx, "log %s", "info")
				So(buf.Len(), ShouldEqual, 0)

				Warningf(ctx, "log %s %s", "warning", "attention")
				So(buf.String(), ShouldEqual, "WARNING: log warning attention\n")
				buf.Reset()

				Errorf(ctx, "%s log", "error")
				So(buf.String(), ShouldEqual, "ERROR: error log\n")
				buf.Reset()
			})
		})
	})
}
