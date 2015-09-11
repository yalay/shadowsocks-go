package util

import (
	"testing"
)

func TestCheckHostValid(t *testing.T) {
	host := "baidu.com"
	if !CheckHostValid(host) {
		t.Errorf("Check host failed:%s", host)
	}

	host = "www.baidu.com"
	if !CheckHostValid(host) {
		t.Errorf("Check host failed:%s", host)
	}

	host = "www.bai-du.com"
	if !CheckHostValid(host) {
		t.Errorf("Check host failed:%s", host)
	}

	host = "dd"
	if CheckHostValid(host) {
		t.Errorf("Check host failed:%s", host)
	}

}
