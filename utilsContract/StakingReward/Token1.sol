// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
/*
质押代币

部署：
1、初始化代币1000*10^18
*/
contract Token1 is ERC20 {
    constructor(uint256 initialSupply) ERC20("Gold1", "GLD1") {
        _mint(msg.sender, initialSupply);
    }

    function min(address to, uint256 amount) public {
        _mint(to, amount);
    }
}
