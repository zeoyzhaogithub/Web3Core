# ERC20Pausable

## 介绍

ERC20Pausable 是基于标准 ERC20 代币的扩展合约，允许合约管理员在紧急情况下暂停代币转账功能。它继承自 OpenZeppelin 的 ERC20 和 Pausable 合约，增加了对代币转移、授权等关键操作的可控暂停能力。

## 核心功能

### 暂停转账

禁止所有 transfer(), transferFrom() 操作

禁止 approve() 和 increaseAllowance() 等授权操作

不影响：余额查询、代币信息读取等非状态变更操作

### 恢复转账

解除暂停后，所有操作恢复正常

### 权限控制

仅授权地址（如合约所有者）可触发暂停/恢复

## 应用场景

场景 |说明
|-|-|
安全应急| 发现合约漏洞时紧急冻结资产转移
合规审查| 配合监管要求临时冻结可疑交易
系统升级 |升级期间暂停交易避免状态不一致
法律纠纷| 法院要求冻结特定资产时快速响应
代币锁仓期 |配合锁仓机制防止未解锁代币转移

## 如何使用（OpenZeppelin v5 示例）

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Pausable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract PausableToken is ERC20, ERC20Pausable, Ownable {
    constructor() ERC20("SafeToken", "SFT") Ownable(msg.sender) {
        _mint(msg.sender, 1000000 * 10**decimals());
    }

    // 暂停代币转账（仅Owner）
    function pause() public onlyOwner {
        _pause();
    }

    // 恢复代币转账（仅Owner）
    function unpause() public onlyOwner {
        _unpause();
    }

    // 重写更新逻辑，加入暂停检查
    function _update(
        address from,
        address to,
        uint256 amount
    ) internal override(ERC20, ERC20Pausable) {
        super._update(from, to, amount); // 自动检查暂停状态
    }
}
```

## 关键注意事项

### 权限风险

确保 pause()/unpause() 调用者高度可信（建议多签钱包）

避免过度中心化，明确暂停权使用条件

### 用户信任影响

需在文档中披露暂停机制的存在

滥用暂停功能可能导致项目信任崩塌

### 技术限制

⚠️ 无法冻结：已授权的第三方转账（需配合 ERC20Permit 限制）
⚠️ 不冻结余额：仅阻止转移，账户余额仍可查询

### 事件透明度

暂停/恢复操作需触发事件供链上追踪

建议前端明确展示合约暂停状态

## 完整操作流程

```mermid
sequenceDiagram
    participant Owner
    participant Contract
    participant User

    Owner->>Contract: pause()
    Contract-->>Contract: 标记 paused=true
    Contract-->>Owner: 触发 Paused 事件

    User->>Contract: transfer(100 tokens)
    Contract->>Contract: 检查 paused==true
    Contract-->>User: 回滚交易，显示 "Pausable: paused"
    
    Owner->>Contract: unpause()
    Contract-->>Contract: 标记 paused=false
    Contract-->>Owner: 触发 Unpaused 事件
    
    User->>Contract: transfer(100 tokens)
    Contract->>User: 转账成功
```

## 📌 最佳实践

配合 Timelock 合约使用，使暂停操作延迟生效，避免管理员作恶
