// if page down => app...
window.onload = function () {
    
    const btn_clean = document.getElementById("btn-clean");
const textarea = document.getElementById("textarea");
const save = document.getElementById("save");
const notif = document.getElementById("saving");

// if in user device content for textarea: get...
if (typeof(Storage) !== "undefined") {

    try {
      let stock = localStorage.getItem("stockacartweb");
      if (stock != null && stock != "") {
        textarea.value = stock;
      } else {
        textarea.value = "";
      }
      notif.innerText = "";
    } catch (error) {
      notif.innerText = "";
    }
  } 

  // if txet area void: button claen: display:none
if (textarea.value == "") {
    btn_clean.style.display = "none";
}  

// clean... for textarea
btn_clean.addEventListener("click", () => {
    textarea.value = ""
    btn_clean.style.display = "none"
})

textarea.addEventListener("input", () => {
    if (textarea.value !== "")
        btn_clean.style.display = "block"
    else
        btn_clean.style.display = "none"
})

// for save accepte for content of textarea in user device
save.addEventListener("click", () => {

    if (typeof(Storage) !== "undefined") {

  try {
    localStorage.setItem("stockacartweb",textarea.value);
    notif.innerText = "Is saved";
  } catch (error) {
    notif.innerText = "It seems that the local save feature is not currently enabled on this browser.";
  }
} else {
    notif.innerText = "It seems that the local save feature is not supported on the device or browser.";
}

})


};

