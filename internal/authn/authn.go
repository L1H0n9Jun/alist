package authn

import (
	"fmt"
	"net/http"
	"net/url"

	"alist/internal/conf"
	"alist/internal/setting"
	"alist/server/common"

	"github.com/go-webauthn/webauthn/webauthn"
)

func NewAuthnInstance(r *http.Request) (*webauthn.WebAuthn, error) {
	siteUrl, err := url.Parse(common.GetApiUrl(r))
	if err != nil {
		return nil, err
	}
	return webauthn.New(&webauthn.Config{
		RPDisplayName: setting.GetStr(conf.SiteTitle),
		RPID:          siteUrl.Hostname(),
		//RPOrigin:      siteUrl.String(),
		RPOrigins: []string{fmt.Sprintf("%s://%s", siteUrl.Scheme, siteUrl.Host)},
		// RPOrigin: "http://localhost:5173"
	})
}
