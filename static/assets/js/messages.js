/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/

// Send message function.
function sendMessage(element) {
     if (!activeRoom._id) { return; }
     if (element.value.length == 0) { return; }
     fetch(`/api/create-message`, {
          method: 'post',
          headers: {
               "Content-type": "application/json; charset=UTF-8",
               "X-Authentication-Token": user.token
          },
          body: `{"to" : "${activeRoom._id}" , "content" : "${element.value}"}`
     }).then(r => {
          element.value = ``;
     }).catch(e => { console.log(e); });
}

// This function handle new message.
function handleNewMessage(message) {
     if (!activeRoom._id) { return; }
     activeRoom.messages.push(message);
     let messagesEl = document.getElementById("messages_id");
     if (messagesEl == null) {
          return;
     }
     // Show message.
     messagesEl.innerHTML += messageTemplate(message);
     scrollToBottom(messagesEl);
}

// Message template.
function messageTemplate(message) {
     return `<div class="p-4 mb-2 ${message.from == user.emailAddress ? 'bg-primary' : 'bg-light'}"><h5 class="text-gray mb-0 ${message.from == user.emailAddress ? 'hidden' : ''}">${message.from}</h5><h4 class="${message.from == user.emailAddress ? 'text-white' : 'text-subtitle'} mb-0">${showText(message.content)}</h4><h5 class="mb-0 ${message.from == user.emailAddress ? 'text-white' : 'text-gray'}">${showDateAndTime(message.createdAt)}</h5></div>`;
}

// Scroll to bottom.
function scrollToBottom(el) {
     el.scrollTop = el.scrollHeight;
}

// This function shows text in DOM.
function showText(content) {
     let el = document.createElement("span");
     el.innerText = content;
     return el.innerHTML;
}

