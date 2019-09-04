// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gapic-generator. DO NOT EDIT.

package osconfig

import (
	"context"
	"fmt"
	"math"
	"time"

	osconfigpb "github.com/GoogleCloudPlatform/osconfig/_internal/gapi-cloud-osconfig-go/google.golang.org/genproto/googleapis/cloud/osconfig/v1alpha2"
	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ExecutePatchJob               []gax.CallOption
	GetPatchJob                   []gax.CallOption
	CancelPatchJob                []gax.CallOption
	ListPatchJobs                 []gax.CallOption
	ReportPatchJobInstanceDetails []gax.CallOption
	ListPatchJobInstanceDetails   []gax.CallOption
	CreateGuestPolicy             []gax.CallOption
	GetGuestPolicy                []gax.CallOption
	ListGuestPolicies             []gax.CallOption
	UpdateGuestPolicy             []gax.CallOption
	DeleteGuestPolicy             []gax.CallOption
	LookupEffectiveGuestPolicies  []gax.CallOption
}

func defaultClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("osconfig.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultCallOptions() *CallOptions {
	retry := map[[2]string][]gax.CallOption{
		{"default", "idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
	}
	return &CallOptions{
		ExecutePatchJob:               retry[[2]string{"default", "non_idempotent"}],
		GetPatchJob:                   retry[[2]string{"default", "idempotent"}],
		CancelPatchJob:                retry[[2]string{"default", "non_idempotent"}],
		ListPatchJobs:                 retry[[2]string{"default", "idempotent"}],
		ReportPatchJobInstanceDetails: retry[[2]string{"default", "non_idempotent"}],
		ListPatchJobInstanceDetails:   retry[[2]string{"default", "idempotent"}],
		CreateGuestPolicy:             retry[[2]string{"default", "non_idempotent"}],
		GetGuestPolicy:                retry[[2]string{"default", "idempotent"}],
		ListGuestPolicies:             retry[[2]string{"default", "idempotent"}],
		UpdateGuestPolicy:             retry[[2]string{"default", "non_idempotent"}],
		DeleteGuestPolicy:             retry[[2]string{"default", "idempotent"}],
		LookupEffectiveGuestPolicies:  retry[[2]string{"default", "non_idempotent"}],
	}
}

// Client is a client for interacting with Cloud OS Config API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type Client struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	client osconfigpb.OsConfigServiceClient

	// The call options for this service.
	CallOptions *CallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new os config service client.
//
// OS Config API
//
// The OS Config service is the server-side component that allows users to
// manage package installations and patch jobs for virtual machines.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:        conn,
		CallOptions: defaultCallOptions(),

		client: osconfigpb.NewOsConfigServiceClient(conn),
	}
	c.setGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *Client) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ExecutePatchJob patch GCE instances by creating and running a PatchJob.
func (c *Client) ExecutePatchJob(ctx context.Context, req *osconfigpb.ExecutePatchJobRequest, opts ...gax.CallOption) (*osconfigpb.PatchJob, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ExecutePatchJob[0:len(c.CallOptions.ExecutePatchJob):len(c.CallOptions.ExecutePatchJob)], opts...)
	var resp *osconfigpb.PatchJob
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ExecutePatchJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetPatchJob get the patch job. This can be used to track the progress of an
// ongoing patch job or review the details of completed jobs.
func (c *Client) GetPatchJob(ctx context.Context, req *osconfigpb.GetPatchJobRequest, opts ...gax.CallOption) (*osconfigpb.PatchJob, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetPatchJob[0:len(c.CallOptions.GetPatchJob):len(c.CallOptions.GetPatchJob)], opts...)
	var resp *osconfigpb.PatchJob
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetPatchJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CancelPatchJob cancel a patch job. The patch job must be active. Canceled patch jobs
// cannot be restarted.
func (c *Client) CancelPatchJob(ctx context.Context, req *osconfigpb.CancelPatchJobRequest, opts ...gax.CallOption) (*osconfigpb.PatchJob, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CancelPatchJob[0:len(c.CallOptions.CancelPatchJob):len(c.CallOptions.CancelPatchJob)], opts...)
	var resp *osconfigpb.PatchJob
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CancelPatchJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListPatchJobs get a page of patch jobs.
func (c *Client) ListPatchJobs(ctx context.Context, req *osconfigpb.ListPatchJobsRequest, opts ...gax.CallOption) *PatchJobIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListPatchJobs[0:len(c.CallOptions.ListPatchJobs):len(c.CallOptions.ListPatchJobs)], opts...)
	it := &PatchJobIterator{}
	req = proto.Clone(req).(*osconfigpb.ListPatchJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*osconfigpb.PatchJob, string, error) {
		var resp *osconfigpb.ListPatchJobsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListPatchJobs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.PatchJobs, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// ReportPatchJobInstanceDetails deprecated:  Should use AgentEndpoint API going forward.  Will be removed
// in v1beta1.
//
// Endpoint used by the agent to report back its state during a patch
// job. This endpoint will also return the patch job's state and
// configurations that the agent needs to know in order to run or stop
// patching.
//
// This endpoint is only used by the agent. Using it in other ways may
// affect the state of the active patch job and prevent the patches from
// being correctly applied to this instance.
func (c *Client) ReportPatchJobInstanceDetails(ctx context.Context, req *osconfigpb.ReportPatchJobInstanceDetailsRequest, opts ...gax.CallOption) (*osconfigpb.ReportPatchJobInstanceDetailsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", req.GetResource()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ReportPatchJobInstanceDetails[0:len(c.CallOptions.ReportPatchJobInstanceDetails):len(c.CallOptions.ReportPatchJobInstanceDetails)], opts...)
	var resp *osconfigpb.ReportPatchJobInstanceDetailsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReportPatchJobInstanceDetails(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListPatchJobInstanceDetails get a page of instances' details for a given patch job.
func (c *Client) ListPatchJobInstanceDetails(ctx context.Context, req *osconfigpb.ListPatchJobInstanceDetailsRequest, opts ...gax.CallOption) *PatchJobInstanceDetailsIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListPatchJobInstanceDetails[0:len(c.CallOptions.ListPatchJobInstanceDetails):len(c.CallOptions.ListPatchJobInstanceDetails)], opts...)
	it := &PatchJobInstanceDetailsIterator{}
	req = proto.Clone(req).(*osconfigpb.ListPatchJobInstanceDetailsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*osconfigpb.PatchJobInstanceDetails, string, error) {
		var resp *osconfigpb.ListPatchJobInstanceDetailsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListPatchJobInstanceDetails(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.PatchJobInstanceDetails, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// CreateGuestPolicy create an OS Config Guest Policy.
func (c *Client) CreateGuestPolicy(ctx context.Context, req *osconfigpb.CreateGuestPolicyRequest, opts ...gax.CallOption) (*osconfigpb.GuestPolicy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateGuestPolicy[0:len(c.CallOptions.CreateGuestPolicy):len(c.CallOptions.CreateGuestPolicy)], opts...)
	var resp *osconfigpb.GuestPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateGuestPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetGuestPolicy get an OS Config GuestPolicy.
func (c *Client) GetGuestPolicy(ctx context.Context, req *osconfigpb.GetGuestPolicyRequest, opts ...gax.CallOption) (*osconfigpb.GuestPolicy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetGuestPolicy[0:len(c.CallOptions.GetGuestPolicy):len(c.CallOptions.GetGuestPolicy)], opts...)
	var resp *osconfigpb.GuestPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetGuestPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListGuestPolicies get a page of OS Config GuestPolicies.
func (c *Client) ListGuestPolicies(ctx context.Context, req *osconfigpb.ListGuestPoliciesRequest, opts ...gax.CallOption) *GuestPolicyIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListGuestPolicies[0:len(c.CallOptions.ListGuestPolicies):len(c.CallOptions.ListGuestPolicies)], opts...)
	it := &GuestPolicyIterator{}
	req = proto.Clone(req).(*osconfigpb.ListGuestPoliciesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*osconfigpb.GuestPolicy, string, error) {
		var resp *osconfigpb.ListGuestPoliciesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListGuestPolicies(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.GuestPolicies, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// UpdateGuestPolicy update an OS Config GuestPolicy.
func (c *Client) UpdateGuestPolicy(ctx context.Context, req *osconfigpb.UpdateGuestPolicyRequest, opts ...gax.CallOption) (*osconfigpb.GuestPolicy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "guest_policy.name", req.GetGuestPolicy().GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateGuestPolicy[0:len(c.CallOptions.UpdateGuestPolicy):len(c.CallOptions.UpdateGuestPolicy)], opts...)
	var resp *osconfigpb.GuestPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateGuestPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteGuestPolicy delete an OS Config GuestPolicy.
func (c *Client) DeleteGuestPolicy(ctx context.Context, req *osconfigpb.DeleteGuestPolicyRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteGuestPolicy[0:len(c.CallOptions.DeleteGuestPolicy):len(c.CallOptions.DeleteGuestPolicy)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteGuestPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// LookupEffectiveGuestPolicies lookup the guest policies that are assigned to a GCE VM instance. This
// lookup will merge all policies that are assigned to the instance. This is
// usually called by the agent running on the instance, but it can also be
// called by users to see what configs are assigned to this instance.
func (c *Client) LookupEffectiveGuestPolicies(ctx context.Context, req *osconfigpb.LookupEffectiveGuestPoliciesRequest, opts ...gax.CallOption) (*osconfigpb.LookupEffectiveGuestPoliciesResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "instance", req.GetInstance()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.LookupEffectiveGuestPolicies[0:len(c.CallOptions.LookupEffectiveGuestPolicies):len(c.CallOptions.LookupEffectiveGuestPolicies)], opts...)
	var resp *osconfigpb.LookupEffectiveGuestPoliciesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.LookupEffectiveGuestPolicies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GuestPolicyIterator manages a stream of *osconfigpb.GuestPolicy.
type GuestPolicyIterator struct {
	items    []*osconfigpb.GuestPolicy
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*osconfigpb.GuestPolicy, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *GuestPolicyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *GuestPolicyIterator) Next() (*osconfigpb.GuestPolicy, error) {
	var item *osconfigpb.GuestPolicy
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *GuestPolicyIterator) bufLen() int {
	return len(it.items)
}

func (it *GuestPolicyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// PatchJobInstanceDetailsIterator manages a stream of *osconfigpb.PatchJobInstanceDetails.
type PatchJobInstanceDetailsIterator struct {
	items    []*osconfigpb.PatchJobInstanceDetails
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*osconfigpb.PatchJobInstanceDetails, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PatchJobInstanceDetailsIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PatchJobInstanceDetailsIterator) Next() (*osconfigpb.PatchJobInstanceDetails, error) {
	var item *osconfigpb.PatchJobInstanceDetails
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PatchJobInstanceDetailsIterator) bufLen() int {
	return len(it.items)
}

func (it *PatchJobInstanceDetailsIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// PatchJobIterator manages a stream of *osconfigpb.PatchJob.
type PatchJobIterator struct {
	items    []*osconfigpb.PatchJob
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*osconfigpb.PatchJob, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PatchJobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PatchJobIterator) Next() (*osconfigpb.PatchJob, error) {
	var item *osconfigpb.PatchJob
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PatchJobIterator) bufLen() int {
	return len(it.items)
}

func (it *PatchJobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
