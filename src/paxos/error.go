package paxos

import "errors"

var (
	KeyNotExists = errors.New("Key not exist")
	RndInvalid = errors.New("Rnd is invalid")
	PaxosNodeLessThanQuorum = errors.New("Paxos nodes size less than quorum")
	CurrentRndNumInvalid = errors.New("Current round number invalid")
	Phrase1Failed = errors.New("Phrase 1 failed")
)
