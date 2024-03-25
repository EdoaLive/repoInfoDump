package repoInfoDump

import (
	"errors"
	"runtime/debug"
	"slices"
)

type Values struct {
	Name     string
	Commit   string
	Time     string
	Modified bool
}

func GetString() string {
	v, _ := New()
	return v.String()
}

func New() (Values, error) {
	var v Values

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return v, errors.New("can't read build info")
	}

	v.Name = bi.Path

	if rev, ok := getBuildSettingValue(bi.Settings, "vcs.revision"); ok {
		v.Commit = rev
	}
	if vcsTime, ok := getBuildSettingValue(bi.Settings, "vcs.time"); ok {
		v.Time = vcsTime
	}
	if vcsMod, ok := getBuildSettingValue(bi.Settings, "vcs.modified"); ok && vcsMod == "true" {
		v.Modified = true
	}
	return v, nil
}

func (v Values) String() string {
	info := "Starting " + v.Name
	if v.Commit != "" {
		info += " commit: " + v.Commit[:7]
	}
	if v.Time != "" {
		info += " (" + v.Time + ")"
	}
	if v.Modified {
		info += " (REPO MODIFIED)"
	}
	return info
}

func getBuildSettingValue(settings []debug.BuildSetting, key string) (string, bool) {
	idx := slices.IndexFunc(settings, func(c debug.BuildSetting) bool { return c.Key == key })
	if idx == -1 {
		return "", false
	}
	return settings[idx].Value, true
}
