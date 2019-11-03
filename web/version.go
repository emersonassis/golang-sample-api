package web

//Version ...
type Version struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
	Patch string `json:"patch"`
}

//VersionRequest ...
type VersionRequest struct {
	Version
}

//VersionResponse ...
type VersionResponse struct {
	ResponseBodyJSONDefault
	Version *Version `json:"version"`
}

//VersionsResponse ...
type VersionsResponse struct {
	ResponseBodyJSONDefault
	Versions []*Version `json:"versions"`
}
