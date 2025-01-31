/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type FixedAddressAPI interface {

	/*
			FixedAddressCreate Create the fixed address.

			Use this method to create a __FixedAddress__ object.
		The __FixedAddress__ object reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so that it can match that client.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiFixedAddressCreateRequest
	*/
	FixedAddressCreate(ctx context.Context) ApiFixedAddressCreateRequest

	// FixedAddressCreateExecute executes the request
	//  @return IpamsvcCreateFixedAddressResponse
	FixedAddressCreateExecute(r ApiFixedAddressCreateRequest) (*IpamsvcCreateFixedAddressResponse, *http.Response, error)

	/*
			FixedAddressDelete Move the fixed address to the recycle bin.

			Use this method to move a __FixedAddress__ object to the recycle bin.
		The __FixedAddress__ object reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so that it can match that client.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiFixedAddressDeleteRequest
	*/
	FixedAddressDelete(ctx context.Context, id string) ApiFixedAddressDeleteRequest

	// FixedAddressDeleteExecute executes the request
	FixedAddressDeleteExecute(r ApiFixedAddressDeleteRequest) (*http.Response, error)

	/*
			FixedAddressList Retrieve fixed addresses.

			Use this method to retrieve __FixedAddress__ objects.
		The __FixedAddress__ object reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so that it can match that client.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiFixedAddressListRequest
	*/
	FixedAddressList(ctx context.Context) ApiFixedAddressListRequest

	// FixedAddressListExecute executes the request
	//  @return IpamsvcListFixedAddressResponse
	FixedAddressListExecute(r ApiFixedAddressListRequest) (*IpamsvcListFixedAddressResponse, *http.Response, error)

	/*
			FixedAddressRead Retrieve the fixed address.

			Use this method to retrieve a __FixedAddress__ object.
		The __FixedAddress__ object reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so that it can match that client.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiFixedAddressReadRequest
	*/
	FixedAddressRead(ctx context.Context, id string) ApiFixedAddressReadRequest

	// FixedAddressReadExecute executes the request
	//  @return IpamsvcReadFixedAddressResponse
	FixedAddressReadExecute(r ApiFixedAddressReadRequest) (*IpamsvcReadFixedAddressResponse, *http.Response, error)

	/*
			FixedAddressUpdate Update the fixed address.

			Use this method to update a __FixedAddress__ object.
		The __FixedAddress__ object reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so that it can match that client.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiFixedAddressUpdateRequest
	*/
	FixedAddressUpdate(ctx context.Context, id string) ApiFixedAddressUpdateRequest

	// FixedAddressUpdateExecute executes the request
	//  @return IpamsvcUpdateFixedAddressResponse
	FixedAddressUpdateExecute(r ApiFixedAddressUpdateRequest) (*IpamsvcUpdateFixedAddressResponse, *http.Response, error)
}

// FixedAddressAPIService FixedAddressAPI service
type FixedAddressAPIService internal.Service

type ApiFixedAddressCreateRequest struct {
	ctx        context.Context
	ApiService FixedAddressAPI
	body       *IpamsvcFixedAddress
	inherit    *string
}

func (r ApiFixedAddressCreateRequest) Body(body IpamsvcFixedAddress) ApiFixedAddressCreateRequest {
	r.body = &body
	return r
}

// This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none
func (r ApiFixedAddressCreateRequest) Inherit(inherit string) ApiFixedAddressCreateRequest {
	r.inherit = &inherit
	return r
}

func (r ApiFixedAddressCreateRequest) Execute() (*IpamsvcCreateFixedAddressResponse, *http.Response, error) {
	return r.ApiService.FixedAddressCreateExecute(r)
}

/*
FixedAddressCreate Create the fixed address.

Use this method to create a __FixedAddress__ object.
The __FixedAddress__ object reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so that it can match that client.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFixedAddressCreateRequest
*/
func (a *FixedAddressAPIService) FixedAddressCreate(ctx context.Context) ApiFixedAddressCreateRequest {
	return ApiFixedAddressCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IpamsvcCreateFixedAddressResponse
func (a *FixedAddressAPIService) FixedAddressCreateExecute(r ApiFixedAddressCreateRequest) (*IpamsvcCreateFixedAddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcCreateFixedAddressResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "FixedAddressAPIService.FixedAddressCreate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/fixed_address"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	if r.inherit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inherit", r.inherit, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFixedAddressDeleteRequest struct {
	ctx        context.Context
	ApiService FixedAddressAPI
	id         string
}

func (r ApiFixedAddressDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.FixedAddressDeleteExecute(r)
}

/*
FixedAddressDelete Move the fixed address to the recycle bin.

Use this method to move a __FixedAddress__ object to the recycle bin.
The __FixedAddress__ object reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so that it can match that client.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiFixedAddressDeleteRequest
*/
func (a *FixedAddressAPIService) FixedAddressDelete(ctx context.Context, id string) ApiFixedAddressDeleteRequest {
	return ApiFixedAddressDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *FixedAddressAPIService) FixedAddressDeleteExecute(r ApiFixedAddressDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "FixedAddressAPIService.FixedAddressDelete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/fixed_address/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFixedAddressListRequest struct {
	ctx        context.Context
	ApiService FixedAddressAPI
	filter     *string
	orderBy    *string
	fields     *string
	offset     *int32
	limit      *int32
	pageToken  *string
	torderBy   *string
	tfilter    *string
	inherit    *string
}

// A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |
func (r ApiFixedAddressListRequest) Filter(filter string) ApiFixedAddressListRequest {
	r.filter = &filter
	return r
}

// A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.
func (r ApiFixedAddressListRequest) OrderBy(orderBy string) ApiFixedAddressListRequest {
	r.orderBy = &orderBy
	return r
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiFixedAddressListRequest) Fields(fields string) ApiFixedAddressListRequest {
	r.fields = &fields
	return r
}

// The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r ApiFixedAddressListRequest) Offset(offset int32) ApiFixedAddressListRequest {
	r.offset = &offset
	return r
}

// The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r ApiFixedAddressListRequest) Limit(limit int32) ApiFixedAddressListRequest {
	r.limit = &limit
	return r
}

// The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r ApiFixedAddressListRequest) PageToken(pageToken string) ApiFixedAddressListRequest {
	r.pageToken = &pageToken
	return r
}

// This parameter is used for sorting by tags.
func (r ApiFixedAddressListRequest) TorderBy(torderBy string) ApiFixedAddressListRequest {
	r.torderBy = &torderBy
	return r
}

// This parameter is used for filtering by tags.
func (r ApiFixedAddressListRequest) Tfilter(tfilter string) ApiFixedAddressListRequest {
	r.tfilter = &tfilter
	return r
}

// This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none
func (r ApiFixedAddressListRequest) Inherit(inherit string) ApiFixedAddressListRequest {
	r.inherit = &inherit
	return r
}

func (r ApiFixedAddressListRequest) Execute() (*IpamsvcListFixedAddressResponse, *http.Response, error) {
	return r.ApiService.FixedAddressListExecute(r)
}

/*
FixedAddressList Retrieve fixed addresses.

Use this method to retrieve __FixedAddress__ objects.
The __FixedAddress__ object reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so that it can match that client.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFixedAddressListRequest
*/
func (a *FixedAddressAPIService) FixedAddressList(ctx context.Context) ApiFixedAddressListRequest {
	return ApiFixedAddressListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IpamsvcListFixedAddressResponse
func (a *FixedAddressAPIService) FixedAddressListExecute(r ApiFixedAddressListRequest) (*IpamsvcListFixedAddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcListFixedAddressResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "FixedAddressAPIService.FixedAddressList")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/fixed_address"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_filter", r.filter, "")
	}
	if r.orderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_order_by", r.orderBy, "")
	}
	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.offset != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_offset", r.offset, "")
	}
	if r.limit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_limit", r.limit, "")
	}
	if r.pageToken != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_token", r.pageToken, "")
	}
	if r.torderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_torder_by", r.torderBy, "")
	}
	if r.tfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_tfilter", r.tfilter, "")
	}
	if r.inherit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inherit", r.inherit, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFixedAddressReadRequest struct {
	ctx        context.Context
	ApiService FixedAddressAPI
	id         string
	fields     *string
	inherit    *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiFixedAddressReadRequest) Fields(fields string) ApiFixedAddressReadRequest {
	r.fields = &fields
	return r
}

// This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none
func (r ApiFixedAddressReadRequest) Inherit(inherit string) ApiFixedAddressReadRequest {
	r.inherit = &inherit
	return r
}

func (r ApiFixedAddressReadRequest) Execute() (*IpamsvcReadFixedAddressResponse, *http.Response, error) {
	return r.ApiService.FixedAddressReadExecute(r)
}

/*
FixedAddressRead Retrieve the fixed address.

Use this method to retrieve a __FixedAddress__ object.
The __FixedAddress__ object reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so that it can match that client.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiFixedAddressReadRequest
*/
func (a *FixedAddressAPIService) FixedAddressRead(ctx context.Context, id string) ApiFixedAddressReadRequest {
	return ApiFixedAddressReadRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IpamsvcReadFixedAddressResponse
func (a *FixedAddressAPIService) FixedAddressReadExecute(r ApiFixedAddressReadRequest) (*IpamsvcReadFixedAddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcReadFixedAddressResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "FixedAddressAPIService.FixedAddressRead")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/fixed_address/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.inherit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inherit", r.inherit, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFixedAddressUpdateRequest struct {
	ctx        context.Context
	ApiService FixedAddressAPI
	id         string
	body       *IpamsvcFixedAddress
	inherit    *string
}

func (r ApiFixedAddressUpdateRequest) Body(body IpamsvcFixedAddress) ApiFixedAddressUpdateRequest {
	r.body = &body
	return r
}

// This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none
func (r ApiFixedAddressUpdateRequest) Inherit(inherit string) ApiFixedAddressUpdateRequest {
	r.inherit = &inherit
	return r
}

func (r ApiFixedAddressUpdateRequest) Execute() (*IpamsvcUpdateFixedAddressResponse, *http.Response, error) {
	return r.ApiService.FixedAddressUpdateExecute(r)
}

/*
FixedAddressUpdate Update the fixed address.

Use this method to update a __FixedAddress__ object.
The __FixedAddress__ object reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so that it can match that client.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiFixedAddressUpdateRequest
*/
func (a *FixedAddressAPIService) FixedAddressUpdate(ctx context.Context, id string) ApiFixedAddressUpdateRequest {
	return ApiFixedAddressUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IpamsvcUpdateFixedAddressResponse
func (a *FixedAddressAPIService) FixedAddressUpdateExecute(r ApiFixedAddressUpdateRequest) (*IpamsvcUpdateFixedAddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcUpdateFixedAddressResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "FixedAddressAPIService.FixedAddressUpdate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/fixed_address/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	if r.inherit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inherit", r.inherit, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
