# 质押收益

- 质押收益介绍
- Defi质押的操作流程
- 合约实现

## 质押收益介绍

**Defi：单币质押、流动性质押、借贷等**

**PoS质押：以太坊通过质押ETH来选择验证人，参与新的区块产生的验证**

**其它：GameFi、NFT质押**

### **常见Defi质押介绍**

#### **单币质押：**

**平台：AAVE、Compound、LaunchPad（Staking）**

**原理：用户存入一种代币，获取利息或质押收益**

**举例：相当于去中心化银行存款，在LaunchPad中存入C2N代币，获取C2N代币奖励**

#### **流动性质押：**

**平台：Uniswap、PancakeSwap**

**原理：用户提供两种代币组成的交易对，换取流动性代币**

**举例：在Uniswap中提供交易币对，获取交易币对的LP Token，通过LP token获得奖励或参与其它活动**

### **借贷质押：**

**平台：AAVE、Compound**

**原理：用户存入抵押资产，借出另一种资产，同时获取利息或治理代币利息**

**举例：在AAVE中存入ETH赚取利息，同时获取AAVE代币奖励**

## **质押机制和流程**

1. **两种IERC20：质押代币、奖励代币（可以是同种代币）**
2. **奖励机制设置**
3. 用户质押代币

![](./jpg/WeChat1a7574a1e8339c422c524b4efd258880.jpg)

## 合约实现**（理解**变量含义、安全性**）**

- [x]  构造函数（质押代币、奖励代币、管理员）
- [x]  管理员modifier
- [x]  设置持续时间、设置奖励速率（notifyRewardAmount）
- [x]  **更新收益modifier**
- [x]  质押、撤回、**获取收益函数**
- [x]  状态查询函数

### 设置奖励金额和分配速率

```solidity
function notifyRewardAmount(uint256 _amount) external onlyOwner updateReward(address(0)) {}
```

- 当当前时间超过上一个奖励周期结束时间(`finishAt`)时，创建一个全新的奖励周期
- 当当前时间还在上一个奖励周期内时，将剩余奖励与新奖励合并计算新的分配速率
- 更新奖励结束时间(`finishAt`)为当前时间加上持续时间(`duration`)

1. 当前时间：1600，finishAt：1500，amount：1000，duration：1000
    1. rewardRate：1
2. 当前时间：2000，finishAt：2600，amount：1000，duration：1000
    1. remainingRewards：600
    2. rewardRate：1.6

### 在关键操作前更新用户奖励状态

```solidity
modifier updateReward(address _account) {}

function rewardPerToken() public view returns (uint256) {}

function lastTimeRewardApplicable() public view returns (uint256) {}
```

时间轴：
t0 ---- t1 (用户A质押) ---- t2 (用户B质押) ---- t3 (现在) —— t4(用户A质押)

全局 rewardPerTokenStored 发展：
t0: 0
t1: 1.2
t2: 1.8
t3: 2.5

t4: 2.8

用户A：

- 质押时(t1) userRewardPerTokenPaid[A] = 1.2
- earned[A] = 质押金额 × (2.5 - 1.2) + rewards[A]

用户B：

- 质押时(t2) userRewardPerTokenPaid[B] = 1.8
- earned[B] = 质押金额 × (2.5 - 1.8) + rewards[B]

### 计算用户当前可领取的奖励总额

```solidity
function earned(address _account) public view returns (uint256) {}
```

质押时，会按照当前用户的质押数量，先以当前质押的数额，将奖励存入用户的rewards中去，然后重置用户的奖励计算参数（因为质押的总金额变了，每个质押代币对应的奖励数额也变了）

每次的质押、撤回，都会重新计算用户的rewards[account]
