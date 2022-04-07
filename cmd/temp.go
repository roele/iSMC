// Copyright (C) 2019  Dinko Korunic
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

//go:build darwin

package cmd

import (
	"fmt"

	"github.com/dkorunic/iSMC/hid"
	"github.com/dkorunic/iSMC/smc"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// rootCmd represents temp command.
var tempCmd = &cobra.Command{
	Use:     "temp",
	Aliases: []string{"temperature", "tmp"},
	Short:   "Display temperature sensors",
	Run: func(cmd *cobra.Command, args []string) {
		t := table.NewWriter()
		defer t.Render()

		setupTable(t)
		fmt.Println("Temperature:")
		smc.PrintTemp(t)
		hid.PrintTemp(t)
	},
}

func init() {
	rootCmd.AddCommand(tempCmd)
}
