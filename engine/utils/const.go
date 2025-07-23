package utils

import "github.com/logrusorgru/aurora"

const (
	PasswdPath = "/etc/passwd"
	ShadowPath = "/etc/shadow"
	RootPath   = "/"
	CronPath   = "/etc/crontab"
)

var (
	Sad = aurora.Red("SAD")
	Res = aurora.Green("YAY")
	Bsh = aurora.Cyan("CMD")
)
