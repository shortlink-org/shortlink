import { useState } from 'react';
import './App.css';

function App() {
  const [links, setLinks] = useState([]);

  const addURLToContainer = () => {
    // Send a message to the background script to execute the content script
    chrome.runtime.sendMessage({ type: 'GET_LINKS' }, (response) => {
      // Update the state with the received links
      setLinks(response.links);
    });
  };

  return (
    <>
      <h2>ShortLink</h2>

      <button id="addLinkButton" onClick={addURLToContainer}>Parse links</button>

      {/* Render the links in a table */}
      <table>
        <thead>
          <tr>
            <th>URL</th>
            <th>Text</th>
          </tr>
        </thead>
        <tbody>
          {links.map((link, index) => (
            <tr key={index}>
              <td>{link.href}</td>
              <td>{link.text}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );
}

export default App;
