/*
Copyright © 2023 Chihiro Hasegawa <encry1024@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/owlinux1000/gcstree/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gcstree <bucket>",
	Short: "A tree command for Google Cloud Storage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if ok, _ := cmd.Flags().GetBool("version"); ok {
			fmt.Printf("gcstree v%s\n", internal.GCSTREE_VERSION)
			os.Exit(0)
		}
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		bucket := args[0]
		var ctx context.Context = context.Background()
		gcsTree, err := internal.NewGCSTree(ctx, bucket)
		if err != nil {
			log.Fatal(err)
		}
		result, err := gcsTree.Tree()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "show the gcstree version")
}
