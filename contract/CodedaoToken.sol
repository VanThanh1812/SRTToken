pragma solidity ^0.4.18;


// ----------------------------------------------------------------------------

// 'CDS' 'Codedao Network Crystal' token contract

//

// Symbol      : CNC

// Name        : Codedao Network Crystal

// Total supply:

// Decimals    : 18

//

// Enjoy.

//

// (c) BokkyPooBah / Bok Consulting Pty Ltd 2017. The MIT Licence.

// ----------------------------------------------------------------------------



// ----------------------------------------------------------------------------

// Safe maths

// ----------------------------------------------------------------------------

library SafeMath {

    function add(uint a, uint b) internal pure returns (uint c) {

        c = a + b;

        require(c >= a);

    }

    function sub(uint a, uint b) internal pure returns (uint c) {

        require(b <= a);

        c = a - b;

    }

    function mul(uint a, uint b) internal pure returns (uint c) {

        c = a * b;

        require(a == 0 || c / a == b);

    }

    function div(uint a, uint b) internal pure returns (uint c) {

        require(b > 0);

        c = a / b;

    }

}



// ----------------------------------------------------------------------------

// ERC Token Standard #20 Interface

// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md

// ----------------------------------------------------------------------------

contract ERC20Interface {

    function totalSupply() public constant returns (uint);

    function balanceOf(address tokenOwner) public constant returns (uint balance);

    function allowance(address tokenOwner, address spender) public constant returns (uint remaining);

    function transfer(address to, uint tokens) public returns (bool success);

    function approve(address spender, uint tokens) public returns (bool success);

    function transferFrom(address from, address to, uint tokens) public returns (bool success);


    event Transfer(address indexed from, address indexed to, uint tokens);

    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);

}

// ----------------------------------------------------------------------------

// Owned contract

// ----------------------------------------------------------------------------
contract Ownable {
    address public owner;

    /**
     * @dev The Ownable constructor sets the original `owner` of the contract to the sender
     * account.
     */
    function Ownable() {
        owner = msg.sender;
    }


    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }


    /**
     * @dev Allows the current owner to transfer control of the contract to a newOwner.
     * @param newOwner The address to transfer ownership to.
     */
    function transferOwnership(address newOwner) onlyOwner {
        if (newOwner != address(0)) {
            owner = newOwner;
        }
    }

}

contract CodedaoNetwork is ERC20Interface, Ownable {
    using SafeMath for uint;

    string public symbol;
    string public  name;
    uint8 public decimals;
    uint public totalSupply;

    address tokenSaleContract;

    mapping(address => uint) balances;

    mapping(address => mapping(address => uint)) allowed;

    modifier validDestination( address to ) {
        require(to != address(0x0));
        require(to != address(this) );
        _;
    }

    modifier onlyWhenTransferEnabled() {
        require( msg.sender == tokenSaleContract );
        _;
    }

    event Earn(address indexed to, uint tokens);

constructor () public {
symbol = "CNC";
name = "Codedao Network Crystal Token";
decimals = 0;
totalSupply = 0;
balances[owner] = totalSupply;
tokenSaleContract = msg.sender;
}

function () payable onlyOwner public {
uint256 _issued = (msg.value*1000)/10**18;
totalSupply += _issued;
balances[msg.sender] += _issued;
}

// ------------------------------------------------------------------------

// Total supply

// ------------------------------------------------------------------------

function totalSupply() public constant returns (uint) {

return totalSupply;

}


// ------------------------------------------------------------------------

// Get the token balance for account `tokenOwner`

// ------------------------------------------------------------------------

function balanceOf(address tokenOwner) public constant returns (uint balance) {

return balances[tokenOwner];

}

// ------------------------------------------------------------------------

// Transfer the balance from token owner's account to `to` account

// - Owner's account must have sufficient balance to transfer

// - 0 value transfers are allowed

// ------------------------------------------------------------------------

function transfer(address to, uint tokens) public validDestination(to) onlyWhenTransferEnabled returns (bool success) {

balances[msg.sender] = balances[msg.sender].sub(tokens);

balances[to] = balances[to].add(tokens);

Transfer(msg.sender, to, tokens);

return true;

}

// ------------------------------------------------------------------------

// Token owner can approve for `spender` to transferFrom(...) `tokens`

// from the token owner's account

//

// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md

// recommends that there are no checks for the approval double-spend attack

// as this should be implemented in user interfaces

// ------------------------------------------------------------------------

function approve(address spender, uint tokens) public returns (bool success) {

allowed[msg.sender][spender] = tokens;

Approval(msg.sender, spender, tokens);

return true;

}

// ------------------------------------------------------------------------

// Transfer `tokens` from the `from` account to the `to` account

//

// The calling account must already have sufficient tokens approve(...)-d

// for spending from the `from` account and

// - From account must have sufficient balance to transfer

// - Spender must have sufficient allowance to transfer

// - 0 value transfers are allowed

// ------------------------------------------------------------------------

function transferFrom(address from, address to, uint tokens) public  validDestination(to) onlyWhenTransferEnabled returns (bool success) {

balances[from] = balances[from].sub(tokens);

allowed[from][msg.sender] = allowed[from][msg.sender].sub(tokens);

balances[to] = balances[to].add(tokens);

Transfer(from, to, tokens);

return true;

}

// Earn `tokens` when share

function earn(address to, uint tokens) public validDestination(to) onlyWhenTransferEnabled returns (bool success) {
assert(tokens > 0);

balances[to] = balances[to].add(tokens);

totalSupply = totalSupply.add(tokens);

// allow owner handler
allowed[to][tokenSaleContract] = allowed[to][tokenSaleContract].add(tokens);

Earn(to, tokens);

return true;
}

// ------------------------------------------------------------------------

// Returns the amount of tokens approved by the owner that can be

// transferred to the spender's account

// ------------------------------------------------------------------------

function allowance(address tokenOwner, address spender) public constant returns (uint remaining) {

return allowed[tokenOwner][spender];

}

}
