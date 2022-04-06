package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"

	"github.com/YaleUniversity/packer-provisioner-goss/v3/provisioner/goss"
	"github.com/YaleUniversity/packer-provisioner-goss/v3/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterProvisioner(plugin.DEFAULT_NAME, new(goss.Provisioner))
	pps.SetVersion(version.PluginVersion)
	if err := pps.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
