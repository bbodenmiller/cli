package v6

import (
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/translatableerror"
	"code.cloudfoundry.org/cli/command/v6/shared"
)

//go:generate counterfeiter . DeleteBuildpackActor

type DeleteBuildpackActor interface {
	CloudControllerAPIVersion() string
}

type DeleteBuildpackCommand struct {
	RequiredArgs    flag.BuildpackName `positional-args:"yes"`
	Force           bool               `short:"f" description:"Force deletion without confirmation"`
	Stack           string             `short:"s" description:"Specify stack to disambiguate buildpacks with the same name. Required when buildpack name is ambiguous"`
	usage           interface{}        `usage:"CF_NAME delete-buildpack BUILDPACK [-f] [-s STACK]"`
	relatedCommands interface{}        `related_commands:"buildpacks"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       DeleteBuildpackActor
}

func (cmd *DeleteBuildpackCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	ccClient, uaaClient, err := shared.GetNewClientsAndConnectToCF(config, ui)
	if err != nil {
		return err
	}
	cmd.Actor = v2action.NewActor(ccClient, uaaClient, config)
	return nil
}

func (cmd DeleteBuildpackCommand) Execute(args []string) error {
	return translatableerror.UnrefactoredCommandError{}
}
