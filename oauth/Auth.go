package oauth

import (
	"fmt"
	"net/http"

	"github.com/realpamisa/gmail-extractor/decode"
	"github.com/realpamisa/gmail-extractor/internal"

	"github.com/realpamisa/gmail-extractor/pkg"
)

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := internal.GoogleOauthConfig.AuthCodeURL(internal.OauthStateString)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {

	content, err := internal.ExtractEmail(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	pkg.CsvSave(decode.AppendEmail, decode.C)
	fmt.Fprintf(w, "Done", content)

}
