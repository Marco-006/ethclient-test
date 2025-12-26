// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract counter{
    uint64 public num;

    function addNum() public {
        num++;
    }

    function getNum() public  view returns (uint64) {
        return num;
    }

    //  solcjs --bin task2.sol
    //  solcjs --abi task2.sol
    // abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=store.go

    // 0xc8d62406c1c828ebeb8c2edba30c0b752cf94c805395e820d10be7a3ce2d0388
    // 0xc8d62406c1c828ebeb8c2edba30c0b752cf94c805395e820d10be7a3ce2d0388
}