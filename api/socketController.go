/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package api

import (
	"net/http"
)

// HandleWebSocket request handle listen to /.
func HandleWebSocket(res http.ResponseWriter, req *http.Request) {
	var user = userObject{Token: req.FormValue("token")}
	if invalidUserToken(&user) {
		return
	}
	socket, _ := upgrader.Upgrade(res, req, nil)
	if err := saveUserSocket(user.EmailAddress, socket); err != nil {
		return
	}
	go userSocketHandler(user.EmailAddress, socket)
}
