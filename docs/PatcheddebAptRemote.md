# PatcheddebAptRemote

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
**Policy** | Pointer to [**Policy762Enum**](Policy762Enum.md) | The policy to use when downloading content. The possible values include: &#39;immediate&#39;, &#39;on_demand&#39;, and &#39;streamed&#39;. &#39;immediate&#39; is the default. | [optional] 
**TotalTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**ConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockReadTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**Headers** | Pointer to **[]map[string]interface{}** | Headers for aiohttp.Clientsession | [optional] 
**RateLimit** | Pointer to **NullableInt32** | Limits requests per second for each concurrent downloader | [optional] 
**Distributions** | Pointer to **string** | Whitespace separated list of distributions to sync. The distribution is the path from the repository root to the \&quot;Release\&quot; file you want to access. This is often, but not always, equal to either the codename or the suite of the release you want to sync. If the repository you are trying to sync uses \&quot;flat repository format\&quot;, the distribution must end with a \&quot;/\&quot;. Based on \&quot;/etc/apt/sources.list\&quot; syntax. | [optional] 
**Components** | Pointer to **NullableString** | Whitespace separatet list of components to sync. If none are supplied, all that are available will be synchronized. Leave blank for repositores using \&quot;flat repository format\&quot;. | [optional] 
**Architectures** | Pointer to **NullableString** | Whitespace separated list of architectures to sync If none are supplied, all that are available will be synchronized. A list of valid architecture specification strings can be found by running \&quot;dpkg-architecture -L\&quot;. A sync will download the intersection of the list of architectures provided via this field and those provided by the relevant \&quot;Release\&quot; file. Architecture&#x3D;\&quot;all\&quot; is always synchronized and does not need to be provided here. | [optional] 
**SyncSources** | Pointer to **bool** | Sync source packages | [optional] 
**SyncUdebs** | Pointer to **bool** | Sync installer packages | [optional] 
**SyncInstaller** | Pointer to **bool** | Sync installer files | [optional] 
**Gpgkey** | Pointer to **NullableString** | Gpg public key to verify origin releases against | [optional] 
**IgnoreMissingPackageIndices** | Pointer to **bool** | By default, upstream repositories that declare architectures and corresponding package indices in their Release files without actually publishing them, will fail to synchronize. Set this flag to True to allow the synchronization of such \&quot;partial mirrors\&quot; instead. Alternatively, you could make your remote filter by architectures for which the upstream repository does have indices. | [optional] 

## Methods

### NewPatcheddebAptRemote

`func NewPatcheddebAptRemote() *PatcheddebAptRemote`

NewPatcheddebAptRemote instantiates a new PatcheddebAptRemote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatcheddebAptRemoteWithDefaults

`func NewPatcheddebAptRemoteWithDefaults() *PatcheddebAptRemote`

NewPatcheddebAptRemoteWithDefaults instantiates a new PatcheddebAptRemote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatcheddebAptRemote) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatcheddebAptRemote) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatcheddebAptRemote) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatcheddebAptRemote) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *PatcheddebAptRemote) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatcheddebAptRemote) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatcheddebAptRemote) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatcheddebAptRemote) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCaCert

`func (o *PatcheddebAptRemote) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *PatcheddebAptRemote) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *PatcheddebAptRemote) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *PatcheddebAptRemote) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *PatcheddebAptRemote) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *PatcheddebAptRemote) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetClientCert

`func (o *PatcheddebAptRemote) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *PatcheddebAptRemote) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *PatcheddebAptRemote) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *PatcheddebAptRemote) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### SetClientCertNil

`func (o *PatcheddebAptRemote) SetClientCertNil(b bool)`

 SetClientCertNil sets the value for ClientCert to be an explicit nil

### UnsetClientCert
`func (o *PatcheddebAptRemote) UnsetClientCert()`

UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
### GetClientKey

`func (o *PatcheddebAptRemote) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *PatcheddebAptRemote) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *PatcheddebAptRemote) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *PatcheddebAptRemote) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### SetClientKeyNil

`func (o *PatcheddebAptRemote) SetClientKeyNil(b bool)`

 SetClientKeyNil sets the value for ClientKey to be an explicit nil

### UnsetClientKey
`func (o *PatcheddebAptRemote) UnsetClientKey()`

UnsetClientKey ensures that no value is present for ClientKey, not even an explicit nil
### GetTlsValidation

`func (o *PatcheddebAptRemote) GetTlsValidation() bool`

GetTlsValidation returns the TlsValidation field if non-nil, zero value otherwise.

### GetTlsValidationOk

`func (o *PatcheddebAptRemote) GetTlsValidationOk() (*bool, bool)`

GetTlsValidationOk returns a tuple with the TlsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsValidation

`func (o *PatcheddebAptRemote) SetTlsValidation(v bool)`

SetTlsValidation sets TlsValidation field to given value.

### HasTlsValidation

`func (o *PatcheddebAptRemote) HasTlsValidation() bool`

HasTlsValidation returns a boolean if a field has been set.

### GetProxyUrl

`func (o *PatcheddebAptRemote) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *PatcheddebAptRemote) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *PatcheddebAptRemote) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *PatcheddebAptRemote) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### SetProxyUrlNil

`func (o *PatcheddebAptRemote) SetProxyUrlNil(b bool)`

 SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil

### UnsetProxyUrl
`func (o *PatcheddebAptRemote) UnsetProxyUrl()`

UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
### GetProxyUsername

`func (o *PatcheddebAptRemote) GetProxyUsername() string`

GetProxyUsername returns the ProxyUsername field if non-nil, zero value otherwise.

### GetProxyUsernameOk

`func (o *PatcheddebAptRemote) GetProxyUsernameOk() (*string, bool)`

GetProxyUsernameOk returns a tuple with the ProxyUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUsername

`func (o *PatcheddebAptRemote) SetProxyUsername(v string)`

SetProxyUsername sets ProxyUsername field to given value.

### HasProxyUsername

`func (o *PatcheddebAptRemote) HasProxyUsername() bool`

HasProxyUsername returns a boolean if a field has been set.

### SetProxyUsernameNil

`func (o *PatcheddebAptRemote) SetProxyUsernameNil(b bool)`

 SetProxyUsernameNil sets the value for ProxyUsername to be an explicit nil

### UnsetProxyUsername
`func (o *PatcheddebAptRemote) UnsetProxyUsername()`

UnsetProxyUsername ensures that no value is present for ProxyUsername, not even an explicit nil
### GetProxyPassword

`func (o *PatcheddebAptRemote) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *PatcheddebAptRemote) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *PatcheddebAptRemote) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *PatcheddebAptRemote) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### SetProxyPasswordNil

`func (o *PatcheddebAptRemote) SetProxyPasswordNil(b bool)`

 SetProxyPasswordNil sets the value for ProxyPassword to be an explicit nil

### UnsetProxyPassword
`func (o *PatcheddebAptRemote) UnsetProxyPassword()`

UnsetProxyPassword ensures that no value is present for ProxyPassword, not even an explicit nil
### GetUsername

`func (o *PatcheddebAptRemote) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatcheddebAptRemote) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatcheddebAptRemote) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatcheddebAptRemote) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *PatcheddebAptRemote) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *PatcheddebAptRemote) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *PatcheddebAptRemote) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatcheddebAptRemote) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatcheddebAptRemote) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatcheddebAptRemote) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PatcheddebAptRemote) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PatcheddebAptRemote) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPulpLabels

`func (o *PatcheddebAptRemote) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatcheddebAptRemote) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatcheddebAptRemote) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatcheddebAptRemote) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetDownloadConcurrency

`func (o *PatcheddebAptRemote) GetDownloadConcurrency() int32`

GetDownloadConcurrency returns the DownloadConcurrency field if non-nil, zero value otherwise.

### GetDownloadConcurrencyOk

`func (o *PatcheddebAptRemote) GetDownloadConcurrencyOk() (*int32, bool)`

GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadConcurrency

`func (o *PatcheddebAptRemote) SetDownloadConcurrency(v int32)`

SetDownloadConcurrency sets DownloadConcurrency field to given value.

### HasDownloadConcurrency

`func (o *PatcheddebAptRemote) HasDownloadConcurrency() bool`

HasDownloadConcurrency returns a boolean if a field has been set.

### SetDownloadConcurrencyNil

`func (o *PatcheddebAptRemote) SetDownloadConcurrencyNil(b bool)`

 SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil

### UnsetDownloadConcurrency
`func (o *PatcheddebAptRemote) UnsetDownloadConcurrency()`

UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
### GetMaxRetries

`func (o *PatcheddebAptRemote) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *PatcheddebAptRemote) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *PatcheddebAptRemote) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *PatcheddebAptRemote) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *PatcheddebAptRemote) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *PatcheddebAptRemote) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetPolicy

`func (o *PatcheddebAptRemote) GetPolicy() Policy762Enum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PatcheddebAptRemote) GetPolicyOk() (*Policy762Enum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PatcheddebAptRemote) SetPolicy(v Policy762Enum)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PatcheddebAptRemote) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetTotalTimeout

`func (o *PatcheddebAptRemote) GetTotalTimeout() float64`

GetTotalTimeout returns the TotalTimeout field if non-nil, zero value otherwise.

### GetTotalTimeoutOk

`func (o *PatcheddebAptRemote) GetTotalTimeoutOk() (*float64, bool)`

GetTotalTimeoutOk returns a tuple with the TotalTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeout

`func (o *PatcheddebAptRemote) SetTotalTimeout(v float64)`

SetTotalTimeout sets TotalTimeout field to given value.

### HasTotalTimeout

`func (o *PatcheddebAptRemote) HasTotalTimeout() bool`

HasTotalTimeout returns a boolean if a field has been set.

### SetTotalTimeoutNil

`func (o *PatcheddebAptRemote) SetTotalTimeoutNil(b bool)`

 SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil

### UnsetTotalTimeout
`func (o *PatcheddebAptRemote) UnsetTotalTimeout()`

UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
### GetConnectTimeout

`func (o *PatcheddebAptRemote) GetConnectTimeout() float64`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *PatcheddebAptRemote) GetConnectTimeoutOk() (*float64, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *PatcheddebAptRemote) SetConnectTimeout(v float64)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *PatcheddebAptRemote) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### SetConnectTimeoutNil

`func (o *PatcheddebAptRemote) SetConnectTimeoutNil(b bool)`

 SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil

### UnsetConnectTimeout
`func (o *PatcheddebAptRemote) UnsetConnectTimeout()`

UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
### GetSockConnectTimeout

`func (o *PatcheddebAptRemote) GetSockConnectTimeout() float64`

GetSockConnectTimeout returns the SockConnectTimeout field if non-nil, zero value otherwise.

### GetSockConnectTimeoutOk

`func (o *PatcheddebAptRemote) GetSockConnectTimeoutOk() (*float64, bool)`

GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockConnectTimeout

`func (o *PatcheddebAptRemote) SetSockConnectTimeout(v float64)`

SetSockConnectTimeout sets SockConnectTimeout field to given value.

### HasSockConnectTimeout

`func (o *PatcheddebAptRemote) HasSockConnectTimeout() bool`

HasSockConnectTimeout returns a boolean if a field has been set.

### SetSockConnectTimeoutNil

`func (o *PatcheddebAptRemote) SetSockConnectTimeoutNil(b bool)`

 SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil

### UnsetSockConnectTimeout
`func (o *PatcheddebAptRemote) UnsetSockConnectTimeout()`

UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
### GetSockReadTimeout

`func (o *PatcheddebAptRemote) GetSockReadTimeout() float64`

GetSockReadTimeout returns the SockReadTimeout field if non-nil, zero value otherwise.

### GetSockReadTimeoutOk

`func (o *PatcheddebAptRemote) GetSockReadTimeoutOk() (*float64, bool)`

GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockReadTimeout

`func (o *PatcheddebAptRemote) SetSockReadTimeout(v float64)`

SetSockReadTimeout sets SockReadTimeout field to given value.

### HasSockReadTimeout

`func (o *PatcheddebAptRemote) HasSockReadTimeout() bool`

HasSockReadTimeout returns a boolean if a field has been set.

### SetSockReadTimeoutNil

`func (o *PatcheddebAptRemote) SetSockReadTimeoutNil(b bool)`

 SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil

### UnsetSockReadTimeout
`func (o *PatcheddebAptRemote) UnsetSockReadTimeout()`

UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
### GetHeaders

`func (o *PatcheddebAptRemote) GetHeaders() []map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PatcheddebAptRemote) GetHeadersOk() (*[]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PatcheddebAptRemote) SetHeaders(v []map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PatcheddebAptRemote) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRateLimit

`func (o *PatcheddebAptRemote) GetRateLimit() int32`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *PatcheddebAptRemote) GetRateLimitOk() (*int32, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *PatcheddebAptRemote) SetRateLimit(v int32)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *PatcheddebAptRemote) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### SetRateLimitNil

`func (o *PatcheddebAptRemote) SetRateLimitNil(b bool)`

 SetRateLimitNil sets the value for RateLimit to be an explicit nil

### UnsetRateLimit
`func (o *PatcheddebAptRemote) UnsetRateLimit()`

UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
### GetDistributions

`func (o *PatcheddebAptRemote) GetDistributions() string`

GetDistributions returns the Distributions field if non-nil, zero value otherwise.

### GetDistributionsOk

`func (o *PatcheddebAptRemote) GetDistributionsOk() (*string, bool)`

GetDistributionsOk returns a tuple with the Distributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributions

`func (o *PatcheddebAptRemote) SetDistributions(v string)`

SetDistributions sets Distributions field to given value.

### HasDistributions

`func (o *PatcheddebAptRemote) HasDistributions() bool`

HasDistributions returns a boolean if a field has been set.

### GetComponents

`func (o *PatcheddebAptRemote) GetComponents() string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *PatcheddebAptRemote) GetComponentsOk() (*string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *PatcheddebAptRemote) SetComponents(v string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *PatcheddebAptRemote) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *PatcheddebAptRemote) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *PatcheddebAptRemote) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetArchitectures

`func (o *PatcheddebAptRemote) GetArchitectures() string`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *PatcheddebAptRemote) GetArchitecturesOk() (*string, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *PatcheddebAptRemote) SetArchitectures(v string)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *PatcheddebAptRemote) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### SetArchitecturesNil

`func (o *PatcheddebAptRemote) SetArchitecturesNil(b bool)`

 SetArchitecturesNil sets the value for Architectures to be an explicit nil

### UnsetArchitectures
`func (o *PatcheddebAptRemote) UnsetArchitectures()`

UnsetArchitectures ensures that no value is present for Architectures, not even an explicit nil
### GetSyncSources

`func (o *PatcheddebAptRemote) GetSyncSources() bool`

GetSyncSources returns the SyncSources field if non-nil, zero value otherwise.

### GetSyncSourcesOk

`func (o *PatcheddebAptRemote) GetSyncSourcesOk() (*bool, bool)`

GetSyncSourcesOk returns a tuple with the SyncSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSources

`func (o *PatcheddebAptRemote) SetSyncSources(v bool)`

SetSyncSources sets SyncSources field to given value.

### HasSyncSources

`func (o *PatcheddebAptRemote) HasSyncSources() bool`

HasSyncSources returns a boolean if a field has been set.

### GetSyncUdebs

`func (o *PatcheddebAptRemote) GetSyncUdebs() bool`

GetSyncUdebs returns the SyncUdebs field if non-nil, zero value otherwise.

### GetSyncUdebsOk

`func (o *PatcheddebAptRemote) GetSyncUdebsOk() (*bool, bool)`

GetSyncUdebsOk returns a tuple with the SyncUdebs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUdebs

`func (o *PatcheddebAptRemote) SetSyncUdebs(v bool)`

SetSyncUdebs sets SyncUdebs field to given value.

### HasSyncUdebs

`func (o *PatcheddebAptRemote) HasSyncUdebs() bool`

HasSyncUdebs returns a boolean if a field has been set.

### GetSyncInstaller

`func (o *PatcheddebAptRemote) GetSyncInstaller() bool`

GetSyncInstaller returns the SyncInstaller field if non-nil, zero value otherwise.

### GetSyncInstallerOk

`func (o *PatcheddebAptRemote) GetSyncInstallerOk() (*bool, bool)`

GetSyncInstallerOk returns a tuple with the SyncInstaller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncInstaller

`func (o *PatcheddebAptRemote) SetSyncInstaller(v bool)`

SetSyncInstaller sets SyncInstaller field to given value.

### HasSyncInstaller

`func (o *PatcheddebAptRemote) HasSyncInstaller() bool`

HasSyncInstaller returns a boolean if a field has been set.

### GetGpgkey

`func (o *PatcheddebAptRemote) GetGpgkey() string`

GetGpgkey returns the Gpgkey field if non-nil, zero value otherwise.

### GetGpgkeyOk

`func (o *PatcheddebAptRemote) GetGpgkeyOk() (*string, bool)`

GetGpgkeyOk returns a tuple with the Gpgkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgkey

`func (o *PatcheddebAptRemote) SetGpgkey(v string)`

SetGpgkey sets Gpgkey field to given value.

### HasGpgkey

`func (o *PatcheddebAptRemote) HasGpgkey() bool`

HasGpgkey returns a boolean if a field has been set.

### SetGpgkeyNil

`func (o *PatcheddebAptRemote) SetGpgkeyNil(b bool)`

 SetGpgkeyNil sets the value for Gpgkey to be an explicit nil

### UnsetGpgkey
`func (o *PatcheddebAptRemote) UnsetGpgkey()`

UnsetGpgkey ensures that no value is present for Gpgkey, not even an explicit nil
### GetIgnoreMissingPackageIndices

`func (o *PatcheddebAptRemote) GetIgnoreMissingPackageIndices() bool`

GetIgnoreMissingPackageIndices returns the IgnoreMissingPackageIndices field if non-nil, zero value otherwise.

### GetIgnoreMissingPackageIndicesOk

`func (o *PatcheddebAptRemote) GetIgnoreMissingPackageIndicesOk() (*bool, bool)`

GetIgnoreMissingPackageIndicesOk returns a tuple with the IgnoreMissingPackageIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMissingPackageIndices

`func (o *PatcheddebAptRemote) SetIgnoreMissingPackageIndices(v bool)`

SetIgnoreMissingPackageIndices sets IgnoreMissingPackageIndices field to given value.

### HasIgnoreMissingPackageIndices

`func (o *PatcheddebAptRemote) HasIgnoreMissingPackageIndices() bool`

HasIgnoreMissingPackageIndices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


