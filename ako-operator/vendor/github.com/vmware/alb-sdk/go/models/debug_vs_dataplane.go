// Copyright © 2025 Broadcom Inc. and/or its subsidiaries. All Rights Reserved.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// DebugVsDataplane debug vs dataplane
// swagger:model DebugVsDataplane
type DebugVsDataplane struct {

	//  Enum options - DEBUG_VS_TCP_CONNECTION, DEBUG_VS_TCP_PKT, DEBUG_VS_TCP_APP, DEBUG_VS_TCP_APP_PKT, DEBUG_VS_TCP_RETRANSMIT, DEBUG_VS_TCP_TIMER, DEBUG_VS_TCP_CONN_ERROR, DEBUG_VS_TCP_PKT_ERROR, DEBUG_VS_TCP_REXMT, DEBUG_VS_TCP_ALL, DEBUG_VS_CREDIT, DEBUG_VS_PROXY_CONNECTION, DEBUG_VS_PROXY_PKT, DEBUG_VS_PROXY_ERR, DEBUG_VS_UDP, DEBUG_VS_UDP_PKT, DEBUG_VS_HM, DEBUG_VS_HM_ERR, DEBUG_VS_HM_PKT, DEBUG_VS_HTTP_CORE, DEBUG_VS_HTTP_ALL, DEBUG_VS_CONFIG, DEBUG_VS_EVENTS, DEBUG_VS_HTTP_RULES, DEBUG_VS_HM_EXT, DEBUG_VS_SSL, DEBUG_VS_WAF, DEBUG_VS_DNS, DEBUG_VS_PAA, DEBUG_VS_AUTH, DEBUG_VS_BOT, DEBUG_VS_DATASCRIPT, DEBUG_VS_LBACTION, DEBUG_VS_CONTENT_REWRITE, DEBUG_VS_CONNPOOL, DEBUG_VS_WAF_MEMORY, DEBUG_VS_SCTP_PKT, DEBUG_VS_SCTP_PKT_DETAIL, DEBUG_VS_ALL, DEBUG_VS_ERROR, DEBUG_VS_NONE. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	// Required: true
	Flag *string `json:"flag"`
}
