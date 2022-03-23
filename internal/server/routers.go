package server

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"net/url"
	"shorturl/internal/storage"
	"shorturl/pkg/shortener"
	"strings"

	"github.com/gin-gonic/gin"
)

// Route interface describes the functions that routers have, for now
type Route interface {
	Get(c *gin.Context)
}

// ShortRouterPath is the path used for GET method
const ShortRouterPath = "/s/:link"

// UnrollRouterPath is the path used for GET method
const UnrollRouterPath = "/u/:link"

// ShortRouter - router itself
type ShortRouter struct {
	storage storage.Model
}

// UnrollRouter router itself
type UnrollRouter struct {
	storage storage.Model
}

// Interface implementation

// Get returns the shortened link
func (s *ShortRouter) Get(c *gin.Context) {
	str := s.prepareLink(c.Param("link"))

	if !isValidURL(str) {
		handleError(c, http.StatusBadRequest)
		return
	}

	if isDuplicate(s.storage, str) {
		link, err := getExistingShortLink(s.storage, str)
		if err != nil {
			handleError(c, http.StatusNotFound)
			return
		}
		s.sendResponse(c, s.appendHostname(c, link))
		return
	}

	responseLink, err := s.makeAndStore(c, str)
	if err != nil {
		handleError(c, http.StatusBadGateway)
		return
	}

	s.sendResponse(c, responseLink)
}

func (s *ShortRouter) prepareLink(link string) string {
	link = strings.ReplaceAll(link, "|", "/")
	link = strings.ReplaceAll(link, "{", "?")
	link = strings.ReplaceAll(link, "}", "=")
	link = strings.ReplaceAll(link, "[", "&")
	return link
}

func (s *ShortRouter) makeAndStore(c *gin.Context, link string) (string, error) {
	hash := shortener.Make(link)
	full := s.appendHostname(c, hash)
	err := s.storage.Store(hash, full)
	return full, err
}

func (s *ShortRouter) appendHostname(c *gin.Context, hash string) string {
	return c.Request.Host + "/u/" + hash
}

func (s *ShortRouter) sendResponse(c *gin.Context, full string) {
	data := fmt.Sprintf("{\"link\": \"%s\"}", full)
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(data))
}

// Get redirects you to the original link
func (s *UnrollRouter) Get(c *gin.Context) {
	link := c.Param("link")

	if !isValidURL(addSchemePrefixIfNeeded(link)) {
		handleError(c, http.StatusBadRequest)
		return
	}

	found, err := s.storage.Get(bson.M{"short": link})
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	found.URL = addSchemePrefixIfNeeded(found.URL)

	c.Redirect(http.StatusPermanentRedirect, found.URL)
	c.Abort()
}

func addSchemePrefixIfNeeded(str string) string {
	if !strings.HasPrefix(str, "http://") || !strings.HasPrefix(str, "https://") {
		str = "https://" + str
	}

	return str
}

func handleError(c *gin.Context, HTTPStatus int) {
	c.AbortWithStatus(HTTPStatus)
	c.Done()
}

func isDuplicate(s storage.Model, link string) bool {
	return s.Contains(link)
}

func getExistingShortLink(s storage.Model, link string) (string, error) {
	ent, err := s.Get(bson.M{"url": link})
	return ent.Short, err
}

func isValidURL(link string) bool {
	u, err := url.Parse(link)
	return err == nil && (u.Host != "" || u.Path != "")
}
