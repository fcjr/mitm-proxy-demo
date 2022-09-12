package proxy

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fcjr/mitm-proxy-demo/utils"
	"github.com/google/martian/v3"
	martianlog "github.com/google/martian/v3/log"
	"github.com/google/martian/v3/mitm"
)

func Serve(port int, authorityName string) {
	martianlog.SetLevel(martianlog.Silent)

	// listen proxy
	l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		log.Fatalf(err.Error())
	}

	crt, privKey, err := mitm.NewAuthority(authorityName, fmt.Sprintf("The %s Company", authorityName), 365*24*time.Hour)
	if err != nil {
		log.Fatalf(err.Error())
	}

	crtPath := "./crt.pem"
	if err = utils.WriteCertToFile(crt, crtPath); err != nil {
		log.Fatalf(err.Error())
	}
	if err := utils.InstallCert(crtPath); err != nil {
		log.Fatalf(err.Error())
	}

	mitmConf, err := mitm.NewConfig(crt, privKey)
	mitmConf.SetOrganization(authorityName)
	if err != nil {
		utils.UninstallCert(authorityName)
		log.Fatalf(err.Error())
	}

	logger := NewRequestLogger()

	proxy := martian.NewProxy()
	proxy.SetMITM(mitmConf)
	proxy.SetRequestModifier(logger)

	fmt.Printf("MITM Proxy Demo listening on: %s", l.Addr().String())
	if err := proxy.Serve(l); err != nil {
		utils.UninstallCert(authorityName)
		log.Fatalf(err.Error())
	}
}
