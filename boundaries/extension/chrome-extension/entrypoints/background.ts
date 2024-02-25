export default defineBackground(() => {
  console.log('Hello background!', { id: browser.runtime.id });

  // Listen for a message from the popup
  chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
    if (request.type === 'GET_LINKS') {
      // Execute the content script in the current tab
      chrome.tabs.query({ active: true, currentWindow: true }, (tabs) => {
        const tab = tabs[0];
        if (tab) {
          chrome.scripting.executeScript(
            {
              target: { tabId: tab.id },
              function: grabURL,
            },
            (results) => {
              // Send the links back to the popup
              sendResponse({ links: results[0].result });
            }
          );
        }
      });

      // Indicate that the response will be sent asynchronously
      return true;
    }
  });
});

// This function will be stringified and injected into the current tab
function grabURL() {
  const links = Array.from(document.querySelectorAll('a')).map((link) => ({
    href: link.href,
    text: link.textContent,
  }));
  return links;
}
