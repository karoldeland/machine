  
package main

import (
	"github.com/karoldeland/machine/drivers/azure"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(azure.NewDriver("", ""))
}
