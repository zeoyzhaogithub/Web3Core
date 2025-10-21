# 哈希函数（ keccak256）

在 Solidity 中，哈希函数（主要是 keccak256）是智能合约开发的核心工具，广泛应用于数据验证、安全机制和状态管理等领域。以下是主要应用场景及示例：

## 1. 数据完整性验证

场景：验证输入数据是否匹配预存储的哈希值（如密码、文件指纹）。

```solidity
mapping(address => bytes32) public userPasswords;

function setPassword(bytes32 hashedPassword) external {
    userPasswords[msg.sender] = hashedPassword;
}

function verifyPassword(string memory password) external view returns (bool) {
    bytes32 submittedHash = keccak256(abi.encodePacked(password));
    return userPasswords[msg.sender] == submittedHash;
}
```

## 2. 唯一标识生成

场景：为复杂数据创建唯一ID（如交易对、NFT元数据）。

```solidity
function generateTokenId(
    address creator,
    string memory name,
    uint256 serial
) external pure returns (bytes32) {
    return keccak256(abi.encodePacked(creator, name, serial));
}
```

## 3. 数字签名验证

场景：验证链下签名（遵循 EIP-191 标准）。

```solidity
function verifySignature(
    address signer,
    bytes32 messageHash,
    bytes memory signature
) external pure returns (bool) {
    bytes32 ethSignedHash = keccak256(
        abi.encodePacked("\x19Ethereum Signed Message:\n32", messageHash)
    );

    (bytes32 r, bytes32 s, uint8 v) = splitSignature(signature);
    return ecrecover(ethSignedHash, v, r, s) == signer;
}
```

## 4. 伪随机数生成

场景：生成难以预测的随机数（需注意矿工可操控风险）。

```solidity
uint256 private nonce;

function random() internal returns (uint256) {
    return uint256(keccak256(abi.encodePacked(
        block.timestamp,
        msg.sender,
        nonce++
    )));
}
```

## 5. 状态通道（Commit-Reveal 模式）

场景：隐藏敏感信息（如竞标价格），后期揭示验证。

```solidity
struct Commit {
    bytes32 commitment;
    bool revealed;
}

mapping(address => Commit) public commits;

// 提交承诺
function commitBid(bytes32 hashedBid) external {
    commits[msg.sender] = Commit(hashedBid, false);
}

// 揭示并验证
function revealBid(uint256 bid, bytes32 salt) external {
    require(keccak256(abi.encodePacked(bid, salt)) == commits[msg.sender].commitment);
    commits[msg.sender].revealed = true;
    // 处理有效出价
}
```

## 6. Merkle 树验证

场景：高效验证大型数据集中的成员资格（如空投白名单）。

```solidity
bytes32 public merkleRoot;

function claimAirdrop(
    uint256 amount,
    bytes32[] calldata merkleProof
) external {
    bytes32 leaf = keccak256(abi.encodePacked(msg.sender, amount));
    require(verifyMerkleProof(merkleProof, merkleRoot, leaf), "Invalid proof");
    // 发放代币
}

function verifyMerkleProof(
    bytes32[] memory proof,
    bytes32 root,
    bytes32 leaf
) internal pure returns (bool) {
    bytes32 currentHash = leaf;
    for (uint256 i = 0; i < proof.length; i++) {
        currentHash = currentHash < proof[i]
            ? keccak256(abi.encodePacked(currentHash, proof[i]))
            : keccak256(abi.encodePacked(proof[i], currentHash));
    }
    return currentHash == root;
}
```

## 7. 合约地址预测

场景：预先计算由 CREATE2 创建的合约地址。

```solidity
function predictAddress(
    address deployer,
    bytes32 salt,
    bytes memory bytecode
) external pure returns (address) {
    bytes32 hash = keccak256(
        abi.encodePacked(
            bytes1(0xff),
            deployer,
            salt,
            keccak256(bytecode)
        )
    );
    return address(uint160(uint256(hash)));
}
```

## 关键注意事项

### 哈希碰撞

使用 abi.encodePacked 时，不同参数组合可能产生相同哈希（如 ("a", "bc") 与 ("ab", "c")）。可通过 abi.encode 或添加分隔符解决。

### 随机数安全性

基于 block.timestamp/blockhash 的随机数易被矿工操控，关键场景应使用预言机（如 Chainlink VRF）。

### 签名重放攻击

签名消息应包含 address(this) 和 nonce 防止跨合约重放。

### Gas 优化

多次哈希计算可能消耗大量 Gas，建议链下预处理哈希值。

通过合理应用哈希函数，开发者可构建更安全、高效的智能合约，尤其在身份验证、数据验证和状态管理等场景中发挥关键作用。
