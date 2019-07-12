package deadbolt

var Version string

func GetVersion() string {
	if Version == "" {
		return "version defined in ./script/release build process"
	}
	return Version
}
