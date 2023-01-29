chrome.runtime.onMessage.addListener(function (message, sender, sendResponse) {
    addURLToContainer(message);
    sendResponse("OK");
});

function addURLToContainer(links) {
    // Create the list
    var list = document.createElement("ul");

    // Loop through the array of items
    for (var i = 0; i < links.length; i++) {
        // Create a list item for each item in the array
        var item = document.createElement("li");

        // Set the item's inner HTML to include the link and text
        var newElement = document.createElement('a');
        newElement.innerHTML = "<a href='" + links[i].href + "'>" + links[i].text + "</a>";
        item.appendChild(newElement);

        // Add the item to the list
        list.appendChild(item);
    }

    // Add the list to the page
    document.body.appendChild(list);
}
