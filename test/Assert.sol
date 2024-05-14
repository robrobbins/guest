// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import './IAssert.sol';

contract Assert is IAssert {

  /// @param a keccack hash of some value
  /// @param b keccack hash of some value
  function equal(bytes32 a, bytes32 b) external pure returns (bool) {
    return(a == b);
  }

  /// @param a string of some value
  /// @param b string hash of some value
  function equal(string calldata a, string calldata b) external pure returns (bool) {
    return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
  }
}
