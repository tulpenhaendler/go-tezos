# \AccountApi

All URIs are relative to *http://betaapi.tezex.info/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](AccountApi.md#GetAccount) | **Get** /account/{account} | Get Account
[**GetAccountBalance**](AccountApi.md#GetAccountBalance) | **Get** /account/{account}/balance | Get Account Balance
[**GetAccountLastSeen**](AccountApi.md#GetAccountLastSeen) | **Get** /account/{account}/last_seen | Get last active date
[**GetAccountOperationCount**](AccountApi.md#GetAccountOperationCount) | **Get** /account/{account}/operations_count | Get operation count of Account
[**GetAccountTransactionCount**](AccountApi.md#GetAccountTransactionCount) | **Get** /account/{account}/transaction_count | Get transaction count of Account
[**GetDelegationsForAccount**](AccountApi.md#GetDelegationsForAccount) | **Get** /account/{account}/delegations | Get Delegations of this account
[**GetDelegationsToAccount**](AccountApi.md#GetDelegationsToAccount) | **Get** /account/{account}/delegated | Get Delegations to this account
[**GetEndorsementsForAccount**](AccountApi.md#GetEndorsementsForAccount) | **Get** /account/{account}/endorsements | Get Endorsements this Account has made
[**GetTransactionForAccountIncoming**](AccountApi.md#GetTransactionForAccountIncoming) | **Get** /account/{account}/transactions/incoming | Get Transaction
[**GetTransactionForAccountOutgoing**](AccountApi.md#GetTransactionForAccountOutgoing) | **Get** /account/{account}/transactions/outgoing | Get Transaction


# **GetAccount**
> Account GetAccount($account)

Get Account

Get Acccount


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| The account | 

### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountBalance**
> string GetAccountBalance($account)

Get Account Balance

Get Balance


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| The account | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountLastSeen**
> time.Time GetAccountLastSeen($account)

Get last active date

Get LastSeen Date


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| The account | 

### Return type

[**time.Time**](time.Time.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountOperationCount**
> int32 GetAccountOperationCount($account)

Get operation count of Account

Get Operation Count


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| The account | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountTransactionCount**
> int32 GetAccountTransactionCount($account)

Get transaction count of Account

Get Transaction Count


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| The account | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDelegationsForAccount**
> []Delegation GetDelegationsForAccount($account, $before)

Get Delegations of this account

Get Delegations this Account has made


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| The account for which to retrieve Delegations | 
 **before** | **int32**| Only Return Delegations before this blocklevel | [optional] 

### Return type

[**[]Delegation**](Delegation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDelegationsToAccount**
> []Delegation GetDelegationsToAccount($account, $before)

Get Delegations to this account

Get that have been made to this Account


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| The account to which delegations have been made | 
 **before** | **int32**| Only Return Delegations before this blocklevel | [optional] 

### Return type

[**[]Delegation**](Delegation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEndorsementsForAccount**
> []Endorsement GetEndorsementsForAccount($account, $before)

Get Endorsements this Account has made

Get Endorsements this Account has made


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| The account for which to retrieve Endorsements | 
 **before** | **int32**| Only Return Delegations before this blocklevel | [optional] 

### Return type

[**[]Endorsement**](Endorsement.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionForAccountIncoming**
> Transactions GetTransactionForAccountIncoming($account, $before)

Get Transaction

Get incoming Transactions for a specific Account


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| The account for which to retrieve incoming Transactions | 
 **before** | **int32**| Only Return transactions before this blocklevel | [optional] 

### Return type

[**Transactions**](Transactions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionForAccountOutgoing**
> Transactions GetTransactionForAccountOutgoing($account, $before)

Get Transaction

Get outgoing Transactions for a specific Account


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| The account for which to retrieve outgoing Transactions | 
 **before** | **int32**| Only return transactions before this blocklevel | [optional] 

### Return type

[**Transactions**](Transactions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

