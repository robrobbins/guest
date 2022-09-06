// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;

interface IERC20 {
  event Transfer(address indexed from, address indexed to, uint value);
  event Approval(address indexed owner, address indexed spender, uint value);

  function name() external view returns (string calldata);
  function totalSupply() external view returns (uint256);
  function balanceOf(address) external view returns (uint256);
  function transfer(address, uint256) external returns (bool);
  function allowance(address, address) external view returns (uint256);
  function approve(address, uint256) external returns (bool);
  function transferFrom(address, address, uint256) external returns (bool);
}

