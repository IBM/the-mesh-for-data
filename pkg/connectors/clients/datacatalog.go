// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"io"

	pb "github.com/mesh-for-data/mesh-for-data/pkg/connectors/protobuf"
)

// DataCatalog is an interface of a facade to a data catalog.
type DataCatalog interface {
	pb.DataCatalogServiceServer
	io.Closer
}
