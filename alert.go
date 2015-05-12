// 
// Go Alert
// Copyright (c) 2015 Brian W. Wolter, All rights reserved.
// 
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
// 
//   * Redistributions of source code must retain the above copyright notice, this
//     list of conditions and the following disclaimer.
// 
//   * Redistributions in binary form must reproduce the above copyright notice,
//     this list of conditions and the following disclaimer in the documentation
//     and/or other materials provided with the distribution.
//     
//   * Neither the names of Brian W. Wolter nor the names of the contributors may
//     be used to endorse or promote products derived from this software without
//     specific prior written permission.
//     
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
// BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
// LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
// OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
// OF THE POSSIBILITY OF SUCH DAMAGE.
// 

package alt

import (
  "fmt"
  "log"
  "flag"
  "github.com/kisielk/raven-go/raven"
)

var debug bool = false
var sentryDSN string
var sentryName string
var sentry *raven.Client

/**
 * Init
 */
func init() {
  var err error
  
  flag.BoolVar  (&debug,      "debug",        false,    "Enable debugging mode.")
  flag.StringVar(&sentryDSN,  "sentry:dsn",   "",       "Log errors to Sentry at the specified API key/DSN (e.g., https://ABC:XYZ@app.getsentry.com/12345).")
  flag.StringVar(&sentryName, "sentry:name",  "main",   "Identify this logger Sentry with the specified name.")
  
  if sentryDSN != "" {
    sentry, err = raven.NewClient(sentryDSN)
    if err != nil {
      panic(err)
    }
  }
  
}

/**
 * Log for debugging
 */
func Debugf(f string, a ...interface{}) {
  if debug {
    log.Printf(f, a...)
  }
}

/**
 * Log for debugging
 */
func Debug(m string) {
  if debug {
    log.Print(m)
  }
}

/**
 * Log to sentry
 */
func Errorf(f string, a ...interface{}) {
  Error(fmt.Sprintf(f, a...))
}

/**
 * Log to sentry
 */
func Error(m string) {
  log.Println(m)
  if sentry != nil {
    sentry.Capture(&raven.Event{Message: m, Logger:sentryName})
  }
}
