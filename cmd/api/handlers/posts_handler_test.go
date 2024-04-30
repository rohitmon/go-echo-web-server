package handlers

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndtoEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndtoEndSuite))
}

func (s *EndtoEndSuite) TestPostHandler() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:1323/posts")
	s.Equal(http.StatusOK, r.StatusCode)
}

func (s *EndtoEndSuite) TestPostHandlerNoResult() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:1323/posts/5678")
	s.Equal(http.StatusOK, r.StatusCode)
	b, _ := ioutil.ReadAll(r.Body)
	s.JSONEq(`{"status": 200, "data": []}`, string(b))
}
