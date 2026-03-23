# 四周学习计划：OpenZeppelin、链上支付/归集、Uniswap 基础

---

## 1. 具体学习内容（按周）

### 第 1 周：ERC20、权限与安全基线

| 时段 | 主题 | 具体学什么 |
|------|------|------------|
| Day 1-2 | ERC20 与授权 | `transfer` / `transferFrom` / `approve`（可选 `permit`）、无限授权风险；`_mint` / `_burn` / `_update`（或 OZ 版本差异下的内部转账钩子）；decimals 与常见扩展（如 `ERC20Burnable`） |
| Day 3-4 | 权限模型 | `Ownable` vs `AccessControl`；`DEFAULT_ADMIN_ROLE`、自定义 `bytes32` 角色、`grantRole` / `revokeRole`；多角色与最小权限拆分思路 |
| Day 5-7 | 安全基线 | `Pausable`（暂停粒度：全停 vs 按函数拆分）；`ReentrancyGuard` 与外部调用顺序；（可选）`Counters` 或自增 id 写法；**SafeERC20** 动机（低层 `call`、非标准 ERC20） |

**本周编码侧建议最小闭环**：一个可 mint 的测试代币 + 一个带 `SWEEPER` 角色、可 `pause` 的「金库/归集」入口合约骨架（函数可先仅占位）。

---

### 第 2 周：OZ 工具、SafeERC20 实践、合约骨架深化

| 时段 | 主题 | 具体学什么 |
|------|------|------------|
| Day 8-9 | ERC20 扩展（按需） | `ERC20Snapshot` / `ERC20Votes`：**何时需要**（审计面、Gas、存储）；与「支付卡」无强绑定时可略读或延后 |
| Day 10-11 | 工具库 | `SafeERC20` 全接口习惯用法；`Math`（`mulDiv` 等）；`Address`（合约码大小、`functionCall` 场景） |
| Day 12-14 | 支付/归集骨架 | `IERC20` + `AccessControl` + `Pausable` + `ReentrancyGuard` 组合；事件设计（`Deposit` / `Withdraw` / `Sweep` 等占位）；**一种**归集路径写通：`transfer` 入金 **或** `approve + transferFrom` sweep（二选一先跑通） |

---

### 第 3 周：Uniswap V2 与「资金搬运」相关读源码

| 时段 | 主题 | 具体学什么 |
|------|------|------------|
| Day 15-16 | Uniswap V2 Core | `Factory` + `Pair`：`mint` / `burn` / `swap` / `sync` / `skim` 语义；**恒定乘积**与价格如何随储备变化（与本笔记 `uniswap/notes/` 对照） |
| Day 17-18 | Uniswap V2 Periphery | `Router` 调用链；`TransferHelper`（`safeTransfer` / `safeApprove` 历史习惯，对照 `SafeERC20`） |
| Day 19-21 | 综合阅读 | 结合 `uniswap/notes/` 顺序阅读；整理「链上资金归集」与 **AMM 多余余额处理**（如 `skim`）**概念差异**，避免业务误判 |

---

### 第 4 周：测试、脚本、可重复部署

| 时段 | 主题 | 具体学什么 |
|------|------|------------|
| Day 22-24 | 合约收尾 | 补齐权限边界用例；Gas  obvious 热点（循环、存储读写）；（可选）简单升级 / 非升级两套策略对比（仅笔记层面也可） |
| Day 25-27 | 测试 | **Foundry**：单测 + 模糊测试入门；**Hardhat**：与前端/脚本习惯一致的集成测试；至少覆盖：正常路径、暂停、`revert`、权限错误 |
| Day 28-29 | 部署脚本 | `forge script` 或 Hardhat deploy；多网络参数（`.env` / `foundry.toml` / `hardhat.config`）；一次「本地 / 测试网」可重复部署流程 |

---

## 2. 根据学习内容：前置基础 & 需攻克的点

> 说明：若你原意是「需要**产出**什么」，请直接看 **第 5 节** 的目录结构；本节对应「先把哪些基础补齐、哪些坑要专门练」。

### 2.1 建议具备的前置（否则会卡壳）

| 模块 | 前置 |
|------|------|
| ERC20 / SafeERC20 | 会读 `external` / `public`、`view`；理解 ** allowance** 与 **`transferFrom` 调用方** |
| AccessControl | 理解 `msg.sender`、**角色只存地址**、admin 链（谁给谁授角色） |
| ReentrancyGuard | 理解 **外部调用点**（尤其对 `token.call`、任意代币） |
| Uniswap V2 | 基本 **储备金** 概念；会读 `interface` 与事件；会用区块浏览器对照一笔 `swap` |
| 测试 / 脚本 | 会跑 `forge test`、`npx hardhat test`；会配 RPC 与私钥**不进仓库** |

### 2.2 学习过程中建议专门「攻克」的难点

| 难点 | 为什么要攻克 |
|------|----------------|
| `approve` + `transferFrom` 与用户心理模型不一致 | 归集、第三方代付、批量操作都绕不开 |
| 暂停粒度与「谁可停」 | 真实系统常因一锅 `pause()` 导致纠纷；需在练习里拆开想 |
| 非标准 ERC20（假返回、手续费币） | `SafeERC20` 存在的理由；否则归集容易「以为成功其实没转」 |
| Pair `swap` 的 `amount0Out/amount1Out` 与回调顺序 | 读 V2 时最容易浮在表面；需对照源码一行行跟 |
| Foundry vs Hardhat 分工 | 避免两周后两套工具都半吊子；建议 **测试主用 Foundry、脚本/集成按队伍习惯二选一深钻** |

---

## 3. 何时初始化项目？笔记怎么记？

### 3.1 工程（合约仓库）什么时候建？

| 阶段 | 建议 |
|------|------|
| **第 1 周 Day 1** | **就初始化一个空仓库**（单一 `learning-payments-lab` 即可）：统一放合约、测试、脚本。早建 = 早养成「改一点就跑测试」的习惯。 |
| **不要等「学完再建」** | OZ / Uniswap 以读文档与读源码为主时，也仍然需要本地小合约验证 `SafeERC20`、`AccessControl` 行为；没有工程很容易只停留在复制粘贴。 |
| **第 3 周** | 可以不新建仓库，在**同一仓库**下加 `reference/uniswap-v2` 子模块或单独 clone 只读；笔记里写「对照行号」比大段粘贴源码更清晰。 |

### 3.2 笔记什么时候写？

| 做法 | 建议 |
|------|------|
| **随学随记（推荐）** | 每学完一个小topic，在本仓库 `solidity_learning_notes/` 或你自建的 `notes/`（见第 5 节）里补一页：标题 + 要点列表 + 链接；耗时 10～20 分钟。 |
| **周末整理** | 把零散的 bullet 合并成「一周一篇」总览，防止月后找不到。 |
| **代码旁注** | 复杂逻辑在仓库 `docs/decisions/` 用短 ADR（1/4 页）记「为什么这样设计」，比长篇教程有效。 |

**结论**：**第 1 周开始同时维护「一个练习工程 + 笔记目录」**；是否每天写长文无所谓，但要 **有固定落盘路径**。

---

## 4. 学习资料清单

### 4.1 官方文档（优先）

- Solidity：https://docs.soliditylang.org/  
- OpenZeppelin Contracts：https://docs.openzeppelin.com/contracts  
- OpenZeppelin Wizard（生成模板）：https://wizard.openzeppelin.com/  
- Hardhat：https://hardhat.org/docs  
- Foundry Book：https://book.getfoundry.sh/  

### 4.2 Uniswap V2（文档 + 数学 + 源码）

- Uniswap V2 概述（文档站，含概念与接口）：https://docs.uniswap.org/contracts/v2/overview  
- 白皮书（恒定乘积与手续费直觉，可与源码对照）：https://uniswap.org/whitepaper.pdf  
- v2-core（Factory / Pair）：https://github.com/Uniswap/v2-core  
- v2-periphery（Router 等）：https://github.com/Uniswap/v2-periphery  

### 4.3 本仓库内笔记路径（对齐学习顺序）

- OpenZeppelin：`solidity_learning_notes/openzeppelin/notes/`  
- Uniswap：`solidity_learning_notes/uniswap/notes/`（阅读顺序见 `solidity_learning_notes/uniswap/readme.md`）  
- Solidity 主线：`solidity_learning_notes/solidity/`（按需查阅 ERC、代理、框架章节）

### 4.4 补充读物（选）

- EIP-20：https://eips.ethereum.org/EIPS/eip-20  
- Solidity by Example（短例）：https://solidity-by-example.org/  

---

## 5. 建议的学习产出目录结构

以下内容可放在 **独立练习仓库** 根目录（与 `solidity_learning_notes` 并排），或放在本 mono-repo 的 `examples/` 下 —— **二选一，不要散落成多个无名文件夹**。

```text
learning-payments-lab/
├── README.md                 # 本周目标、如何运行测试/脚本
├── foundry.toml              # 若用 Foundry
├── hardhat.config.ts         # 若用 Hardhat（可与 Foundry 同仓）
├── contracts/
│   ├── token/                # 练习用 ERC20
│   ├── vault/                # 金库 / sweep 练习
│   └── playground/           # 随手试 OZ API 的小合约
├── test/                     # forge 或 hardhat 测试
├── script/                   # 部署与归集 smoke 脚本
├── docs/
│   ├── week-01.md … week-04.md   # 每周一页复盘（链接到下面 notes）
│   └── decisions/                 # 可选：ADR 简短设计决策
└── notes/                    # 可选：与 solidity_learning_notes 二选一
    ├── 01-erc20-allowance.md
    ├── 02-access-control.md
    ├── 03-pausable-reentrancy.md
    ├── 04-safeerc20-sweep.md
    └── 05-uniswap-v2-readthrough.md
```

若你 **只愿意维护一处笔记**：可直接把 `docs/week-*.md` 写成指向 `solidity_learning_notes/plan.md` 与 `openzeppelin/`、`uniswap/` 子笔记的索引，避免重复粘贴。

---

## 6. Uniswap V2：只读源码够不够？还要什么？到什么程度算「学会」？

### 6.1 是否只需要读源码？

**对多数人：不够。** 源码回答「怎么写」，但往往不够回答「为什么这样、边界条件、和链上真实行为是否一致」。

| 学习方式 | 作用 |
|----------|------|
| 读 v2-core / v2-periphery 源码 | 掌握实现细节与调用关系（主路径） |
| 白皮书 + 恒定乘积 / 手续费数学 | 理解价格、储备、输出的**因果**，避免死记代码 |
| 官方/社区文档与图解 | 对齐 Factory / Pair / Router **职责与调用顺序** |
| 自建笔记 + 资金流/调用图 | 把「谁转给谁」画清楚，比长时间盯着代码高效 |
| 浏览器看真实 `swap` / 加流动性 交易 | 用事件与状态变化**验证书本与源码理解** |
| 本地 fork + 小段脚本 | 改储备、跑一笔 swap，验证对 `getAmountOut` 等的直觉 |

**建议组合**：**源码 + 数学/文档 + 链上观测 + 小实验**，并与本仓库 `uniswap/notes/` 顺序对照。

### 6.2 掌握程度自检（L1 → L3）

| 层级 | 标准（能自测） |
|------|----------------|
| **L1：能讲清楚** | 说清 Pair 中 `mint` / `burn` / `swap` / `sync` / `skim` 各自解决什么问题；在恒定乘积下能根据储备与一侧输入**推导或核对**另一侧输出关系（不必死背公式）；理解 Factory、`createPair`、`INIT_CODE_HASH` 与 Router 的关系 |
| **L2：能跟调用链** | 从 Router 的 `swap*` / `addLiquidity` 跟到 Pair，说明代币流经顺序与最终余额；理解**滑点、最小输出、deadline** 为何存在；知道非标准 ERC20、fee-on-transfer 与 `TransferHelper` / `SafeERC20` 动机的关联 |
| **L3：能安全集成** | 能写**对接型**合约：拉报价、发起 swap、加/撤流动性（按需）、读储备；权限最小化、重入与回调顺序有数、关键路径有测试或 fork 验证 |

**「学会」落地标准**：对大多数人指 **达到 L2～L3**——能独立做集成与排错，而不是背下每一行实现。

### 6.3 「能写合格合约」指什么？

需先明确目标类型：

| 目标 | 需要学过的 V2 深度 |
|------|---------------------|
| **对接 Router 做 swap、路由、归集/换币** | 重点：**Router + Pair 接口 + 安全转账**；Pair 内部可精读 `swap` 相关段即可 |
| **流动性策略、自算报价、牵涉回调** | 接近 **Pair `swap` 细节 + 数学** |
| **从零做 DEX 或深度魔改 Pair** | 需 **几乎通读 Core + 完整测试与审计意识**；只读源码只是起点 |

对常见场景：**合格 = 会正确集成 + 会测 + 知道边界与代币怪异行为**，而不是重写整套 Uniswap。

### 6.4 与第 3 周计划如何对齐

建议每条笔记配套：**（1）对照源码 （2）对照一笔链上交易 （3）本地一句/小段验证**。能独立画出「入金 → swap → 出金」的资金流图且不自相矛盾，一般即达到 **L2**，可开始写「对接型」合约。

---

## 附录（可选）：架构选型自检表

仅在学习后期、需要收口「长期选哪种合约形态」时使用。

| 待确认 | 可选方向 | 备注 |
|--------|----------|------|
| 余额模型 | 链上 ERC20 / 内部账本 / 链下账 + 链上结算 | 影响合约边界 |
| 归集路径 | 用户直转金库 / `approve + sweep` / 暂存再批量转 | 影响 Gas、权限与用户体验 |
| 风控 | 全 pause / 分功能 pause / 额度限制 / 黑名单 | 影响 `AccessControl` 设计 |
| 升级 | 不可升级 / UUPS / Transparent | 影响存储布局与运维；第 4 周再定也可 |
