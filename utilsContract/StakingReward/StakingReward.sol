// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "@openzeppelin/contracts/interfaces/IERC20.sol";

/*
质押收益

部署
1、先部署token1得到地址t1
2、部署token2得到地址t2
3、部署StakingRewards，将t1,t2传入deploy得到地址s1
4、设置奖励时间1000，执行setRewardsDuration
5、给质押收益设置奖励代币，在token2中，执行mint，给s1设置代币1000
6、设置奖励参数，setRewardsDuration,
7、设置奖励速率notifyRewardAmount 1000

切换账户a3进行质押
8、在token1合约中，执行mint对a3设置代币1000
在质押代币之前，必须先approve质押合约
9、在token1中，执行approve，给s1设置代币100
10、在StakingRewards中执行setRewardsDuration质押合约
11、用之前部署合约的账户，查看earned奖励
12、在token2中，执行balanceOf查看当前账户的token2的数量,为0
13、执行getReard收获奖励，每执行一次，步骤12的值会更新

# 总结：
updateReward属于更新的整个过程
notifyRewardAmount设置奖励的整个过程
质押，收益，奖励这部分内容比较常规
*/
contract StakingRewards {
    IERC20 public immutable stakingToken; // 质押token
    IERC20 public immutable rewardsToken; //收益token

    address public owner; // 管理员

    uint256 public duration; // 持续时间
    uint256 public finishAt;
    uint256 public updateAt;
    uint256 public rewardRate;
    uint256 public rewardPerTokenStored;

    mapping(address => uint256) public userRewardPerTokenPaid; // 质押用户当前每个代币的奖励数量
    mapping(address => uint256) public rewards; // 用户奖励的数值

    uint256 public totalSupply; // 总代币数量
    mapping(address => uint256) public balanceOf; // 质押用户的代币数额

    modifier onlyOwner() {
        require(msg.sender == owner, "not owner");
        _;
    }

    /*
    @function 更新每个用户当前的奖励数额，在质押，撤回，重新设置奖励参数时都会调用
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

    */
    modifier updateReward(address _account) {
        rewardPerTokenStored = rewardPerToken();
        updateAt = lastTimeRewardApplicable();

        if (_account != address(0)) {
            rewards[_account] = earned(_account);
            userRewardPerTokenPaid[_account] = rewardPerTokenStored;
        }
        _;
    }

    constructor(address _stakingToken, address _rewardsToken) {
        owner = msg.sender;
        stakingToken = IERC20(_stakingToken);
        rewardsToken = IERC20(_rewardsToken);
    }

    function rewardPerToken() public view returns (uint256) {
        if (totalSupply == 0) return rewardPerTokenStored;
        return
            rewardPerTokenStored +
            ((rewardRate * lastTimeRewardApplicable() - updateAt) * 1e18) /
            totalSupply;
    }

    function lastTimeRewardApplicable() public view returns (uint256) {
        return _min(block.timestamp, finishAt);
    }

    /*
    计算用户当前可领取的奖励总额
    */
    function earned(address _account) public view returns (uint256) {
        return
            (balanceOf[_account] *
                (rewardPerToken() - userRewardPerTokenPaid[_account])) /
            1e18 +
            rewards[_account];
    }

    function _min(uint256 x, uint256 y) private pure returns (uint256) {
        return x <= y ? x : y;
    }

    /*
    @function设置奖励金额和分配速率
        1. 当当前时间超过上一个奖励周期结束时间(finishAt)时，创建一个全新的奖励周期
        2. 当当前时间还在上一个奖励周期内时，将剩余奖励与新奖励合并计算新的分配速率
        3. 更新奖励结束时间（finishAt）为当前时间加上持续时间(duration)
          1. 当前时间：1600，finishAt：1500，amount：1000，duration：1000
    1. rewardRate：1
2. 当前时间：2000，finishAt：2600，amount：1000，duration：1000
    1. remainingRewards：600
    2. rewardRate：1.6
    */
    function notifyRewardAmount(
        uint256 _amount
    ) external onlyOwner updateReward(address(0)) {
        if (block.timestamp > finishAt) {
            rewardRate = _amount / duration;
        } else {
            uint256 remainingRewards = rewardRate *
                (finishAt - block.timestamp);
            rewardRate = (remainingRewards + _amount) / duration;
        }
        require(rewardRate > 0, "reward rate = 0");
        require(
            rewardRate * duration <= rewardsToken.balanceOf(address(this)),
            "reward amount >balance"
        );

        finishAt = block.timestamp + duration;
        updateAt = block.timestamp;
    }

    /*
     @function 质押
     @modify updateReward 需要更新奖励参数
     */
    function stake(uint _amount) public payable updateReward(msg.sender) {
        require(_amount > 0, "amount should be more than zero");
        // 将代币装入合约地址
        stakingToken.transferFrom(msg.sender, address(this), _amount);
        // 记录当前用户的质押金额
        balanceOf[msg.sender] += _amount;
        // 更新总的质押金额
        totalSupply += _amount;
    }

    /*
    @ function 撤销函数
    */
    function withdraw(
        uint256 _amount
    ) external payable updateReward(msg.sender) {
        require(_amount > 0, "withdraw amount should be more than zero");

        // 记录当前用户的质押金额
        balanceOf[msg.sender] -= _amount;
        // 更新总的质押金额
        totalSupply -= _amount;
        // 将代币转给用户
        stakingToken.transfer(msg.sender, _amount);
    }

    function getReward() external updateReward(msg.sender) {
        require(block.timestamp == finishAt, "no reward yet.");

        // 获取用户存储的奖励数额
        uint256 reward = rewards[msg.sender];

        // 不用减法，防止数据下溢
        if (reward > 0) {
            rewards[msg.sender] = 0;
            // 将奖励转给用户的账户
            rewardsToken.transfer(owner, reward);
        }
    }

    function setRewardsDuration(
        uint256 _duration
    ) public onlyOwner updateReward(msg.sender) {
        // 当前的时间周期还没结束时，不能设置周期
        require(finishAt < block.timestamp, "already rewarded.");
        duration = _duration;
    }

    // 设置奖励参数
    function mint() public payable {}
}
