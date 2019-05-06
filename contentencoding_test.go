package http_ece

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContentEncoding_Padding(t *testing.T) {
	assert.Equal(t, 2, AESGCM.Padding())
	assert.Equal(t, 1, AES128GCM.Padding())
	assert.Equal(t, 0, ContentEncoding(0).Padding())
	assert.Equal(t, 0, ContentEncoding(3).Padding())
}

func TestContentEncoding_String(t *testing.T) {
	assert.Equal(t, "AESGCM", AESGCM.String())
	assert.Equal(t, "AES128GCM", AES128GCM.String())
	assert.Equal(t, "ContentEncoding(0)", ContentEncoding(0).String())
	assert.Equal(t, "ContentEncoding(3)", ContentEncoding(3).String())
}