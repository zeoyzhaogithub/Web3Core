# Alchemy

## 一、Alchemy 是什么？

Alchemy 是一个区块链基础设施即服务（IaaS）平台，为开发者提供构建和扩展去中心化应用程序（dApps）所需的核心工具和服务。你可以把它想象成区块链世界的“AWS”或“Google Cloud”，专门服务于 Web3 开发。

### 核心功能

#### 1. 超级节点服务

提供高性能、高可靠性、可扩展的区块链节点访问（JSON-RPC 和 WebSocket 端点）。开发者无需自己搭建和维护节点

#### 2.增强型 API

NFT API： 强大的 NFT 数据检索（所有权、元数据、属性、交易历史等）。

Transfers API： 追踪代币（ERC-20, ERC-721, ERC-1155）转账。

Notify API： 基于 Webhook 的实时事件通知（地址活动、合约事件、挖矿交易等）。

Mempool API： 实时访问交易内存池数据。

Debug API： 深入分析交易失败原因。

Trace API： 追踪智能合约内部调用。

#### 3.开发者工具

仪表盘： 监控应用性能、请求量、错误率。

Composer： 可视化构建和测试 RPC 请求。

Mempool Visualizer： 可视化内存池交易。

Explorer： 查看区块、交易、地址信息（类似 Etherscan，但更侧重开发者视角）。

#### 4.SDK

提供多种语言的 SDK（如 Javascript, Python）简化集成

#### 5.强大的文档和教程： 丰富的学习资源

## 🧪 二、如何使用 Alchemy 连接和操作区块链测试网？

Alchemy 支持众多区块链的测试网。以下以以太坊的 Sepolia 测试网为例，说明基本步骤：

### 步骤 1：创建 Alchemy 账户

访问 <https://www.alchemy.com/>

点击 “Sign Up” 注册一个免费账户。

### 步骤 2：创建 App 并获取 API Key (HTTP URL)

登录后进入仪表盘。

点击 “Create App”。

填写 App 信息：

- Name: 给你的应用起个名字（例如 “My Sepolia Test App”）。
- Description: 可选描述。
- Environment: 选择 “Development”。
- Chain: 选择 “Ethereum”。
- Network: 选择 “Sepolia”。

点击 “Create App”。创建成功后，你的 App 会出现在仪表盘列表中。

点击你的 App 名字进入详情页。

点击右上角的 “View Key”。你会看到你的 HTTPS URL，格式类似：<https://eth-sepolia.g.alchemy.com/v2/your-unique-api-key。这就是你连接> Sepolia 测试网的关键凭证！ 🔑 请妥善保管，不要泄露。

### 步骤 3：配置你的开发环境

#### 在代码中使用（例如 web3.js / ethers.js）

```javascript
// 使用 ethers.js 示例
const { ethers } = require("ethers");

// 用你的 Alchemy Sepolia URL 替换下面的占位符
const alchemyUrl = "<https://eth-sepolia.g.alchemy.com/v2/your-unique-api-key>";
const provider = new ethers.providers.JsonRpcProvider(alchemyUrl);

// 现在你可以使用 provider 与 Sepolia 测试网交互了
async function main() {
  const blockNumber = await provider.getBlockNumber();
  console.log("Current Sepolia block number:", blockNumber);
}
main();
```

#### 在 MetaMask 中使用

打开 MetaMask 钱包。

确保网络选择器当前显示的是 “Sepolia Test Network”。如果不是，请添加或切换到 Sepolia。

在 MetaMask 设置中找到 “Networks” -> “Sepolia”。

将 “New RPC URL” 替换为你的 Alchemy Sepolia URL (<https://eth-sepolia.g.alchemy.com/v2/your-unique-api-key)。>

保存设置。现在你的 MetaMask 钱包将通过 Alchemy 节点连接 Sepolia 测试网，速度和可靠性通常比公共节点更好。

### 步骤 4：获取测试网代币（Testnet Faucet）

大多数测试网（如 Sepolia）需要免费的测试代币（ETH 或对应链的 Gas 代币）来支付交易 Gas 费。

Alchemy 提供了强大的 Sepolia Faucet！ 🚰

- 访问 Alchemy Sepolia Faucet: <https://sepoliafaucet.com/>
- 输入你的 Sepolia 测试网钱包地址（MetaMask 中当前显示的那个地址）。
- 完成人机验证（通常是 Captcha）。
- 点击 “Send Me ETH”。稍等片刻（通常几秒到几分钟），你的钱包就会收到 0.5 Sepolia ETH。

其他测试网（如 Goerli - 已弃用但部分项目还在用、Mumbai/Polygon, Arbitrum Goerli 等）也有各自的 Faucet，Alchemy 通常也提供链接或自己运营部分 Faucet。

## 💡 三、使用实例展示：通过 Alchemy 在 Sepolia 测试网上部署一个简单的 NFT 合约并铸造一个 NFT

### 前提条件

已完成上述步骤（有 Alchemy 账户、App、API Key）。

MetaMask 配置好 Sepolia 网络并使用 Alchemy 节点 URL。

通过 Alchemy Faucet 获得了 Sepolia ETH。

安装了 Node.js, npm/yarn。

安装了开发框架（这里用 Hardhat）。

### 步骤

#### 初始化 Hardhat 项目

```bash
mkdir my-nft-project
cd my-nft-project
npm init -y
npm install --save-dev hardhat
npx hardhat

# 选择 "Create a JavaScript project"，接受默认选项

npm install --save-dev @nomicfoundation/hardhat-toolbox
npm install @openzeppelin/contracts # 用于标准 NFT 合约
```

#### 配置 Hardhat (hardhat.config.js)

```javascript
require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config(); // 推荐使用 dotenv 管理密钥

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.20", // 使用与合约匹配的 Solidity 版本
  networks: {
    sepolia: {
      url: process.env.ALCHEMY_SEPOLIA_URL, // 从 .env 文件读取
      accounts: [process.env.SEPOLIA_PRIVATE_KEY] // 从 .env 文件读取（你的 MetaMask Sepolia 账户私钥）
    }
  }
};
```

#### 创建 .env 文件

```text
ALCHEMY_SEPOLIA_URL="<https://eth-sepolia.g.alchemy.com/v2/your-unique-api-key>"
SEPOLIA_PRIVATE_KEY="your-sepolia-metamask-account-private-key" # 注意：极其敏感！不要提交到Git！
```

#### 编写简单 NFT 合约 (contracts/MyNFT.sol)

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyNFT is ERC721, Ownable {
    uint256 private _nextTokenId;

    constructor()
        ERC721("MyNFT", "MNFT")
        Ownable(msg.sender)
    {}

    function safeMint(address to) public onlyOwner {
        uint256 tokenId = _nextTokenId++;
        _safeMint(to, tokenId);
    }
}
```

#### 编写部署脚本 (scripts/deploy.js)

```javascript
const hre = require("hardhat");

async function main() {
  const MyNFT = await hre.ethers.getContractFactory("MyNFT");
  const myNFT = await MyNFT.deploy();

  await myNFT.waitForDeployment();

  console.log(
    `MyNFT deployed to ${myNFT.target}`
  );
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
```

#### 部署合约到 Sepolia 测试网

```bash
npx hardhat run scripts/deploy.js --network sepolia
```

执行此命令。

Hardhat 会使用 Alchemy 提供的节点连接 Sepolia 网络。

它会使用你 .env 文件中指定的私钥账户发起部署交易（需要支付 Gas，所以确保有 Sepolia ETH）。

稍等片刻，控制台会输出部署成功的合约地址，例如：MyNFT deployed to 0x1234567890AbCdEf1234567890aBcDeF12345678。复制这个地址。

#### 7.铸造一个 NFT (scripts/mint.js)

```javascript
const hre = require("hardhat");

async function main() {
  // 替换为你的合约部署地址
  const contractAddress = "0x1234567890AbCdEf1234567890aBcDeF12345678";
  const MyNFT = await hre.ethers.getContractAt("MyNFT", contractAddress);

  // 替换为你想要接收 NFT 的钱包地址（可以是部署者自己或其他地址）
  const recipientAddress = "0xYourRecipientAddressHere";

  // 调用 safeMint 函数
  const mintTx = await MyNFT.safeMint(recipientAddress);
  await mintTx.wait(); // 等待交易确认

  console.log(`Successfully minted an NFT to ${recipientAddress}!`);
  console.log(`Transaction hash: ${mintTx.hash}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
```

#### 8.执行铸造脚本

```bash
npx hardhat run scripts/mint.js --network sepolia
```

执行此命令。

Hardhat 再次通过 Alchemy 连接 Sepolia。

它调用你之前部署的合约的 safeMint 方法，向指定地址铸造一个 NFT。

控制台会输出交易哈希和成功信息。

#### 9.验证结果

打开一个支持 Sepolia 的区块浏览器，如 Sepolia Etherscan。

粘贴你部署合约的交易哈希或合约地址，查看合约是否成功部署。

粘贴你铸造 NFT 的交易哈希，查看 Mint 操作是否成功。

查看接收 NFT 的地址，确认 NFT 是否已到达（在区块浏览器的 “Token” 或 “NFTs” 部分查看）。

## ✅ 总结

Alchemy 是一个强大的平台，通过提供高性能节点、增强型 API 和开发者工具，极大地简化了与各种区块链（包括主网和测试网）的交互过程。

使用 Alchemy 连接测试网（如 Sepolia）的核心是：注册账号 -> 创建 App -> 获取对应网络的 API Key (URL) -> 在代码或 MetaMask 中使用该 URL。

Alchemy Faucet 是获取测试网代币（如 Sepolia ETH）的便捷途径。

通过 Hardhat（或其他框架）结合 Alchemy URL，你可以轻松地在测试网上部署智能合约、调用合约函数、发送交易，进行完整的 dApp 开发和测试。
