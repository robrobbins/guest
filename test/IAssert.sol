// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

interface IAssert {
  function equal(bytes32, bytes32) external pure returns (bool);
  function equal(string calldata, string calldata) external pure returns (bool);
}
