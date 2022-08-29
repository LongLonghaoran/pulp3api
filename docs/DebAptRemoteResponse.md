# DebAptRemoteResponse

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
**Policy** | Pointer to [**Policy762Enum**](Policy762Enum.md) | The policy to use when downloading content. The possible values include: &#39;immediate&#39;, &#39;on_demand&#39;, and &#39;streamed&#39;. &#39;immediate&#39; is the default. | [optional] 
**TotalTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**ConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockReadTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**Headers** | Pointer to **[]map[string]interface{}** | Headers for aiohttp.Clientsession | [optional] 
**RateLimit** | Pointer to **NullableInt32** | Limits requests per second for each concurrent downloader | [optional] 
**Distributions** | **string** | Whitespace separated list of distributions to sync. The distribution is the path from the repository root to the \&quot;Release\&quot; file you want to access. This is often, but not always, equal to either the codename or the suite of the release you want to sync. If the repository you are trying to sync uses \&quot;flat repository format\&quot;, the distribution must end with a \&quot;/\&quot;. Based on \&quot;/etc/apt/sources.list\&quot; syntax. | 
**Components** | Pointer to **NullableString** | Whitespace separatet list of components to sync. If none are supplied, all that are available will be synchronized. Leave blank for repositores using \&quot;flat repository format\&quot;. | [optional] 
**Architectures** | Pointer to **NullableString** | Whitespace separated list of architectures to sync If none are supplied, all that are available will be synchronized. A list of valid architecture specification strings can be found by running \&quot;dpkg-architecture -L\&quot;. A sync will download the intersection of the list of architectures provided via this field and those provided by the relevant \&quot;Release\&quot; file. Architecture&#x3D;\&quot;all\&quot; is always synchronized and does not need to be provided here. | [optional] 
**SyncSources** | Pointer to **bool** | Sync source packages | [optional] 
**SyncUdebs** | Pointer to **bool** | Sync installer packages | [optional] 
**SyncInstaller** | Pointer to **bool** | Sync installer files | [optional] 
**Gpgkey** | Pointer to **NullableString** | Gpg public key to verify origin releases against | [optional] 
**IgnoreMissingPackageIndices** | Pointer to **bool** | By default, upstream repositories that declare architectures and corresponding package indices in their Release files without actually publishing them, will fail to synchronize. Set this flag to True to allow the synchronization of such \&quot;partial mirrors\&quot; instead. Alternatively, you could make your remote filter by architectures for which the upstream repository does have indices. | [optional] 

## Methods

### NewDebAptRemoteResponse

`func NewDebAptRemoteResponse(name string, url string, distributions string, ) *DebAptRemoteResponse`

NewDebAptRemoteResponse instantiates a new DebAptRemoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebAptRemoteResponseWithDefaults

`func NewDebAptRemoteResponseWithDefaults() *DebAptRemoteResponse`

NewDebAptRemoteResponseWithDefaults instantiates a new DebAptRemoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *DebAptRemoteResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *DebAptRemoteResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *DebAptRemoteResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *DebAptRemoteResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *DebAptRemoteResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *DebAptRemoteResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *DebAptRemoteResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *DebAptRemoteResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *DebAptRemoteResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DebAptRemoteResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DebAptRemoteResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *DebAptRemoteResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DebAptRemoteResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DebAptRemoteResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCaCert

`func (o *DebAptRemoteResponse) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *DebAptRemoteResponse) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *DebAptRemoteResponse) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *DebAptRemoteResponse) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *DebAptRemoteResponse) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *DebAptRemoteResponse) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetClientCert

`func (o *DebAptRemoteResponse) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *DebAptRemoteResponse) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *DebAptRemoteResponse) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *DebAptRemoteResponse) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### SetClientCertNil

`func (o *DebAptRemoteResponse) SetClientCertNil(b bool)`

 SetClientCertNil sets the value for ClientCert to be an explicit nil

### UnsetClientCert
`func (o *DebAptRemoteResponse) UnsetClientCert()`

UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
### GetTlsValidation

`func (o *DebAptRemoteResponse) GetTlsValidation() bool`

GetTlsValidation returns the TlsValidation field if non-nil, zero value otherwise.

### GetTlsValidationOk

`func (o *DebAptRemoteResponse) GetTlsValidationOk() (*bool, bool)`

GetTlsValidationOk returns a tuple with the TlsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsValidation

`func (o *DebAptRemoteResponse) SetTlsValidation(v bool)`

SetTlsValidation sets TlsValidation field to given value.

### HasTlsValidation

`func (o *DebAptRemoteResponse) HasTlsValidation() bool`

HasTlsValidation returns a boolean if a field has been set.

### GetProxyUrl

`func (o *DebAptRemoteResponse) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *DebAptRemoteResponse) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *DebAptRemoteResponse) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *DebAptRemoteResponse) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### SetProxyUrlNil

`func (o *DebAptRemoteResponse) SetProxyUrlNil(b bool)`

 SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil

### UnsetProxyUrl
`func (o *DebAptRemoteResponse) UnsetProxyUrl()`

UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
### GetPulpLabels

`func (o *DebAptRemoteResponse) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *DebAptRemoteResponse) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *DebAptRemoteResponse) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *DebAptRemoteResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *DebAptRemoteResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *DebAptRemoteResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *DebAptRemoteResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *DebAptRemoteResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetDownloadConcurrency

`func (o *DebAptRemoteResponse) GetDownloadConcurrency() int32`

GetDownloadConcurrency returns the DownloadConcurrency field if non-nil, zero value otherwise.

### GetDownloadConcurrencyOk

`func (o *DebAptRemoteResponse) GetDownloadConcurrencyOk() (*int32, bool)`

GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadConcurrency

`func (o *DebAptRemoteResponse) SetDownloadConcurrency(v int32)`

SetDownloadConcurrency sets DownloadConcurrency field to given value.

### HasDownloadConcurrency

`func (o *DebAptRemoteResponse) HasDownloadConcurrency() bool`

HasDownloadConcurrency returns a boolean if a field has been set.

### SetDownloadConcurrencyNil

`func (o *DebAptRemoteResponse) SetDownloadConcurrencyNil(b bool)`

 SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil

### UnsetDownloadConcurrency
`func (o *DebAptRemoteResponse) UnsetDownloadConcurrency()`

UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
### GetMaxRetries

`func (o *DebAptRemoteResponse) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *DebAptRemoteResponse) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *DebAptRemoteResponse) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *DebAptRemoteResponse) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *DebAptRemoteResponse) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *DebAptRemoteResponse) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetPolicy

`func (o *DebAptRemoteResponse) GetPolicy() Policy762Enum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *DebAptRemoteResponse) GetPolicyOk() (*Policy762Enum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *DebAptRemoteResponse) SetPolicy(v Policy762Enum)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *DebAptRemoteResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetTotalTimeout

`func (o *DebAptRemoteResponse) GetTotalTimeout() float64`

GetTotalTimeout returns the TotalTimeout field if non-nil, zero value otherwise.

### GetTotalTimeoutOk

`func (o *DebAptRemoteResponse) GetTotalTimeoutOk() (*float64, bool)`

GetTotalTimeoutOk returns a tuple with the TotalTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeout

`func (o *DebAptRemoteResponse) SetTotalTimeout(v float64)`

SetTotalTimeout sets TotalTimeout field to given value.

### HasTotalTimeout

`func (o *DebAptRemoteResponse) HasTotalTimeout() bool`

HasTotalTimeout returns a boolean if a field has been set.

### SetTotalTimeoutNil

`func (o *DebAptRemoteResponse) SetTotalTimeoutNil(b bool)`

 SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil

### UnsetTotalTimeout
`func (o *DebAptRemoteResponse) UnsetTotalTimeout()`

UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
### GetConnectTimeout

`func (o *DebAptRemoteResponse) GetConnectTimeout() float64`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *DebAptRemoteResponse) GetConnectTimeoutOk() (*float64, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *DebAptRemoteResponse) SetConnectTimeout(v float64)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *DebAptRemoteResponse) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### SetConnectTimeoutNil

`func (o *DebAptRemoteResponse) SetConnectTimeoutNil(b bool)`

 SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil

### UnsetConnectTimeout
`func (o *DebAptRemoteResponse) UnsetConnectTimeout()`

UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
### GetSockConnectTimeout

`func (o *DebAptRemoteResponse) GetSockConnectTimeout() float64`

GetSockConnectTimeout returns the SockConnectTimeout field if non-nil, zero value otherwise.

### GetSockConnectTimeoutOk

`func (o *DebAptRemoteResponse) GetSockConnectTimeoutOk() (*float64, bool)`

GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockConnectTimeout

`func (o *DebAptRemoteResponse) SetSockConnectTimeout(v float64)`

SetSockConnectTimeout sets SockConnectTimeout field to given value.

### HasSockConnectTimeout

`func (o *DebAptRemoteResponse) HasSockConnectTimeout() bool`

HasSockConnectTimeout returns a boolean if a field has been set.

### SetSockConnectTimeoutNil

`func (o *DebAptRemoteResponse) SetSockConnectTimeoutNil(b bool)`

 SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil

### UnsetSockConnectTimeout
`func (o *DebAptRemoteResponse) UnsetSockConnectTimeout()`

UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
### GetSockReadTimeout

`func (o *DebAptRemoteResponse) GetSockReadTimeout() float64`

GetSockReadTimeout returns the SockReadTimeout field if non-nil, zero value otherwise.

### GetSockReadTimeoutOk

`func (o *DebAptRemoteResponse) GetSockReadTimeoutOk() (*float64, bool)`

GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockReadTimeout

`func (o *DebAptRemoteResponse) SetSockReadTimeout(v float64)`

SetSockReadTimeout sets SockReadTimeout field to given value.

### HasSockReadTimeout

`func (o *DebAptRemoteResponse) HasSockReadTimeout() bool`

HasSockReadTimeout returns a boolean if a field has been set.

### SetSockReadTimeoutNil

`func (o *DebAptRemoteResponse) SetSockReadTimeoutNil(b bool)`

 SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil

### UnsetSockReadTimeout
`func (o *DebAptRemoteResponse) UnsetSockReadTimeout()`

UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
### GetHeaders

`func (o *DebAptRemoteResponse) GetHeaders() []map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DebAptRemoteResponse) GetHeadersOk() (*[]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DebAptRemoteResponse) SetHeaders(v []map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DebAptRemoteResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRateLimit

`func (o *DebAptRemoteResponse) GetRateLimit() int32`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *DebAptRemoteResponse) GetRateLimitOk() (*int32, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *DebAptRemoteResponse) SetRateLimit(v int32)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *DebAptRemoteResponse) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### SetRateLimitNil

`func (o *DebAptRemoteResponse) SetRateLimitNil(b bool)`

 SetRateLimitNil sets the value for RateLimit to be an explicit nil

### UnsetRateLimit
`func (o *DebAptRemoteResponse) UnsetRateLimit()`

UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
### GetDistributions

`func (o *DebAptRemoteResponse) GetDistributions() string`

GetDistributions returns the Distributions field if non-nil, zero value otherwise.

### GetDistributionsOk

`func (o *DebAptRemoteResponse) GetDistributionsOk() (*string, bool)`

GetDistributionsOk returns a tuple with the Distributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributions

`func (o *DebAptRemoteResponse) SetDistributions(v string)`

SetDistributions sets Distributions field to given value.


### GetComponents

`func (o *DebAptRemoteResponse) GetComponents() string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DebAptRemoteResponse) GetComponentsOk() (*string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DebAptRemoteResponse) SetComponents(v string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *DebAptRemoteResponse) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *DebAptRemoteResponse) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *DebAptRemoteResponse) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetArchitectures

`func (o *DebAptRemoteResponse) GetArchitectures() string`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *DebAptRemoteResponse) GetArchitecturesOk() (*string, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *DebAptRemoteResponse) SetArchitectures(v string)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *DebAptRemoteResponse) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### SetArchitecturesNil

`func (o *DebAptRemoteResponse) SetArchitecturesNil(b bool)`

 SetArchitecturesNil sets the value for Architectures to be an explicit nil

### UnsetArchitectures
`func (o *DebAptRemoteResponse) UnsetArchitectures()`

UnsetArchitectures ensures that no value is present for Architectures, not even an explicit nil
### GetSyncSources

`func (o *DebAptRemoteResponse) GetSyncSources() bool`

GetSyncSources returns the SyncSources field if non-nil, zero value otherwise.

### GetSyncSourcesOk

`func (o *DebAptRemoteResponse) GetSyncSourcesOk() (*bool, bool)`

GetSyncSourcesOk returns a tuple with the SyncSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSources

`func (o *DebAptRemoteResponse) SetSyncSources(v bool)`

SetSyncSources sets SyncSources field to given value.

### HasSyncSources

`func (o *DebAptRemoteResponse) HasSyncSources() bool`

HasSyncSources returns a boolean if a field has been set.

### GetSyncUdebs

`func (o *DebAptRemoteResponse) GetSyncUdebs() bool`

GetSyncUdebs returns the SyncUdebs field if non-nil, zero value otherwise.

### GetSyncUdebsOk

`func (o *DebAptRemoteResponse) GetSyncUdebsOk() (*bool, bool)`

GetSyncUdebsOk returns a tuple with the SyncUdebs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUdebs

`func (o *DebAptRemoteResponse) SetSyncUdebs(v bool)`

SetSyncUdebs sets SyncUdebs field to given value.

### HasSyncUdebs

`func (o *DebAptRemoteResponse) HasSyncUdebs() bool`

HasSyncUdebs returns a boolean if a field has been set.

### GetSyncInstaller

`func (o *DebAptRemoteResponse) GetSyncInstaller() bool`

GetSyncInstaller returns the SyncInstaller field if non-nil, zero value otherwise.

### GetSyncInstallerOk

`func (o *DebAptRemoteResponse) GetSyncInstallerOk() (*bool, bool)`

GetSyncInstallerOk returns a tuple with the SyncInstaller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncInstaller

`func (o *DebAptRemoteResponse) SetSyncInstaller(v bool)`

SetSyncInstaller sets SyncInstaller field to given value.

### HasSyncInstaller

`func (o *DebAptRemoteResponse) HasSyncInstaller() bool`

HasSyncInstaller returns a boolean if a field has been set.

### GetGpgkey

`func (o *DebAptRemoteResponse) GetGpgkey() string`

GetGpgkey returns the Gpgkey field if non-nil, zero value otherwise.

### GetGpgkeyOk

`func (o *DebAptRemoteResponse) GetGpgkeyOk() (*string, bool)`

GetGpgkeyOk returns a tuple with the Gpgkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgkey

`func (o *DebAptRemoteResponse) SetGpgkey(v string)`

SetGpgkey sets Gpgkey field to given value.

### HasGpgkey

`func (o *DebAptRemoteResponse) HasGpgkey() bool`

HasGpgkey returns a boolean if a field has been set.

### SetGpgkeyNil

`func (o *DebAptRemoteResponse) SetGpgkeyNil(b bool)`

 SetGpgkeyNil sets the value for Gpgkey to be an explicit nil

### UnsetGpgkey
`func (o *DebAptRemoteResponse) UnsetGpgkey()`

UnsetGpgkey ensures that no value is present for Gpgkey, not even an explicit nil
### GetIgnoreMissingPackageIndices

`func (o *DebAptRemoteResponse) GetIgnoreMissingPackageIndices() bool`

GetIgnoreMissingPackageIndices returns the IgnoreMissingPackageIndices field if non-nil, zero value otherwise.

### GetIgnoreMissingPackageIndicesOk

`func (o *DebAptRemoteResponse) GetIgnoreMissingPackageIndicesOk() (*bool, bool)`

GetIgnoreMissingPackageIndicesOk returns a tuple with the IgnoreMissingPackageIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMissingPackageIndices

`func (o *DebAptRemoteResponse) SetIgnoreMissingPackageIndices(v bool)`

SetIgnoreMissingPackageIndices sets IgnoreMissingPackageIndices field to given value.

### HasIgnoreMissingPackageIndices

`func (o *DebAptRemoteResponse) HasIgnoreMissingPackageIndices() bool`

HasIgnoreMissingPackageIndices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


