package repoInfoDump

import (
	"errors"
	"path"
	"runtime/debug"
	"slices"
)

type Values struct {
	Name     string
	Commit   string
	Time     string
	Modified bool
}

// Version can be set before calling these methods,
// or it can be set with build time parameter, e.g.
// -ldflags "-X github.com/EdoaLive/repoInfoDump.Version=v0.1.3"
var Version = ""

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

	v.Name = path.Base(bi.Main.Path)

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
	info := v.Name
	if Version != "" {
		info += " " + Version
	}
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
