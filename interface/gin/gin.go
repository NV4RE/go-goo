package http_gin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"github.com/nv4re/go-goo/entity/errors"
)

type GinServer struct {
	host string
	port int
	gin  *gin.Engine
}

type GinResponse struct {
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"`
}

func (gs *GinServer) Start() error {
	sys := gs.gin.Group("/system")
	{
		sys.GET("/health", func(c *gin.Context) { systemHealth(c, gs) })
	}

	err := gs.gin.Run(fmt.Sprintf(`%s:%d`, gs.host, gs.port))

	if err != nil {
		return errors.DeserializationFailed
	}
	return nil
}

func NewGinServer(host string, port int) *GinServer {
	return &GinServer{
		host,
		port,
		gin.New(),
	}
}

func parseBody(c *gin.Context, data interface{}) error {
	rawJSONFlow, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return errors.DeserializationFailed
	}

	err = json.Unmarshal(rawJSONFlow, data)
	if err != nil {
		return errors.DeserializationFailed
	}
	return nil
}
