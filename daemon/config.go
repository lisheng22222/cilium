// Copyright 2016-2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net"
	"os"
	"runtime"
	"sync"

	"github.com/cilium/cilium/api/v1/models"
	"github.com/cilium/cilium/daemon/options"
	"github.com/cilium/cilium/pkg/labels"
	"github.com/cilium/cilium/pkg/option"
)

const (
	// AllowLocalhostAuto defaults to policy except when running in
	// Kubernetes where it then defaults to "always"
	AllowLocalhostAuto = "auto"

	// AllowLocalhostAlways always allows the local stack to reach local
	// endpoints
	AllowLocalhostAlways = "always"

	// AllowLocalhostPolicy requires a policy rule to allow the local stack
	// to reach particular endpoints or policy enforcement must be
	// disabled.
	AllowLocalhostPolicy = "policy"
)

// Config is the configuration used by Daemon.
type Config struct {
	BpfDir         string     // BPF template files directory
	LibDir         string     // Cilium library files directory
	RunDir         string     // Cilium runtime directory
	NAT46Prefix    *net.IPNet // NAT46 IPv6 Prefix
	Device         string     // Receive device
	HostV4Addr     net.IP     // Host v4 address of the snooping device
	HostV6Addr     net.IP     // Host v6 address of the snooping device
	DockerEndpoint string     // Docker endpoint
	IPv4Disabled   bool       // Disable IPv4 allocation
	K8sEndpoint    string     // Kubernetes endpoint
	K8sCfgPath     string     // Kubeconfig path
	LBInterface    string     // Set with name of the interface to loadbalance packets from

	EnablePolicyMU sync.RWMutex // Protects the variable below
	EnablePolicy   string       // Whether policy enforcement is enabled.

	Tunnel string // Tunnel mode

	ValidLabelPrefixesMU sync.RWMutex           // Protects the 2 variables below
	ValidLabelPrefixes   *labels.LabelPrefixCfg // Label prefixes used to filter from all labels

	DryMode       bool // Do not create BPF maps, devices, ..
	RestoreState  bool // RestoreState restores the state from previous running daemons.
	KeepConfig    bool // Keep configuration of existing endpoints when starting up.
	KeepTemplates bool // Do not overwrite the template files

	// AllowLocalhost defines when to allows the local stack to local endpoints
	// values: { auto | always | policy }
	AllowLocalhost string

	// alwaysAllowLocalhost is set based on the value of AllowLocalhost and
	// is either set to true when localhost can always reach local
	// endpoints or false when policy should be evaluated
	alwaysAllowLocalhost bool

	// StateDir is the directory where runtime state of endpoints is stored
	StateDir string

	// Options changeable at runtime
	Opts *option.BoolOptions

	// Mutex for serializing configuration updates to the daemon.
	ConfigPatchMutex sync.RWMutex

	// Monitor contains the configuration for the node monitor.
	Monitor *models.MonitorStatus
}

func NewConfig() *Config {
	return &Config{
		Opts:    option.NewBoolOptions(&options.Library),
		Monitor: &models.MonitorStatus{Cpus: int64(runtime.NumCPU()), Npages: 64, Pagesize: int64(os.Getpagesize()), Lost: 0, Unknown: 0},
	}
}

// IsK8sEnabled checks if Cilium is being used in tandem with Kubernetes.
func (c *Config) IsK8sEnabled() bool {
	return c.K8sEndpoint != "" || c.K8sCfgPath != ""
}

func (c *Config) IsLBEnabled() bool {
	return c.LBInterface != ""
}
