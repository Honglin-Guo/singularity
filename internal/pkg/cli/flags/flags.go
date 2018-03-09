/*
  Copyright (c) 2018, Sylabs, Inc. All rights reserved.

  This software is licensed under a 3-clause BSD license.  Please
  consult LICENSE file distributed with the sources of this project regarding
  your rights to use or distribute this software.
*/
package sflags

import (
	"log"
	"os/user"

	"github.com/spf13/pflag"
)

// flags.go contains flag variables for all commands to use
var (
	BindPaths   []string
	HomePath    string
	OverlayPath string
	ScratchPath string
	WorkdirPath string
	PwdPath     string
	ShellPath   string
	Hostname    string

	IsBoot       bool
	IsFakeroot   bool
	IsCleanEnv   bool
	IsContained  bool
	IsContainAll bool
	Nvidia       bool

	NetNamespace  bool
	UtsNamespace  bool
	UserNamespace bool
	PidNamespace  bool
	IpcNamespace  bool

	AllowSUID bool
	KeepPrivs bool
	NoPrivs   bool
	AddCaps   []string
	DropCaps  []string
)

var Flags *pflag.FlagSet = pflag.NewFlagSet("SFlags", pflag.ExitOnError)

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return usr.HomeDir
}

func init() {
	initPathVars()
	initBoolVars()
	initNamespaceVars()
	initPrivilegeVars()
}

// initPathVars initializes flags that take a string argument
func initPathVars() {
	// -B|--bind
	Flags.StringSliceVarP(&BindPaths, "bind", "B", []string{}, "A user-bind path specification.  spec has the format src[:dest[:opts]], where src and dest are outside and inside paths.  If dest is not given, it is set equal to src.  Mount options ('opts') may be specified as 'ro' (read-only) or 'rw' (read/write, which is the default). Multiple bind paths can be given by a comma separated list.")
	Flags.SetAnnotation("bind", "argtag", []string{"<spec>"})

	// -H|--home
	Flags.StringVarP(&HomePath, "home", "H", getHomeDir(), "A home directory specification.  spec can either be a src path or src:dest pair.  src is the source path of the home directory outside the container and dest overrides the home directory within the container.")
	Flags.SetAnnotation("home", "argtag", []string{"<spec>"})

	// -o|--overlay
	Flags.StringVarP(&OverlayPath, "overlay", "o", "", "Use a persistent overlayFS via a writable image.")
	Flags.SetAnnotation("overlay", "argtag", []string{"<path>"})

	// -S|--scratch
	Flags.StringVarP(&ScratchPath, "scratch", "S", "", "Include a scratch directory within the container that is linked to a temporary dir (use -W to force location)")
	Flags.SetAnnotation("scratch", "argtag", []string{"<path>"})

	// -W|--workdir
	Flags.StringVarP(&WorkdirPath, "workdir", "W", "", "Working directory to be used for /tmp, /var/tmp and $HOME (if -c/--contain was also used)")
	Flags.SetAnnotation("workdir", "argtag", []string{"<path>"})

	// -s|--shell
	Flags.StringVarP(&ScratchPath, "shell", "s", "", "Path to program to use for interactive shell")
	Flags.SetAnnotation("shell", "argtag", []string{"<path>"})

	// --pwd
	Flags.StringVar(&ScratchPath, "pwd", "", "Include a scratch directory within the container that is linked to a temporary dir (use -W to force location).")
	Flags.SetAnnotation("scratch", "argtag", []string{"<path>"})

	// --hostname
	Flags.StringVar(&Hostname, "hostname", "", "Set container hostname")
	Flags.SetAnnotation("hostname", "argtag", []string{"<name>"})
}

// initBoolVars initializes flags that take a boolean argument
func initBoolVars() {
	// --boot
	Flags.BoolVar(&IsBoot, "boot", false, "Execute /sbin/init to boot container (root only)")

	// -f|--fakeroot
	Flags.BoolVarP(&IsFakeroot, "fakeroot", "f", false, "Run container in new user namespace as uid 0")

	// -e|--cleanenv
	Flags.BoolVarP(&IsCleanEnv, "cleanenv", "e", false, "Clean environment before running container")

	// -c|--contain
	Flags.BoolVarP(&IsContained, "contain", "c", false, "Use minimal /dev and empty other directories (e.g. /tmp and $HOME) instead of sharing filesystems from your host.")

	// -C|--containall
	Flags.BoolVarP(&IsContainAll, "containall", "C", false, "Contain not only file systems, but also PID, IPC, and environment")

	// --nv
	Flags.BoolVar(&Nvidia, "nv", false, "Enable experimental Nvidia support")
}

// initNamespaceVars initializes flags that take toggle namespace support
func initNamespaceVars() {
	// -p|--pid
	Flags.BoolVarP(&PidNamespace, "pid", "p", false, "Run container in a new PID namespace")

	// -i|--ipc
	Flags.BoolVarP(&IpcNamespace, "ipc", "i", false, "Run container in a new IPC namespace")

	// -n|--net
	Flags.BoolVarP(&NetNamespace, "net", "n", false, "Run container in a new network namespace (loopback is the only network device active).")

	// --uts
	Flags.BoolVar(&UtsNamespace, "uts", false, "Run container in a new UTS namespace")

	// -u|--userns
	Flags.BoolVarP(&UserNamespace, "userns", "u", false, "Run container in a new user namespace, allowing Singularity to run completely unprivileged on recent kernels. This may not support every feature of Singularity.")

}

// initPrivilegeVars initializes flags that manipulate privileges
func initPrivilegeVars() {
	// --keep-privs
	Flags.BoolVar(&KeepPrivs, "keep-privs", false, "Let root user keep privileges in container")

	// --no-privs
	Flags.BoolVar(&NoPrivs, "no-privs", true, "Drop all privileges from root user in container")

	// --add-caps
	Flags.StringSliceVar(&AddCaps, "add-caps", []string{}, "A comma separated capability list to add")

	// --drop-caps
	Flags.StringSliceVar(&DropCaps, "drop-caps", []string{}, "A comma separated capability list to drop")

	// --allow-setuid
	Flags.BoolVar(&AllowSUID, "allow-setuid", false, "Allow setuid binaries in container (root only)")
}
