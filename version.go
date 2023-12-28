package version

import (
	"fmt"
	"runtime"
)

var (
	Version   string
	CommitId  string
	BuildTime string
)

type VersionInfo struct {
	BuildTime string
	Version   string
	Os        string
	Arch      string
	GoVersion string
	CommitId  string
}

func String() string {
	format := "\tBuild Time: %s \n\tGO Version:%s \n\tGit Hash:%s "
	return fmt.Sprintf(format, BuildTime, "1", CommitId)
}

func GetVersion() VersionInfo {
	return VersionInfo{
		BuildTime: BuildTime,
		Version:   Version,
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		GoVersion: runtime.Version(),
		CommitId:  CommitId,
	}
}
