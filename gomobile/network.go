// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin linux windows

// An app that paints green if golang.org is reachable when the app first
// starts, or red otherwise.
//
// In order to access the network from the Android app, its AndroidManifest.xml
// file must include the permission to access the network.
//
//   http://developer.android.com/guide/topics/manifest/manifest-intro.html#perms
//
// The gomobile tool auto-generates a default AndroidManifest file by default
// unless the package directory contains the AndroidManifest.xml. Users can
// customize app behavior, such as permissions and app name, by providing
// the AndroidManifest file. This is irrelevent to iOS.
//
// Note: This demo is an early preview of Go 1.5. In order to build this
// program as an Android APK using the gomobile tool.
//
// See http://godoc.org/golang.org/x/mobile/cmd/gomobile to install gomobile.
//
// Get the network example and use gomobile to build or install it on your device.
//
//   $ go get -d golang.org/x/mobile/example/network
//   $ gomobile build golang.org/x/mobile/example/network # will build an APK
//
//   # plug your Android device to your computer or start an Android emulator.
//   # if you have adb installed on your machine, use gomobile install to
//   # build and deploy the APK to an Android target.
//   $ gomobile install golang.org/x/mobile/example/network
//
// Switch to your device or emulator to start the network application from
// the launcher.
// You can also run the application on your desktop by running the command
// below. (Note: It currently doesn't work on Windows.)
//   $ go install golang.org/x/mobile/example/network && network
package main

import (
	"net/http"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/gl"
)

func main() {
	// checkNetwork runs only once when the app first loads.
	//Keyword go, starts a goroutine(a lightweight thread of execution.), which is managed by golang run-time.
	go checkNetwork()

	// In summary this App draws a green screen if you can connect to http://golang.org/ else prints a red screen.
	app.Main(func(a app.App) {
		//Open graphics context
		var glctx gl.Context
		det, sz := determined, size.Event{}

		//(If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.)
		for {
			//The select statement lets a goroutine wait on multiple communication operations.
			//A select blocks until one of its cases can run, then it executes that case.
			select {
			case <-det:
				a.Send(paint.Event{})
				det = nil
			// Here are the different events that can occur in our app; There is the lifecyle event, paint, and size events
			case e := <-a.Events():
				switch e := a.Filter(e).(type) {
				case lifecycle.Event:
					glctx, _ = e.DrawContext.(gl.Context)
				case size.Event:
					sz = e
				case paint.Event:
					if glctx == nil {
						continue
					}
					onDraw(glctx, sz)
					a.Publish()
				}
			}
		}
	})
}

var (
	//make - allows to create dynamically-sized arrays
	//Channels are a typed conduit through which you can send and receive values with the channel operator, <-.(Data flows the direction of the arrow)
	determined = make(chan struct{})
	ok         = false
)

//Func to check network.
//If can connect to http://golang.org/ set ok to true.
func checkNetwork() {
	defer close(determined)

	//_ is a blank identifier. It avoids having to declare all the variables for the returns values. (Get Method returns two values)
	_, err := http.Get("http://golang.org/")
	if err != nil {
		return
	}
	ok = true
}

//Func to set and clear the color
func onDraw(glctx gl.Context, sz size.Event) {
	select {
	case <-determined:
		if ok {
			//ClearColor(red, green, blue, alpha float32)
			glctx.ClearColor(0, 1, 0, 1)
		} else {
			glctx.ClearColor(1, 0, 0, 1)
		}
	default:
		glctx.ClearColor(0, 0, 0, 1)
	}
	glctx.Clear(gl.COLOR_BUFFER_BIT)
}
