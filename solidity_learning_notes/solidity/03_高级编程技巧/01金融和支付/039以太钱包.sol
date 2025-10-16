// SPDX-License-Identifier: MIT
// 以太钱包
// 申请solidity 版本号 ^表示使用大于等于0.8.30但是不超过0.9的版本号，部署主网的时候需要使用固定版本号
pragma solidity ^0.8.7;

contract EtherWallet {
    address payable public owner;

    // 初始化owner
    constructor() {
        owner = payable(msg.sender);
    }

    // 接收以太
    receive() external payable {}

    // withdraw 用以接收以太
    function withdraw(uint _amount) external {
        require(msg.sender == owner, "caller is not the owner");
        // 调用transfer 函数，其中第二个参数是地址（可以使用msg.sender）、第三个参数是需要转账的金额
        payable(msg.sender).transfer(_amount);

        // (bool sent, ) = msg.sender.call{value:_amount}("");
        // require(sent, "tx failed");
    }

    // 查询余额
    function getBalance() external view returns (uint) {
        return address(this).balance;
    }
}
