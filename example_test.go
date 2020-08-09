package jsoncr_test

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/thedevop1/jsoncr"
)

func ExampleRemove() {
	const j = `// JSON with commments
{	"configDesc": "configuration data", // config comment
	/* configBool:
		 false (default)
	*/
	"configBool": true					// override default
}`
	condenseJ, err := jsoncr.Remove(strings.NewReader(j))
	if err != nil {
		fmt.Println(err)
		return
	}
	var z interface{}
	err = json.Unmarshal(condenseJ, &z)
	if err != nil {
		fmt.Println("error Unmarshal", err)
		return
	}
	fmt.Println(z)

	// Output:
	// map[configBool:true configDesc:configuration data]
}
