// Copyright 2015 Yury Kozyrev. All rights reserved.
package logger

import (
	"io"
	"os"
    "github.com/op/go-logging"
)

var (
	defaultFormat = "%{time:2006/01/02 15:04:05} %{level:.9s}: [%{shortfunc}] %{message}"
)

type Options struct {
	Out    io.Writer
	Format string
}

func (o *Options) applyDefault() {
	if nil == o.Out {
		o.Out = os.Stderr
	}
	if o.Format == "" {
		o.Format = defaultFormat
	}
}


func NewLogger(options *Options) (LoggerInterface, error) {
	options.applyDefault()
	log, _ := logging.GetLogger("logger")
	var format, err = logging.NewStringFormatter(options.Format)
	if err != nil {
		return nil, err
	}
	backend := logging.NewLogBackend(options.Out, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)

	return log, nil
}





