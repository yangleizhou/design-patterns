package behavioral

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	var downloaderHTTP Downloader = newHTTPDownloader()
	downloaderHTTP.TemplateMethod("http://example.com/xxx.zip")

	fmt.Println()
	var downloaderFTP Downloader = newFTPDownloader()
	downloaderFTP.TemplateMethod("http://example.com/xxx.zip")

	fmt.Println()
	var downloaderP2P Downloader = newP2PDownloader()
	downloaderP2P.TemplateMethod("http://example.com/xxx.zip")
}
