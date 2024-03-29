// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import "encoding/json"

type PrimitiveSpec struct {
	Service   string `json:"service"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Profile   string `json:"profile"`
	Config    []byte `json:"config"`
}

func (s PrimitiveSpec) UnmarshalConfig(config any) error {
	if s.Config == nil {
		return nil
	}
	return json.Unmarshal(s.Config, config)
}
