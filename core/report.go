package core

type Report []*ReportEntry

type ReportEntry struct {
	Path string `json:"path"`
	Line int    `json:"line"`
}
