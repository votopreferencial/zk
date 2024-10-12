// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract STVContract {

    struct Candidate {
        string name;
        uint voteCount;
    }

    struct Voter {
        bool voted;
        uint[] preferences;
    }

    address public owner;
    mapping(address => Voter) public voters;
    Candidate[] public candidates;
    uint public quorum;
    uint public totalVotes;

    constructor(string[] memory candidateNames, uint _quorum) {
        owner = msg.sender;
        quorum = _quorum;
        for (uint i = 0; i < candidateNames.length; i++) {
            candidates.push(Candidate({
                name: candidateNames[i],
                voteCount: 0
            }));
        }
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    modifier hasNotVoted() {
        require(!voters[msg.sender].voted, "You have already voted");
        _;
    }

    function vote(uint[] memory preferences) public hasNotVoted {
        require(preferences.length == candidates.length, "Invalid vote format");
        
        voters[msg.sender].voted = true;
        voters[msg.sender].preferences = preferences;
        totalVotes++;
    }

function calculateSTVTally() public view onlyOwner returns (string memory) {
    require(totalVotes >= quorum, "Quorum não atingido");

    uint[] memory tally = new uint[](candidates.length);
    uint winnerIndex = 0;

    for (uint i = 0; i < candidates.length; i++) {
        tally[i] = 0;
    }

    // Correção: Iterar sobre os endereços dos eleitores, não sobre totalVotes
    address[] memory voterAddresses = new address[](totalVotes);
    uint voterCount = 0;
    for (uint i = 0; i < voters.length; i++) {
        if (voters[i].voted) {
            voterAddresses[voterCount] = address(uint160(i));
            voterCount++;
        }
    }

    for (uint i = 0; i < voterCount; i++) {
        Voter storage voter = voters[voterAddresses[i]];
        uint firstPreference = voter.preferences[0];
        tally[firstPreference]++;
    }

    for (uint i = 0; i < tally.length; i++) {
        if (tally[i] > tally[winnerIndex]) {
            winnerIndex = i;
        }
    }

    return candidates[winnerIndex].name;
}