package commands

import (
	"github.com/codegangsta/cli"
	"github.com/docker/machine/log"
)

func cmdRegenerateCerts(c *cli.Context) {
	force := c.Bool("force")
	if force || confirmInput("Regenerate TLS machine certs?  Warning: this is irreversible.") {
		log.Infof("Regenerating TLS certificates")

		// Update with new SubjectAltNames if provided
		if len(c.GlobalStringSlice("tls-san")) > 0 {
			machines, err := getHosts(c)
			if err != nil {
				log.Fatal(err)
			}

			if len(machines) == 0 {
				log.Fatal(ErrNoMachineSpecified)
			}

			for i := range machines {
				machines[i].HostOptions.AuthOptions.ServerCertSANs = c.GlobalStringSlice("tls-san")
				if err := machines[i].SaveConfig(); err != nil {
					log.Fatal(err)
				}
			}
		}

		if err := runActionWithContext("configureAuth", c); err != nil {
			log.Fatal(err)
		}
	}
}
