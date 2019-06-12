package btccli

import (
	"fmt"
	"testing"

	"github.com/lemon-sunxiansong/btccli/btcjson"
)

func TestCliCreatemultisig(t *testing.T) {
	closeChan, err := BitcoindRegtest()
	trueThenFailNow(t, err != nil, "Failed to start btcd", err)
	defer func() {
		closeChan <- struct{}{}
	}()

	type addrinfo struct {
		addr, privkey, pubkey string
	}
	var addrs [5]addrinfo
	{ //获取几个新地址
		for i := 0; i < len(addrs); i++ {
			add, err := CliGetnewaddress(nil, nil)
			trueThenFailNow(t, err != nil, "Failed to get new address", err)
			addrs[i] = addrinfo{addr: add}

			info, err := CliGetAddressInfo(add)
			trueThenFailNow(t, err != nil, "Failed to get address info", err)
			addrs[i].pubkey = info.Pubkey

			privkey, err := CliDumpprivkey(add)
			trueThenFailNow(t, err != nil, "Failed to dump privkey", err)
			addrs[i].privkey = privkey
		}
		fmt.Println("addrs", addrs)
	}

	var multisigResp btcjson.CreateMultiSigResult
	{ //create multisig address
		var keys []string
		for _, info := range addrs {
			keys = append(keys, info.pubkey)
		}
		multisigResp, err = CliCreatemultisig(4, keys, nil)
		trueThenFailNow(t, err != nil, "Failed to create multi sig", err)
		fmt.Println("multisig address:", jsonStr(multisigResp))
	}

	{
		info, err := CliGetAddressInfo(multisigResp.Address)
		trueThenFailNow(t, err != nil, "Failed to get address info", err)
		fmt.Println("multisigAddres info", ToJsonIndent(info))
	}
	{
		vRes, err := CliValidateaddress(multisigResp.Address)
		trueThenFailNow(t, err != nil, "Failed to validate address info", err)
		fmt.Println("validate multisig address", ToJsonIndent(vRes))
	}

}

func TestCliAddmultisigaddress(t *testing.T) {
	closeChan, err := BitcoindRegtest()
	trueThenFailNow(t, err != nil, "Failed to start btcd", err)
	defer func() {
		closeChan <- struct{}{}
	}()

	type addrinfo struct {
		addr, privkey, pubkey string
	}
	var addrs [5]addrinfo
	{ //获取几个新地址
		for i := 0; i < len(addrs); i++ {
			add, err := CliGetnewaddress(nil, nil)
			trueThenFailNow(t, err != nil, "Failed to get new address", err)
			addrs[i] = addrinfo{addr: add}

			info, err := CliGetAddressInfo(add)
			trueThenFailNow(t, err != nil, "Failed to get address info", err)
			addrs[i].pubkey = info.Pubkey

			privkey, err := CliDumpprivkey(add)
			trueThenFailNow(t, err != nil, "Failed to dump privkey", err)
			addrs[i].privkey = privkey
		}
		fmt.Println("addrs", addrs)
	}

	var multisigResp btcjson.CreateMultiSigResult
	{ //create multisig address
		var keys []string
		for _, info := range addrs {
			keys = append(keys, info.pubkey)
		}

		multisigResp, err = CliAddmultisigaddress(btcjson.AddMultisigAddressCmd{
			NRequired: 3, Keys: keys,
		})
		trueThenFailNow(t, err != nil, "Failed to add multi sig address", err)
		fmt.Println("multisig address:", jsonStr(multisigResp))
	}

	{
		info, err := CliGetAddressInfo(multisigResp.Address)
		trueThenFailNow(t, err != nil, "Failed to get address info", err)
		fmt.Println("multisigAddres info", ToJsonIndent(info))
	}
}
