// Copyright (C) 2015  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package hrpc

import (
	"bytes"
	"context"

	"github.com/tsuna/gohbase/pb"
	"google.golang.org/protobuf/proto"
)

// DisableTable represents a DisableTable HBase call
type DisableTable struct {
	base
}

// NewDisableTable creates a new DisableTable request that will disable the
// given table in HBase. For use by the admin client.
func NewDisableTable(ctx context.Context, table []byte) *DisableTable {
	return &DisableTable{
		base{
			table:    table,
			ctx:      ctx,
			resultch: make(chan RPCResult, 1),
		},
	}
}

// Name returns the name of this RPC call.
func (dt *DisableTable) Name() string {
	return "DisableTable"
}

// Description returns the description of this RPC call.
func (dt *DisableTable) Description() string {
	return dt.Name()
}

// ToProto converts the RPC into a protobuf message
func (dt *DisableTable) ToProto() proto.Message {
	namespace := []byte("default")
	table := dt.table
	if i := bytes.Index(table, []byte(":")); i > -1 {
		namespace = table[:i]
		table = table[i+1:]
	}
	return &pb.DisableTableRequest{
		TableName: &pb.TableName{
			Namespace: namespace,
			Qualifier: table,
		},
	}
}

// NewResponse creates an empty protobuf message to read the response of this
// RPC.
func (dt *DisableTable) NewResponse() proto.Message {
	return &pb.DisableTableResponse{}
}
