// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import '../IAssert.sol';
import './IMulti.sol';

contract MultiTest {
  IAssert private asserts;
  IMulti private multi;

  address private user1;
  address private user2;


  /// @notice addresses of test targets are passed to the constructor by a unit test's `dep.go`
  constructor(address assertAddr, address multiAddr, address u1, address u2) {
    asserts = IAssert(assertAddr);
    multi = IMulti(multiAddr);
    user1 = u1;
    user2 = u2;
  }

  function transferTest() internal returns (bool) {
    
  }
}
