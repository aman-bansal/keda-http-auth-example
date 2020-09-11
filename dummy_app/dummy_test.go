package main_test

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	fmt.Print(base64.StdEncoding.EncodeToString([]byte("basicAuth")))
}
