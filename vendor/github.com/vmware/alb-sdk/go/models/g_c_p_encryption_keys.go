// Copyright © 2025 Broadcom Inc. and/or its subsidiaries. All Rights Reserved.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// GCPEncryptionKeys g c p encryption keys
// swagger:model GCPEncryptionKeys
type GCPEncryptionKeys struct {

	// CMEK Resource ID to encrypt Google Cloud Storage Bucket. This Bucket is used to upload Service Engine raw image. Field introduced in 18.2.10, 20.1.2. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	GcsBucketKmsKeyID *string `json:"gcs_bucket_kms_key_id,omitempty"`

	// CMEK Resource ID to encrypt Service Engine raw image. The raw image is a Google Cloud Storage Object. Field introduced in 18.2.10, 20.1.2. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	GcsObjectsKmsKeyID *string `json:"gcs_objects_kms_key_id,omitempty"`

	// CMEK Resource ID to encrypt Service Engine Disks. Field introduced in 18.2.10, 20.1.2. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	SeDiskKmsKeyID *string `json:"se_disk_kms_key_id,omitempty"`

	// CMEK Resource ID to encrypt Service Engine GCE Image. Field introduced in 18.2.10, 20.1.2. Allowed in Enterprise edition with any value, Enterprise with Cloud Services edition.
	SeImageKmsKeyID *string `json:"se_image_kms_key_id,omitempty"`
}
