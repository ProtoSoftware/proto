package backend

import "os"

// Detects if the binary is a preview version. (Identified by Version string)
func IsPreview() bool {
	Debug("IsPreview: Checking for preview version, detected: " + Version)
	return Version == "0.0.1-next" && os.Getenv("PROTO_PREVIEW_CONSENT") != "true"
}
