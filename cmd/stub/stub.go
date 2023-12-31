package stub

import (
	"fmt"
	"github.com/mileapp_screening/utils/stub"

	"github.com/spf13/cobra"
)

func MakeStub() {
	stub.MakeStubs()
}

func GenerateFromStub(module string) {
	stub.Stubs(module)
}

func GenerateFromTemplateStub(templateType string, templateName string, name string) {
	stub.TemplateStub(templateType, templateName, name)
}

var TemplateCmd = &cobra.Command{
	Use:   "template [COMANDS]",
	Short: "Generate from template or make template",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you must specify a comand like generate")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

var UsecaseCmd = &cobra.Command{
	Use:   "usecase [NAME]",
	Short: "Generate usecase from template",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you must give a name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		GenerateFromTemplateStub("usecase", "default", args[0])
	},
}

var RepositoryCmd = &cobra.Command{
	Use:   "repository [NAME]",
	Short: "Generate repository from template",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you must give a name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		GenerateFromTemplateStub("repository", "default", args[0])
	},
}

var HandlerCmd = &cobra.Command{
	Use:   "handler [NAME]",
	Short: "Generate handler from template",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you must give a name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		GenerateFromTemplateStub("handler", "default", args[0])
	},
}

func init() {
	TemplateCmd.AddCommand(UsecaseCmd)
	TemplateCmd.AddCommand(RepositoryCmd)
	TemplateCmd.AddCommand(HandlerCmd)
}
