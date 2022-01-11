/*
   Copyright © 2021 The CDI Authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/container-orchestrated-devices/container-device-interface/pkg/cdi"
)

var specDirs []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cdi",
	Short: "Inpsect and interact with the CDI Registry",
	Long: `
The 'cdi' utility allows you to inspect and interact with the
CDI Registry. Various commands are available for listing CDI
Spec files, vendors, classes, devices, validating the content
of the registry, injecting devices into OCI Specs, and for
monitoring changes in the Registry.

See cdi --help for a list of available commands. You can get
additional help about <command> by using 'cdi <command> -h'.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initSpecDirs)
	rootCmd.PersistentFlags().StringSliceVarP(&specDirs, "spec-dirs", "d", nil, "directories to scan for CDI Spec files")
}

func initSpecDirs() {
	if len(specDirs) > 0 {
		cdi.GetRegistry(cdi.WithSpecDirs(specDirs...))
		if len(cdi.GetRegistry().GetErrors()) > 0 {
			cdiPrintRegistryErrors()
		}
	}
}
