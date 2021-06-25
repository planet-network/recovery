package main

import (
	"github.com/spf13/cobra"
)

func createCommandsStructure() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "recovery",
		Short: "reference i2i functions",
		Long:  `reference i2i functions`,
		Run:   nil,
	}

	rootCmd.AddCommand(createRevokeCommand())
	rootCmd.AddCommand(createStoreCommand())

	return rootCmd
}

func createRevokeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "revoke",
		Short: "revoke signing key scenario",
		Long:  `revoke signing key scenario`,
		Run:   revoke,
	}
}

func createStoreCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "store",
		Short: "store encrypted keychain",
		Long:  `store encrypted keychain`,
		Run:   store,
	}
}
