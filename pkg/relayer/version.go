package relayer

const (
	Version           = "0.1.0-rc.1"
	VersionMajor      = 0
	VersionMinor      = 1
	VersionPatch      = 0
	VersionPrerelease = "rc.1"
)

// GetVersion returns the SDK version string.
func GetVersion() string {
	return Version
}
