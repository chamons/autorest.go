// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package enum

import (
	"fmt"
)

func EnumStableWithAddition() {
	// This code was written before the "/ka" path was added to enum.json
	var item Get0ItemsItem
	item = Get0ItemsItemEtag
	fmt.Print(item)

	var item2 Head0ItemsItem
	item2 = Head0ItemsItemLocked
	fmt.Print(item2)
}
