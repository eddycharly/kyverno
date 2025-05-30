/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2 "github.com/kyverno/kyverno/pkg/client/clientset/versioned/typed/kyverno/v2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKyvernoV2 struct {
	*testing.Fake
}

func (c *FakeKyvernoV2) CleanupPolicies(namespace string) v2.CleanupPolicyInterface {
	return newFakeCleanupPolicies(c, namespace)
}

func (c *FakeKyvernoV2) ClusterCleanupPolicies() v2.ClusterCleanupPolicyInterface {
	return newFakeClusterCleanupPolicies(c)
}

func (c *FakeKyvernoV2) PolicyExceptions(namespace string) v2.PolicyExceptionInterface {
	return newFakePolicyExceptions(c, namespace)
}

func (c *FakeKyvernoV2) UpdateRequests(namespace string) v2.UpdateRequestInterface {
	return newFakeUpdateRequests(c, namespace)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKyvernoV2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
