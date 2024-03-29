// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package util

import "time"

func GetCurrentTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format(time.RFC3339)
}
