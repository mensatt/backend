package utils

import (
	"errors"
	"runtime/debug"
)

type VCSBuildInfo struct {
	CommitTime string
	Commit     string
	Modified   string
}

func GetVscBuildInfo() (*VCSBuildInfo, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, errors.New("failed to read build info")
	}
	var vscInfo VCSBuildInfo
	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			vscInfo.Commit = setting.Value
		case "vcs.modified":
			vscInfo.Modified = setting.Value
		case "vcs.time":
			vscInfo.CommitTime = setting.Value
		}
	}

	if vscInfo.Commit == "" {
		return nil, errors.New("failed to read vsc.revision")
	}
	if vscInfo.CommitTime == "" {
		return nil, errors.New("failed to read vsc.time")
	}
	if vscInfo.Modified == "" {
		return nil, errors.New("failed to read vsc.modified")
	}

	return &vscInfo, nil
}
