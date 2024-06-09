package helper

var version string = "1.0.0"
var Application string = "Golang"

func SayHello(name string) string {
	return "Hello, " + name
}

func getVersion() string {
	return version
}

func Version() string {
	return "Version: " + getVersion() + " - " + Application
}
