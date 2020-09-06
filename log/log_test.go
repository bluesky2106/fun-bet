package log

import (
	"testing"

	"github.com/bluesky2106/fun-bet/config"
	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	assert := assert.New(t)

	InitLogger(config.EnvDevelopment)
	assert.NotNil(GetLogger())
}
