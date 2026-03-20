# solidity_learning_notes 目录规划

当前建议采用“主线 + 专题”结构：  
- `solidity/` 保留通用主线（语法、EVM、合约设计、安全基础、标准）  
- `openzeppelin/` 独立为专题目录（笔记与代码同目录管理）  
- `uniswap/` 独立为专题目录（协议机制与实战代码同目录管理）  

---

## 目录结构

```text
solidity_learning_notes/
├── readme.md
├── solidity/
│   ├── 00_前言/
│   ├── 01_solidity基础/
│   ├── 02_合约设计和功能/
│   ├── 03_高级编程技巧/
│   ├── 04_智能合约进阶/
│   ├── 05_高级合约应用/
│   ├── 06_ERC标准和代理升级/
│   ├── 07_开发框架_hardhat/
│   └── 08_常用库和框架/
│       └── 02_foundry/
├── openzeppelin/
│   ├── readme.md
│   ├── notes/
│   │   ├── 00_总览与安装.md
│   │   ├── 01_ERC20_ERC721_ERC1155.md
│   │   ├── 02_AccessControl_Ownable.md
│   │   ├── 03_Pausable_ReentrancyGuard.md
│   │   ├── 04_Upgradeable_Proxy.md
│   │   └── 05_Governor_治理.md
│   ├── examples/
│   │   ├── erc20-access-control/
│   │   ├── nft-mint-role/
│   │   └── upgradeable-uups-demo/
│   └── assets/
│       ├── images/
│       └── diagrams/
└── uniswap/
    ├── readme.md
    ├── notes/
    │   ├── 00_总览.md
    │   ├── 01_AMM与恒定乘积.md
    │   ├── 02_核心合约Factory_Pair.md
    │   ├── 03_Router与Swap路径.md
    │   ├── 04_LP与手续费机制.md
    │   ├── 05_价格预言机与TWAP.md
    │   ├── 06_V2与V3关键差异.md
    │   └── 07_安全与常见坑.md
    ├── examples/
    │   ├── v2-swap-demo/
    │   ├── add-remove-liquidity-demo/
    │   └── twap-demo/
    └── assets/
        ├── images/
        └── diagrams/
```

---

## 迁移策略（最小改动）

1. `solidity/08_常用库和框架/01openzeppelin/` 的文档逐步迁移到 `openzeppelin/notes/`。  
2. `solidity/09_uniswap详解/` 的文档逐步迁移到 `uniswap/notes/`。  
3. 对应 `.sol` 示例迁移到各自 `examples/`，每个示例目录添加 `README.md`（目标、依赖、运行步骤）。  
4. 暂不一次性重命名全部历史文件，先新增结构并按学习进度逐步迁移。  



