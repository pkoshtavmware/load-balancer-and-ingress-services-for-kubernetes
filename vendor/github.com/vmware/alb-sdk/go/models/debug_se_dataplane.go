// Copyright © 2025 Broadcom Inc. and/or its subsidiaries. All Rights Reserved.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// DebugSeDataplane debug se dataplane
// swagger:model DebugSeDataplane
type DebugSeDataplane struct {

	//  Enum options - DEBUG_DISPATCHER_FLOW, DEBUG_DISPATCHER_FLOW_DETAIL, DEBUG_DISPATCHER_FLOW_ALL, DEBUG_CONFIG, DEBUG_IP, DEBUG_IP_PKT_IN, DEBUG_IP_PKT_OUT, DEBUG_ARP, DEBUG_ARP_PKT_IN, DEBUG_ARP_PKT_OUT, DEBUG_ETHERNET, DEBUG_ETHERNET_PKT_IN, DEBUG_ETHERNET_PKT_OUT, DEBUG_ICMP, DEBUG_PCAP_RX, DEBUG_PCAP_TX, DEBUG_PCAP_DROP, DEBUG_PCAP_ALL, DEBUG_MISC, DEBUG_CRUD, DEBUG_POOL, DEBUG_PCAP_DOS, DEBUG_PCAP_HM, DEBUG_SE_APP, DEBUG_UDP, DEBUG_SE_VS_HB, DEBUG_ND, DEBUG_ERROR, DEBUG_NONE, DEBUG_ALL, DEBUG_STRICT, DEBUG_FLOW_MIRROR, DEBUG_NAT, DEBUG_NAT_ERROR, DEBUG_NAT_PKT, DEBUG_NAT_END2END, DEBUG_NAT_ALL, DEBUG_PCAP_SE_IPC, DEBUG_PCAP_NAT, DEBUG_PCAP_ROUTING. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	// Required: true
	Flag *string `json:"flag"`
}
