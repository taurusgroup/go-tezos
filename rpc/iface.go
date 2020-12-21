package rpc

import "encoding/json"

// IFace is an interface mocking a GoTezos object.
type IFace interface {
	ActiveChains() (ActiveChains, error)
	BakingRights(input BakingRightsInput) (*BakingRights, error)
	ContractBalance(input ContractBalanceInput) (string, error)
	BallotList(blockhash string) (BallotList, error)
	Ballots(blockhash string) (Ballots, error)
	BigMap(input BigMapInput) ([]byte, error)
	Block(id interface{}) (*Block, error)
	Blocks(input BlocksInput) ([][]string, error)
	Bootstrap() (Bootstrap, error)
	ChainID() (string, error)
	Checkpoint() (Checkpoint, error)
	Commit() (string, error)
	Connections() (Connections, error)
	Constants(ConstantsInput) (Constants, error)
	ContractStorage(input ContractStorageInput) (*json.RawMessage, error)
	ContractCounter(input ContractCounterInput) (int, error)
	CurrentPeriodKind(blockhash string) (string, error)
	CurrentProposal(blockhash string) (string, error)
	CurrentQuorum(blockhash string) (int, error)
	Cycle(cycle int) (Cycle, error)
	// Delegate(input DelegateInput) (Delegate, error)
	Delegates(input DelegatesInput) ([]string, error)
	// DelegatedContracts(input DelegateDelegatedContractsInput) ([]string, error)
	DeleteInvalidBlock(blockHash string) error
	EndorsingRights(input EndorsingRightsInput) (*EndorsingRights, error)
	ForgeOperation(input ForgeOperationInput) (string, error)
	GetFA12Allowance(input GetFA12AllowanceInput) (string, error)
	GetFA12Balance(input GetFA12BalanceInput) (string, error)
	GetFA12Supply(input GetFA12SupplyInput) (string, error)
	Head() (*Block, error)
	InjectionBlock(input InjectionBlockInput) ([]byte, error)
	InjectionOperation(input InjectionOperationInput) (string, error)
	InvalidBlock(blockHash string) (InvalidBlock, error)
	InvalidBlocks() ([]InvalidBlock, error)
	OperationHashes(blockhash string) ([][]string, error)
	PreapplyOperations(input PreapplyOperationsInput) ([]Operations, error)
	Proposals(blockhash string) (Proposals, error)
	RunOperation(input RunOperationInput) (Operations, error)
	StakingBalance(input StakingBalanceInput) (int, error)
	UnforgeOperation(input UnforgeOperationInput) ([]Operations, error)
	UserActivatedProtocolOverrides() (UserActivatedProtocolOverrides, error)
	Version() (Version, error)
}
