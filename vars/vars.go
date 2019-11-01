package vars

import "time"

const (
	DefaultPollInterval          = 900
	DefaultMaxConcurrentIndexers = 2
	DefaultPollEnabled           = true
	DefaultVcs                   = "git"
	DefaultBaseUrl               = "{url}/blob/master/{path}{anchor}"
	DefaultAnchor                = "#L{line}"
)

var (
	REPO_PATH    string
	MAX_INDEXERS int

	HTTP_HOST string
	HTTP_PORT int

	MAX_Concurrency_REPOS int

	DEBUG_MODE_SCAN bool

	DEBUG_MODE_WEB bool

	PAGE_SIZE = 10

	TIME_OUT time.Duration = 60 * 4

	Exts map[string]bool
)

