/*
Copyright 2024.

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

package v1alpha1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProxySpec defines the desired state of Proxy
type ProxySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Proxies []apiextensionsv1.JSON `json:"proxies"`
}

// ProxyStatus defines the observed state of Proxy
type ProxyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Proxy is the Schema for the proxies API
type Proxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProxySpec   `json:"spec,omitempty"`
	Status ProxyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ProxyList contains a list of Proxy
type ProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Proxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Proxy{}, &ProxyList{})
}

type ProxyBaseConfig struct {
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Transport   ProxyTransport    `json:"transport,omitempty"`
	// metadata info for each proxy
	Metadatas    map[string]string  `json:"metadatas,omitempty"`
	LoadBalancer LoadBalancerConfig `json:"loadBalancer,omitempty"`
	HealthCheck  HealthCheckConfig  `json:"healthCheck,omitempty"`
	ProxyBackend `json:",inline"`
}

type ProxyTransport struct {
	// UseEncryption controls whether or not communication with the server will
	// be encrypted. Encryption is done using the tokens supplied in the server
	// and client configuration.
	UseEncryption bool `json:"useEncryption,omitempty"`
	// UseCompression controls whether or not communication with the server
	// will be compressed.
	UseCompression bool `json:"useCompression,omitempty"`
	// BandwidthLimit limit the bandwidth
	// 0 means no limit
	BandwidthLimit string `json:"bandwidthLimit,omitempty"`
	// BandwidthLimitMode specifies whether to limit the bandwidth on the
	// client or server side. Valid values include "client" and "server".
	// By default, this value is "client".
	BandwidthLimitMode string `json:"bandwidthLimitMode,omitempty"`
	// ProxyProtocolVersion specifies which protocol version to use. Valid
	// values include "v1", "v2", and "". If the value is "", a protocol
	// version will be automatically selected. By default, this value is "".
	ProxyProtocolVersion string `json:"proxyProtocolVersion,omitempty"`
}

type LoadBalancerConfig struct {
	// Group specifies which group the is a part of. The server will use
	// this information to load balance proxies in the same group. If the value
	// is "", this will not be in a group.
	Group string `json:"group"`
	// GroupKey specifies a group key, which should be the same among proxies
	// of the same group.
	GroupKey string `json:"groupKey,omitempty"`
}

// HealthCheckConfig configures health checking. This can be useful for load
// balancing purposes to detect and remove proxies to failing services.
type HealthCheckConfig struct {
	// Type specifies what protocol to use for health checking.
	// Valid values include "tcp", "http", and "". If this value is "", health
	// checking will not be performed.
	//
	// If the type is "tcp", a connection will be attempted to the target
	// server. If a connection cannot be established, the health check fails.
	//
	// If the type is "http", a GET request will be made to the endpoint
	// specified by HealthCheckURL. If the response is not a 200, the health
	// check fails.
	Type string `json:"type"` // tcp | http
	// TimeoutSeconds specifies the number of seconds to wait for a health
	// check attempt to connect. If the timeout is reached, this counts as a
	// health check failure. By default, this value is 3.
	TimeoutSeconds int `json:"timeoutSeconds,omitempty"`
	// MaxFailed specifies the number of allowed failures before the
	// is stopped. By default, this value is 1.
	MaxFailed int `json:"maxFailed,omitempty"`
	// IntervalSeconds specifies the time in seconds between health
	// checks. By default, this value is 10.
	IntervalSeconds int `json:"intervalSeconds"`
	// Path specifies the path to send health checks to if the
	// health check type is "http".
	Path string `json:"path,omitempty"`
	// HTTPHeaders specifies the headers to send with the health request, if
	// the health check type is "http".
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`
}

type HTTPHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProxyBackend struct {
	// LocalIP specifies the IP address or host name of the backend.
	LocalIP string `json:"localIP,omitempty"`
	// LocalPort specifies the port of the backend.
	LocalPort int `json:"localPort,omitempty"`

	// Plugin specifies what plugin should be used for handling connections. If this value
	// is set, the LocalIP and LocalPort values will be ignored.
	Plugin TypedClientPluginOptions `json:"plugin,omitempty"`
}

type TypedClientPluginOptions struct {
	Type string `json:"type"`
	// ClientPluginOptions
}
