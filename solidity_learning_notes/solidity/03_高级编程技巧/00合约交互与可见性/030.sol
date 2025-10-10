// SPDX-License-Identifier: MIT
pragma solidity 0.8.3;
/*

2. 继承顺序的规则
◦ 从最基础的合约到派⽣合约的顺序
◦ 最基础的合约是继承最少的合约
3. 继承顺序⽰例
◦ ⽰例1：合约Z继承Y和X
■ Z继承Y和X，Y继承X，X没有继承其他合约
■ 顺序：X->Y->Z
◦ ⽰例2：合约Z继承B，B继承A，A继承X
■ 顺序：X ->Y-> A-> B-> Z
4. 多重继承的语法
◦ 使⽤ is 关键字声明继承
◦ 继承顺序：从最基础到最派⽣
◦ ⽰例：合约Z继承X和Y
contract Z is X, Y
5. 函数重写
◦ 重写单个⽗合约的函数
◦ 重写多个⽗合约的函数
■ 使⽤ override 关键字
■ 在括号内列出继承的合约
■
。⽰例：
▪ override (X, Y ) override (Y, X )

*/

// 继承顺序： 最上层的最优先

/*
7 X
8 / |
9 Y |
10 \ |
11 Z
12
13 // X , Y, Z
14
15 X
16 / \
17 Y A
18 | |
19 | B
20 \ /
21 Z
22
23 // X, Y, A, B, Z
24 */

contract X {
    function foo() public pure virtual returns (string memory) {
        return "X";
    }
    function bar() public pure virtual returns (string memory) {
        return "X";
    }

    // more code here
    function x() public pure returns (string memory) {
        return "X";
    }
}

contract Y is X {
    function foo() public pure virtual override returns (string memory) {
        return "Y";
    }
    function bar() public pure virtual override returns (string memory) {
        return "Y";
    }

    // more code here
    function y() public pure returns (string memory) {
        return "Y";
    }
}

contract Z is X, Y {
    function foo() public pure override(X, Y) returns (string memory) {
        return "Z";
    }
    function bar() public pure override(Y, X) returns (string memory) {
        return "Z";
    }
}
