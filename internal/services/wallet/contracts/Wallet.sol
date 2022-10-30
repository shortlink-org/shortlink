// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Wallet {
  uint256 balance;
  address public admin;

  constructor() {
    admin = msg.sender;
    balance = 0;
    updateBalance();
  }

  function updateBalance() internal {
    balance += msg.value;
  }

  function Winthdraw(uint256 _amt) public {
    require(msg.sender == admin, "You are not the admin");
    balance -= _amt;
  }

  function Deposit(uint256 amt) public returns(uint256) {
    balance += amt;
    return balance;
  }

  function Balance() public view returns(uint256) {
    return balance;
  }
}
