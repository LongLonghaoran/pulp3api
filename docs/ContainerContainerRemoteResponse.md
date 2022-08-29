# ContainerContainerRemoteResponse

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
**Policy** | Pointer to [**NullablePolicy762Enum**](Policy762Enum.md) |          immediate - All manifests and blobs are downloaded and saved during a sync.         on_demand - Only tags and manifests are downloaded. Blobs are not                     downloaded until they are requested for the first time by a client.         streamed - Blobs are streamed to the client with every request and never saved.          | [optional] 
**TotalTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**ConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockReadTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**Headers** | Pointer to **[]map[string]interface{}** | Headers for aiohttp.Clientsession | [optional] 
**RateLimit** | Pointer to **NullableInt32** | Limits requests per second for each concurrent downloader | [optional] 
**UpstreamName** | **string** | Name of the upstream repository | 
**IncludeTags** | Pointer to **[]string** |              A list of tags to include during sync.             Wildcards *, ? are recognized.             &#39;include_tags&#39; is evaluated before &#39;exclude_tags&#39;.              | [optional] 
**ExcludeTags** | Pointer to **[]string** |              A list of tags to exclude during sync.             Wildcards *, ? are recognized.             &#39;exclude_tags&#39; is evaluated after &#39;include_tags&#39;.              | [optional] 
**Sigstore** | Pointer to **string** | A URL to a sigstore to download image signatures from | [optional] 

## Methods

### NewContainerContainerRemoteResponse

`func NewContainerContainerRemoteResponse(name string, url string, upstreamName string, ) *ContainerContainerRemoteResponse`

NewContainerContainerRemoteResponse instantiates a new ContainerContainerRemoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerContainerRemoteResponseWithDefaults

`func NewContainerContainerRemoteResponseWithDefaults() *ContainerContainerRemoteResponse`

NewContainerContainerRemoteResponseWithDefaults instantiates a new ContainerContainerRemoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ContainerContainerRemoteResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContainerContainerRemoteResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContainerContainerRemoteResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContainerContainerRemoteResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ContainerContainerRemoteResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ContainerContainerRemoteResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ContainerContainerRemoteResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ContainerContainerRemoteResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *ContainerContainerRemoteResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerContainerRemoteResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerContainerRemoteResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *ContainerContainerRemoteResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContainerContainerRemoteResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContainerContainerRemoteResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCaCert

`func (o *ContainerContainerRemoteResponse) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *ContainerContainerRemoteResponse) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *ContainerContainerRemoteResponse) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *ContainerContainerRemoteResponse) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *ContainerContainerRemoteResponse) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *ContainerContainerRemoteResponse) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetClientCert

`func (o *ContainerContainerRemoteResponse) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *ContainerContainerRemoteResponse) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *ContainerContainerRemoteResponse) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *ContainerContainerRemoteResponse) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### SetClientCertNil

`func (o *ContainerContainerRemoteResponse) SetClientCertNil(b bool)`

 SetClientCertNil sets the value for ClientCert to be an explicit nil

### UnsetClientCert
`func (o *ContainerContainerRemoteResponse) UnsetClientCert()`

UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
### GetTlsValidation

`func (o *ContainerContainerRemoteResponse) GetTlsValidation() bool`

GetTlsValidation returns the TlsValidation field if non-nil, zero value otherwise.

### GetTlsValidationOk

`func (o *ContainerContainerRemoteResponse) GetTlsValidationOk() (*bool, bool)`

GetTlsValidationOk returns a tuple with the TlsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsValidation

`func (o *ContainerContainerRemoteResponse) SetTlsValidation(v bool)`

SetTlsValidation sets TlsValidation field to given value.

### HasTlsValidation

`func (o *ContainerContainerRemoteResponse) HasTlsValidation() bool`

HasTlsValidation returns a boolean if a field has been set.

### GetProxyUrl

`func (o *ContainerContainerRemoteResponse) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *ContainerContainerRemoteResponse) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *ContainerContainerRemoteResponse) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *ContainerContainerRemoteResponse) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### SetProxyUrlNil

`func (o *ContainerContainerRemoteResponse) SetProxyUrlNil(b bool)`

 SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil

### UnsetProxyUrl
`func (o *ContainerContainerRemoteResponse) UnsetProxyUrl()`

UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
### GetPulpLabels

`func (o *ContainerContainerRemoteResponse) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ContainerContainerRemoteResponse) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ContainerContainerRemoteResponse) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ContainerContainerRemoteResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *ContainerContainerRemoteResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *ContainerContainerRemoteResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *ContainerContainerRemoteResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *ContainerContainerRemoteResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetDownloadConcurrency

`func (o *ContainerContainerRemoteResponse) GetDownloadConcurrency() int32`

GetDownloadConcurrency returns the DownloadConcurrency field if non-nil, zero value otherwise.

### GetDownloadConcurrencyOk

`func (o *ContainerContainerRemoteResponse) GetDownloadConcurrencyOk() (*int32, bool)`

GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadConcurrency

`func (o *ContainerContainerRemoteResponse) SetDownloadConcurrency(v int32)`

SetDownloadConcurrency sets DownloadConcurrency field to given value.

### HasDownloadConcurrency

`func (o *ContainerContainerRemoteResponse) HasDownloadConcurrency() bool`

HasDownloadConcurrency returns a boolean if a field has been set.

### SetDownloadConcurrencyNil

`func (o *ContainerContainerRemoteResponse) SetDownloadConcurrencyNil(b bool)`

 SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil

### UnsetDownloadConcurrency
`func (o *ContainerContainerRemoteResponse) UnsetDownloadConcurrency()`

UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
### GetMaxRetries

`func (o *ContainerContainerRemoteResponse) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *ContainerContainerRemoteResponse) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *ContainerContainerRemoteResponse) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *ContainerContainerRemoteResponse) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *ContainerContainerRemoteResponse) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *ContainerContainerRemoteResponse) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetPolicy

`func (o *ContainerContainerRemoteResponse) GetPolicy() Policy762Enum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ContainerContainerRemoteResponse) GetPolicyOk() (*Policy762Enum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ContainerContainerRemoteResponse) SetPolicy(v Policy762Enum)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ContainerContainerRemoteResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *ContainerContainerRemoteResponse) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *ContainerContainerRemoteResponse) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetTotalTimeout

`func (o *ContainerContainerRemoteResponse) GetTotalTimeout() float64`

GetTotalTimeout returns the TotalTimeout field if non-nil, zero value otherwise.

### GetTotalTimeoutOk

`func (o *ContainerContainerRemoteResponse) GetTotalTimeoutOk() (*float64, bool)`

GetTotalTimeoutOk returns a tuple with the TotalTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeout

`func (o *ContainerContainerRemoteResponse) SetTotalTimeout(v float64)`

SetTotalTimeout sets TotalTimeout field to given value.

### HasTotalTimeout

`func (o *ContainerContainerRemoteResponse) HasTotalTimeout() bool`

HasTotalTimeout returns a boolean if a field has been set.

### SetTotalTimeoutNil

`func (o *ContainerContainerRemoteResponse) SetTotalTimeoutNil(b bool)`

 SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil

### UnsetTotalTimeout
`func (o *ContainerContainerRemoteResponse) UnsetTotalTimeout()`

UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
### GetConnectTimeout

`func (o *ContainerContainerRemoteResponse) GetConnectTimeout() float64`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *ContainerContainerRemoteResponse) GetConnectTimeoutOk() (*float64, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *ContainerContainerRemoteResponse) SetConnectTimeout(v float64)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *ContainerContainerRemoteResponse) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### SetConnectTimeoutNil

`func (o *ContainerContainerRemoteResponse) SetConnectTimeoutNil(b bool)`

 SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil

### UnsetConnectTimeout
`func (o *ContainerContainerRemoteResponse) UnsetConnectTimeout()`

UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
### GetSockConnectTimeout

`func (o *ContainerContainerRemoteResponse) GetSockConnectTimeout() float64`

GetSockConnectTimeout returns the SockConnectTimeout field if non-nil, zero value otherwise.

### GetSockConnectTimeoutOk

`func (o *ContainerContainerRemoteResponse) GetSockConnectTimeoutOk() (*float64, bool)`

GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockConnectTimeout

`func (o *ContainerContainerRemoteResponse) SetSockConnectTimeout(v float64)`

SetSockConnectTimeout sets SockConnectTimeout field to given value.

### HasSockConnectTimeout

`func (o *ContainerContainerRemoteResponse) HasSockConnectTimeout() bool`

HasSockConnectTimeout returns a boolean if a field has been set.

### SetSockConnectTimeoutNil

`func (o *ContainerContainerRemoteResponse) SetSockConnectTimeoutNil(b bool)`

 SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil

### UnsetSockConnectTimeout
`func (o *ContainerContainerRemoteResponse) UnsetSockConnectTimeout()`

UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
### GetSockReadTimeout

`func (o *ContainerContainerRemoteResponse) GetSockReadTimeout() float64`

GetSockReadTimeout returns the SockReadTimeout field if non-nil, zero value otherwise.

### GetSockReadTimeoutOk

`func (o *ContainerContainerRemoteResponse) GetSockReadTimeoutOk() (*float64, bool)`

GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockReadTimeout

`func (o *ContainerContainerRemoteResponse) SetSockReadTimeout(v float64)`

SetSockReadTimeout sets SockReadTimeout field to given value.

### HasSockReadTimeout

`func (o *ContainerContainerRemoteResponse) HasSockReadTimeout() bool`

HasSockReadTimeout returns a boolean if a field has been set.

### SetSockReadTimeoutNil

`func (o *ContainerContainerRemoteResponse) SetSockReadTimeoutNil(b bool)`

 SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil

### UnsetSockReadTimeout
`func (o *ContainerContainerRemoteResponse) UnsetSockReadTimeout()`

UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
### GetHeaders

`func (o *ContainerContainerRemoteResponse) GetHeaders() []map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ContainerContainerRemoteResponse) GetHeadersOk() (*[]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ContainerContainerRemoteResponse) SetHeaders(v []map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ContainerContainerRemoteResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRateLimit

`func (o *ContainerContainerRemoteResponse) GetRateLimit() int32`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *ContainerContainerRemoteResponse) GetRateLimitOk() (*int32, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *ContainerContainerRemoteResponse) SetRateLimit(v int32)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *ContainerContainerRemoteResponse) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### SetRateLimitNil

`func (o *ContainerContainerRemoteResponse) SetRateLimitNil(b bool)`

 SetRateLimitNil sets the value for RateLimit to be an explicit nil

### UnsetRateLimit
`func (o *ContainerContainerRemoteResponse) UnsetRateLimit()`

UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
### GetUpstreamName

`func (o *ContainerContainerRemoteResponse) GetUpstreamName() string`

GetUpstreamName returns the UpstreamName field if non-nil, zero value otherwise.

### GetUpstreamNameOk

`func (o *ContainerContainerRemoteResponse) GetUpstreamNameOk() (*string, bool)`

GetUpstreamNameOk returns a tuple with the UpstreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamName

`func (o *ContainerContainerRemoteResponse) SetUpstreamName(v string)`

SetUpstreamName sets UpstreamName field to given value.


### GetIncludeTags

`func (o *ContainerContainerRemoteResponse) GetIncludeTags() []string`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *ContainerContainerRemoteResponse) GetIncludeTagsOk() (*[]string, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *ContainerContainerRemoteResponse) SetIncludeTags(v []string)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *ContainerContainerRemoteResponse) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### SetIncludeTagsNil

`func (o *ContainerContainerRemoteResponse) SetIncludeTagsNil(b bool)`

 SetIncludeTagsNil sets the value for IncludeTags to be an explicit nil

### UnsetIncludeTags
`func (o *ContainerContainerRemoteResponse) UnsetIncludeTags()`

UnsetIncludeTags ensures that no value is present for IncludeTags, not even an explicit nil
### GetExcludeTags

`func (o *ContainerContainerRemoteResponse) GetExcludeTags() []string`

GetExcludeTags returns the ExcludeTags field if non-nil, zero value otherwise.

### GetExcludeTagsOk

`func (o *ContainerContainerRemoteResponse) GetExcludeTagsOk() (*[]string, bool)`

GetExcludeTagsOk returns a tuple with the ExcludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeTags

`func (o *ContainerContainerRemoteResponse) SetExcludeTags(v []string)`

SetExcludeTags sets ExcludeTags field to given value.

### HasExcludeTags

`func (o *ContainerContainerRemoteResponse) HasExcludeTags() bool`

HasExcludeTags returns a boolean if a field has been set.

### SetExcludeTagsNil

`func (o *ContainerContainerRemoteResponse) SetExcludeTagsNil(b bool)`

 SetExcludeTagsNil sets the value for ExcludeTags to be an explicit nil

### UnsetExcludeTags
`func (o *ContainerContainerRemoteResponse) UnsetExcludeTags()`

UnsetExcludeTags ensures that no value is present for ExcludeTags, not even an explicit nil
### GetSigstore

`func (o *ContainerContainerRemoteResponse) GetSigstore() string`

GetSigstore returns the Sigstore field if non-nil, zero value otherwise.

### GetSigstoreOk

`func (o *ContainerContainerRemoteResponse) GetSigstoreOk() (*string, bool)`

GetSigstoreOk returns a tuple with the Sigstore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigstore

`func (o *ContainerContainerRemoteResponse) SetSigstore(v string)`

SetSigstore sets Sigstore field to given value.

### HasSigstore

`func (o *ContainerContainerRemoteResponse) HasSigstore() bool`

HasSigstore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


