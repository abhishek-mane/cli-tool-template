package commands

type category struct {
	General string
	Git     string
	Docker  string
}

// Category :
var Category = category{
	General: "general",
	Git:     "git",
	Docker:  "docker",
}
