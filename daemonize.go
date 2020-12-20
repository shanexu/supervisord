// +build !windows

package main

import (
	daemon "github.com/ochinchina/go-daemon"
	"go.uber.org/zap"
)

// Deamonize run this process in daemon mode
func Deamonize(proc func()) {
	context := daemon.Context{LogFileName: "/dev/stdout"}

	child, err := context.Reborn()
	if err != nil {
		context := daemon.Context{}
		child, err = context.Reborn()
		if err != nil {
			zap.S().Fatalw("Unable to run", "error", err)
		}
	}
	if child != nil {
		return
	}
	defer context.Release()
	proc()
}
