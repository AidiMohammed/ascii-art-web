// if page load : instructions...
window.onload = function () {
  // for download text file in user device content copy of the current result for ASCI ART:
  const download = document.getElementById("down");
  const result = document.getElementById("res");

  download.addEventListener('click', function () {
    
    const textContent = result.innerText;
    // like a fake file...
    const blob = new Blob([textContent], { type: 'text/plain' });
    const link = document.createElement('a');
    // write in memory or...
    link.href = URL.createObjectURL(blob);
    link.download = 'content.txt'; 

    // for download:
    link.click();

    // library memory or stock...
    URL.revokeObjectURL(link.href);
})

};