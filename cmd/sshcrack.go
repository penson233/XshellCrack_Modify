package cmd

import (
	"fmt"
	"gorun/ssh"

	"github.com/spf13/cobra"
)

var (
	path string
	USID string
	Name string
	Directory string
	xshFile []string
)
var sshCrack = &cobra.Command{
	Use:   "SshCrack",
	Short: "The realization of xshell password cracking",
	PreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Lookup("USID").Value.String() !="" && cmd.Flags().Lookup("username").Value.String() !=""{
			fmt.Println("[*]find this localname is :" + cmd.Flags().Lookup("USID").Value.String())
			fmt.Println("[*]find this   SID     is :" + cmd.Flags().Lookup("username").Value.String())
			fmt.Println("[*]crack ssh loading...................")
		}else{
			localname, SID := ssh.GetUserSid()
			fmt.Println("[*]find this localname is :" + localname)
			fmt.Println("[*]find this   SID     is :" + SID)
			fmt.Println("[*]crack ssh loading...................")
		}

	},
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Lookup("Path").Value.String() == "" && cmd.Flags().Lookup("directory").Value.String() == ""{
			m := ssh.Run()
			for _, v := range m {
				ssh.InitXSh(v)
			}
		} else {
			if cmd.Flags().Lookup("USID").Value.String() !="" && cmd.Flags().Lookup("username").Value.String() !=""{
				if cmd.Flags().Lookup("Path").Value.String()!=""{
					File := cmd.Flags().Lookup("Path").Value.String()
					name := cmd.Flags().Lookup("username").Value.String()
					usid := cmd.Flags().Lookup("USID").Value.String()
					ssh.RunFile2(File,name,usid)
				}else if cmd.Flags().Lookup("directory").Value.String() !=""{
					path := cmd.Flags().Lookup("directory").Value.String();
					name := cmd.Flags().Lookup("username").Value.String()
					usid := cmd.Flags().Lookup("USID").Value.String()
					xshFile :=ssh.GetFileName(path);
					m:=ssh.ReadXshFile(xshFile);
					for _, v := range m {
						ssh.InitXSh2(v,name,usid);
					}
				}
			}else if cmd.Flags().Lookup("Path").Value.String() !=""{
				File := cmd.Flags().Lookup("Path").Value.String()
				ssh.RunFile(File)
			}else if cmd.Flags().Lookup("directory").Value.String() !=""{

				path := cmd.Flags().Lookup("directory").Value.String();
				xshFile :=ssh.GetFileName(path);
				m:=ssh.ReadXshFile(xshFile);
				for _, v := range m {
					ssh.InitXSh(v)
				}

			}

		}
		// fmt.Printf("cmd.Flags().Lookup(\"path\").Value: %v\n", cmd.Flags().Lookup("Path").Value)
	},
}

func init() {
	rootCmd.AddCommand(sshCrack)
	sshCrack.Flags().StringVarP(&path, "Path", "P", "", "eg -P C:\\xxxx\\xxx.xsh")
	sshCrack.Flags().StringVarP(&USID, "USID", "I", "", "eg -I S-1xxxxxx")
	sshCrack.Flags().StringVarP(&Name, "username", "u", "", "eg -u test")
	sshCrack.Flags().StringVarP(&Directory, "directory", "d", "", "eg -d C:\\xxxx\\")

}
