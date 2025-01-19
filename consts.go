package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var (
	GuildID        = "147313959819542528"
	discordUsers   = make(map[string]bool)
	dbUsers        = make(map[string]bool)
	dbPath         string
	wLogger        *logrus.Entry
	mLogger        *logrus.Entry
	sLogger        *logrus.Entry
	dLogger        *logrus.Entry
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)
