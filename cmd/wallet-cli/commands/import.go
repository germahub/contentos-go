package commands

import (
	"bufio"
	"context"
	"fmt"
	"github.com/coschain/cobra"
	"github.com/coschain/contentos-go/cmd/wallet-cli/commands/utils"
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet"
	"github.com/coschain/contentos-go/prototype"
	"github.com/coschain/contentos-go/rpc/pb"
	"os"
	"strings"
)

var importForceFlag bool

var ImportCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "import",
		Short:   "import an account",
		Example: "import [name] [privkey]",
		Args:    cobra.ExactArgs(2),
		Run:     importAccount,
	}
	cmd.Flags().BoolVarP(&importForceFlag, "force", "f", false, "import --force")
	return cmd
}

var ImportFromMnemonicCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use: "import_from_mnemonic",
		Short: "import an account from mnemonic",
		Example: "import_from_mnemonic [name]",
		Args: cobra.ExactArgs(1),
		Run: importFromMnemonic,
	}
	cmd.Flags().BoolVarP(&importForceFlag, "force", "f", false, "import_from_mnemonic --force")
	return cmd
}

func importAccount(cmd *cobra.Command, args []string) {
	defer func() {
		importForceFlag = false
	}()
	c := cmd.Context["rpcclient"]
	client := c.(grpcpb.ApiServiceClient)
	w := cmd.Context["wallet"]
	r := cmd.Context["preader"]
	preader := r.(utils.PasswordReader)
	mywallet := w.(wallet.Wallet)
	name := args[0]
	privKeyStr := args[1]
	passphrase, err := utils.GetPassphrase(preader)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !importForceFlag {
		// just try to load or reload, if the name exist then we can find it in next step
		_ = mywallet.Load(name)
		ok := mywallet.IsExist(name)
		if ok {
			fmt.Println(fmt.Sprintf("the account: %s has been in your local keychain, please load it or you can import -f",
				name))
			return
		}
	}
	privKey, err := prototype.PrivateKeyFromWIF(privKeyStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	pubKey, err := privKey.PubKey()
	if err != nil {
		fmt.Println(err)
		return
	}
	pubKeyStr := pubKey.ToWIF()
	req := &grpcpb.GetAccountByNameRequest{AccountName: &prototype.AccountName{Value: name}}
	resp, err := client.GetAccountByName(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		pubkey := resp.Info.PublicKey
		is_exist := pubKeyStr == pubkey.ToWIF()

		if is_exist {
			// the pubkey and account name should be check by api
			err = mywallet.Create(name, passphrase, pubKeyStr, privKeyStr)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(fmt.Sprintf("pubkey %s doesn't match %s", pubKeyStr, name))
		}
		//_ = is_exist
		//err = mywallet.Create(name, passphrase, pubKeyStr, privKeyStr)
		//if err != nil {
		//	fmt.Println(err)
		//}
	}

}

func importFromMnemonic(cmd *cobra.Command, args []string) {
	defer func() {
		importForceFlag = false
	}()
	c := cmd.Context["rpcclient"]
	client := c.(grpcpb.ApiServiceClient)
	w := cmd.Context["wallet"]
	mywallet := w.(wallet.HDWallet)
	r := cmd.Context["preader"]
	preader := r.(utils.PasswordReader)
	name := args[0]

	if !importForceFlag {
		// just try to load or reload, if the name exist then we can find it in next step
		_ = mywallet.Load(name)
		ok := mywallet.IsExist(name)
		if ok {
			fmt.Println(fmt.Sprintf("the account: %s has been in your local keychain, please load it or you can import -f",
				name))
			return
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter mnemonic > ")
	mnemonic, _ := reader.ReadString('\n')
	// remove \n
	mnemonic = strings.Replace(mnemonic, "\n", "", -1)

	passphrase, err := utils.GetPassphrase(preader)
	if err != nil {
		fmt.Println(err)
		return
	}

	pubKeyStr, privKeyStr, err := mywallet.GenerateFromMnemonic(mnemonic)
	if err != nil {
		fmt.Println(err)
		return
	}

	req := &grpcpb.GetAccountByNameRequest{AccountName: &prototype.AccountName{Value: name}}
	resp, err := client.GetAccountByName(context.Background(), req)

	if err != nil {
		fmt.Println(err)
	} else {
		pubkey := resp.Info.PublicKey
		isExist := pubKeyStr == pubkey.ToWIF()
		if isExist {
			err = mywallet.Create(name, passphrase, pubKeyStr, privKeyStr)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(fmt.Sprintf("pubkey %s doesn't match %s", pubKeyStr, name))
		}
	}
}

