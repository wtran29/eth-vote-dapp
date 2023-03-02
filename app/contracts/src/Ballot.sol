// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;


contract Ballot {

    string public candidate;
    // Constructor
    constructor () public {
        candidate = "Candidate 1";
    }
}