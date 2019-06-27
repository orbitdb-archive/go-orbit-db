package http

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestBuildURL(t *testing.T) {
    url := BuildURL("http://", "localhost", ":3000")
    assert.Equal(t, "http://localhost:3000", url, "Expecting `http://localhost:3000`")
}
