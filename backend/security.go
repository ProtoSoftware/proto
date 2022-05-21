package backend

import "os"

func IsPreview() bool {
	return Version == "0.0.1-next" && os.Getenv("PROTO_PREVIEW_CONSENT") != "true"
}
