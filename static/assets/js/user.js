/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
// All functions related to user.
// User and default variables defined.
let user = { emailAddress: "", token: 0, socket: { id: 0 } };
let socketURL = "ws://localhost:3000"

// Get user email from local storage.
if (localStorage.getItem("user-data") != null) {
     let u = JSON.parse(localStorage.getItem("user-data"));
     user.emailAddress = u.emailAddress;
     user.token = u.token;
}

// Initialize web socket.
function initializeWebSocket() {
     if (!isEmail(user.emailAddress) || user.socket.readyState < 2) { return; }
     user.socket = new WebSocket(`${socketURL}/ws?token=${user.token}`);

     // On open.
     user.socket.onopen = (e) => { console.log("connected"); }

     // On error.
     user.socket.onerror = (e) => {
          window.location = "/#/login";
     }

     // This function run on message.
     user.socket.onmessage = (response) => {
          let message = {};
          try {
               message = JSON.parse(response.data);
          } catch (e) { return; }
          console.log(message);
          switch (message.type) {
               case "userJoined":
                    handleUserJoined(message);
                    break;
               case "userLeft":
                    handleUserLeft(message);
                    break;
               case "message":
                    handleNewMessage(message.message);
                    break;     
               case "roomUnavailable":
                    handleRoomUnavailable(message);
                    break;     
          }
     }
}

// Check if user is logged.
function isLogged() {
     if (!isEmail(user.emailAddress) || user.token == 0) {
          // User not logged in.
          window.location = "/#/login";
          return false;
     }
     return true;
}

// User sign in function to register user's email address.
function signIn(e) {
     e.preventDefault();
     let emailAddress = e.target.emailAddress.value;
     fetch(`/api/sign-in`, {
          method: 'post',
          headers: {
               "Content-type": "application/json; charset=UTF-8"
          },
          body: `{"emailAddress" : "${emailAddress}"}`
     }).then(response => response.json()).then(json => {
          user.emailAddress = emailAddress;
          user.token = json.token;
          localStorage.setItem("user-data", JSON.stringify({ emailAddress: user.emailAddress, token: user.token }));
          window.location = '/#/';
     }).catch(e => { console.log(e); });
}