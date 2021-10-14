/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package api

import (
	"encoding/json"
	"net/http"
)

// HandleSignIn function listen to post /api/sign-in.
func HandleSignIn(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var user userObject
	if json.NewDecoder(req.Body).Decode(&user) != nil {
		http.Error(res, `{"error" : "Invalid parameters"}`, 400)
		return
	}
	if inValidEmailAddress(user.EmailAddress) {
		http.Error(res, `{"error" : "Invalid email address"}`, 400)
		return
	}
	res.Write([]byte(userJSON(user)))
}
