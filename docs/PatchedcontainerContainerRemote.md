# PatchedcontainerContainerRemote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name for this remote. | [optional] 
**Url** | Pointer to **string** | The URL of an external content source. | [optional] 
**CaCert** | Pointer to **NullableString** | A PEM encoded CA certificate used to validate the server certificate presented by the remote server. | [optional] 
**ClientCert** | Pointer to **NullableString** | A PEM encoded client certificate used for authentication. | [optional] 
**ClientKey** | Pointer to **NullableString** | A PEM encoded private key used for authentication. | [optional] 
**TlsValidation** | Pointer to **bool** | If True, TLS peer validation must be performed. | [optional] 
**ProxyUrl** | Pointer to **NullableString** | The proxy URL. Format: scheme://host:port | [optional] 
**ProxyUsername** | Pointer to **NullableString** | The username to authenticte to the proxy. | [optional] 
**ProxyPassword** | Pointer to **NullableString** | The password to authenticate to the proxy. Extra leading and trailing whitespace characters are not trimmed. | [optional] 
**Username** | Pointer to **NullableString** | The username to be used for authentication when syncing. | [optional] 
**Password** | Pointer to **NullableString** | The password to be used for authentication when syncing. Extra leading and trailing whitespace characters are not trimmed. | [optional] 
**PulpLabels** | Pointer to **map[string]interface{}** |  | [optional] 
**DownloadConcurrency** | Pointer to **NullableInt32** | Total number of simultaneous connections. If not set then the default value will be used. | [optional] 
**MaxRetries** | Pointer to **NullableInt32** | Maximum number of retry attempts after a download failure. If not set then the default value (3) will be used. | [optional] 
**Policy** | Pointer to [**NullablePolicy762Enum**](Policy762Enum.md) |          immediate - All manifests and blobs are downloaded and saved during a sync.         on_demand - Only tags and manifests are downloaded. Blobs are not                     downloaded until they are requested for the first time by a client.         streamed - Blobs are streamed to the client with every request and never saved.          | [optional] 
**TotalTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**ConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockReadTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**Headers** | Pointer to **[]map[string]interface{}** | Headers for aiohttp.Clientsession | [optional] 
**RateLimit** | Pointer to **NullableInt32** | Limits requests per second for each concurrent downloader | [optional] 
**UpstreamName** | Pointer to **string** | Name of the upstream repository | [optional] 
**IncludeTags** | Pointer to **[]string** |              A list of tags to include during sync.             Wildcards *, ? are recognized.             &#39;include_tags&#39; is evaluated before &#39;exclude_tags&#39;.              | [optional] 
**ExcludeTags** | Pointer to **[]string** |              A list of tags to exclude during sync.             Wildcards *, ? are recognized.             &#39;exclude_tags&#39; is evaluated after &#39;include_tags&#39;.              | [optional] 
**Sigstore** | Pointer to **string** | A URL to a sigstore to download image signatures from | [optional] 

## Methods

### NewPatchedcontainerContainerRemote

`func NewPatchedcontainerContainerRemote() *PatchedcontainerContainerRemote`

NewPatchedcontainerContainerRemote instantiates a new PatchedcontainerContainerRemote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedcontainerContainerRemoteWithDefaults

`func NewPatchedcontainerContainerRemoteWithDefaults() *PatchedcontainerContainerRemote`

NewPatchedcontainerContainerRemoteWithDefaults instantiates a new PatchedcontainerContainerRemote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedcontainerContainerRemote) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedcontainerContainerRemote) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedcontainerContainerRemote) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedcontainerContainerRemote) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedcontainerContainerRemote) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedcontainerContainerRemote) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedcontainerContainerRemote) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedcontainerContainerRemote) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCaCert

`func (o *PatchedcontainerContainerRemote) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *PatchedcontainerContainerRemote) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *PatchedcontainerContainerRemote) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *PatchedcontainerContainerRemote) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *PatchedcontainerContainerRemote) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *PatchedcontainerContainerRemote) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetClientCert

`func (o *PatchedcontainerContainerRemote) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *PatchedcontainerContainerRemote) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *PatchedcontainerContainerRemote) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *PatchedcontainerContainerRemote) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### SetClientCertNil

`func (o *PatchedcontainerContainerRemote) SetClientCertNil(b bool)`

 SetClientCertNil sets the value for ClientCert to be an explicit nil

### UnsetClientCert
`func (o *PatchedcontainerContainerRemote) UnsetClientCert()`

UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
### GetClientKey

`func (o *PatchedcontainerContainerRemote) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *PatchedcontainerContainerRemote) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *PatchedcontainerContainerRemote) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *PatchedcontainerContainerRemote) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### SetClientKeyNil

`func (o *PatchedcontainerContainerRemote) SetClientKeyNil(b bool)`

 SetClientKeyNil sets the value for ClientKey to be an explicit nil

### UnsetClientKey
`func (o *PatchedcontainerContainerRemote) UnsetClientKey()`

UnsetClientKey ensures that no value is present for ClientKey, not even an explicit nil
### GetTlsValidation

`func (o *PatchedcontainerContainerRemote) GetTlsValidation() bool`

GetTlsValidation returns the TlsValidation field if non-nil, zero value otherwise.

### GetTlsValidationOk

`func (o *PatchedcontainerContainerRemote) GetTlsValidationOk() (*bool, bool)`

GetTlsValidationOk returns a tuple with the TlsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsValidation

`func (o *PatchedcontainerContainerRemote) SetTlsValidation(v bool)`

SetTlsValidation sets TlsValidation field to given value.

### HasTlsValidation

`func (o *PatchedcontainerContainerRemote) HasTlsValidation() bool`

HasTlsValidation returns a boolean if a field has been set.

### GetProxyUrl

`func (o *PatchedcontainerContainerRemote) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *PatchedcontainerContainerRemote) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *PatchedcontainerContainerRemote) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *PatchedcontainerContainerRemote) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### SetProxyUrlNil

`func (o *PatchedcontainerContainerRemote) SetProxyUrlNil(b bool)`

 SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil

### UnsetProxyUrl
`func (o *PatchedcontainerContainerRemote) UnsetProxyUrl()`

UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
### GetProxyUsername

`func (o *PatchedcontainerContainerRemote) GetProxyUsername() string`

GetProxyUsername returns the ProxyUsername field if non-nil, zero value otherwise.

### GetProxyUsernameOk

`func (o *PatchedcontainerContainerRemote) GetProxyUsernameOk() (*string, bool)`

GetProxyUsernameOk returns a tuple with the ProxyUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUsername

`func (o *PatchedcontainerContainerRemote) SetProxyUsername(v string)`

SetProxyUsername sets ProxyUsername field to given value.

### HasProxyUsername

`func (o *PatchedcontainerContainerRemote) HasProxyUsername() bool`

HasProxyUsername returns a boolean if a field has been set.

### SetProxyUsernameNil

`func (o *PatchedcontainerContainerRemote) SetProxyUsernameNil(b bool)`

 SetProxyUsernameNil sets the value for ProxyUsername to be an explicit nil

### UnsetProxyUsername
`func (o *PatchedcontainerContainerRemote) UnsetProxyUsername()`

UnsetProxyUsername ensures that no value is present for ProxyUsername, not even an explicit nil
### GetProxyPassword

`func (o *PatchedcontainerContainerRemote) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *PatchedcontainerContainerRemote) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *PatchedcontainerContainerRemote) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *PatchedcontainerContainerRemote) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### SetProxyPasswordNil

`func (o *PatchedcontainerContainerRemote) SetProxyPasswordNil(b bool)`

 SetProxyPasswordNil sets the value for ProxyPassword to be an explicit nil

### UnsetProxyPassword
`func (o *PatchedcontainerContainerRemote) UnsetProxyPassword()`

UnsetProxyPassword ensures that no value is present for ProxyPassword, not even an explicit nil
### GetUsername

`func (o *PatchedcontainerContainerRemote) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedcontainerContainerRemote) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedcontainerContainerRemote) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedcontainerContainerRemote) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *PatchedcontainerContainerRemote) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *PatchedcontainerContainerRemote) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *PatchedcontainerContainerRemote) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedcontainerContainerRemote) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedcontainerContainerRemote) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedcontainerContainerRemote) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PatchedcontainerContainerRemote) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PatchedcontainerContainerRemote) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPulpLabels

`func (o *PatchedcontainerContainerRemote) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatchedcontainerContainerRemote) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatchedcontainerContainerRemote) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatchedcontainerContainerRemote) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetDownloadConcurrency

`func (o *PatchedcontainerContainerRemote) GetDownloadConcurrency() int32`

GetDownloadConcurrency returns the DownloadConcurrency field if non-nil, zero value otherwise.

### GetDownloadConcurrencyOk

`func (o *PatchedcontainerContainerRemote) GetDownloadConcurrencyOk() (*int32, bool)`

GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadConcurrency

`func (o *PatchedcontainerContainerRemote) SetDownloadConcurrency(v int32)`

SetDownloadConcurrency sets DownloadConcurrency field to given value.

### HasDownloadConcurrency

`func (o *PatchedcontainerContainerRemote) HasDownloadConcurrency() bool`

HasDownloadConcurrency returns a boolean if a field has been set.

### SetDownloadConcurrencyNil

`func (o *PatchedcontainerContainerRemote) SetDownloadConcurrencyNil(b bool)`

 SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil

### UnsetDownloadConcurrency
`func (o *PatchedcontainerContainerRemote) UnsetDownloadConcurrency()`

UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
### GetMaxRetries

`func (o *PatchedcontainerContainerRemote) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *PatchedcontainerContainerRemote) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *PatchedcontainerContainerRemote) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *PatchedcontainerContainerRemote) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *PatchedcontainerContainerRemote) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *PatchedcontainerContainerRemote) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetPolicy

`func (o *PatchedcontainerContainerRemote) GetPolicy() Policy762Enum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PatchedcontainerContainerRemote) GetPolicyOk() (*Policy762Enum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PatchedcontainerContainerRemote) SetPolicy(v Policy762Enum)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PatchedcontainerContainerRemote) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *PatchedcontainerContainerRemote) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *PatchedcontainerContainerRemote) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetTotalTimeout

`func (o *PatchedcontainerContainerRemote) GetTotalTimeout() float64`

GetTotalTimeout returns the TotalTimeout field if non-nil, zero value otherwise.

### GetTotalTimeoutOk

`func (o *PatchedcontainerContainerRemote) GetTotalTimeoutOk() (*float64, bool)`

GetTotalTimeoutOk returns a tuple with the TotalTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeout

`func (o *PatchedcontainerContainerRemote) SetTotalTimeout(v float64)`

SetTotalTimeout sets TotalTimeout field to given value.

### HasTotalTimeout

`func (o *PatchedcontainerContainerRemote) HasTotalTimeout() bool`

HasTotalTimeout returns a boolean if a field has been set.

### SetTotalTimeoutNil

`func (o *PatchedcontainerContainerRemote) SetTotalTimeoutNil(b bool)`

 SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil

### UnsetTotalTimeout
`func (o *PatchedcontainerContainerRemote) UnsetTotalTimeout()`

UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
### GetConnectTimeout

`func (o *PatchedcontainerContainerRemote) GetConnectTimeout() float64`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *PatchedcontainerContainerRemote) GetConnectTimeoutOk() (*float64, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *PatchedcontainerContainerRemote) SetConnectTimeout(v float64)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *PatchedcontainerContainerRemote) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### SetConnectTimeoutNil

`func (o *PatchedcontainerContainerRemote) SetConnectTimeoutNil(b bool)`

 SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil

### UnsetConnectTimeout
`func (o *PatchedcontainerContainerRemote) UnsetConnectTimeout()`

UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
### GetSockConnectTimeout

`func (o *PatchedcontainerContainerRemote) GetSockConnectTimeout() float64`

GetSockConnectTimeout returns the SockConnectTimeout field if non-nil, zero value otherwise.

### GetSockConnectTimeoutOk

`func (o *PatchedcontainerContainerRemote) GetSockConnectTimeoutOk() (*float64, bool)`

GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockConnectTimeout

`func (o *PatchedcontainerContainerRemote) SetSockConnectTimeout(v float64)`

SetSockConnectTimeout sets SockConnectTimeout field to given value.

### HasSockConnectTimeout

`func (o *PatchedcontainerContainerRemote) HasSockConnectTimeout() bool`

HasSockConnectTimeout returns a boolean if a field has been set.

### SetSockConnectTimeoutNil

`func (o *PatchedcontainerContainerRemote) SetSockConnectTimeoutNil(b bool)`

 SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil

### UnsetSockConnectTimeout
`func (o *PatchedcontainerContainerRemote) UnsetSockConnectTimeout()`

UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
### GetSockReadTimeout

`func (o *PatchedcontainerContainerRemote) GetSockReadTimeout() float64`

GetSockReadTimeout returns the SockReadTimeout field if non-nil, zero value otherwise.

### GetSockReadTimeoutOk

`func (o *PatchedcontainerContainerRemote) GetSockReadTimeoutOk() (*float64, bool)`

GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockReadTimeout

`func (o *PatchedcontainerContainerRemote) SetSockReadTimeout(v float64)`

SetSockReadTimeout sets SockReadTimeout field to given value.

### HasSockReadTimeout

`func (o *PatchedcontainerContainerRemote) HasSockReadTimeout() bool`

HasSockReadTimeout returns a boolean if a field has been set.

### SetSockReadTimeoutNil

`func (o *PatchedcontainerContainerRemote) SetSockReadTimeoutNil(b bool)`

 SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil

### UnsetSockReadTimeout
`func (o *PatchedcontainerContainerRemote) UnsetSockReadTimeout()`

UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
### GetHeaders

`func (o *PatchedcontainerContainerRemote) GetHeaders() []map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PatchedcontainerContainerRemote) GetHeadersOk() (*[]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PatchedcontainerContainerRemote) SetHeaders(v []map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PatchedcontainerContainerRemote) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRateLimit

`func (o *PatchedcontainerContainerRemote) GetRateLimit() int32`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *PatchedcontainerContainerRemote) GetRateLimitOk() (*int32, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *PatchedcontainerContainerRemote) SetRateLimit(v int32)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *PatchedcontainerContainerRemote) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### SetRateLimitNil

`func (o *PatchedcontainerContainerRemote) SetRateLimitNil(b bool)`

 SetRateLimitNil sets the value for RateLimit to be an explicit nil

### UnsetRateLimit
`func (o *PatchedcontainerContainerRemote) UnsetRateLimit()`

UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
### GetUpstreamName

`func (o *PatchedcontainerContainerRemote) GetUpstreamName() string`

GetUpstreamName returns the UpstreamName field if non-nil, zero value otherwise.

### GetUpstreamNameOk

`func (o *PatchedcontainerContainerRemote) GetUpstreamNameOk() (*string, bool)`

GetUpstreamNameOk returns a tuple with the UpstreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamName

`func (o *PatchedcontainerContainerRemote) SetUpstreamName(v string)`

SetUpstreamName sets UpstreamName field to given value.

### HasUpstreamName

`func (o *PatchedcontainerContainerRemote) HasUpstreamName() bool`

HasUpstreamName returns a boolean if a field has been set.

### GetIncludeTags

`func (o *PatchedcontainerContainerRemote) GetIncludeTags() []string`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *PatchedcontainerContainerRemote) GetIncludeTagsOk() (*[]string, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *PatchedcontainerContainerRemote) SetIncludeTags(v []string)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *PatchedcontainerContainerRemote) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### SetIncludeTagsNil

`func (o *PatchedcontainerContainerRemote) SetIncludeTagsNil(b bool)`

 SetIncludeTagsNil sets the value for IncludeTags to be an explicit nil

### UnsetIncludeTags
`func (o *PatchedcontainerContainerRemote) UnsetIncludeTags()`

UnsetIncludeTags ensures that no value is present for IncludeTags, not even an explicit nil
### GetExcludeTags

`func (o *PatchedcontainerContainerRemote) GetExcludeTags() []string`

GetExcludeTags returns the ExcludeTags field if non-nil, zero value otherwise.

### GetExcludeTagsOk

`func (o *PatchedcontainerContainerRemote) GetExcludeTagsOk() (*[]string, bool)`

GetExcludeTagsOk returns a tuple with the ExcludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeTags

`func (o *PatchedcontainerContainerRemote) SetExcludeTags(v []string)`

SetExcludeTags sets ExcludeTags field to given value.

### HasExcludeTags

`func (o *PatchedcontainerContainerRemote) HasExcludeTags() bool`

HasExcludeTags returns a boolean if a field has been set.

### SetExcludeTagsNil

`func (o *PatchedcontainerContainerRemote) SetExcludeTagsNil(b bool)`

 SetExcludeTagsNil sets the value for ExcludeTags to be an explicit nil

### UnsetExcludeTags
`func (o *PatchedcontainerContainerRemote) UnsetExcludeTags()`

UnsetExcludeTags ensures that no value is present for ExcludeTags, not even an explicit nil
### GetSigstore

`func (o *PatchedcontainerContainerRemote) GetSigstore() string`

GetSigstore returns the Sigstore field if non-nil, zero value otherwise.

### GetSigstoreOk

`func (o *PatchedcontainerContainerRemote) GetSigstoreOk() (*string, bool)`

GetSigstoreOk returns a tuple with the Sigstore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigstore

`func (o *PatchedcontainerContainerRemote) SetSigstore(v string)`

SetSigstore sets Sigstore field to given value.

### HasSigstore

`func (o *PatchedcontainerContainerRemote) HasSigstore() bool`

HasSigstore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


