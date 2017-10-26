/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package aggregation

// HPM client data collection configuration
type ClientTypeCollectionConfiguration struct {

	// The client type for which this data collection frequency setting applies
	ClientType string `json:"client_type"`

	// The set of data collection type configurations, one for each data collection type
	DataTypeConfigurations []DataTypeCollectionConfiguration `json:"data_type_configurations"`
}
