// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package cli

import (
	"fmt"
	"strings"

	"github.com/singularityware/singularity/src/docs"
	"github.com/singularityware/singularity/src/pkg/libexec"
	"github.com/spf13/cobra"
)

const (
	// SyCloudLibrary holds sylabs cloud library base URI
	// for more info refer to https://cloud.sylabs.io/library
	SyCloudLibrary = "library"
	// Shub holds singularity hub base URI
	// for more info refer to https://singularity-hub.org/
	Shub = "shub"
)

var (
	// PullLibraryURI holds the base URI to a Sylabs library API instance
	PullLibraryURI string
)

func init() {
	PullCmd.Flags().SetInterspersed(false)

	PullCmd.Flags().StringVar(&PullLibraryURI, "library", "https://library.sylabs.io", "")
	PullCmd.Flags().BoolVarP(&force, "force", "F", false, "overwrite an image file if it exists")

	SingularityCmd.AddCommand(PullCmd)
}

// PullCmd singularity pull
var PullCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:   cobra.RangeArgs(1, 2),
	PreRun: sylabsToken,
	Run: func(cmd *cobra.Command, args []string) {
		var uri string
		if len(args) == 2 {
			uri = args[1]
		} else {
			uri = args[0]
		}
		BaseURI := strings.Split(uri, "://")
		switch BaseURI[0] {
		case SyCloudLibrary:
			if len(args) == 2 {
				libexec.PullImage(args[0], uri, PullLibraryURI, force, authToken)
				return
			}
			libexec.PullImage("", args[0], PullLibraryURI, force, authToken)
		case Shub:
			fmt.Println("Shub not yet supported")
		default:
			fmt.Println(BaseURI[0], "Not a supported URI")
		}
	},

	Use:     docs.PullUse,
	Short:   docs.PullShort,
	Long:    docs.PullLong,
	Example: docs.PullExample,
}
