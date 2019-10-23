package elfinder

const (
	APIVERSION    = 2.1050
	UPLOADMAXSIZE = "10M"
)

const (
	defaultZipMaxSize = 1024 * 1024 * 1024 // 1G
	defaultTmpPath    = "/tmp"
)

type Option struct {
	zipMaxSize int64
	zipTmpPath string
}
