// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateTagRequest wrapper for the UpdateTag operation
type UpdateTagRequest struct {

	// The OCID of the tag namespace.
	TagNamespaceId *string `mandatory:"true" contributesTo:"path" name:"tagNamespaceId"`

	// The name of the tag.
	TagName *string `mandatory:"true" contributesTo:"path" name:"tagName"`

	// Request object for updating a tag.
	UpdateTagDetails `contributesTo:"body"`
}

func (request UpdateTagRequest) String() string {
	return common.PointerString(request)
}

// UpdateTagResponse wrapper for the UpdateTag operation
type UpdateTagResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Tag instance
	Tag `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateTagResponse) String() string {
	return common.PointerString(response)
}
