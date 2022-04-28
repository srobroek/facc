package cli

import (
	"github.com/pterm/pcli"
	"github.com/spf13/cobra"
)

func Execute() error {

	cmd, err := newRootCmd()
	if err != nil {
		return err
	}
	pcli.SetRepo("srobroek/facc")
	pcli.SetRootCmd(cmd)
	//pcli.Setup()

	cobra.OnInitialize()

	if err = cmd.Execute(); err != nil {
		return err
	}

	return nil
}

// func debug(format string, v ...interface{}) {
// 	if settings.GetBool("debug") {
// 		format = fmt.Sprintf("[debug] %s\n", format)
// 		log.Output(2, fmt.Sprintf(format, v...))
// 	}
// }
