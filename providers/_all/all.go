// Package all is simply a container to reference all known provider implementations for easy import into other packages
package all

import (
	// Define all known providers here. They should each register themselves with the providers package via init function.
	_ "github.com/StackExchange/dnscontrol/providers/activedir"
	_ "github.com/StackExchange/dnscontrol/providers/bind"
	_ "github.com/StackExchange/dnscontrol/providers/cloudflare"
	_ "github.com/StackExchange/dnscontrol/providers/coredns"
	_ "github.com/StackExchange/dnscontrol/providers/digitalocean"
	_ "github.com/StackExchange/dnscontrol/providers/dnsimple"
	_ "github.com/StackExchange/dnscontrol/providers/exoscale"
	_ "github.com/StackExchange/dnscontrol/providers/gandi"
	_ "github.com/StackExchange/dnscontrol/providers/gcloud"
	_ "github.com/StackExchange/dnscontrol/providers/hexonet"
	_ "github.com/StackExchange/dnscontrol/providers/linode"
	_ "github.com/StackExchange/dnscontrol/providers/namecheap"
	_ "github.com/StackExchange/dnscontrol/providers/namedotcom"
	_ "github.com/StackExchange/dnscontrol/providers/ns1"
	_ "github.com/StackExchange/dnscontrol/providers/octodns"
	_ "github.com/StackExchange/dnscontrol/providers/opensrs"
	_ "github.com/StackExchange/dnscontrol/providers/ovh"
	_ "github.com/StackExchange/dnscontrol/providers/route53"
	_ "github.com/StackExchange/dnscontrol/providers/softlayer"
	_ "github.com/StackExchange/dnscontrol/providers/vultr"
)
