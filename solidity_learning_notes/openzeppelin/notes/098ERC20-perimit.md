# ERC20 permit

ERC20 的 permit 功能是 EIP-2612 标准的核心扩展，它通过链下签名实现代币授权，无需用户预先发送交易调用 approve。以下是关键点解析：

## 一、permit 的作用

### 免 Gas 预授权

用户通过链下签名授权第三方（如合约或服务商），第三方凭签名调用 permit 完成链上授权，用户无需支付 approve 的 Gas 费。

### 提升用户体验

可与实际交易（如转账、兑换）合并为单笔交易，避免两步操作（先 approve 再执行）。

## 二、permit 函数结构

```solidity
function permit(
    address owner,      // 代币持有者
    address spender,    // 被授权方
    uint256 value,      // 授权数量
    uint256 deadline,   // 过期时间戳
    uint8 v, bytes32 r, bytes32 s  // ECDSA 签名
) external;
```

## 三、核心实现机制

### 签名内容

用户对以下结构化数据签名（符合 EIP-712 标准）：

```solidity
struct Permit {
    address owner;
    address spender;
    uint256 value;
    uint256 nonce;    // 防重放攻击
    uint256 deadline;
}
```

### 关键验证步骤

合约内部验证逻辑：

```solidity
require(block.timestamp <= deadline, "Expired");
require(owner != address(0), "Invalid owner");

bytes32 digest = keccak256(
    abi.encodePacked(
        "\x19\x01",
        DOMAIN_SEPARATOR,  // 链/合约唯一标识
        keccak256(abi.encode(
            PERMIT_TYPEHASH,
            owner,
            spender,
            value,
            nonces[owner]++, // 递增 nonce
            deadline
        ))
    )
);
require(ecrecover(digest, v, r, s) == owner, "Invalid signature");

_approve(owner, spender, value); // 执行授权
```

### 安全组件

nonce：每个地址独立计数器，防止签名重放。
DOMAIN_SEPARATOR：包含链 ID、合约地址，防止跨链重放。
deadline：限制签名有效期。

## 四、使用场景示例

### DEX 交易（如 Uniswap）

用户签名授权 DEX 使用代币，DEX 在单笔交易中完成：

- 调用 permit 获取授权
- 调用 transferFrom 执行兑换

### 借贷协议

用户签名授权存款，合约在单笔交易中完成：

- permit 授权代币
- 存入代币到资金池

## 五、开发者注意事项

### 依赖 EIP-712

需实现 DOMAIN_SEPARATOR 和结构化数据哈希。

### 安全风险

- 前端需安全处理签名请求（防钓鱼）。
- 确保 deadline 在链上验证。

### OpenZeppelin 支持

使用现成库（如 ERC20Permit）：

```solidity
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract MyToken is ERC20, ERC20Permit {
    constructor() ERC20("Token", "TKN") ERC20Permit("Token") {}
}
```

## 六、用户视角

### 签名操作

钱包弹出签名请求（非交易），显示授权细节（金额、接收方、有效期）。

### 优势

- 节省 Gas 费（避免单独 approve）
- 操作原子性（授权+交易一步完成）

总结
permit 将 ERC20 授权流程从链上迁移至链下，通过密码学签名实现安全授权，显著优化 DeFi 交互体验。开发时应优先使用审计库（如 OpenZeppelin），用户需谨慎验证签名内容。

------

ERC20permit在ERC20的基础上，结合EIP-2612，添加了一个permit函数，允许用户通过EIP-712进行链下签名授权

- 授权你仅需用户在链下签名，减少一笔交易
- 签名之后，用户可以委托第三方进行后续交易，不需要持有ETH：用户A可以将签名发送给拥有gas的第三方，委托B来执行后续交易

链下加签，链上验签过程
![](./WeChat061e49fbefebad511e2d69a6a9160d40.jpg)

## 流程

1. 在openzeppelin的官网上，搜索ERC20permit，选择ERC20permit
<https://wizard.openzeppelin.com/>
生成代币协议，并将代币协议复制到hardhat中的contracts文件夹中
2. 用npx hardhat compile 编译代币协议

uniswap的样例代码在github上：<https://github.com/Uniswap/v3-periphery/test>
test文件夹下shared/permit.ts
直接复制在test文件夹下MyPermit.ts文件

3. 复制之后发现部分引入不存在，

```bash
yarn add ethers@^5.7.0
yarn add --dev @typechain/hardhat @typechain/ethers-v5

# 重新编译合约
npx hardhat clean
npx hardhat compile
```

4. 在MyPermit.ts文件，修改代码符合要求，并在改文件编写测试代码

5. 运行测试代码,测试通过之后将代码部署到本地测试网

```bash
npx hardhat test test --network localhost
```

在node窗口拿到Contract address:    0xb7f8bc63bbcad18155201308c8f3540b07f84f5e

6. 获取部署地址在remix上进行测试
   6.1 创建文件MyPermit.sol
   6.2 部署合约
   6.3 编译合约,环境选择 dev hardhat provider
       ![](./WeChat2e02a964030a842a7da0a9c89ce8c925.jpg)
    6.4 将前面的Contract address填入AT address，并点击
    6.5 此时在deployed contracts中查看合约mypermit
