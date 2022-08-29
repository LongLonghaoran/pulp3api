# Remote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name for this remote. | 
**Url** | **string** | The URL of an external content source. | 
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
**Policy** | Pointer to [**NullablePolicyDb6Enum**](PolicyDb6Enum.md) | The policy to use when downloading content. | [optional] 
**TotalTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**ConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockReadTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**Headers** | Pointer to **[]map[string]interface{}** | Headers for aiohttp.Clientsession | [optional] 
**RateLimit** | Pointer to **NullableInt32** | Limits requests per second for each concurrent downloader | [optional] 

## Methods

### NewRemote

`func NewRemote(name string, url string, ) *Remote`

NewRemote instantiates a new Remote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteWithDefaults

`func NewRemoteWithDefaults() *Remote`

NewRemoteWithDefaults instantiates a new Remote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Remote) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Remote) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Remote) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *Remote) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Remote) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Remote) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCaCert

`func (o *Remote) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *Remote) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *Remote) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *Remote) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *Remote) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *Remote) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetClientCert

`func (o *Remote) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *Remote) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *Remote) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *Remote) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### SetClientCertNil

`func (o *Remote) SetClientCertNil(b bool)`

 SetClientCertNil sets the value for ClientCert to be an explicit nil

### UnsetClientCert
`func (o *Remote) UnsetClientCert()`

UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
### GetClientKey

`func (o *Remote) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *Remote) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *Remote) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *Remote) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### SetClientKeyNil

`func (o *Remote) SetClientKeyNil(b bool)`

 SetClientKeyNil sets the value for ClientKey to be an explicit nil

### UnsetClientKey
`func (o *Remote) UnsetClientKey()`

UnsetClientKey ensures that no value is present for ClientKey, not even an explicit nil
### GetTlsValidation

`func (o *Remote) GetTlsValidation() bool`

GetTlsValidation returns the TlsValidation field if non-nil, zero value otherwise.

### GetTlsValidationOk

`func (o *Remote) GetTlsValidationOk() (*bool, bool)`

GetTlsValidationOk returns a tuple with the TlsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsValidation

`func (o *Remote) SetTlsValidation(v bool)`

SetTlsValidation sets TlsValidation field to given value.

### HasTlsValidation

`func (o *Remote) HasTlsValidation() bool`

HasTlsValidation returns a boolean if a field has been set.

### GetProxyUrl

`func (o *Remote) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *Remote) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *Remote) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *Remote) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### SetProxyUrlNil

`func (o *Remote) SetProxyUrlNil(b bool)`

 SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil

### UnsetProxyUrl
`func (o *Remote) UnsetProxyUrl()`

UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
### GetProxyUsername

`func (o *Remote) GetProxyUsername() string`

GetProxyUsername returns the ProxyUsername field if non-nil, zero value otherwise.

### GetProxyUsernameOk

`func (o *Remote) GetProxyUsernameOk() (*string, bool)`

GetProxyUsernameOk returns a tuple with the ProxyUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUsername

`func (o *Remote) SetProxyUsername(v string)`

SetProxyUsername sets ProxyUsername field to given value.

### HasProxyUsername

`func (o *Remote) HasProxyUsername() bool`

HasProxyUsername returns a boolean if a field has been set.

### SetProxyUsernameNil

`func (o *Remote) SetProxyUsernameNil(b bool)`

 SetProxyUsernameNil sets the value for ProxyUsername to be an explicit nil

### UnsetProxyUsername
`func (o *Remote) UnsetProxyUsername()`

UnsetProxyUsername ensures that no value is present for ProxyUsername, not even an explicit nil
### GetProxyPassword

`func (o *Remote) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *Remote) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *Remote) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *Remote) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### SetProxyPasswordNil

`func (o *Remote) SetProxyPasswordNil(b bool)`

 SetProxyPasswordNil sets the value for ProxyPassword to be an explicit nil

### UnsetProxyPassword
`func (o *Remote) UnsetProxyPassword()`

UnsetProxyPassword ensures that no value is present for ProxyPassword, not even an explicit nil
### GetUsername

`func (o *Remote) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Remote) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Remote) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Remote) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *Remote) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *Remote) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *Remote) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Remote) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Remote) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Remote) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *Remote) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *Remote) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPulpLabels

`func (o *Remote) GetPulpLabels() map[string]interface{}`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *Remote) GetPulpLabelsOk() (*map[string]interface{}, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *Remote) SetPulpLabels(v map[string]interface{})`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *Remote) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetDownloadConcurrency

`func (o *Remote) GetDownloadConcurrency() int32`

GetDownloadConcurrency returns the DownloadConcurrency field if non-nil, zero value otherwise.

### GetDownloadConcurrencyOk

`func (o *Remote) GetDownloadConcurrencyOk() (*int32, bool)`

GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadConcurrency

`func (o *Remote) SetDownloadConcurrency(v int32)`

SetDownloadConcurrency sets DownloadConcurrency field to given value.

### HasDownloadConcurrency

`func (o *Remote) HasDownloadConcurrency() bool`

HasDownloadConcurrency returns a boolean if a field has been set.

### SetDownloadConcurrencyNil

`func (o *Remote) SetDownloadConcurrencyNil(b bool)`

 SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil

### UnsetDownloadConcurrency
`func (o *Remote) UnsetDownloadConcurrency()`

UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
### GetMaxRetries

`func (o *Remote) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *Remote) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *Remote) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *Remote) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *Remote) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *Remote) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetPolicy

`func (o *Remote) GetPolicy() PolicyDb6Enum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *Remote) GetPolicyOk() (*PolicyDb6Enum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *Remote) SetPolicy(v PolicyDb6Enum)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *Remote) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *Remote) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *Remote) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetTotalTimeout

`func (o *Remote) GetTotalTimeout() float64`

GetTotalTimeout returns the TotalTimeout field if non-nil, zero value otherwise.

### GetTotalTimeoutOk

`func (o *Remote) GetTotalTimeoutOk() (*float64, bool)`

GetTotalTimeoutOk returns a tuple with the TotalTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeout

`func (o *Remote) SetTotalTimeout(v float64)`

SetTotalTimeout sets TotalTimeout field to given value.

### HasTotalTimeout

`func (o *Remote) HasTotalTimeout() bool`

HasTotalTimeout returns a boolean if a field has been set.

### SetTotalTimeoutNil

`func (o *Remote) SetTotalTimeoutNil(b bool)`

 SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil

### UnsetTotalTimeout
`func (o *Remote) UnsetTotalTimeout()`

UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
### GetConnectTimeout

`func (o *Remote) GetConnectTimeout() float64`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *Remote) GetConnectTimeoutOk() (*float64, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *Remote) SetConnectTimeout(v float64)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *Remote) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### SetConnectTimeoutNil

`func (o *Remote) SetConnectTimeoutNil(b bool)`

 SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil

### UnsetConnectTimeout
`func (o *Remote) UnsetConnectTimeout()`

UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
### GetSockConnectTimeout

`func (o *Remote) GetSockConnectTimeout() float64`

GetSockConnectTimeout returns the SockConnectTimeout field if non-nil, zero value otherwise.

### GetSockConnectTimeoutOk

`func (o *Remote) GetSockConnectTimeoutOk() (*float64, bool)`

GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockConnectTimeout

`func (o *Remote) SetSockConnectTimeout(v float64)`

SetSockConnectTimeout sets SockConnectTimeout field to given value.

### HasSockConnectTimeout

`func (o *Remote) HasSockConnectTimeout() bool`

HasSockConnectTimeout returns a boolean if a field has been set.

### SetSockConnectTimeoutNil

`func (o *Remote) SetSockConnectTimeoutNil(b bool)`

 SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil

### UnsetSockConnectTimeout
`func (o *Remote) UnsetSockConnectTimeout()`

UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
### GetSockReadTimeout

`func (o *Remote) GetSockReadTimeout() float64`

GetSockReadTimeout returns the SockReadTimeout field if non-nil, zero value otherwise.

### GetSockReadTimeoutOk

`func (o *Remote) GetSockReadTimeoutOk() (*float64, bool)`

GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockReadTimeout

`func (o *Remote) SetSockReadTimeout(v float64)`

SetSockReadTimeout sets SockReadTimeout field to given value.

### HasSockReadTimeout

`func (o *Remote) HasSockReadTimeout() bool`

HasSockReadTimeout returns a boolean if a field has been set.

### SetSockReadTimeoutNil

`func (o *Remote) SetSockReadTimeoutNil(b bool)`

 SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil

### UnsetSockReadTimeout
`func (o *Remote) UnsetSockReadTimeout()`

UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
### GetHeaders

`func (o *Remote) GetHeaders() []map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Remote) GetHeadersOk() (*[]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Remote) SetHeaders(v []map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Remote) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRateLimit

`func (o *Remote) GetRateLimit() int32`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *Remote) GetRateLimitOk() (*int32, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *Remote) SetRateLimit(v int32)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *Remote) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### SetRateLimitNil

`func (o *Remote) SetRateLimitNil(b bool)`

 SetRateLimitNil sets the value for RateLimit to be an explicit nil

### UnsetRateLimit
`func (o *Remote) UnsetRateLimit()`

UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


