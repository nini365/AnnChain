/*
 * This file is part of The AnnChain.
 *
 * The AnnChain is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The AnnChain is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The www.annchain.io.  If not, see <http://www.gnu.org/licenses/>.
 */


package tools

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_Crypto(t *testing.T) {
	text := "aa"
	pwd := "aa"
	ret, err := Encrypt([]byte(text), []byte(pwd))
	fmt.Println("Encrypt:", ret, ",err:", err, ",hex:", hex.EncodeToString(ret))
	deret, err := Decrypt(ret, []byte(pwd))
	fmt.Println("Decrypt:", string(deret), ",err:", err)
}

func Test_ToPubkey(t *testing.T) {
	priv := "9D2EBF7939A41458DC01500736D816F0804A5F467615011B9A51C1A6626E1FAD1B020FAA48FEE672B9C5AA75AEDC072EB57BD02143E0ECE55C0D525290A382B8"
	fmt.Println("pubkey", ED25519Pubkey(priv))
}