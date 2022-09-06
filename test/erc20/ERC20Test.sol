// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;

import '../IAssert.sol';
import './IERC20.sol';

contract ERC20Test {
  IAssert private asserts;
  IERC20 private token;

  /// @notice addresses of test targets are passed to the constructor by a unit test's `dep.go`
  constructor(address assertAddr, address tokenAddr) {
    asserts = IAssert(assertAddr);
    token = IERC20(tokenAddr);
  }

  function nameTest() internal view {
    require(asserts.equal(token.name(), 'Example Tizoken'), 'name mismatch'); 
  }

  function run() external view {
    nameTest();
  }
}
