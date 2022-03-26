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
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogging(t *testing.T) {
	Convey("Level methods work", t, func() {
		So(Debug.String(), ShouldEqual, "debug")
		So(Info.String(), ShouldEqual, "info")
		So(Warning.String(), ShouldEqual, "warning")
		So(Error.String(), ShouldEqual, "error")
	})

	Convey("Top-level logging methods work", t, func() {
		ctx := context.Background()
		Convey("Default Null Logger", func() {
			So(func() { Debugf(ctx, "foo %s", "bar") }, ShouldNotPanic)
			So(func() { Infof(ctx, "foo %s", "bar") }, ShouldNotPanic)
			So(func() { Warningf(ctx, "foo %s", "bar") }, ShouldNotPanic)
			So(func() { Errorf(ctx, "foo %s", "bar") }, ShouldNotPanic)
		})
	})
}
