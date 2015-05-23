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
  "github.com/bww/raven-go/raven"
)

var config Config
var sentry *raven.Client

/**
 * Alert configuration
 */
type Config struct {
  Debug       bool
  SentryDSN   string
  Name        string
  Tags        map[string]interface{}  // tags sent with every event
}

/**
 * Init
 */
func Init(c Config) {
  var err error
  
  if c.SentryDSN != "" {
    if c.Name == "" {
      c.Name = "main"
    }
    sentry, err = raven.NewClient(c.SentryDSN)
    if err != nil {
      panic(err)
    }
  }
  
  config = c
}

/**
 * Log for debugging
 */
func Debugf(f string, a ...interface{}) {
  if config.Debug {
    log.Printf(f, a...)
  }
}

/**
 * Log for debugging
 */
func Debug(m string) {
  if config.Debug {
    log.Print(m)
  }
}

/**
 * Log to sentry
 */
func Errorf(f string, a ...interface{}) {
  Error(fmt.Sprintf(f, a...), nil, nil)
}

/**
 * Log to sentry
 */
func Error(m string, tags, extra map[string]interface{}) {
  log.Println(m)
  
  if tags != nil && len(tags) > 0 {
    var t string
    var i int
    for k, v := range tags {
      if i > 0 {
        t += fmt.Sprintf(", %s = %v", k, v)
      }else{
        t += fmt.Sprintf("%s = %v", k, v)
      }
      i++
    }
    log.Printf("  # %s", t)
  }
  
  if config.Tags != nil && len(config.Tags) > 0 {
    if tags == nil {
      tags = make(map[string]interface{})
    }
    for k, v := range config.Tags {
      tags[k] = v
    }
  }
  
  if extra != nil {
    var w, l int
    for k, _ := range extra {
      if l = len(k); l > w { w = l }
    }
    for k, v := range extra {
      log.Printf(fmt.Sprintf("  > %%%ds: %%v", w), k, v)
    }
  }
  
  if sentry != nil {
    sentry.Capture(&raven.Event{Message:m, Logger:config.Name, Tags:tags, Extra:extra})
  }
}
