/*
Copyright © 2022 windvalley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package vault

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/windvalley/gossh/internal/pkg/aes"
	"github.com/windvalley/gossh/pkg/util"
)

// EncryptCmd represents the vault encrypt command
var EncryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt sensitive content",
	Long: `
Encrypt sensitive content.`,
	Example: `
    # Encrypt plaintext by asking for vault password.
    $ gossh vault encrypt your-sensitive-content

    # Encrypt plaintext by vault password file.
    $ gossh vault encrypt your-sensitive-content -V /path/vault-password-file`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			util.CheckErr("requires one arg to represent the plaintxt to be encrypted")
		}

		if len(args) > 1 {
			util.CheckErr("to many args, only need one")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		vaultPass := getVaultConfirmPassword()
		encryptContent, err := aes.AES256Encode(args[0], vaultPass)
		if err != nil {
			err = fmt.Errorf("encrypt failed: %w", err)
		}
		util.CheckErr(err)

		fmt.Println(encryptContent)
	},
}

func init() {
	Cmd.AddCommand(EncryptCmd)
}
