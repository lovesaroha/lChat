"use-strict";

/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/

let activeRoom = {};

// Room page.
appRoutes["/#/room"] = function (urlParameters) {
     if (!isLogged()) { return; }
     // Check active room.
     if (!isRoomMember()) {
          // Room not found or not a member.
          window.location = "/#/";
          return;
     }
     view.innerHTML = document.getElementById("roomPageTemplate_id").innerHTML;
     document.getElementById("roomName_id").innerText = `Room-${activeRoom._id}`;
     activeRoom.messages = [];
     showMembers();
}

// Set default.
function setDefault() {
     rooms = {};
     activeRoom = {};
     localStorage.clear();
     user = { emailAddress: "", token: 0, socket: { id: 0 } };
}

// Handle user joined.
function handleUserJoined(message) {
     if (!window.location.hash.includes(message.roomID) || !activeRoom._id) { return; }
     if (activeRoom.members.includes(message.emailAddress)) { return; }
     activeRoom.members.push(message.emailAddress);
     showMembers();
}

// Handle user left.
function handleUserLeft(message) {
     if (!window.location.hash.includes(message.roomID) || !activeRoom._id) { return; }
     if (!activeRoom.members.includes(message.emailAddress)) { return; }
     activeRoom.members.splice(activeRoom.members.indexOf(message.emailAddress), 1);
     showMembers();
}


