# 区块链生态参与方研究资料目录结构
# Blockchain Ecosystem Participants Research Directory Structure

本文档定义了区块链生态参与方深入分析与学习研究资料的推荐目录结构。

This document defines the recommended directory structure for in-depth analysis and research materials on blockchain ecosystem participants.

---

## 一、顶层目录结构 / Top-Level Directory Structure

```
blockchain-ecosystem-participants/          # 区块链生态参与方研究
├── 00-overview/                            # 00-总览
│   ├── 00.1-ecosystem-overview.md          # 00.1-生态概述
│   ├── 00.2-layered-architecture.md        # 00.2-分层架构
│   ├── 00.3-participant-classification.md  # 00.3-参与方分类
│   └── img/                                # 图表资源
│
├── layer-01-infrastructure/                # layer-01-基础层
│   ├── 01.1-overview.md                    # 01.1-层级概述
│   ├── 01.2-node-operators/                # 01.2-节点运营商
│   ├── 01.3-miners-validators/             # 01.3-矿工验证者
│   ├── 01.4-rpc-providers/                 # 01.4-RPC服务商
│   ├── 01.5-protocol-developers/           # 01.5-协议开发者
│   └── 01.6-blockchain-networks/           # 01.6-区块链网络
│
├── layer-02-assets-issuance/                # layer-02-资产与发行层
│   ├── 02.1-overview.md                    # 02.1-层级概述
│   ├── 02.2-native-assets/                 # 02.2-原生资产
│   ├── 02.3-token-issuers/                 # 02.3-通证发行商
│   ├── 02.4-stablecoin-issuers/            # 02.4-稳定币发行商 ⭐
│   └── 02.5-custodians/                    # 02.5-资产托管方
│
├── layer-03-interaction-application/       # layer-03-交互与应用层 ⭐
│   ├── 03.1-overview.md                    # 03.1-层级概述
│   ├── 03.2-cex/                           # 03.2-中心化交易所
│   ├── 03.3-dex/                           # 03.3-去中心化交易所
│   ├── 03.4-wallet-providers/              # 03.4-钱包服务商
│   ├── 03.5-defi-protocols/                # 03.5-DeFi协议
│   ├── 03.6-bridge-cross-chain/           # 03.6-桥与跨链服务
│   └── 03.7-application-platforms/         # 03.7-应用平台 ⭐
│       └── p2p-platform/                   # P2P平台（我们的项目）
│
├── layer-04-auxiliary-services/            # layer-04-辅助与服务层
│   ├── 04.1-overview.md                    # 04.1-层级概述
│   ├── 04.2-data-analytics/                # 04.2-数据分析机构
│   ├── 04.3-audit-security/                # 04.3-审计与安全机构
│   ├── 04.4-developer-tools/              # 04.4-开发者工具
│   ├── 04.5-media-community/               # 04.5-媒体与社区
│   ├── 04.6-kyc-providers/                # 04.6-KYC服务商
│   └── 04.7-price-providers/              # 04.7-价格服务商
│
├── layer-05-regulation-compliance/         # layer-05-监管与合规层
│   ├── 05.1-overview.md                   # 05.1-层级概述
│   ├── 05.2-regulatory-agencies/           # 05.2-各国监管机构
│   ├── 05.3-international-orgs/            # 05.3-国际组织
│   └── 05.4-compliance-providers/         # 05.4-合规服务商
│
├── cross-layer-analysis/                   # cross-layer-跨层分析
│   ├── relationships/                     # 参与方关系
│   ├── interactions/                      # 交互机制
│   ├── workflows/                         # 业务流程
│   └── case-studies/                      # 案例研究
│
├── resources/                              # resources-参考资料
│   ├── authoritative-reports/              # 权威报告
│   ├── industry-standards/                # 行业标准
│   ├── research-papers/                   # 研究论文
│   └── official-docs/                     # 官方文档
│
└── tools/                                  # tools-工具与模板
    ├── templates/                          # 文档模板
    ├── diagrams/                           # 图表模板
    └── checklists/                         # 检查清单
```

---

## 二、各层级详细目录结构 / Detailed Directory Structure by Layer

### 2.1 基础层 / Layer 01: Infrastructure

```
layer-01-infrastructure/
├── 01.1-overview.md                       # 层级概述
├── 01.2-node-operators/                   # 节点运营商
│   ├── overview.md                        # 概述
│   ├── major-providers/                   # 主要服务商
│   │   ├── infura.md                      # Infura
│   │   ├── alchemy.md                     # Alchemy
│   │   └── quicknode.md                   # QuickNode
│   ├── technical-analysis/                # 技术分析
│   └── market-analysis/                   # 市场分析
│
├── 01.3-miners-validators/                # 矿工/验证者
│   ├── overview.md                        # 概述
│   ├── pow-miners/                        # PoW矿工
│   ├── pos-validators/                    # PoS验证者
│   ├── consensus-mechanisms/              # 共识机制
│   └── economics/                         # 经济学分析
│
├── 01.4-rpc-providers/                    # RPC服务商
│   ├── overview.md                        # 概述
│   ├── providers/                         # 服务商列表
│   └── api-comparison/                    # API对比
│
├── 01.5-protocol-developers/              # 协议开发者
│   ├── overview.md                        # 概述
│   ├── ethereum-foundation/               # 以太坊基金会
│   ├── tron-foundation/                   # TRON基金会
│   └── governance/                        # 治理机制
│
└── 01.6-blockchain-networks/              # 区块链网络
    ├── overview.md                        # 概述
    ├── ethereum/                          # 以太坊
    ├── base/                              # Base
    ├── tron/                              # TRON
    ├── polygon/                           # Polygon
    └── comparison/                        # 网络对比
```

### 2.2 资产与发行层 / Layer 02: Assets & Issuance

```
layer-02-assets-issuance/
├── 02.1-overview.md                       # 层级概述
├── 02.2-native-assets/                    # 原生资产
│   ├── overview.md                        # 概述
│   ├── bitcoin/                           # 比特币
│   ├── ethereum/                          # 以太坊
│   └── other-native-assets/               # 其他原生资产
│
├── 02.3-token-issuers/                    # 通证发行商
│   ├── overview.md                        # 概述
│   ├── erc20-tokens/                      # ERC-20代币
│   ├── trc20-tokens/                      # TRC-20代币
│   └── issuance-mechanisms/               # 发行机制
│
├── 02.4-stablecoin-issuers/               # 稳定币发行商 ⭐
│   ├── overview.md                        # 概述
│   ├── tether-usdt/                       # Tether (USDT)
│   │   ├── overview.md                    # 概述
│   │   ├── issuance-mechanism.md          # 发行机制
│   │   ├── reserve-transparency.md         # 储备透明度
│   │   ├── regulatory-compliance.md       # 监管合规
│   │   └── market-analysis.md             # 市场分析
│   ├── circle-usdc/                       # Circle (USDC)
│   ├── makerdao-dai/                      # MakerDAO (DAI)
│   ├── paxos-busd/                        # Paxos (BUSD)
│   ├── comparison/                        # 稳定币对比
│   └── regulatory-landscape/             # 监管环境
│
└── 02.5-custodians/                       # 资产托管方
    ├── overview.md                        # 概述
    ├── coinbase-custody/                  # Coinbase Custody
    ├── bitgo/                             # BitGo
    └── custody-solutions/                 # 托管解决方案
```

### 2.3 交互与应用层 / Layer 03: Interaction & Application ⭐

```
layer-03-interaction-application/
├── 03.1-overview.md                       # 层级概述
├── 03.2-cex/                              # 中心化交易所
│   ├── overview.md                        # 概述
│   ├── binance/                           # 币安
│   ├── coinbase/                          # Coinbase
│   ├── okx/                               # OKX
│   ├── business-model/                    # 商业模式
│   └── regulatory-compliance/            # 监管合规
│
├── 03.3-dex/                              # 去中心化交易所
│   ├── overview.md                        # 概述
│   ├── uniswap/                           # Uniswap
│   ├── curve/                             # Curve
│   ├── pancakeswap/                       # PancakeSwap
│   └── amm-mechanisms/                    # AMM机制
│
├── 03.4-wallet-providers/                 # 钱包服务商
│   ├── overview.md                        # 概述
│   ├── custodial-wallets/                 # 托管钱包
│   ├── non-custodial-wallets/             # 非托管钱包
│   │   ├── metamask/                      # MetaMask
│   │   ├── trust-wallet/                  # Trust Wallet
│   │   └── privy/                         # Privy
│   └── hardware-wallets/                 # 硬件钱包
│
├── 03.5-defi-protocols/                   # DeFi协议
│   ├── overview.md                        # 概述
│   ├── lending/                           # 借贷协议
│   │   ├── aave/                          # Aave
│   │   └── compound/                      # Compound
│   ├── derivatives/                       # 衍生品
│   │   └── dydx/                          # dYdX
│   └── yield-farming/                     # 收益挖矿
│
├── 03.6-bridge-cross-chain/               # 桥与跨链服务
│   ├── overview.md                        # 概述
│   ├── polygon-bridge/                    # Polygon Bridge
│   ├── arbitrum-bridge/                   # Arbitrum Bridge
│   └── cross-chain-mechanisms/            # 跨链机制
│
└── 03.7-application-platforms/             # 应用平台 ⭐
    ├── overview.md                        # 概述
    ├── p2p-platform/                      # P2P平台（我们的项目）
    │   ├── overview.md                    # 平台概述
    │   ├── positioning.md                 # 生态定位
    │   ├── core-features/                 # 核心功能
    │   │   ├── p2p-matching.md            # P2P撮合
    │   │   ├── escrow-service.md          # 托管服务
    │   │   ├── dispute-resolution.md      # 争议处理
    │   │   └── kyc-risk-control.md        # KYC风控
    │   ├── technical-architecture/         # 技术架构
    │   │   ├── frontend/                  # 前端
    │   │   ├── backend/                   # 后端
    │   │   ├── smart-contracts/           # 智能合约
    │   │   └── admin-platform/            # 管理平台
    │   ├── dependencies/                  # 依赖关系
    │   ├── competitive-analysis/          # 竞争分析
    │   └── business-workflows/            # 业务流程
    ├── nft-markets/                       # NFT市场
    │   └── opensea/                       # OpenSea
    └── gaming-platforms/                  # 游戏平台
        └── axie-infinity/                 # Axie Infinity
```

### 2.4 辅助与服务层 / Layer 04: Auxiliary & Services

```
layer-04-auxiliary-services/
├── 04.1-overview.md                       # 层级概述
├── 04.2-data-analytics/                   # 数据分析机构
│   ├── overview.md                        # 概述
│   ├── chainalysis/                       # Chainalysis
│   ├── nansen/                            # Nansen
│   └── analytics-tools/                   # 分析工具
│
├── 04.3-audit-security/                   # 审计与安全机构
│   ├── overview.md                        # 概述
│   ├── certik/                            # CertiK
│   ├── openzeppelin/                      # OpenZeppelin
│   └── security-standards/                # 安全标准
│
├── 04.4-developer-tools/                  # 开发者工具
│   ├── overview.md                        # 概述
│   ├── infura/                            # Infura
│   ├── alchemy/                           # Alchemy
│   ├── the-graph/                         # The Graph
│   └── tool-comparison/                   # 工具对比
│
├── 04.5-media-community/                  # 媒体与社区
│   ├── overview.md                        # 概述
│   ├── coindesk/                          # CoinDesk
│   ├── cointelegraph/                     # CoinTelegraph
│   └── community-platforms/               # 社区平台
│
├── 04.6-kyc-providers/                    # KYC服务商
│   ├── overview.md                        # 概述
│   ├── sumsub/                            # Sumsub
│   ├── jumio/                             # Jumio
│   ├── onfido/                            # Onfido
│   └── kyc-solutions/                     # KYC解决方案
│
└── 04.7-price-providers/                  # 价格服务商
    ├── overview.md                        # 概述
    ├── pyth-network/                      # Pyth Network
    ├── chainlink/                         # Chainlink
    ├── coingecko/                         # CoinGecko
    └── price-oracles/                     # 价格预言机
```

### 2.5 监管与合规层 / Layer 05: Regulation & Compliance

```
layer-05-regulation-compliance/
├── 05.1-overview.md                       # 层级概述
├── 05.2-regulatory-agencies/              # 各国监管机构
│   ├── overview.md                        # 概述
│   ├── united-states/                     # 美国
│   │   ├── sec/                           # SEC
│   │   └── cftc/                          # CFTC
│   ├── european-union/                    # 欧盟
│   │   └── mica/                          # MiCA法规
│   ├── singapore/                         # 新加坡
│   │   └── mas/                           # MAS
│   ├── uae/                               # 阿联酋
│   │   └── vara/                          # VARA
│   └── other-jurisdictions/              # 其他司法管辖区
│
├── 05.3-international-orgs/               # 国际组织
│   ├── overview.md                        # 概述
│   ├── fatf/                              # FATF
│   └── bis/                               # BIS
│
└── 05.4-compliance-providers/             # 合规服务商
    ├── overview.md                        # 概述
    ├── kyc-aml-solutions/                 # KYC/AML解决方案
    └── compliance-tools/                  # 合规工具
```

---

## 三、跨层分析目录 / Cross-Layer Analysis Directory

```
cross-layer-analysis/
├── relationships/                         # 参与方关系
│   ├── layer-relationships.md             # 层级关系
│   ├── participant-interactions.md        # 参与方交互
│   └── dependency-mapping.md              # 依赖关系映射
│
├── interactions/                          # 交互机制
│   ├── api-interfaces.md                  # API接口
│   ├── smart-contract-calls.md            # 智能合约调用
│   └── on-chain-off-chain.md              # 链上/链下交互
│
├── workflows/                             # 业务流程
│   ├── trading-flows/                    # 交易流程
│   ├── asset-issuance-flows/              # 资产发行流程
│   └── compliance-flows/                 # 合规流程
│
└── case-studies/                          # 案例研究
    ├── stablecoin-ecosystem/              # 稳定币生态案例
    ├── exchange-ecosystem/                # 交易所生态案例
    └── defi-ecosystem/                    # DeFi生态案例
```

---

## 四、参考资料目录 / Resources Directory

```
resources/
├── authoritative-reports/                 # 权威报告
│   ├── caict-blockchain-whitepaper/      # 中国信通院区块链白皮书
│   ├── tbi-web3-report/                  # 可信区块链Web3报告
│   ├── messari-reports/                   # Messari报告
│   └── other-reports/                     # 其他报告
│
├── industry-standards/                    # 行业标准
│   ├── technical-standards/               # 技术标准
│   └── regulatory-standards/             # 监管标准
│
├── research-papers/                       # 研究论文
│   ├── academic-papers/                   # 学术论文
│   └── industry-research/                 # 行业研究
│
└── official-docs/                         # 官方文档
    ├── protocol-docs/                     # 协议文档
    ├── api-docs/                          # API文档
    └── regulatory-docs/                    # 监管文档
```

---

## 五、命名规范 / Naming Conventions

### 5.1 文件夹命名规范

- **使用小写字母和连字符**：`layer-01-infrastructure`
- **层级编号**：使用 `layer-01` 到 `layer-05` 格式
- **子分类编号**：使用 `01.1`, `01.2` 等格式
- **英文命名**：优先使用英文，便于国际化
- **中文注释**：在目录文件中提供中文对照

### 5.2 文件命名规范

- **概述文件**：`overview.md`
- **具体参与方**：使用参与方名称，如 `tether-usdt.md`
- **分析文档**：使用描述性名称，如 `issuance-mechanism.md`
- **图表文件**：存放在 `img/` 子目录中

### 5.3 文档结构规范

每个参与方的详细分析文档应包含：
- **概述** (Overview)
- **核心功能** (Core Functions)
- **技术架构** (Technical Architecture)
- **商业模式** (Business Model)
- **监管合规** (Regulatory Compliance)
- **市场分析** (Market Analysis)
- **与其他参与方的关系** (Relationships)

---

## 六、使用建议 / Usage Recommendations

### 6.1 文件组织

1. **按层级分类**：所有资料按5层模型分类存放
2. **重点突出**：使用 ⭐ 标记重要参与方（如稳定币发行商、P2P平台）
3. **保持更新**：定期更新各参与方的分析文档
4. **版本控制**：使用Git管理文档版本

### 6.2 内容维护

1. **统一格式**：使用统一的Markdown格式
2. **图表资源**：所有图表存放在对应的 `img/` 目录
3. **交叉引用**：使用相对路径引用其他文档
4. **中英文对照**：重要术语提供中英文对照

### 6.3 协作规范

1. **文档模板**：使用 `tools/templates/` 中的模板
2. **检查清单**：参考 `tools/checklists/` 中的清单
3. **定期审查**：定期审查和更新文档内容

---

**文档版本**：v1.0  
**创建日期**：2026-01-24  
**最后更新**：2026-01-27  


