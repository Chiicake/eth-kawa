// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract BasicERC20 {
    // 代币名称
    string public name;
    // 代币符号
    string public symbol;
    // 小数位数，通常为18
    uint8 public decimals = 18;
    // 总供应量
    uint256 public totalSupply;

    // 存储每个地址的代币余额
    mapping(address => uint256) public balanceOf;

    // 存储授权额度：owner允许spender花费的代币数量
    mapping(address => mapping(address => uint256)) public allowance;

    // 转账事件
    event Transfer(address indexed from, address indexed to, uint256 value);

    // 授权事件
    event Approval(address indexed owner, address indexed spender, uint256 value);

    // 构造函数：初始化代币名称、符号和初始供应量
    constructor(
        string memory _name,
        string memory _symbol,
        uint256 _initialSupply
    ) {
        name = _name;
        symbol = _symbol;
        // 考虑小数位数，实际发行量 = 初始供应 * 10^decimals
        totalSupply = _initialSupply * (10 **uint256(decimals));
        // 将初始代币分配给合约部署者
        balanceOf[msg.sender] = totalSupply;
        emit Transfer(address(0), msg.sender, totalSupply);
    }

    // 转账功能：从调用者地址向to地址转账value数量的代币
    function transfer(address to, uint256 value) public returns (bool success) {
        // 检查接收地址不为零地址
        require(to != address(0), "ERC20: transfer to the zero address");
        // 检查发送者余额是否充足
        require(balanceOf[msg.sender] >= value, "ERC20: insufficient balance");

        // 更新余额
        balanceOf[msg.sender] -= value;
        balanceOf[to] += value;

        // 触发转账事件
        emit Transfer(msg.sender, to, value);
        return true;
    }

    // 授权功能：允许spender从调用者地址花费value数量的代币
    function approve(address spender, uint256 value) public returns (bool success) {
        // 检查授权地址不为零地址
        require(spender != address(0), "ERC20: approve to the zero address");

        // 更新授权额度
        allowance[msg.sender][spender] = value;

        // 触发授权事件
        emit Approval(msg.sender, spender, value);
        return true;
    }

    // 代转账功能：从sender地址向recipient地址转账value数量的代币（需先授权）
    function transferFrom(
        address sender,
        address recipient,
        uint256 value
    ) public returns (bool success) {
        // 检查接收地址不为零地址
        require(recipient != address(0), "ERC20: transfer to the zero address");
        // 检查发送者余额是否充足
        require(balanceOf[sender] >= value, "ERC20: insufficient balance");
        // 检查授权额度是否充足
        require(allowance[sender][msg.sender] >= value, "ERC20: allowance exceeded");

        // 更新余额 
        balanceOf[sender] -= value;
        balanceOf[recipient] += value;
        // 更新授权额度
        allowance[sender][msg.sender] -= value;

        // 触发转账事件
        emit Transfer(sender, recipient, value);
        return true;
    }
}
