/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2019 HereweTech Co.LTD
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file runtime.go
 * @package common
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 05/23/2019
 */

package common

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

// GlobalRuntime : Global runtime context
type GlobalRuntime struct {
	// Application : Application name
	Application string

	// Config : Viper config
	Config *viper.Viper

	// Waiter : Global waiter
	Waiter *sync.WaitGroup

	// Logger : Global logger
	Logger *log.Logger
}

// NewRuntime : Create a global runtime
func NewRuntime(application string) *GlobalRuntime {
	var (
		w sync.WaitGroup
		c = viper.New()
		l = log.New(os.Stdout, fmt.Sprintf("[%s] - ", application), log.LstdFlags|log.Ldate|log.Ltime)
	)

	c.SetEnvPrefix(application)
	c.AutomaticEnv()
	c.SetConfigName("config")
	c.SetConfigType("json")
	c.AddConfigPath(fmt.Sprintf("/etc/%s", application))
	c.AddConfigPath(fmt.Sprintf("$HOME/.%s", application))
	c.AddConfigPath(".")
	err := c.ReadInConfig()
	if err != nil {
		l.Printf("Load configuration file failed: %s\n", err)
	}

	return &GlobalRuntime{
		Application: application,
		Config:      c,
		Waiter:      &w,
		Logger:      l,
	}
}

// Wait : GlobalRuntime.Waiter wait
func (r *GlobalRuntime) Wait() {
	r.Waiter.Wait()
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
