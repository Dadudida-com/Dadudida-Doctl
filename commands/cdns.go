/*
Copyright 2018 The Doctl Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"errors"
	"fmt"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/doctl/commands/displayers"
	"github.com/digitalocean/doctl/do"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

// CDN creates the CDN command
func CDN() *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Use:   "cdn",
			Short: "cdn commands",
			Long:  "cdn is used to access cdn commands",
		},
	}

	CmdBuilder(cmd, RunCDNList, "list", "list cdn", Writer,
		aliasOpt("ls"), displayerType(&displayers.CDN{}))

	cmdCDNCreate := CmdBuilder(cmd, RunCDNCreate, "create <cdn-origin>", "create a cdn", Writer,
		aliasOpt("c"), displayerType(&displayers.CDN{}))
	AddIntFlag(cmdCDNCreate, doctl.ArgCDNTTL, "", 3600, "CDN ttl")

	cmdRunCDNDelete := CmdBuilder(cmd, RunCDNDelete, "delete <cdn-id>", "delete a cdn", Writer,
		aliasOpt("rm"))
	AddBoolFlag(cmdRunCDNDelete, doctl.ArgForce, doctl.ArgShortForce, false, "Force cdn delete")

	CmdBuilder(cmd, RunCDNGet, "get <cdn-id>", "get a cdn", Writer, aliasOpt("g"),
		displayerType(&displayers.CDN{}))

	cmdCDNUpdate := CmdBuilder(cmd, RunCDNUpdateTTL, "update <cdn-id>", "update a cdn", Writer,
		aliasOpt("u"), displayerType(&displayers.CDN{}))
	AddIntFlag(cmdCDNUpdate, doctl.ArgCDNTTL, "", 3600, "cdn ttl")

	cmdCDNFlushCache := CmdBuilder(cmd, RunCDNFlushCache, "flush <cdn-id>", "flush cdn cache", Writer,
		aliasOpt("fc"))
	AddStringSliceFlag(cmdCDNFlushCache, doctl.ArgCDNFiles, "", []string{"*"}, "cdn files")

	return cmd
}

// RunCDNList returns a list of CDNs.
func RunCDNList(c *CmdConfig) error {
	cdns, err := c.CDNs().List()
	if err != nil {
		return err
	}

	return c.Display(&displayers.CDN{CDNs: cdns})
}

// RunCDNGet returns an individual CDN.
func RunCDNGet(c *CmdConfig) error {
	if len(c.Args) == 0 {
		return doctl.NewMissingArgsErr(c.NS)
	}

	id := c.Args[0]
	item, err := c.CDNs().Get(id)
	if err != nil {
		return err
	}

	return c.Display(&displayers.CDN{CDNs: []do.CDN{*item}})
}

// RunCDNCreate creates a cdn.
func RunCDNCreate(c *CmdConfig) error {
	if len(c.Args) == 0 {
		return doctl.NewMissingArgsErr(c.NS)
	}

	origin := c.Args[0]

	ttl, err := c.Doit.GetInt(c.NS, doctl.ArgCDNTTL)
	if err != nil {
		return err
	}
	if ttl <= 0 {
		return errors.New("ttl cannot be zero or negative")
	}

	createCDN := &godo.CDNCreateRequest{
		Origin: origin,
		TTL:    uint32(ttl),
	}

	item, err := c.CDNs().Create(createCDN)
	if err != nil {
		return err
	}

	return c.Display(&displayers.CDN{CDNs: []do.CDN{*item}})
}

// RunCDNUpdateTTL updates an individual cdn ttl
func RunCDNUpdateTTL(c *CmdConfig) error {
	if len(c.Args) == 0 {
		return doctl.NewMissingArgsErr(c.NS)
	}

	id := c.Args[0]

	cs := c.CDNs()

	ttl, err := c.Doit.GetInt(c.NS, doctl.ArgCDNTTL)
	if err != nil {
		return err
	}
	if ttl <= 0 {
		return errors.New("ttl cannot be zero or negative")
	}

	updateCDN := &godo.CDNUpdateRequest{TTL: uint32(ttl)}

	item, err := cs.UpdateTTL(id, updateCDN)
	if err != nil {
		return err
	}

	return c.Display(&displayers.CDN{CDNs: []do.CDN{*item}})
}

// RunCDNDelete deletes a cdn.
func RunCDNDelete(c *CmdConfig) error {
	if len(c.Args) == 0 {
		return doctl.NewMissingArgsErr(c.NS)
	}

	force, err := c.Doit.GetBool(c.NS, doctl.ArgForce)
	if err != nil {
		return err
	}

	if force || AskForConfirm("delete cdn") == nil {
		id := c.Args[0]
		return c.CDNs().Delete(id)
	}

	return fmt.Errorf("operation aborted")
}

// RunCDNFlushCache flushes the cache of an individual cdn
func RunCDNFlushCache(c *CmdConfig) error {
	if len(c.Args) == 0 {
		return doctl.NewMissingArgsErr(c.NS)
	}

	id := c.Args[0]

	files, err := c.Doit.GetStringSlice(c.NS, doctl.ArgCDNFiles)
	if err != nil {
		return err
	}

	flushCDN := &godo.CDNFlushCacheRequest{Files: files}

	return c.CDNs().FlushCache(id, flushCDN)
}
