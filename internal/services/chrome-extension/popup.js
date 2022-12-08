const grabBtn = document.getElementById("addLinkButton");
grabBtn.addEventListener("click",() => {
    chrome.tabs.query({active: true}, (tabs) => {
        const tab = tabs[0];
        if (tab) {
            chrome.scripting.executeScript(
                {
                    target: {tabId: tab.id, allFrames: true},
                    func: grabURL
                },
                onResult
            )
        } else {
            alert("There are no active tabs")
        }
    })
})

/*
    * This function is executed in the context of the active tab.
    * It grabs the URL of the page and returns it to the popup.
    * @return {string} The URL of the page.
 */
function grabURL() {
    const links = document.querySelectorAll("a")
    return Array.from(links).map(link => ({href: link.href, text: link.innerText}))
}

function onResult(frames) {
    console.log(frames)
    if (!frames || !frames.length) {
        alert("Could not retrieve links from specified page")
        return
    }

    const links = frames.map(frame => frame.result)
                        .reduce((r1,r2)=>r1.concat(r2))

    window.navigator.clipboard
        .writeText(links.join("\n"))
        .then(() => alert("Links copied to clipboard"))
        .catch(() => alert("Could not copy links to clipboard"))

    window.close()
}