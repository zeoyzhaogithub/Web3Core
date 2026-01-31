# ECDSA 库

## 介绍

OpenZeppelin 中的 ECDSA 库提供了基于椭圆曲线数字签名算法（ECDSA）的实用工具，主要用于在智能合约中验证链下生成的签名。以下是详细说明：

### 核心功能

#### 地址恢复

通过消息哈希（hash）和签名（signature）恢复出签名者的以太坊地址。

#### 签名验证

比较恢复出的地址与预期地址，验证签名的有效性。

#### 防篡改消息

支持添加以太坊专属前缀（\x19Ethereum Signed Message:\n），防止签名被用于其他场景。

## 应用场景

### 无 Gas 交易

用户链下签名，代理人支付 Gas 提交交易（如 MetaTx）。

### 去中心化交易所

验证用户挂单/取消订单的签名。

### 空投认领

用户签名证明资格，合约验证后发放代币。

### 多签授权

多个签名者联合签署交易后触发合约操作。

### 身份验证

证明用户拥有某个私钥（如登录系统）。

## 使用步骤

### 1. 链下生成签名（JavaScript 示例）

```javascript
const { ethers } = require("ethers");

// 1. 构造待签名的消息
const message = "Hello, Contract";
const messageHash = ethers.utils.id(message); // keccak256

// 2. 添加以太坊前缀并签名
const signer = new ethers.Wallet(privateKey);
const signature = await signer.signMessage(ethers.utils.arrayify(messageHash));
```

### 2. 链上验证签名（Solidity 代码）

```solidity
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract SignatureVerifier {
    using ECDSA for bytes32;

    function verify(
        address expectedSigner,
        string memory message,
        bytes memory signature
    ) public pure returns (bool) {
        // 1. 计算消息哈希
        bytes32 messageHash = keccak256(abi.encodePacked(message));
        
        // 2. 添加以太坊前缀后哈希
        bytes32 ethSignedMessageHash = messageHash.toEthSignedMessageHash();
        
        // 3. 恢复签名者地址
        address recoveredSigner = ethSignedMessageHash.recover(signature);
        
        // 4. 验证匹配
        return recoveredSigner == expectedSigner;
    }
}
```

## 完整示例合约：代币空投认领

```solidity
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract Airdrop {
    using ECDSA for bytes32;

    address public owner;
    IERC20 public token;
    mapping(address => bool) public hasClaimed; // 防止重复领取
    
    constructor(address tokenAddress) {
        owner = msg.sender;
        token = IERC20(tokenAddress);
    }

    function claimAirdrop(
        uint256 amount,
        bytes calldata signature
    ) external {
        require(!hasClaimed[msg.sender], "Already claimed");
        
        // 1. 构造合约预期的消息
        bytes32 message = keccak256(abi.encodePacked(msg.sender, amount));
        
        // 2. 验证签名是否来自owner
        address signer = message.toEthSignedMessageHash().recover(signature);
        require(signer == owner, "Invalid signature");
        
        // 3. 标记并发放代币
        hasClaimed[msg.sender] = true;
        token.transfer(msg.sender, amount);
    }
}
```

## 关键注意事项

### 重放攻击防护

在消息中包含用户地址和唯一标识（如 nonce、链 ID）。

使用映射记录已使用的签名（示例中的 hasClaimed）。

### 消息构造一致性

链下签名与链上验证的消息必须完全一致（编码、顺序、类型）。

### 前缀安全机制

始终使用 .toEthSignedMessageHash() 添加前缀，避免签名被恶意复用。

### 签名歧义问题

ECDSA 有少数签名可能产生歧义（高位 S 值），OpenZeppelin 会自动处理。

### 智能合约签名者

如果签名者可能是合约（非 EOA），需兼容 EIP-1271 标准。

### 签名格式

确保签名为 65 字节（r + s + v），v 值必须为 27 或 28。

## 最佳实践

### 使用 EIP-712

对结构化数据签名（如域名、版本），提升用户体验和安全性。

### 链 ID 隔离

在消息中包含 block.chainid，防止跨链重放攻击。

### 避免硬编码地址

恢复出的地址应动态验证（如检查白名单），而非固定地址。

通过合理使用 OpenZeppelin 的 ECDSA 工具，可显著提升合约的安全性与用户体验，同时实现复杂的链下-链上交互逻辑。

加签与验签
![](./WeChatddd74f2a569ebb5f021b070024fef0c2.jpg)

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;
import "@openzeppelin/contracts/utils/cryptography/ECDSA.so
l";
import "@openzeppelin/contracts/utils/cryptography/MessageH
ashUtils.sol";
contract VerifySignature {
 
 using ECDSA for bytes32;
 using MessageHashUtils for bytes32;
 
 function recover(string memory str,bytes memory signatu
re) external pure returns(address){
 bytes32 hash=keccak256(bytes(str));
 return hash.toEthSignedMessageHash().recover(signatu
re);
 
 }
}
```
