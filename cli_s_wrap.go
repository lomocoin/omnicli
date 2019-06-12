package btccli

import (
	"encoding/json"
	"github.com/lemon-sunxiansong/btccli/btcjson"
	"os/exec"
	"strconv"
)

// CliSendtoaddress https://bitcoin.org/en/developer-reference#sendtoaddress
func CliSendtoaddress(cmd *btcjson.SendToAddressCmd) (string, error) {
	args := []string{
		CmdParamRegtest,
		"sendtoaddress",
		cmd.Address,
		strconv.FormatFloat(cmd.Amount, 'f', 6, 64),
	}
	if cmd.Comment != nil {
		args = append(args, *cmd.Comment)
	} else {
		args = append(args, "")
	}

	if cmd.CommentTo != nil {
		args = append(args, *cmd.CommentTo)
	} else {
		args = append(args, "")
	}
	//TODO support other params
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	//TODO validate hex
	return cmdPrint, nil
}

// CliSendrawtransaction https://bitcoin.org/en/developer-reference#sendrawtransaction
func CliSendrawtransaction(cmd btcjson.SendRawTransactionCmd) (string, error) {
	args := []string{
		CmdParamRegtest,
		"sendrawtransaction",
		cmd.HexTx,
	}
	if cmd.AllowHighFees != nil {
		args = append(args, strconv.FormatBool(*cmd.AllowHighFees))
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	//TODO validate hex
	return cmdPrint, nil
}

// CliSignrawtransactionwithkey https://bitcoin.org/en/developer-reference#signrawtransactionwithkey
func CliSignrawtransactionwithkey(cmd btcjson.SignRawTransactionCmd) (btcjson.SignRawTransactionResult, error) {
	args := []string{
		CmdParamRegtest,
		"signrawtransactionwithkey",
		cmd.RawTx,
		ToJson(cmd.PrivKeys),
		IfOrString(len(cmd.Prevtxs) > 0, ToJson(cmd.Prevtxs), ""),
	}
	if cmd.Sighashtype != nil {
		args = append(args, *cmd.Sighashtype)
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	var res btcjson.SignRawTransactionResult
	err := json.Unmarshal([]byte(cmdPrint), &res)
	return res, err
}