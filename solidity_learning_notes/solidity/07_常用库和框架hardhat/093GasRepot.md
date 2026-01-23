# Gas Report

## 本地调试怎么计算gas消费？

---

- gas reporter

`yarn add -D hardhat-gas-reporter`

- 添加到hardhat.config.ts

`import "hardhat-gas-reporter"`

- 在hardhat配置中开启gas-reporter

```json
const config: Config = {
  ...,
  networks: {
    hardhat: {
      chainId: 31337
    },
    ...
  },
  gasReporter: {
    enabled: true
  }
};
```

## 测试操作gas消耗

---

```json
pragma solidity ^0.8.24;

contract TestGas {
 uint a;
 uint b;
 uint c;
 
 function test1() public {
  a++;
 }

 function test2() public {
  a++;
  b++;
 }
 
 function test3() public {
  a++;
  b++;
  c++;
 }
 
 function test4() public {
  c = a + b;
 }
 
 function test5() public {
  test4();
  b = a + c;
 }
 
}
```

## Test Gas

---

```json
const TG = await ethers.getContractFactory("TestGas");
const tg = await TG.deploy();
await tg.waitForDeployment();

for(let i = 0; i < 10; ++i){
 await tg.test1();
 await tg.test2();
 await tg.test3();
 await tg.test4();
 await tg.test5();
}
```

和测试代码一样，运行

`npx hardhat test`

## 核心要点

---

- 注意你的每个操作，他们都会有gas消耗
- 在一些编程语言里面，你可能很爱做copy操作，js（deepClone），在solidity里面别这么做，这种开销非常大。
- 转换一下你的开发思路，你在solidity里面做的所有操作都是真实在花钱的，以前你可能觉得内存、cpu够用就吃满，但在solidity里别这样想。【array举例】immutation

-----

要获取智能合约的 Gas 报告（即执行合约函数所需的 Gas 消耗分析），具体方法取决于你的开发工具链。以下是主流工具的详细操作指南：

## 1. 使用 Hardhat

### 步骤

#### 安装插件

```bash
npm install --save-dev hardhat-gas-reporter
```

#### 启用插件

在 hardhat.config.js 中添加：

```javascript
require("hardhat-gas-reporter");
module.exports = {
  gasReporter: {
    enabled: true,
    currency: "USD",  // 可选：显示美元成本（需 API Key）
    token: "ETH",     // 可选：主网币种（ETH/BNB等）
    gasPrice: 20,     // 可选：自定义 Gas 价格（单位：Gwei）
  }
};
```

#### 运行测试

```bash
npx hardhat test
```

**输出示例：**

```text
·-------------|--------|-----------|---------·
|  Contract·  Method   ·  Min        ·  Max   │
·-----------|----------|--------|----------·
|MyContract · mintNFT ·  105214  · 132541  gas │
·----------|---------|----------|--------------·
```

## 2. 使用 Foundry

Foundry 内置 Gas 报告功能，无需额外插件。

### 步骤

#### 运行测试并生成报告

```bash
forge test --gas-report
```

#### 输出示例

```text

| Contract      | Method    | Min     | Max     | Avg     |
|---------------|-----------|---------|---------|---------|
| MyContract    | mintNFT   | 105214  | 132541  | 118877  |
```

## 3. 使用 Truffle

### 步骤

#### 安装插件

```bash
npm install truffle-gas-reporter
```

#### 添加配置

在 truffle-config.js 中添加：

```javascript
module.exports = {
  plugins: ["truffle-gas-reporter"],
  mocha: {
    reporter: 'eth-gas-reporter',
    reporterOptions: {
      currency: 'USD',
      gasPrice: 20
    }
  }
};
```

#### 运行测试

```bash
truffle test
```

## 4. 手动获取 Gas 消耗（适用于任何环境）

在合约中使用 gasleft() 函数：

```solidity
function myFunction() public {
    uint256 startGas = gasleft();
    // ...执行操作...
    uint256 gasUsed = startGas - gasleft();
    console.log("Gas used:", gasUsed);
}
```

## 关键注意事项

### Gas 价格波动

Gas 消耗量是固定的，但 Gas 成本（费用）随网络拥堵变化。工具中显示的 USD 成本是估算值。

### 优化建议

减少存储操作（SSTORE 消耗最高）

使用 calldata 替代 memory

避免循环中的动态数组扩容

### 测试网 vs 主网

Gas 消耗在测试网和主网一致，但实际费用取决于当前网络的 Gas 价格。

## 高级场景

比较优化效果：在 Hardhat/Foundry 中运行测试时，工具会自动对比修改前后的 Gas 变化。

CI/CD 集成：将 Gas 报告导出为 JSON/CSV（Hardhat 配置 outputFile 选项）。

通过以上方法，你可以精准分析合约函数的 Gas 消耗，为优化和部署提供关键数据支持。
