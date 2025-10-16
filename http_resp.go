package etherscan

// Account Module Response Types

type RespGetEthBalance string

// RespBridgeTx represents a bridge transaction
// Example:
//
//	{
//	    "hash": "0x79eaf9951f474d5fd78bae5e3d6e089b54e61b81d272ce6f98e8ac6b56ec0f93",
//	    "blockNumber": "42107370",
//	    "timeStamp": "1757792370",
//	    "from": "0xfa98b60e02a61b6590f073cad56e68326652d094",
//	    "address": "0x1545c4ccf40a5e89ac1482a32485f62d369560d3",
//	    "amount": "1000000000000000000",
//	    "tokenName": "",
//	    "symbol": "",
//	    "contractAddress": "0x7301cfa0e1756b71869e93d4e4dca5c7d0eb0aa6",
//	    "divisor": ""
//	}
type RespBridgeTx struct {
	Hash            string `json:"hash" bson:"hash"`
	BlockNumber     string `json:"blockNumber" bson:"blockNumber"`
	TimeStamp       string `json:"timeStamp" bson:"timeStamp"`
	From            string `json:"from" bson:"from"`
	Address         string `json:"address" bson:"address"`
	Amount          string `json:"amount" bson:"amount"`
	TokenName       string `json:"tokenName" bson:"tokenName"`
	Symbol          string `json:"symbol" bson:"symbol"`
	ContractAddress string `json:"contractAddress" bson:"contractAddress"`
	Divisor         string `json:"divisor" bson:"divisor"`
}

type RespGetBridgeTxs []RespBridgeTx

// RespEthBalanceEntry represents an Ethereum balance entry
// Example:
//
//	{
//	    "account": "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a",
//	    "balance": "40891626854930000000000"
//	}
type RespEthBalanceEntry struct {
	Account string `json:"account" bson:"account"`
	Balance string `json:"balance" bson:"balance"`
}

type RespGetEthBalances []RespEthBalanceEntry

// RespNormalTx represents a normal transaction
// Example:
//
//	{
//	    "blockNumber": "14923678",
//	    "timeStamp": "1654646411",
//	    "hash": "0xc52783ad354aecc04c670047754f062e3d6d04e8f5b24774472651f9c3882c60",
//	    ...
//	}
type RespNormalTx struct {
	BlockNumber       string `json:"blockNumber" bson:"blockNumber"`
	TimeStamp         string `json:"timeStamp" bson:"timeStamp"`
	Hash              string `json:"hash" bson:"hash"`
	Nonce             string `json:"nonce" bson:"nonce"`
	BlockHash         string `json:"blockHash" bson:"blockHash"`
	TransactionIndex  string `json:"transactionIndex" bson:"transactionIndex"`
	From              string `json:"from" bson:"from"`
	To                string `json:"to" bson:"to"`
	Value             string `json:"value" bson:"value"`
	Gas               string `json:"gas" bson:"gas"`
	GasPrice          string `json:"gasPrice" bson:"gasPrice"`
	IsError           string `json:"isError" bson:"isError"`
	TxReceiptStatus   string `json:"txreceipt_status" bson:"txreceipt_status"`
	Input             string `json:"input" bson:"input"`
	ContractAddress   string `json:"contractAddress" bson:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed" bson:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed" bson:"gasUsed"`
	Confirmations     string `json:"confirmations" bson:"confirmations"`
	MethodID          string `json:"methodId" bson:"methodId"`
	FunctionName      string `json:"functionName" bson:"functionName"`
}

type RespGetNormalTxs []RespNormalTx

// RespInternalTxByAddress represents an internal transaction by address
type RespInternalTxByAddress struct {
	BlockNumber     string `json:"blockNumber" bson:"blockNumber"`
	TimeStamp       string `json:"timeStamp" bson:"timeStamp"`
	Hash            string `json:"hash" bson:"hash"`
	From            string `json:"from" bson:"from"`
	To              string `json:"to" bson:"to"`
	Value           string `json:"value" bson:"value"`
	ContractAddress string `json:"contractAddress" bson:"contractAddress"`
	Input           string `json:"input" bson:"input"`
	Type            string `json:"type" bson:"type"`
	Gas             string `json:"gas" bson:"gas"`
	GasUsed         string `json:"gasUsed" bson:"gasUsed"`
	TraceID         string `json:"traceId" bson:"traceId"`
	IsError         string `json:"isError" bson:"isError"`
	ErrCode         string `json:"errCode" bson:"errCode"`
}

type RespGetInternalTxsByAddress []RespInternalTxByAddress

// RespInternalTxByHash represents an internal transaction by hash
type RespInternalTxByHash struct {
	BlockNumber     string `json:"blockNumber" bson:"blockNumber"`
	TimeStamp       string `json:"timeStamp" bson:"timeStamp"`
	From            string `json:"from" bson:"from"`
	To              string `json:"to" bson:"to"`
	Value           string `json:"value" bson:"value"`
	ContractAddress string `json:"contractAddress" bson:"contractAddress"`
	Input           string `json:"input" bson:"input"`
	Type            string `json:"type" bson:"type"`
	Gas             string `json:"gas" bson:"gas"`
	GasUsed         string `json:"gasUsed" bson:"gasUsed"`
	IsError         string `json:"isError" bson:"isError"`
	ErrCode         string `json:"errCode" bson:"errCode"`
}

type RespGetInternalTxsByHash []RespInternalTxByHash

// RespInternalTxByBlockRange represents an internal transaction by block range
type RespInternalTxByBlockRange struct {
	BlockNumber     string `json:"blockNumber" bson:"blockNumber"`
	TimeStamp       string `json:"timeStamp" bson:"timeStamp"`
	Hash            string `json:"hash" bson:"hash"`
	From            string `json:"from" bson:"from"`
	To              string `json:"to" bson:"to"`
	Value           string `json:"value" bson:"value"`
	ContractAddress string `json:"contractAddress" bson:"contractAddress"`
	Input           string `json:"input" bson:"input"`
	Type            string `json:"type" bson:"type"`
	Gas             string `json:"gas" bson:"gas"`
	GasUsed         string `json:"gasUsed" bson:"gasUsed"`
	TraceID         string `json:"traceId" bson:"traceId"`
	IsError         string `json:"isError" bson:"isError"`
	ErrCode         string `json:"errCode" bson:"errCode"`
}

type RespGetInternalTxsByBlockRange []RespInternalTxByBlockRange

// RespERC20TokenTransfer represents an ERC-20 token transfer
type RespERC20TokenTransfer struct {
	BlockNumber       string `json:"blockNumber" bson:"blockNumber"`
	TimeStamp         string `json:"timeStamp" bson:"timeStamp"`
	Hash              string `json:"hash" bson:"hash"`
	Nonce             string `json:"nonce" bson:"nonce"`
	BlockHash         string `json:"blockHash" bson:"blockHash"`
	From              string `json:"from" bson:"from"`
	ContractAddress   string `json:"contractAddress" bson:"contractAddress"`
	To                string `json:"to" bson:"to"`
	Value             string `json:"value" bson:"value"`
	TokenName         string `json:"tokenName" bson:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol" bson:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal" bson:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex" bson:"transactionIndex"`
	Gas               string `json:"gas" bson:"gas"`
	GasPrice          string `json:"gasPrice" bson:"gasPrice"`
	GasUsed           string `json:"gasUsed" bson:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed" bson:"cumulativeGasUsed"`
	Input             string `json:"input" bson:"input"`
	Confirmations     string `json:"confirmations" bson:"confirmations"`
}

type RespGetERC20TokenTransfers []RespERC20TokenTransfer

// RespERC721TokenTransfer represents an ERC-721 token transfer
type RespERC721TokenTransfer struct {
	BlockNumber       string `json:"blockNumber" bson:"blockNumber"`
	TimeStamp         string `json:"timeStamp" bson:"timeStamp"`
	Hash              string `json:"hash" bson:"hash"`
	Nonce             string `json:"nonce" bson:"nonce"`
	BlockHash         string `json:"blockHash" bson:"blockHash"`
	From              string `json:"from" bson:"from"`
	ContractAddress   string `json:"contractAddress" bson:"contractAddress"`
	To                string `json:"to" bson:"to"`
	TokenID           string `json:"tokenID" bson:"tokenID"`
	TokenName         string `json:"tokenName" bson:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol" bson:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal" bson:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex" bson:"transactionIndex"`
	Gas               string `json:"gas" bson:"gas"`
	GasPrice          string `json:"gasPrice" bson:"gasPrice"`
	GasUsed           string `json:"gasUsed" bson:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed" bson:"cumulativeGasUsed"`
	Input             string `json:"input" bson:"input"`
	Confirmations     string `json:"confirmations" bson:"confirmations"`
}

type RespGetERC721TokenTransfers []RespERC721TokenTransfer

// RespERC1155TokenTransfer represents an ERC-1155 token transfer
type RespERC1155TokenTransfer struct {
	BlockNumber       string `json:"blockNumber" bson:"blockNumber"`
	TimeStamp         string `json:"timeStamp" bson:"timeStamp"`
	Hash              string `json:"hash" bson:"hash"`
	Nonce             string `json:"nonce" bson:"nonce"`
	BlockHash         string `json:"blockHash" bson:"blockHash"`
	TransactionIndex  string `json:"transactionIndex" bson:"transactionIndex"`
	Gas               string `json:"gas" bson:"gas"`
	GasPrice          string `json:"gasPrice" bson:"gasPrice"`
	GasUsed           string `json:"gasUsed" bson:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed" bson:"cumulativeGasUsed"`
	Input             string `json:"input" bson:"input"`
	ContractAddress   string `json:"contractAddress" bson:"contractAddress"`
	From              string `json:"from" bson:"from"`
	To                string `json:"to" bson:"to"`
	TokenID           string `json:"tokenID" bson:"tokenID"`
	TokenValue        string `json:"tokenValue" bson:"tokenValue"`
	TokenName         string `json:"tokenName" bson:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol" bson:"tokenSymbol"`
	Confirmations     string `json:"confirmations" bson:"confirmations"`
}

type RespGetERC1155TokenTransfers []RespERC1155TokenTransfer

// RespAddressFundedBy represents address funding information
type RespAddressFundedBy struct {
	Block          int64  `json:"block" bson:"block"`
	TimeStamp      string `json:"timeStamp" bson:"timeStamp"`
	FundingAddress string `json:"fundingAddress" bson:"fundingAddress"`
	FundingTxn     string `json:"fundingTxn" bson:"fundingTxn"`
	Value          string `json:"value" bson:"value"`
}

// RespBlockValidated represents a validated block
type RespBlockValidated struct {
	BlockNumber string `json:"blockNumber" bson:"blockNumber"`
	TimeStamp   string `json:"timeStamp" bson:"timeStamp"`
	BlockReward string `json:"blockReward" bson:"blockReward"`
}

type RespBlocksValidatedByAddress []RespBlockValidated

// RespBeaconChainWithdrawal represents a beacon chain withdrawal
type RespBeaconChainWithdrawal struct {
	WithdrawalIndex string `json:"withdrawalIndex" bson:"withdrawalIndex"`
	ValidatorIndex  string `json:"validatorIndex" bson:"validatorIndex"`
	Address         string `json:"address" bson:"address"`
	Amount          string `json:"amount" bson:"amount"`
	BlockNumber     string `json:"blockNumber" bson:"blockNumber"`
	Timestamp       string `json:"timestamp" bson:"timestamp"`
}

type RespBeaconChainWithdrawals []RespBeaconChainWithdrawal

type RespEthBalanceByBlockNumber string

// Contract Module Response Types

type RespContractABI string

// RespContractSourceCode represents contract source code information
type RespContractSourceCode struct {
	SourceCode           string `json:"SourceCode" bson:"SourceCode"`
	ABI                  string `json:"ABI" bson:"ABI"`
	ContractName         string `json:"ContractName" bson:"ContractName"`
	CompilerVersion      string `json:"CompilerVersion" bson:"CompilerVersion"`
	OptimizationUsed     string `json:"OptimizationUsed" bson:"OptimizationUsed"`
	Runs                 string `json:"Runs" bson:"Runs"`
	ConstructorArguments string `json:"ConstructorArguments" bson:"ConstructorArguments"`
	EVMVersion           string `json:"EVMVersion" bson:"EVMVersion"`
	Library              string `json:"Library" bson:"Library"`
	LicenseType          string `json:"LicenseType" bson:"LicenseType"`
	Proxy                string `json:"Proxy" bson:"Proxy"`
	Implementation       string `json:"Implementation" bson:"Implementation"`
	SwarmSource          string `json:"SwarmSource" bson:"SwarmSource"`
	SimilarMatch         string `json:"SimilarMatch" bson:"SimilarMatch"`
}

type RespContractSourceCodes []RespContractSourceCode

// RespContractCreationAndCreation represents contract creation information
type RespContractCreationAndCreation struct {
	ContractAddress  string `json:"contractAddress" bson:"contractAddress"`
	ContractCreator  string `json:"contractCreator" bson:"contractCreator"`
	TxHash           string `json:"txHash" bson:"txHash"`
	BlockNumber      string `json:"blockNumber" bson:"blockNumber"`
	Timestamp        string `json:"timestamp" bson:"timestamp"`
	ContractFactory  string `json:"contractFactory" bson:"contractFactory"`
	CreationBytecode string `json:"creationBytecode" bson:"creationBytecode"`
}

type RespContractCreationAndCreations []RespContractCreationAndCreation

type RespVerifySourceCode string

type RespVerifyVyperSourceCode string

type RespVerifyStylusSourceCode string

type RespCheckSourceCodeVerificationStatus string

// RespContractExecutionStatus represents contract execution status
type RespContractExecutionStatus struct {
	IsError        string `json:"isError" bson:"isError"`
	ErrDescription string `json:"errDescription" bson:"errDescription"`
}

// RespCheckTxReceiptStatus represents transaction receipt status
type RespCheckTxReceiptStatus struct {
	Status string `json:"status" bson:"status"`
}

// Block Module Response Types

// UncleReward represents an uncle block reward
type UncleReward struct {
	Miner         string `json:"miner" bson:"miner"`
	UnclePosition string `json:"unclePosition" bson:"unclePosition"`
	Blockreward   string `json:"blockreward" bson:"blockreward"`
}

// RespBlockReward represents block reward information
type RespBlockReward struct {
	BlockNumber          string        `json:"blockNumber" bson:"blockNumber"`
	TimeStamp            string        `json:"timeStamp" bson:"timeStamp"`
	BlockMiner           string        `json:"blockMiner" bson:"blockMiner"`
	BlockReward          string        `json:"blockReward" bson:"blockReward"`
	Uncles               []UncleReward `json:"uncles" bson:"uncles"`
	UncleInclusionReward string        `json:"uncleInclusionReward" bson:"uncleInclusionReward"`
}

// RespBlockTxsCountByBlockNo represents transaction count by block number
type RespBlockTxsCountByBlockNo struct {
	Block            int64 `json:"block" bson:"block"`
	TxsCount         int64 `json:"txsCount" bson:"txsCount"`
	InternalTxsCount int64 `json:"internalTxsCount" bson:"internalTxsCount"`
	ERC20TxsCount    int64 `json:"erc20TxsCount" bson:"erc20TxsCount"`
	ERC721TxsCount   int64 `json:"erc721TxsCount" bson:"erc721TxsCount"`
	ERC1155TxsCount  int64 `json:"erc1155TxsCount" bson:"erc1155TxsCount"`
}

// RespEstimateBlockCountdownTimeByBlockNo represents estimated block countdown time
type RespEstimateBlockCountdownTimeByBlockNo struct {
	CurrentBlock      string `json:"CurrentBlock" bson:"CurrentBlock"`
	CountdownBlock    string `json:"CountdownBlock" bson:"CountdownBlock"`
	RemainingBlock    string `json:"RemainingBlock" bson:"RemainingBlock"`
	EstimateTimeInSec string `json:"EstimateTimeInSec" bson:"EstimateTimeInSec"`
}

type RespBlockNumber int

// RespDailyAvgBlockSize represents daily average block size
type RespDailyAvgBlockSize struct {
	UTCDate        string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp  string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	BlockSizeBytes int64  `json:"blockSize_bytes" bson:"blockSize_bytes"`
}

type RespDailyAvgBlockSizes []RespDailyAvgBlockSize

// RespDailyBlockCountReward represents daily block count and reward
type RespDailyBlockCountReward struct {
	UTCDate         string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp   string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	BlockCount      int64  `json:"blockCount" bson:"blockCount"`
	BlockRewardsEth string `json:"blockRewards_Eth" bson:"blockRewards_Eth"`
}

type RespDailyBlockCountRewards []RespDailyBlockCountReward

// RespDailyBlockReward represents daily block reward
type RespDailyBlockReward struct {
	UTCDate         string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp   string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	BlockRewardsEth string `json:"blockRewards_Eth" bson:"blockRewards_Eth"`
}

type RespDailyBlockRewards []RespDailyBlockReward

// RespDailyAvgTimeBlockMined represents daily average time for block mining
type RespDailyAvgTimeBlockMined struct {
	UTCDate       string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	BlockTimeSec  string `json:"blockTime_sec" bson:"blockTime_sec"`
}

type RespDailyAvgTimeBlockMineds []RespDailyAvgTimeBlockMined

// RespDailyUncleBlockCountAndReward represents daily uncle block count and reward
type RespDailyUncleBlockCountAndReward struct {
	UTCDate              string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp        string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	UncleBlockCount      int64  `json:"uncleBlockCount" bson:"uncleBlockCount"`
	UncleBlockRewardsEth string `json:"uncleBlockRewards_Eth" bson:"uncleBlockRewards_Eth"`
}

type RespDailyUncleBlockCountAndRewards []RespDailyUncleBlockCountAndReward

// Logs Module Response Types

// RespEventLogByAddress represents an event log by address
type RespEventLogByAddress struct {
	Address          string   `json:"address" bson:"address"`
	Topics           []string `json:"topics" bson:"topics"`
	Data             string   `json:"data" bson:"data"`
	BlockNumber      string   `json:"blockNumber" bson:"blockNumber"`
	TimeStamp        string   `json:"timeStamp" bson:"timeStamp"`
	GasPrice         string   `json:"gasPrice" bson:"gasPrice"`
	GasUsed          string   `json:"gasUsed" bson:"gasUsed"`
	LogIndex         string   `json:"logIndex" bson:"logIndex"`
	TransactionHash  string   `json:"transactionHash" bson:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex" bson:"transactionIndex"`
}

type RespEventLogsByAddress []RespEventLogByAddress

// RespEventLogByTopics represents an event log by topics
type RespEventLogByTopics struct {
	Address          string   `json:"address" bson:"address"`
	Topics           []string `json:"topics" bson:"topics"`
	Data             string   `json:"data" bson:"data"`
	BlockNumber      string   `json:"blockNumber" bson:"blockNumber"`
	TimeStamp        string   `json:"timeStamp" bson:"timeStamp"`
	GasPrice         string   `json:"gasPrice" bson:"gasPrice"`
	GasUsed          string   `json:"gasUsed" bson:"gasUsed"`
	LogIndex         string   `json:"logIndex" bson:"logIndex"`
	TransactionHash  string   `json:"transactionHash" bson:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex" bson:"transactionIndex"`
}

type RespEventLogsByTopics []RespEventLogByTopics

// RespEventLogByAddressFilteredByTopics represents an event log by address filtered by topics
type RespEventLogByAddressFilteredByTopics struct {
	Address          string   `json:"address" bson:"address"`
	Topics           []string `json:"topics" bson:"topics"`
	Data             string   `json:"data" bson:"data"`
	BlockNumber      string   `json:"blockNumber" bson:"blockNumber"`
	TimeStamp        string   `json:"timeStamp" bson:"timeStamp"`
	GasPrice         string   `json:"gasPrice" bson:"gasPrice"`
	GasUsed          string   `json:"gasUsed" bson:"gasUsed"`
	LogIndex         string   `json:"logIndex" bson:"logIndex"`
	TransactionHash  string   `json:"transactionHash" bson:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex" bson:"transactionIndex"`
}

type RespEventLogsByAddressFilteredByTopics []RespEventLogByAddressFilteredByTopics

// Geth/Parity Proxy Module Response Types

// RespJsonRpc represents a generic JSON-RPC response
type RespJsonRpc[Result any] struct {
	Jsonrpc string `json:"jsonrpc" bson:"jsonrpc"`
	ID      int64  `json:"id" bson:"id"`
	Result  Result `json:"result" bson:"result"`
}

type RespEthBlockNumberHex = RespJsonRpc[string]

// RespEthBlockInfo represents Ethereum block information
type RespEthBlockInfo struct {
	BaseFeePerGas    string   `json:"baseFeePerGas" bson:"baseFeePerGas"`
	Difficulty       string   `json:"difficulty" bson:"difficulty"`
	ExtraData        string   `json:"extraData" bson:"extraData"`
	GasLimit         string   `json:"gasLimit" bson:"gasLimit"`
	GasUsed          string   `json:"gasUsed" bson:"gasUsed"`
	Hash             string   `json:"hash" bson:"hash"`
	LogsBloom        string   `json:"logsBloom" bson:"logsBloom"`
	Miner            string   `json:"miner" bson:"miner"`
	MixHash          string   `json:"mixHash" bson:"mixHash"`
	Nonce            string   `json:"nonce" bson:"nonce"`
	Number           string   `json:"number" bson:"number"`
	ParentHash       string   `json:"parentHash" bson:"parentHash"`
	ReceiptsRoot     string   `json:"receiptsRoot" bson:"receiptsRoot"`
	Sha3Uncles       string   `json:"sha3Uncles" bson:"sha3Uncles"`
	Size             string   `json:"size" bson:"size"`
	StateRoot        string   `json:"stateRoot" bson:"stateRoot"`
	Timestamp        string   `json:"timestamp" bson:"timestamp"`
	TotalDifficulty  string   `json:"totalDifficulty" bson:"totalDifficulty"`
	Transactions     []string `json:"transactions" bson:"transactions"`
	TransactionsRoot string   `json:"transactionsRoot" bson:"transactionsRoot"`
	Uncles           []string `json:"uncles" bson:"uncles"`
}

type RespEthBlock = RespJsonRpc[RespEthBlockInfo]

type RespEthBlockInfoWithFullTxs struct {
	BaseFeePerGas    string          `json:"baseFeePerGas" bson:"baseFeePerGas"`
	Difficulty       string          `json:"difficulty" bson:"difficulty"`
	ExtraData        string          `json:"extraData" bson:"extraData"`
	GasLimit         string          `json:"gasLimit" bson:"gasLimit"`
	GasUsed          string          `json:"gasUsed" bson:"gasUsed"`
	Hash             string          `json:"hash" bson:"hash"`
	LogsBloom        string          `json:"logsBloom" bson:"logsBloom"`
	Miner            string          `json:"miner" bson:"miner"`
	MixHash          string          `json:"mixHash" bson:"mixHash"`
	Nonce            string          `json:"nonce" bson:"nonce"`
	Number           string          `json:"number" bson:"number"`
	ParentHash       string          `json:"parentHash" bson:"parentHash"`
	ReceiptsRoot     string          `json:"receiptsRoot" bson:"receiptsRoot"`
	Sha3Uncles       string          `json:"sha3Uncles" bson:"sha3Uncles"`
	Size             string          `json:"size" bson:"size"`
	StateRoot        string          `json:"stateRoot" bson:"stateRoot"`
	Timestamp        string          `json:"timestamp" bson:"timestamp"`
	TotalDifficulty  string          `json:"totalDifficulty" bson:"totalDifficulty"`
	Transactions     []RespEthTxInfo `json:"transactions" bson:"transactions"`
	TransactionsRoot string          `json:"transactionsRoot" bson:"transactionsRoot"`
	Uncles           []string        `json:"uncles" bson:"uncles"`
}

type RespEthBlockWithFullTxs = RespJsonRpc[RespEthBlockInfoWithFullTxs]

// RespEthUncleBlockInfo represents Ethereum uncle block information
type RespEthUncleBlockInfo struct {
	BaseFeePerGas    string   `json:"baseFeePerGas" bson:"baseFeePerGas"`
	Difficulty       string   `json:"difficulty" bson:"difficulty"`
	ExtraData        string   `json:"extraData" bson:"extraData"`
	GasLimit         string   `json:"gasLimit" bson:"gasLimit"`
	GasUsed          string   `json:"gasUsed" bson:"gasUsed"`
	Hash             string   `json:"hash" bson:"hash"`
	LogsBloom        string   `json:"logsBloom" bson:"logsBloom"`
	Miner            string   `json:"miner" bson:"miner"`
	MixHash          string   `json:"mixHash" bson:"mixHash"`
	Nonce            string   `json:"nonce" bson:"nonce"`
	Number           string   `json:"number" bson:"number"`
	ParentHash       string   `json:"parentHash" bson:"parentHash"`
	ReceiptsRoot     string   `json:"receiptsRoot" bson:"receiptsRoot"`
	Sha3Uncles       string   `json:"sha3Uncles" bson:"sha3Uncles"`
	Size             string   `json:"size" bson:"size"`
	StateRoot        string   `json:"stateRoot" bson:"stateRoot"`
	Timestamp        string   `json:"timestamp" bson:"timestamp"`
	TransactionsRoot string   `json:"transactionsRoot" bson:"transactionsRoot"`
	Uncles           []string `json:"uncles" bson:"uncles"`
}

type RespEthUncleBlock = RespJsonRpc[RespEthUncleBlockInfo]

type RespEthBlockTxCount = RespJsonRpc[string]

// RespEthTxInfo represents Ethereum transaction information
type RespEthTxInfo struct {
	BlockHash            string `json:"blockHash" bson:"blockHash"`
	BlockNumber          string `json:"blockNumber" bson:"blockNumber"`
	From                 string `json:"from" bson:"from"`
	Gas                  string `json:"gas" bson:"gas"`
	GasPrice             string `json:"gasPrice" bson:"gasPrice"`
	MaxFeePerGas         string `json:"maxFeePerGas" bson:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas" bson:"maxPriorityFeePerGas"`
	Hash                 string `json:"hash" bson:"hash"`
	Input                string `json:"input" bson:"input"`
	Nonce                string `json:"nonce" bson:"nonce"`
	To                   string `json:"to" bson:"to"`
	TransactionIndex     string `json:"transactionIndex" bson:"transactionIndex"`
	Value                string `json:"value" bson:"value"`
	Type                 string `json:"type" bson:"type"`
	AccessList           []any  `json:"accessList" bson:"accessList"`
	ChainID              string `json:"chainId" bson:"chainId"`
	V                    string `json:"v" bson:"v"`
	R                    string `json:"r" bson:"r"`
	S                    string `json:"s" bson:"s"`
}

type RespEthTx = RespJsonRpc[RespEthTxInfo]

type RespEthTxCount = RespJsonRpc[string]

type RespEthSendRawTx = RespJsonRpc[string]

// RespEthTxReceiptLog represents a transaction receipt log
type RespEthTxReceiptLog struct {
	Address          string   `json:"address" bson:"address"`
	Topics           []string `json:"topics" bson:"topics"`
	Data             string   `json:"data" bson:"data"`
	BlockNumber      string   `json:"blockNumber" bson:"blockNumber"`
	TransactionHash  string   `json:"transactionHash" bson:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex" bson:"transactionIndex"`
	BlockHash        string   `json:"blockHash" bson:"blockHash"`
	LogIndex         string   `json:"logIndex" bson:"logIndex"`
	Removed          bool     `json:"removed" bson:"removed"`
}

// RespEthTxReceiptInfo represents Ethereum transaction receipt information
type RespEthTxReceiptInfo struct {
	BlockHash         string                `json:"blockHash" bson:"blockHash"`
	BlockNumber       string                `json:"blockNumber" bson:"blockNumber"`
	ContractAddress   *string               `json:"contractAddress" bson:"contractAddress"`
	CumulativeGasUsed string                `json:"cumulativeGasUsed" bson:"cumulativeGasUsed"`
	EffectiveGasPrice string                `json:"effectiveGasPrice" bson:"effectiveGasPrice"`
	From              string                `json:"from" bson:"from"`
	GasUsed           string                `json:"gasUsed" bson:"gasUsed"`
	Logs              []RespEthTxReceiptLog `json:"logs" bson:"logs"`
	LogsBloom         string                `json:"logsBloom" bson:"logsBloom"`
	Status            string                `json:"status" bson:"status"`
	To                string                `json:"to" bson:"to"`
	TransactionHash   string                `json:"transactionHash" bson:"transactionHash"`
	TransactionIndex  string                `json:"transactionIndex" bson:"transactionIndex"`
	Type              string                `json:"type" bson:"type"`
}

type RespEthTxReceipt = RespJsonRpc[RespEthTxReceiptInfo]

type RespEthCall = RespJsonRpc[string]

type RespEthGetCode = RespJsonRpc[string]

type RespEthGetStorageAt = RespJsonRpc[string]

type RespEthGetGasPrice = RespJsonRpc[string]

type RespEthEstimateGas = RespJsonRpc[string]

// Token Module Response Types

type RespERC20TotalSupply string

type RespERC20AccountBalance string

type RespERC20HistoricalTotalSupply string

type RespERC20HistoricalAccountBalance string

// RespERC20HolderInfo represents ERC-20 token holder information
type RespERC20HolderInfo struct {
	TokenHolderAddress  string `json:"TokenHolderAddress" bson:"TokenHolderAddress"`
	TokenHolderQuantity string `json:"TokenHolderQuantity" bson:"TokenHolderQuantity"`
}

type RespERC20Holders []RespERC20HolderInfo

type RespERC20HolderCount string

// RespTopTokenHolder represents top token holder information
type RespTopTokenHolder struct {
	TokenHolderAddress     string `json:"TokenHolderAddress" bson:"TokenHolderAddress"`
	TokenHolderQuantity    string `json:"TokenHolderQuantity" bson:"TokenHolderQuantity"`
	TokenHolderAddressType string `json:"TokenHolderAddressType" bson:"TokenHolderAddressType"`
}

type RespTopTokenHolders []RespTopTokenHolder

// RespTokenInfo represents token information
type RespTokenInfo struct {
	ContractAddress string `json:"contractAddress" bson:"contractAddress"`
	TokenName       string `json:"tokenName" bson:"tokenName"`
	Symbol          string `json:"symbol" bson:"symbol"`
	Divisor         string `json:"divisor" bson:"divisor"`
	TokenType       string `json:"tokenType" bson:"tokenType"`
	TotalSupply     string `json:"totalSupply" bson:"totalSupply"`
	BlueCheckmark   string `json:"blueCheckmark" bson:"blueCheckmark"`
	Description     string `json:"description" bson:"description"`
	Website         string `json:"website" bson:"website"`
	Email           string `json:"email" bson:"email"`
	Blog            string `json:"blog" bson:"blog"`
	Reddit          string `json:"reddit" bson:"reddit"`
	Slack           string `json:"slack" bson:"slack"`
	Facebook        string `json:"facebook" bson:"facebook"`
	Twitter         string `json:"twitter" bson:"twitter"`
	Bitcointalk     string `json:"bitcointalk" bson:"bitcointalk"`
	Github          string `json:"github" bson:"github"`
	Telegram        string `json:"telegram" bson:"telegram"`
	Wechat          string `json:"wechat" bson:"wechat"`
	Linkedin        string `json:"linkedin" bson:"linkedin"`
	Discord         string `json:"discord" bson:"discord"`
	Whitepaper      string `json:"whitepaper" bson:"whitepaper"`
	TokenPriceUSD   string `json:"tokenPriceUSD" bson:"tokenPriceUSD"`
}

// RespERC20Holding represents an ERC-20 token holding
type RespERC20Holding struct {
	TokenAddress  string `json:"TokenAddress" bson:"TokenAddress"`
	TokenName     string `json:"TokenName" bson:"TokenName"`
	TokenSymbol   string `json:"TokenSymbol" bson:"TokenSymbol"`
	TokenQuantity string `json:"TokenQuantity" bson:"TokenQuantity"`
	TokenDivisor  string `json:"TokenDivisor" bson:"TokenDivisor"`
}

type RespERC20Holdings []RespERC20Holding

// RespNFTHolding represents an NFT holding
type RespNFTHolding struct {
	TokenAddress  string `json:"TokenAddress" bson:"TokenAddress"`
	TokenName     string `json:"TokenName" bson:"TokenName"`
	TokenSymbol   string `json:"TokenSymbol" bson:"TokenSymbol"`
	TokenQuantity string `json:"TokenQuantity" bson:"TokenQuantity"`
}

type RespNFTHoldings []RespNFTHolding

// RespNFTTokenInventory represents an NFT token inventory item
type RespNFTTokenInventory struct {
	TokenAddress string `json:"TokenAddress" bson:"TokenAddress"`
	TokenID      string `json:"TokenId" bson:"TokenId"`
}

type RespNFTTokenInventories []RespNFTTokenInventory

// Gas Tracker Module Response Types

type RespConfirmationTimeEstimation string

// RespGasOracle represents gas oracle information
type RespGasOracle struct {
	LastBlock       string `json:"LastBlock" bson:"LastBlock"`
	SafeGasPrice    string `json:"SafeGasPrice" bson:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice" bson:"ProposeGasPrice"`
	FastGasPrice    string `json:"FastGasPrice" bson:"FastGasPrice"`
	SuggestBaseFee  string `json:"suggestBaseFee" bson:"suggestBaseFee"`
	GasUsedRatio    string `json:"gasUsedRatio" bson:"gasUsedRatio"`
}

// RespDailyAvgGasLimit represents daily average gas limit
type RespDailyAvgGasLimit struct {
	UTCDate       string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	GasLimit      string `json:"gasLimit" bson:"gasLimit"`
}

type RespDailyAvgGasLimits []RespDailyAvgGasLimit

// RespDailyTotalGasUsed represents daily total gas used
type RespDailyTotalGasUsed struct {
	UTCDate       string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	GasUsed       string `json:"gasUsed" bson:"gasUsed"`
}

type RespDailyTotalGasUseds []RespDailyTotalGasUsed

// RespDailyAvgGasPrice represents daily average gas price
type RespDailyAvgGasPrice struct {
	UTCDate        string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp  string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	MaxGasPriceWei string `json:"maxGasPrice_Wei" bson:"maxGasPrice_Wei"`
	MinGasPriceWei string `json:"minGasPrice_Wei" bson:"minGasPrice_Wei"`
	AvgGasPriceWei string `json:"avgGasPrice_Wei" bson:"avgGasPrice_Wei"`
}

type RespDailyAvgGasPrices []RespDailyAvgGasPrice

// Stats Module Response Types

type RespTotalEthSupply string

type RespTotalEth2Supply string

// RespEthPrice represents Ethereum price information
type RespEthPrice struct {
	EthBTC          string `json:"ethbtc" bson:"ethbtc"`
	EthBTCTimestamp string `json:"ethbtc_timestamp" bson:"ethbtc_timestamp"`
	EthUSD          string `json:"ethusd" bson:"ethusd"`
	EthUSDTimestamp string `json:"ethusd_timestamp" bson:"ethusd_timestamp"`
}

// RespEtheumNodeSize represents Ethereum node size information
type RespEtheumNodeSize struct {
	BlockNumber    string `json:"blockNumber" bson:"blockNumber"`
	ChainTimeStamp string `json:"chainTimeStamp" bson:"chainTimeStamp"`
	ChainSize      string `json:"chainSize" bson:"chainSize"`
	ClientType     string `json:"clientType" bson:"clientType"`
	SyncMode       string `json:"syncMode" bson:"syncMode"`
}

type RespEtheumNodesSize []RespEtheumNodeSize

// RespNodeCount represents node count information
type RespNodeCount struct {
	UTCDate        string `json:"UTCDate" bson:"UTCDate"`
	TotalNodeCount string `json:"TotalNodeCount" bson:"TotalNodeCount"`
}

// RespDailyTxFee represents daily transaction fee
type RespDailyTxFee struct {
	UTCDate           string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp     string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	TransactionFeeEth string `json:"transactionFee_Eth" bson:"transactionFee_Eth"`
}

type RespDailyTxFees []RespDailyTxFee

// RespDailyNewAddress represents daily new address count
type RespDailyNewAddress struct {
	UTCDate         string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp   string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	NewAddressCount int64  `json:"newAddressCount" bson:"newAddressCount"`
}

type RespDailyNewAddresses []RespDailyNewAddress

// RespDailyNetworkUtilization represents daily network utilization
type RespDailyNetworkUtilization struct {
	UTCDate            string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp      string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	NetworkUtilization string `json:"networkUtilization" bson:"networkUtilization"`
}

type RespDailyNetworkUtilizations []RespDailyNetworkUtilization

// RespDailyAvgHashrate represents daily average hashrate
type RespDailyAvgHashrate struct {
	UTCDate         string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp   string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	NetworkHashRate string `json:"networkHashRate" bson:"networkHashRate"`
}

type RespDailyAvgHashrates []RespDailyAvgHashrate

// RespDailyTxCount represents daily transaction count
type RespDailyTxCount struct {
	UTCDate          string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp    string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	TransactionCount int64  `json:"transactionCount" bson:"transactionCount"`
}

type RespDailyTxCounts []RespDailyTxCount

// RespDailyAvgDifficulty represents daily average difficulty
type RespDailyAvgDifficulty struct {
	UTCDate           string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp     string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	NetworkDifficulty string `json:"networkDifficulty" bson:"networkDifficulty"`
}

type RespDailyAvgDifficulties []RespDailyAvgDifficulty

// RespEthHistoricalPrice represents Ethereum historical price
type RespEthHistoricalPrice struct {
	UTCDate       string `json:"UTCDate" bson:"UTCDate"`
	UnixTimeStamp string `json:"unixTimeStamp" bson:"unixTimeStamp"`
	Value         string `json:"value" bson:"value"`
}

type RespEthHistoricalPrices []RespEthHistoricalPrice

// Layer 2 Module Response Types

// RespPlasmaDeposit represents a plasma deposit
type RespPlasmaDeposit struct {
	BlockNumber string `json:"blockNumber" bson:"blockNumber"`
	TimeStamp   string `json:"timeStamp" bson:"timeStamp"`
	BlockReward string `json:"blockReward" bson:"blockReward"`
}

type RespPlasmaDeposits []RespPlasmaDeposit

// RespDepositTx represents a deposit transaction
type RespDepositTx struct {
	BlockNumber       string `json:"blockNumber" bson:"blockNumber"`
	TimeStamp         string `json:"timeStamp" bson:"timeStamp"`
	BlockHash         string `json:"blockHash" bson:"blockHash"`
	Hash              string `json:"hash" bson:"hash"`
	Nonce             string `json:"nonce" bson:"nonce"`
	From              string `json:"from" bson:"from"`
	To                string `json:"to" bson:"to"`
	Value             string `json:"value" bson:"value"`
	Gas               string `json:"gas" bson:"gas"`
	GasPrice          string `json:"gasPrice" bson:"gasPrice"`
	Input             string `json:"input" bson:"input"`
	CumulativeGasUsed string `json:"cumulativeGasUsed" bson:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed" bson:"gasUsed"`
	IsError           string `json:"isError" bson:"isError"`
	ErrDescription    string `json:"errDescription" bson:"errDescription"`
	TxReceiptStatus   string `json:"txreceipt_status" bson:"txreceipt_status"`
	QueueIndex        string `json:"queueIndex" bson:"queueIndex"`
	L1TransactionHash string `json:"L1transactionhash" bson:"L1transactionhash"`
	L1TxOrigin        string `json:"L1TxOrigin" bson:"L1TxOrigin"`
	TokenAddress      string `json:"tokenAddress" bson:"tokenAddress"`
	TokenSentFrom     string `json:"tokenSentFrom" bson:"tokenSentFrom"`
	TokenSentTo       string `json:"tokenSentTo" bson:"tokenSentTo"`
	TokenValue        string `json:"tokenValue" bson:"tokenValue"`
}

type RespDepositTxs []RespDepositTx

// RespWithdrawalTx represents a withdrawal transaction
type RespWithdrawalTx struct {
	BlockNumber            string `json:"blockNumber" bson:"blockNumber"`
	TimeStamp              string `json:"timeStamp" bson:"timeStamp"`
	BlockHash              string `json:"blockHash" bson:"blockHash"`
	Hash                   string `json:"hash" bson:"hash"`
	Nonce                  string `json:"nonce" bson:"nonce"`
	From                   string `json:"from" bson:"from"`
	To                     string `json:"to" bson:"to"`
	Value                  string `json:"value" bson:"value"`
	Gas                    string `json:"gas" bson:"gas"`
	GasPrice               string `json:"gasPrice" bson:"gasPrice"`
	Input                  string `json:"input" bson:"input"`
	CumulativeGasUsed      string `json:"cumulativeGasUsed" bson:"cumulativeGasUsed"`
	GasUsed                string `json:"gasUsed" bson:"gasUsed"`
	IsError                string `json:"isError" bson:"isError"`
	ErrDescription         string `json:"errDescription" bson:"errDescription"`
	TxReceiptStatus        string `json:"txreceipt_status" bson:"txreceipt_status"`
	Message                string `json:"message" bson:"message"`
	MessageNonce           string `json:"messageNonce" bson:"messageNonce"`
	Status                 string `json:"status" bson:"status"`
	L1TransactionHash      string `json:"L1transactionhash" bson:"L1transactionhash"`
	TokenAddress           string `json:"tokenAddress" bson:"tokenAddress"`
	WithdrawalType         string `json:"withdrawalType" bson:"withdrawalType"`
	TokenValue             string `json:"tokenValue" bson:"tokenValue"`
	L1TransactionHashProve string `json:"L1transactionhashProve" bson:"L1transactionhashProve"`
}

type RespWithdrawalTxs []RespWithdrawalTx

// API Credit & Chain Info Module Response Types

// RespCreditUsage represents credit usage information
type RespCreditUsage struct {
	CreditsUsed            int64  `json:"creditsUsed" bson:"creditsUsed"`
	CreditsAvailable       int64  `json:"creditsAvailable" bson:"creditsAvailable"`
	CreditLimit            int64  `json:"creditLimit" bson:"creditLimit"`
	LimitInterval          string `json:"limitInterval" bson:"limitInterval"`
	IntervalExpiryTimespan string `json:"intervalExpiryTimespan" bson:"intervalExpiryTimespan"`
}

// RespSupportedChain represents a supported chain
type RespSupportedChain struct {
	ChainName     string `json:"chainname" bson:"chainname"`
	ChainID       string `json:"chainid" bson:"chainid"`
	BlockExplorer string `json:"blockexplorer" bson:"blockexplorer"`
	APIURL        string `json:"apiurl" bson:"apiurl"`
	Status        int64  `json:"status" bson:"status"`
}

// RespSupportedChains represents supported chains information
type RespSupportedChains struct {
	TotalCount int64                `json:"totalcount" bson:"totalcount"`
	Result     []RespSupportedChain `json:"result" bson:"result"`
}

// Address Tagging & Labeling Module Response Types

// RespAddressTag represents address tag information
type RespAddressTag struct {
	Address              string   `json:"address" bson:"address"`
	Nametag              string   `json:"nametag" bson:"nametag"`
	InternalNametag      string   `json:"internal_nametag" bson:"internal_nametag"`
	URL                  string   `json:"url" bson:"url"`
	ShortDescription     string   `json:"shortdescription" bson:"shortdescription"`
	Notes1               string   `json:"notes_1" bson:"notes_1"`
	Notes2               string   `json:"notes_2" bson:"notes_2"`
	Labels               []string `json:"labels" bson:"labels"`
	LabelsSlug           []string `json:"labels_slug" bson:"labels_slug"`
	Reputation           int64    `json:"reputation" bson:"reputation"`
	OtherAttributes      []string `json:"other_attributes" bson:"other_attributes"`
	LastUpdatedTimestamp int64    `json:"lastupdatedtimestamp" bson:"lastupdatedtimestamp"`
}

type RespAddressTags []RespAddressTag

// RespLabelMaster represents label master information
type RespLabelMaster struct {
	LabelName            string `json:"labelname" bson:"labelname"`
	LabelSlug            string `json:"labelslug" bson:"labelslug"`
	ShortDescription     string `json:"shortdescription" bson:"shortdescription"`
	Notes                string `json:"notes" bson:"notes"`
	LastUpdatedTimestamp int64  `json:"lastupdatedtimestamp" bson:"lastupdatedtimestamp"`
}

type RespLabelMasterlist []RespLabelMaster

// RespLatestCSVBatchNumber represents nametag batch information
type RespLatestCSVBatchNumber struct {
	Nametag              string `json:"nametag" bson:"nametag"`
	Batch                string `json:"batch" bson:"batch"`
	LastUpdatedTimestamp int64  `json:"lastUpdatedTimestamp" bson:"lastUpdatedTimestamp"`
}

type RespLatestCSVBatchNumbers []RespLatestCSVBatchNumber
