const btn_clean = document.getElementById("btn-clean");
const textarea = document.getElementById("textarea");

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