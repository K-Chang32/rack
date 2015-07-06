package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"

	"github.com/convox/cli/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/convox/cli/Godeps/_workspace/src/github.com/inconshreveable/go-update"
	"github.com/convox/cli/Godeps/_workspace/src/github.com/inconshreveable/go-update/check"
	"github.com/convox/cli/stdcli"
)

func init() {
	stdcli.RegisterCommand(cli.Command{
		Name:        "update",
		Description: "update the cli",
		Usage:       "",
		Action:      cmdUpdate,
	})
}

func cmdUpdate(c *cli.Context) {
	client, err := updateClient()

	if err != nil {
		stdcli.Error(err)
	}

	params := check.Params{
		AppVersion: Version,
		AppId:      "ap_TKxvw_eIPVyOzl6rKEonCU5DUY",
		Channel:    "stable",
	}

	updater := update.New()
	updater.HTTPClient = client

	r, err := params.CheckForUpdate("https://api.equinox.io/1/Updates", updater)

	if err != nil {
		if err != check.NoUpdateAvailable {
			stdcli.Error(err)
		}
		return
	}

	err, _ = r.Update()

	if err != nil {
		stdcli.Error(err)
		return
	}

	fmt.Printf("Updated to %s\n", r.Version)
}

func updateClient() (*http.Client, error) {
	root, err := Asset("data/root.pem")

	if err != nil {
		return nil, err
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(root)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: pool},
		},
	}

	return client, nil
}
