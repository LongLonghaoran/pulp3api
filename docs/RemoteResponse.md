# RemoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Name** | **string** | A unique name for this remote. | 
**Url** | **string** | The URL of an external content source. | 
**CaCert** | Pointer to **NullableString** | A PEM encoded CA certificate used to validate the server certificate presented by the remote server. | [optional] 
**ClientCert** | Pointer to **NullableString** | A PEM encoded client certificate used for authentication. | [optional] 
**TlsValidation** | Pointer to **bool** | If True, TLS peer validation must be performed. | [optional] 
**ProxyUrl** | Pointer to **NullableString** | The proxy URL. Format: scheme://host:port | [optional] 
**PulpLabels** | Pointer to **map[string]interface{}** |  | [optional] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the most recent update of the remote. | [optional] [readonly] 
**DownloadConcurrency** | Pointer to **NullableInt32** | Total number of simultaneous connections. If not set then the default value will be used. | [optional] 
**MaxRetries** | Pointer to **NullableInt32** | Maximum number of retry attempts after a download failure. If not set then the default value (3) will be used. | [optional] 
**Policy** | Pointer to [**NullablePolicyDb6Enum**](PolicyDb6Enum.md) | The policy to use when downloading content. | [optional] 
**TotalTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**ConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockReadTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**Headers** | Pointer to **[]map[string]interface{}** | Headers for aiohttp.Clientsession | [optional] 
**RateLimit** | Pointer to **NullableInt32** | Limits requests per second for each concurrent downloader | [optional] 

## Methods

### NewRemoteResponse

`func NewRemoteResponse(name string, url string, ) *RemoteResponse`

NewRemoteResponse instantiates a new RemoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteResponseWithDefaults

`func NewRemoteResponseWithDefaults() *RemoteResponse`

NewRemoteResponseWithDefaults instantiates a new RemoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RemoteResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RemoteResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RemoteResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RemoteResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RemoteResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RemoteResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RemoteResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RemoteResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *RemoteResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *RemoteResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RemoteResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RemoteResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCaCert

`func (o *RemoteResponse) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *RemoteResponse) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *RemoteResponse) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *RemoteResponse) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *RemoteResponse) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *RemoteResponse) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetClientCert

`func (o *RemoteResponse) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *RemoteResponse) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *RemoteResponse) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *RemoteResponse) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### SetClientCertNil

`func (o *RemoteResponse) SetClientCertNil(b bool)`

 SetClientCertNil sets the value for ClientCert to be an explicit nil

### UnsetClientCert
`func (o *RemoteResponse) UnsetClientCert()`

UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
### GetTlsValidation

`func (o *RemoteResponse) GetTlsValidation() bool`

GetTlsValidation returns the TlsValidation field if non-nil, zero value otherwise.

### GetTlsValidationOk

`func (o *RemoteResponse) GetTlsValidationOk() (*bool, bool)`

GetTlsValidationOk returns a tuple with the TlsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsValidation

`func (o *RemoteResponse) SetTlsValidation(v bool)`

SetTlsValidation sets TlsValidation field to given value.

### HasTlsValidation

`func (o *RemoteResponse) HasTlsValidation() bool`

HasTlsValidation returns a boolean if a field has been set.

### GetProxyUrl

`func (o *RemoteResponse) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *RemoteResponse) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *RemoteResponse) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *RemoteResponse) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### SetProxyUrlNil

`func (o *RemoteResponse) SetProxyUrlNil(b bool)`

 SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil

### UnsetProxyUrl
`func (o *RemoteResponse) UnsetProxyUrl()`

UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
### GetPulpLabels

`func (o *RemoteResponse) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *RemoteResponse) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *RemoteResponse) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *RemoteResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *RemoteResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *RemoteResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *RemoteResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *RemoteResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetDownloadConcurrency

`func (o *RemoteResponse) GetDownloadConcurrency() int32`

GetDownloadConcurrency returns the DownloadConcurrency field if non-nil, zero value otherwise.

### GetDownloadConcurrencyOk

`func (o *RemoteResponse) GetDownloadConcurrencyOk() (*int32, bool)`

GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadConcurrency

`func (o *RemoteResponse) SetDownloadConcurrency(v int32)`

SetDownloadConcurrency sets DownloadConcurrency field to given value.

### HasDownloadConcurrency

`func (o *RemoteResponse) HasDownloadConcurrency() bool`

HasDownloadConcurrency returns a boolean if a field has been set.

### SetDownloadConcurrencyNil

`func (o *RemoteResponse) SetDownloadConcurrencyNil(b bool)`

 SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil

### UnsetDownloadConcurrency
`func (o *RemoteResponse) UnsetDownloadConcurrency()`

UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
### GetMaxRetries

`func (o *RemoteResponse) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *RemoteResponse) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *RemoteResponse) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *RemoteResponse) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *RemoteResponse) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *RemoteResponse) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetPolicy

`func (o *RemoteResponse) GetPolicy() PolicyDb6Enum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *RemoteResponse) GetPolicyOk() (*PolicyDb6Enum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *RemoteResponse) SetPolicy(v PolicyDb6Enum)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *RemoteResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *RemoteResponse) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *RemoteResponse) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetTotalTimeout

`func (o *RemoteResponse) GetTotalTimeout() float64`

GetTotalTimeout returns the TotalTimeout field if non-nil, zero value otherwise.

### GetTotalTimeoutOk

`func (o *RemoteResponse) GetTotalTimeoutOk() (*float64, bool)`

GetTotalTimeoutOk returns a tuple with the TotalTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeout

`func (o *RemoteResponse) SetTotalTimeout(v float64)`

SetTotalTimeout sets TotalTimeout field to given value.

### HasTotalTimeout

`func (o *RemoteResponse) HasTotalTimeout() bool`

HasTotalTimeout returns a boolean if a field has been set.

### SetTotalTimeoutNil

`func (o *RemoteResponse) SetTotalTimeoutNil(b bool)`

 SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil

### UnsetTotalTimeout
`func (o *RemoteResponse) UnsetTotalTimeout()`

UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
### GetConnectTimeout

`func (o *RemoteResponse) GetConnectTimeout() float64`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *RemoteResponse) GetConnectTimeoutOk() (*float64, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *RemoteResponse) SetConnectTimeout(v float64)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *RemoteResponse) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### SetConnectTimeoutNil

`func (o *RemoteResponse) SetConnectTimeoutNil(b bool)`

 SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil

### UnsetConnectTimeout
`func (o *RemoteResponse) UnsetConnectTimeout()`

UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
### GetSockConnectTimeout

`func (o *RemoteResponse) GetSockConnectTimeout() float64`

GetSockConnectTimeout returns the SockConnectTimeout field if non-nil, zero value otherwise.

### GetSockConnectTimeoutOk

`func (o *RemoteResponse) GetSockConnectTimeoutOk() (*float64, bool)`

GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockConnectTimeout

`func (o *RemoteResponse) SetSockConnectTimeout(v float64)`

SetSockConnectTimeout sets SockConnectTimeout field to given value.

### HasSockConnectTimeout

`func (o *RemoteResponse) HasSockConnectTimeout() bool`

HasSockConnectTimeout returns a boolean if a field has been set.

### SetSockConnectTimeoutNil

`func (o *RemoteResponse) SetSockConnectTimeoutNil(b bool)`

 SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil

### UnsetSockConnectTimeout
`func (o *RemoteResponse) UnsetSockConnectTimeout()`

UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
### GetSockReadTimeout

`func (o *RemoteResponse) GetSockReadTimeout() float64`

GetSockReadTimeout returns the SockReadTimeout field if non-nil, zero value otherwise.

### GetSockReadTimeoutOk

`func (o *RemoteResponse) GetSockReadTimeoutOk() (*float64, bool)`

GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockReadTimeout

`func (o *RemoteResponse) SetSockReadTimeout(v float64)`

SetSockReadTimeout sets SockReadTimeout field to given value.

### HasSockReadTimeout

`func (o *RemoteResponse) HasSockReadTimeout() bool`

HasSockReadTimeout returns a boolean if a field has been set.

### SetSockReadTimeoutNil

`func (o *RemoteResponse) SetSockReadTimeoutNil(b bool)`

 SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil

### UnsetSockReadTimeout
`func (o *RemoteResponse) UnsetSockReadTimeout()`

UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
### GetHeaders

`func (o *RemoteResponse) GetHeaders() []map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *RemoteResponse) GetHeadersOk() (*[]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *RemoteResponse) SetHeaders(v []map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *RemoteResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRateLimit

`func (o *RemoteResponse) GetRateLimit() int32`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *RemoteResponse) GetRateLimitOk() (*int32, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *RemoteResponse) SetRateLimit(v int32)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *RemoteResponse) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### SetRateLimitNil

`func (o *RemoteResponse) SetRateLimitNil(b bool)`

 SetRateLimitNil sets the value for RateLimit to be an explicit nil

### UnsetRateLimit
`func (o *RemoteResponse) UnsetRateLimit()`

UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


