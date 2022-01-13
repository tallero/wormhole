package logs

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var lines = `Program F11YDwLVireDZ7zFgnjo3psyiSCW3oumYsWWaXbqR5bF invoke [1]
Program log: Unstake NFT Call
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3121 of 157104 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]
Program log: Instruction: CloseAccount
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 2297 of 150388 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program F11YDwLVireDZ7zFgnjo3psyiSCW3oumYsWWaXbqR5bF consumed 55122 of 200000 compute units
Program F11YDwLVireDZ7zFgnjo3psyiSCW3oumYsWWaXbqR5bF success
Program F11YDwLVireDZ7zFgnjo3psyiSCW3oumYsWWaXbqR5bF invoke [1]
Program log: Unstake NFT Call
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3121 of 157104 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]
Program log: Instruction: CloseAccount
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 2297 of 150388 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program F11YDwLVireDZ7zFgnjo3psyiSCW3oumYsWWaXbqR5bF consumed 55122 of 200000 compute units
Program F11YDwLVireDZ7zFgnjo3psyiSCW3oumYsWWaXbqR5bF success`

func TestParseLogs(t *testing.T) {
	val, err := ParseLogs(strings.Split(lines, "\n"))
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(val)
	fmt.Println(string(b))
}
