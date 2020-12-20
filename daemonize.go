// +build !windows

package main

import (
	daemon "github.com/ochinchina/go-daemon"

	"github.com/ochinchina/supervisord/log"
)

// Deamonize run this process in daemon mode
func Deamonize(proc func(), log log.Log) {
	context := daemon.Context{LogFileName: "/dev/stdout"}

	child, err := context.Reborn()
	if err != nil {
		context := daemon.Context{}
		child, err = context.Reborn()
		if err != nil {
			log.Fatalw("Unable to run", "error", err)
		}
	}
	if child != nil {
		return
	}
	defer context.Release()
	proc()
}
