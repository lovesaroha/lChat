/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
// All helper functions.
// Color themes defined here.
const themes = [
  {
      normal: "#5468e7",
      dark: "#4c5ed0",
      light: "#98a4f1",
      veryLight: "#eef0fd",
      primaryText: "#ffffff",
      iconSecondary: "#FFD43B"
  }, {
      normal: "#e94c2b",
      dark: "#d24427",
      veryLight: "#fdedea",
      light: "#f29480",
      primaryText: "#ffffff",
      iconSecondary: "#FFD43B"
  }
];

// Choose random color theme.
let colorTheme = themes[Math.floor(Math.random() * themes.length)];

// This function set random color theme.
function setTheme() {
  // Change css values.
  document.documentElement.style.setProperty("--primary", colorTheme.normal);
  document.documentElement.style.setProperty("--icon-secondary", colorTheme.iconSecondary);
}

// Set random theme.
setTheme();

// This function checks if email is valid or not.
function isEmail(string) {
  const emailRegexp = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;
  if (emailRegexp.test(string) == false) {
      return false;
  }
  return true;
}

// This is a showModal function which shows modal based on given options as an argument.  
function showModal(content) {
  let modal = document.getElementById("modal_id");
  if (modal == null) { return; }
  modal.style = "display: block;";
  modal.innerHTML = content;
}

// This is closeModal function which closes modal and remove backdrop from body.
function closeModal() {
  let modal = document.getElementById("modal_id");
  if (modal == null) { return; }
  modal.style = "display: none;";
  modal.innerHTML = ``;
}

// This is closeModal background function which closes modal.
function closeModalBackground(e) {
  if (e.target.id != "modal_id") { return; }
  let modal = document.getElementById("modal_id");
  if (modal == null) { return; }
  modal.style = "display: none;";
  modal.innerHTML = ``;
}

// This function show given date and time.
function showDateAndTime(date) {
     if (typeof date == "string" || typeof date == "number") {
          // Convert it to date object.
          date = new Date(date);
     }
     let hours = date.getHours();
     let minutes = date.getMinutes();
     let time = "am";
     if (hours > 12) {
          hours -= 12;
          time = "pm";
     }
     if (minutes < 10) { minutes = `0${minutes}`; }
     if (isToday(date)) {
          // Given date is today.
          return `${hours}:${minutes} ${time}`;
     }
     let d = date.getDate();
     let day = `${d}th`;
     let month = months[date.getMonth()];
     let year = date.getFullYear();
     switch (d % 10) {
          // Set day values.
          case 1: {
               day = `${d}st`;
               break;
          }
          case 2: {
               day = `${d}nd`;
               break
          }
          case 3: {
               day = `${d}rd`;
               break;
          }
     }
     if (d > 10 && d < 20) { day = `${d}th`; }
     if (new Date().getFullYear() == year) { year = ""; }
     return `${day} ${month} ${year}`;
}

// This function checks if given date is today.
function isToday(date) {
     let today = new Date();
     return (today.getDate() == date.getDate() && today.getMonth() == date.getMonth() && today.getFullYear() == date.getFullYear());
}