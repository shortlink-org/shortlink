import { useState } from 'react';
import './App.css';

function App() {
  const [links, setLinks] = useState([]);
  const [error, setError] = useState(null);

  const addURLToContainer = () => {
    // Send a message to the background script to execute the content script
    chrome.runtime.sendMessage({ type: 'GET_LINKS' }, (response) => {
      // Check for errors
      if (chrome.runtime.lastError) {
        setError(chrome.runtime.lastError.message);
      } else {
        // Update the state with the received links
        setLinks(response.links);
      }
    });
  };

  const saveLink = (link) => {
    // Implement your save logic here
    console.log(`Saving link: ${link.href}`);
  };

  return (
    <>
      <h2 className="text-2xl font-bold mb-4">ShortLink</h2>

      <button id="addLinkButton" onClick={addURLToContainer} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mb-4">
        Parse links
      </button>

      {/* Render the links in a table */}
      <table className="w-full table-auto">
        <thead>
        <tr>
          <th className="px-4 py-2">URL</th>
          <th className="px-4 py-2">Control</th>
        </tr>
        </thead>
        <tbody>
        {links.map((link, index) => (
          <tr key={index} className="border-t border-gray-200">
            <td className="px-4 py-3">
              <a href={link.href} className="text-blue-500 hover:underline">
                {link.text ? link.text : 'Link'}
              </a>
            </td>
            <td className="px-4 py-3">
              <button onClick={() => saveLink(link)} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                Save
              </button>
            </td>
          </tr>
        ))}
        </tbody>
      </table>

      {/* Render any error messages */}
      {error && <div className="mt-4 text-red-500">Error: {error}</div>}
    </>
  );
}

export default App;
