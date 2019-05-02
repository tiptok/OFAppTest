package gitbeego

import (
	"testing"

	"time"

	"github.com/astaxie/beego/session"
)

func TestSession(t *testing.T) {
	sesConfig := &session.ManagerConfig{
		CookieName:      "globalid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}

	globalSes, _ := session.NewManager("file", sesConfig)
	go globalSes.GC()

    //globalSes.SessionStart(w,r)
	time.Sleep(time.Second * 60)
}
