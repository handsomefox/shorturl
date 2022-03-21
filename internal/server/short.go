package server

import (
	"fmt"
	"net/http"
	"shorturl/internal/storage"
	"shorturl/pkg/shortener"

	"github.com/gin-gonic/gin"
)

// ShortRouterPath is the path used for GET method
const ShortRouterPath = "/s/:link"

// ShortRouter - router itself
type ShortRouter struct {
	storage storage.Database
}

// Interface implementation

// UseStorage points to the File for storing the links
func (s *ShortRouter) UseStorage(storage storage.Database) {
	s.storage = storage
}

// Get returns the shortened link
func (s *ShortRouter) Get(c *gin.Context) {
	link := c.Param("link")

	short, full := shortener.Make(link)

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(formatLinkToHtml(full)))
	s.storage.Store(link, short)
}

// Just makes it look better in the browser
func formatLinkToHtml(link string) string {
	onclick := "<script>function copyElementText(id) {var text = document.getElementById(id).innerText;var elem = " +
		"document.createElement(\"textarea\");document.body.appendChild(elem);elem.value" +
		" = text;elem.select();document.execCommand(\"copy\");document.body.removeChild(elem);}" +
		"</script>"

	linkStyle := "\"text-align:center;position:fixed;top: 50%;left: 50%;margin-top: -50px;" +
		"margin-left: -150px;background-color: red;color: white;padding: 1em 1.5em;text-decoration: none;\""

	messageStyle := "\"text-align:center;position:fixed;top: 40%;left: 50%;margin-top: -50px;margin-left: -50px;\""

	return fmt.Sprintf("<p style=%s>Click to copy!</p><p id=\"text\" onclick=\"copyElementText(this.id)\" style=%s>%s</p> %s", messageStyle, linkStyle, link, onclick)
}
