# OpenZeppelin AccessControl 介绍

## 一、功能概述

AccessControl 是 OpenZeppelin 提供的基于角色的权限控制（RBAC）合约，核心功能包括：

### 角色管理

创建自定义角色（如 MINTER_ROLE, ADMIN_ROLE）

授予/撤销账户角色

检查账户是否拥有角色

### 角色继承

支持角色层级关系（如管理员可管理子角色）

### 权限检查

通过修饰器 onlyRole 限制函数访问

## 二、应用场景

### 代币权限控制

铸币权（MINTER_ROLE）

销毁权（BURNER_ROLE）

### 合约管理

升级合约（UPGRADER_ROLE）

暂停合约（PAUSER_ROLE）

### 多签管理

多重管理员共同授权敏感操作

### DAO 治理

将治理角色分配给 DAO 合约

## 三、使用方式

### 1. 基础使用步骤

```solidity
// 导入库
import "@openzeppelin/contracts/access/AccessControl.sol";

contract MyContract is AccessControl {
    // 1. 定义角色（bytes32）
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    constructor() {
        // 2. 初始化角色（通常给部署者管理员权限）
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    // 3. 使用修饰器保护函数
    function mint() public onlyRole(MINTER_ROLE) {
        // 只有 MINTER_ROLE 可调用
    }
}
```

### 2. 关键操作

```solidity
// 授予角色
grantRole(MINTER_ROLE, 0x123...);

// 撤销角色
revokeRole(MINTER_ROLE, 0x123...);

// 检查角色
require(hasRole(MINTER_ROLE, msg.sender), "Unauthorized");

// 设置角色管理员（控制角色分配权）
_setRoleAdmin(MINTER_ROLE, ADMIN_ROLE);
```

## 四、示例合约

### 带权限管理的 ERC20 代币

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract MyToken is ERC20, AccessControl {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant BURNER_ROLE = keccak256("BURNER_ROLE");

    constructor() ERC20("MyToken", "MTK") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender); // 部署者获得默认管理员
        _grantRole(MINTER_ROLE, msg.sender);
        _grantRole(BURNER_ROLE, msg.sender);
    }

    // 只有铸币者能铸造代币
    function mint(address to, uint256 amount) public onlyRole(MINTER_ROLE) {
        _mint(to, amount);
    }

    // 只有销毁者能销毁代币
    function burn(address from, uint256 amount) public onlyRole(BURNER_ROLE) {
        _burn(from, amount);
    }
}
```

## 五、注意事项

### 初始化管理员

```solidity
// 必须初始化 DEFAULT_ADMIN_ROLE
_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
```

否则后续无法管理角色。

### 角色冲突处理

避免使用 DEFAULT_ADMIN_ROLE 直接管理业务逻辑

为不同功能创建独立角色（如分离 MINTER_ROLE 和 BURNER_ROLE）

### 批量操作优化

使用循环批量授权（避免单独调用消耗 Gas）：

```solidity
function batchGrantRole(bytes32 role, address[] calldata accounts) external onlyRole(getRoleAdmin(role)) {
    for(uint i=0; i<accounts.length; i++) {
        grantRole(role, accounts[i]);
    }
}
```

### 撤销管理员权限

撤销自己管理员权限前需确保有其他管理员：

```solidity
revokeRole(DEFAULT_ADMIN_ROLE, msg.sender); // 危险！需确认有其他管理员存在
```

### 接口兼容性

支持 ERC165 标准接口检测：

```solidity
supportsInterface(IAccessControl); // 返回 true
```

### Gas 优化

使用 _grantRole 替代 grantRole 在构造函数中初始化（避免检查）

## 六、最佳实践

最小权限原则：仅授予必要权限

多角色分离：关键操作需多角色共同授权

事件监控：监听 RoleGranted/RoleRevoked 事件跟踪权限变更

紧急恢复：保留超级管理员紧急暂停机制

通过合理使用 AccessControl，可实现企业级智能合约的精细权限管理，大幅提升合约安全性。

![](./WeChat71017fe3e02d92d68fe27ddd91c8e3bb.jpg)
