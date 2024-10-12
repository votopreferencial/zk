package main

import (
	encodingjson
	fmt
	strconv

	github.comhyperledgerfabric-contract-api-gocontractapi
)

type SmartContract struct {
	contractapi.Contract
}

type Candidate struct {
	Name      string `jsonname`
	VoteCount int    `jsonvoteCount`
}

type Ballot struct {
	Voter      string   `jsonvoter`
	Preferences []int    `jsonpreferences`
	Voted      bool     `jsonvoted`
}

type Election struct {
	Candidates []Candidate `jsoncandidates`
	Seats      int         `jsonseats`
	TotalVotes int         `jsontotalVotes`
	Ballots    []Ballot    `jsonballots`
}

func (s SmartContract) InitElection(ctx contractapi.TransactionContextInterface, candidateNames []string, seats int) error {
	var candidates []Candidate
	for _, name = range candidateNames {
		candidates = append(candidates, Candidate{Name name, VoteCount 0})
	}
	election = Election{Candidates candidates, Seats seats, TotalVotes 0, Ballots []Ballot{}}

	electionJSON, err = json.Marshal(election)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(election, electionJSON)
}

func (s SmartContract) CastVote(ctx contractapi.TransactionContextInterface, voter string, preferences []int) error {
	electionJSON, err = ctx.GetStub().GetState(election)
	if err != nil {
		return err
	}
	if electionJSON == nil {
		return fmt.Errorf(election does not exist)
	}

	var election Election
	err = json.Unmarshal(electionJSON, &election)
	if err != nil {
		return err
	}

	for _, b = range election.Ballots {
		if b.Voter == voter {
			return fmt.Errorf(voter has already voted)
		}
	}

	ballot = Ballot{
		Voter      voter,
		Preferences preferences,
		Voted      true,
	}
	election.Ballots = append(election.Ballots, ballot)
	election.TotalVotes++

	electionJSON, err = json.Marshal(election)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(election, electionJSON)
}

func (s SmartContract) CountVotes(ctx contractapi.TransactionContextInterface) error {
	electionJSON, err = ctx.GetStub().GetState(election)
	if err != nil {
		return err
	}
	if electionJSON == nil {
		return fmt.Errorf(election does not exist)
	}

	var election Election
	err = json.Unmarshal(electionJSON, &election)
	if err != nil {
		return err
	}

	quota = election.TotalVotes  (election.Seats + 1)

	for seatsRemaining = election.Seats; seatsRemaining  0; {
		for _, ballot = range election.Ballots {
			firstPref = ballot.Preferences[0]
			election.Candidates[firstPref].VoteCount++
		}

		for i, candidate = range election.Candidates {
			if candidate.VoteCount = quota {
				seatsRemaining--
				candidate.VoteCount = 0
				election.Candidates[i] = candidate
			}
		}
	}

	electionJSON, err = json.Marshal(election)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(election, electionJSON)
}

func (s SmartContract) GetElectionResults(ctx contractapi.TransactionContextInterface) ([]Candidate, error) {
	electionJSON, err = ctx.GetStub().GetState(election)
	if err != nil {
		return nil, err
	}
	if electionJSON == nil {
		return nil, fmt.Errorf(election does not exist)
	}

	var election Election
	err = json.Unmarshal(electionJSON, &election)
	if err != nil {
		return nil, err
	}

	return election.Candidates, nil
}

func main() {
	chaincode, err = contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf(Error creating chaincode %s, err)
	}

	if err = chaincode.Start(); err != nil {
		fmt.Printf(Error starting chaincode %s, err)
	}
}