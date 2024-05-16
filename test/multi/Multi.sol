// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "./IMulti.sol";
import "../erc20/IERC20.sol";

contract Multi is IMulti {
  function transfer(
    address[] calldata recipients,
    uint256[] calldata amounts,
    address token) external returns (bool) {
      uint256 len = recipients.length;

      if(amounts.length != len) {
        revert InputLengthMismatch();
      }

      IERC20 target = IERC20(token);

      for(uint256 i = 0; i < len;) {
        target.transferFrom(msg.sender, recipients[i], amounts[i]);
        unchecked { ++i; }
      }

      emit MultiTransfer(msg.sender, token);
      return true;
    }
}
