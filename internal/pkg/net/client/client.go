// Copyright 2025 fsyyft-ai
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package client

import (
	"context"
	"net"
	"net/http"

	"github.com/hashicorp/go.net/proxy"

	kitlog "github.com/fsyyft-go/kit/log"

	appconf "github.com/fsyyft-ai/eino-wizard/internal/pkg/conf"
)

type (
	Client interface {
		Client() *http.Client
	}
	client struct {
		client *http.Client
	}
)

func NewClient(logger kitlog.Logger, cfg *appconf.Config) (Client, func(), error) {
	clean := func() {}
	if cfg.Network.ProxyProtocol == "socks5" {
		if dialer, err := proxy.SOCKS5("tcp", cfg.Network.ProxyAddress, nil, proxy.Direct); nil != err {
			return nil, clean, err
		} else {
			c := &http.Client{
				Transport: &http.Transport{
					DialContext: func(c context.Context, network, addr string) (net.Conn, error) {
						return dialer.Dial(network, addr)
					},
				},
			}
			return &client{client: c}, clean, nil
		}
	}
	return &client{client: &http.Client{}}, clean, nil
}

func (c *client) Client() *http.Client {
	return c.client
}
