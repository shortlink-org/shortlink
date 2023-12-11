// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping(bytes32 => bytes32) private items;

  constructor(string memory _version) public {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) public {
    items[key] = value;
    emit ItemSet(key, value);
  }
}
