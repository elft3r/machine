package commands

import (
	"encoding/json"
	"fmt"

	"github.com/docker/machine/log"
	"github.com/docker/machine/drivers"

	"github.com/codegangsta/cli"
)

func cmdPrintDriverFlags(c *cli.Context) {
	driverFlags := make(map[string][]cli.Flag)
	for _, driverName := range drivers.GetDriverNames() {
		flags, err := drivers.GetCreateFlagsForDriver(driverName)
		if err != nil {
			log.Fatal(err)
		}
		driverFlags[driverName] = flags
	}

	j, err := json.MarshalIndent(driverFlags, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j)) }

