// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

interface IMulti {
  event MultiTransfer(address indexed caller, address indexed token);

  error InputLengthMismatch();

  function transfer(address[] calldata, uint256[] calldata, address) external returns (bool); 
}
