 SPDX-License-Identifier MIT
pragma solidity ^0.8.0;

contract STVVote {
    struct Candidate {
        string name;
        uint voteCount;
    }

    struct Ballot {
        address voter;
        uint[] preferences;
        bool voted;
    }

    address public owner;
    Candidate[] public candidates;
    mapping(address = Ballot) public ballots;
    uint public totalVotes;
    uint public seats;

    event VoteCast(address voter, uint[] preferences);

    modifier onlyOwner() {
        require(msg.sender == owner, Not authorized.);
        _;
    }

    constructor(string[] memory candidateNames, uint _seats) {
        owner = msg.sender;
        seats = _seats;
        for (uint i = 0; i  candidateNames.length; i++) {
            candidates.push(Candidate({
                name candidateNames[i],
                voteCount 0
            }));
        }
    }

    function vote(uint[] memory preferences) public {
        require(!ballots[msg.sender].voted, Already voted.);
        require(preferences.length == candidates.length, Invalid ballot.);

        ballots[msg.sender] = Ballot({
            voter msg.sender,
            preferences preferences,
            voted true
        });

        totalVotes++;
        emit VoteCast(msg.sender, preferences);
    }

    function countVotes() public onlyOwner {
        uint quota = totalVotes  (seats + 1) + 1;
        uint[] memory remainingSeats = new uint[](candidates.length);

        for (uint i = 0; i  candidates.length; i++) {
            remainingSeats[i] = i;
        }

        while (seats  0) {
            for (uint i = 0; i  totalVotes; i++) {
                address voter = ballots[msg.sender].voter;
                uint firstPref = ballots[voter].preferences[0];
                candidates[firstPref].voteCount++;
            }

            for (uint i = 0; i  candidates.length; i++) {
                if (candidates[i].voteCount = quota) {
                    seats--;
                    candidates[i].voteCount = 0;
                }
            }
        }
    }

    function getCandidates() public view returns (Candidate[] memory) {
        return candidates;
    }
}