# Hardhat基本使用

## 一、Hardhat 概念

Hardhat 是以太坊智能合约开发环境（Development Environment），集成了编译、部署、测试和调试等功能。核心特点：

本地开发网络：内置 Hardhat Network（支持 Solidity 堆栈跟踪、console.log 等）

插件生态：扩展功能（如部署工具、验证工具）

任务系统：自定义自动化工作流

TypeScript 支持：强类型开发体验

## 二、核心功能

编译合约：自动编译 Solidity 代码

本地测试网络：模拟以太坊节点（带调试功能）

运行测试：集成 Mocha/Chai/Waffle

脚本部署：编写部署脚本到任意网络

插件扩展：

- @nomicfoundation/hardhat-toolbox（常用工具包）
- @openzeppelin/hardhat-upgrades（可升级合约）
- hardhat-etherscan（合约验证）

## 三、基本使用流程

### 环境安装

nodejs: [https://nodejs.org/en](https://nodejs.org/en)

yarn: `npm install -g yarn`

### 初始化项目

```bash
mkdir my-project && cd my-project
npm init -y    # 生成配置文件
npm install --save-dev hardhat  # 安装 Hardhat
npx hardhat init  # 选择 TypeScript 初始化项目

# 编写合约

# 编译合约
npx hardhat compile

# 使用typescript 写测试
yarn add -D ts-node typescript 

# 安装测试依赖
yarn add -D @types/mocha

# 运行测试代码
npx hardhat test

# 写部署合约
# 创建scripts文件夹，在里面创建deploy.hello.ts 文件
# 窗口1、先启动node
npx hardhat node
# 然后在另一个终端部署合约
npx hardhat run ./scripts/deploy-hello.ts --network localhost

# 创建webpack配置文件,创建src文件夹，在里面创建index.ts

# 添加webpackk依赖
yarn add -D webpack webpack-cli ts-loader html-webpack-plugin dotenv webpack-dev-server

# 启动服务
npx webpack server

# 对webpack在工程化上的配置内容做补充，合并文件
yarn add -D webpack-merge

## 拆分一下webpack.config.js
- webpack.config.js
  - webpack.common.js 共同
  - webpack.dev.js  开发环境
  - webpack.prod.js 生产环境

# 可以删除掉webpack.config.js

# 将打包命令添加到package.json中
# 注：deploy:local这中间不能有空格
"scripts": {
    "dev": "webpack server --config webpack.dev.js",
    "build": "webpack --config webpack.prod.js",
    "deploy:local": "hardhat run scripts/deploy-counter.ts --network localhost"
} 

# 运行打包命令
npm run build
npm run dev

# 运行部署命令
npm run deploy:local

# 部署之后再node服务中查看合约地址，
# 并将地址填入环境变.env文件中

# 窗口2、重启webpack服务
npx webpack server


```

### 安装依赖

```bash
npm install --save-dev @nomicfoundation/hardhat-toolbox
npm install @openzeppelin/contracts @openzeppelin/contracts-upgradeable
```

在运行测试代码时，测试代码找运行代码的路径是：
./artifacts/contracts/MyContract.sol/MyContract.json

### 配置 hardhat.config.ts

```typescript
import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: "0.8.20",
  networks: {
    hardhat: {}, // 本地网络
    sepolia: {   // 测试网
      url: "<https://sepolia.infura.io/v3/YOUR_KEY>",
      accounts: [process.env.PRIVATE_KEY!]
    }
  },
  etherscan: {    // 验证合约
    apiKey: process.env.ETHERSCAN_KEY
  }
};

export default config;
```

### 目录结构

```text
contracts/   # Solidity 合约
scripts/     # 部署脚本
test/        # 测试脚本
hardhat.config.ts
```

## 四、可升级合约关键流程

### 代理模式原理

代理合约：存储数据，委托调用逻辑合约
逻辑合约：执行业务代码（可替换）
Admin 合约：管理升级权限（透明代理模式）

### 使用 OpenZeppelin 插件

```bash
npm install --save-dev @openzeppelin/hardhat-upgrades
```

在 hardhat.config.ts 中添加：

```typescript
import "@openzeppelin/hardhat-upgrades";
```

## 五、完整可升级合约示例

### 1. 初始合约 V1

contracts/MyContractV1.sol：

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract MyContractV1 is Initializable, OwnableUpgradeable {
    uint256 public value;

    // 构造函数禁用（用 initialize 替代）
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __Ownable_init(msg.sender);
        value = 42;
    }

    function setValue(uint256 _value) public onlyOwner {
        value = _value;
    }
}
```

### 2. 升级合约 V2

contracts/MyContractV2.sol：

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./MyContractV1.sol";

contract MyContractV2 is MyContractV1 {
    // 新增功能
    function increment() public onlyOwner {
        value += 1;
    }
}
```

### 3. 部署脚本 (scripts/deploy_upgradeable.ts)

```typescript
import { ethers, upgrades } from "hardhat";

async function main() {
  // 部署 V1
  const MyContract = await ethers.getContractFactory("MyContractV1");
  const myContract = await upgrades.deployProxy(MyContract, [], {
    initializer: "initialize",
  });

  await myContract.waitForDeployment();
  console.log("Proxy deployed to:", await myContract.getAddress());
}

main();
```

### 4. 升级脚本 (scripts/upgrade.ts)

```typescript
import { ethers, upgrades } from "hardhat";

async function main() {
  const proxyAddress = "0x..."; // 替换为你的代理地址

  const MyContractV2 = await ethers.getContractFactory("MyContractV2");
  const upgraded = await upgrades.upgradeProxy(proxyAddress, MyContractV2);

  console.log("Upgraded to V2 at:", await upgraded.getAddress());
}

main();
```

### 5. 运行命令

```bash

# 部署 V1

npx hardhat run scripts/deploy_upgradeable.ts --network sepolia

# 升级到 V2

npx hardhat run scripts/upgrade.ts --network sepolia
```

## 六、关键注意事项

### 存储兼容性

升级时不可修改已有变量的顺序/类型

新变量必须追加在末尾

### 初始化函数

使用 initialize 替代构造函数

用 initializer 修饰符确保只初始化一次

### 安全最佳实践

使用透明代理模式防止函数冲突

在测试网充分测试后再部署主网

使用 Timelock 控制升级权限

完整代码库参考：OpenZeppelin Upgrades Plugins
