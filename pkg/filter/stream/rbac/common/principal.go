/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"fmt"
	"net"
	"reflect"
	"strconv"

	"github.com/alipay/sofa-mosn/pkg/log"
	"github.com/alipay/sofa-mosn/pkg/types"
	"github.com/envoyproxy/go-control-plane/envoy/config/rbac/v2alpha"
)

type InheritPrincipal interface {
	isInheritPrincipal()
	// A policy matches if and only if at least one of InheritPermission.Match return true
	// AND at least one of InheritPrincipal.Match return true
	Match(cb types.StreamReceiverFilterCallbacks, headers types.HeaderMap) bool
}

func (*PrincipalAny) isInheritPrincipal()      {}
func (*PrincipalSourceIp) isInheritPrincipal() {}
func (*PrincipalHeader) isInheritPrincipal()   {}
func (*PrincipalAndIds) isInheritPrincipal()   {}
func (*PrincipalOrIds) isInheritPrincipal()    {}

// Principal_Any
type PrincipalAny struct {
	Any bool
}

func NewPrincipalAny(principal *v2alpha.Principal_Any) (*PrincipalAny, error) {
	return &PrincipalAny{
		Any: principal.Any,
	}, nil
}

func (principal *PrincipalAny) Match(cb types.StreamReceiverFilterCallbacks, headers types.HeaderMap) bool {
	return principal.Any
}

// Principal_SourceIp
type PrincipalSourceIp struct {
	CidrRange *net.IPNet
}

func NewPrincipalSourceIp(principal *v2alpha.Principal_SourceIp) (*PrincipalSourceIp, error) {
	addressPrefix := principal.SourceIp.AddressPrefix
	prefixLen := principal.SourceIp.PrefixLen.GetValue()
	if _, ipNet, err := net.ParseCIDR(addressPrefix + "/" + strconv.Itoa(int(prefixLen))); err != nil {
		return nil, err
	} else {
		inheritPrincipal := &PrincipalSourceIp{
			CidrRange: ipNet,
		}
		return inheritPrincipal, nil
	}
}

func (principal *PrincipalSourceIp) Match(cb types.StreamReceiverFilterCallbacks, headers types.HeaderMap) bool {
	remoteAddr := cb.Connection().RemoteAddr().String()
	remoteIP, _, err := parseAddr(remoteAddr)
	if err != nil {
		log.DefaultLogger.Errorf(
			"[PrincipalSourceIp.Match] failed to parse remote address in rbac filter, err: ", err)
		return false
	}
	if principal.CidrRange.Contains(remoteIP) {
		return true
	} else {
		return false
	}
}

// Principal_Header
type PrincipalHeader struct {
	Target      string
	Matcher     HeaderMatcher
	InvertMatch bool
}

func NewPrincipalHeader(principal *v2alpha.Principal_Header) (*PrincipalHeader, error) {
	inheritPrincipal := &PrincipalHeader{}
	inheritPrincipal.Target = principal.Header.Name
	inheritPrincipal.InvertMatch = principal.Header.InvertMatch
	if headerMatcher, err := NewHeaderMatcher(principal.Header); err != nil {
		return nil, err
	} else {
		inheritPrincipal.Matcher = headerMatcher
		return inheritPrincipal, nil
	}
}

func (principal *PrincipalHeader) Match(cb types.StreamReceiverFilterCallbacks, headers types.HeaderMap) bool {
	targetValue, found := headerMapper(principal.Target, headers)

	// HeaderMatcherPresentMatch is a little special
	if matcher, ok := principal.Matcher.(*HeaderMatcherPresentMatch); ok {
		// HeaderMatcherPresentMatch matches if and only if header found and PresentMatch is true
		isMatch := found && matcher.PresentMatch
		return principal.InvertMatch != isMatch
	}

	// return false when targetValue is not found, except matcher is `HeaderMatcherPresentMatch`
	if !found {
		return false
	} else {
		isMatch := principal.Matcher.Equal(targetValue)
		// principal.InvertMatch xor isMatch
		return principal.InvertMatch != isMatch
	}
}

// Principal_AndIds
type PrincipalAndIds struct {
	AndIds []InheritPrincipal
}

func NewPrincipalAndIds(principal *v2alpha.Principal_AndIds) (*PrincipalAndIds, error) {
	inheritPrincipal := &PrincipalAndIds{}
	inheritPrincipal.AndIds = make([]InheritPrincipal, len(principal.AndIds.Ids))
	for idx, subPrincipal := range principal.AndIds.Ids {
		if subInheritPrincipal, err := NewInheritPrincipal(subPrincipal); err != nil {
			return nil, err
		} else {
			inheritPrincipal.AndIds[idx] = subInheritPrincipal
		}
	}
	return inheritPrincipal, nil
}

func (principal *PrincipalAndIds) Match(cb types.StreamReceiverFilterCallbacks, headers types.HeaderMap) bool {
	for _, ids := range principal.AndIds {
		if isMatch := ids.Match(cb, headers); isMatch {
			continue
		} else {
			return false
		}
	}
	return true
}

// Principal_OrIds
type PrincipalOrIds struct {
	OrIds []InheritPrincipal
}

func NewPrincipalOrIds(principal *v2alpha.Principal_OrIds) (*PrincipalOrIds, error) {
	inheritPrincipal := &PrincipalOrIds{}
	inheritPrincipal.OrIds = make([]InheritPrincipal, len(principal.OrIds.Ids))
	for idx, subPrincipal := range principal.OrIds.Ids {
		if subInheritPrincipal, err := NewInheritPrincipal(subPrincipal); err != nil {
			return nil, err
		} else {
			inheritPrincipal.OrIds[idx] = subInheritPrincipal
		}
	}
	return inheritPrincipal, nil
}

func (principal *PrincipalOrIds) Match(cb types.StreamReceiverFilterCallbacks, headers types.HeaderMap) bool {
	for _, ids := range principal.OrIds {
		if isMatch := ids.Match(cb, headers); isMatch {
			return true
		} else {
			continue
		}
	}
	return false
}

// Receive the v2alpha.Principal input and convert it to mosn rbac principal
func NewInheritPrincipal(principal *v2alpha.Principal) (InheritPrincipal, error) {
	// Types that are valid to be assigned to Identifier:
	//	*Principal_AndIds (supported)
	//	*Principal_OrIds (supported)
	//	*Principal_Any (supported)
	//	*Principal_SourceIp (supported)
	//	*Principal_Header (supported)
	// TODO: Principal_Authenticated_ & Principal_Metadata support
	//	*Principal_Authenticated_ (unsupported)
	//	*Principal_Metadata (unsupported)
	switch principal.Identifier.(type) {
	case *v2alpha.Principal_Any:
		return NewPrincipalAny(principal.Identifier.(*v2alpha.Principal_Any))
	case *v2alpha.Principal_SourceIp:
		return NewPrincipalSourceIp(principal.Identifier.(*v2alpha.Principal_SourceIp))
	case *v2alpha.Principal_Header:
		return NewPrincipalHeader(principal.Identifier.(*v2alpha.Principal_Header))
	case *v2alpha.Principal_AndIds:
		return NewPrincipalAndIds(principal.Identifier.(*v2alpha.Principal_AndIds))
	case *v2alpha.Principal_OrIds:
		return NewPrincipalOrIds(principal.Identifier.(*v2alpha.Principal_OrIds))
	default:
		return nil, fmt.Errorf("[NewInheritPrincipal] not supported Principal.Identifier type found, detail: %v",
			reflect.TypeOf(principal.Identifier))
	}
}
