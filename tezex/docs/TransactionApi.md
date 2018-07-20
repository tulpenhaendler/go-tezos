# \TransactionApi

All URIs are relative to *http://betaapi.tezex.info/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransaction**](TransactionApi.md#GetTransaction) | **Get** /transaction/{transaction_hash} | Get Transaction
[**GetTransactionsRecent**](TransactionApi.md#GetTransactionsRecent) | **Get** /transactions/recent | Returns the last 50 Transactions
[**TransactionsAll**](TransactionApi.md#TransactionsAll) | **Get** /transactions/all | Get All Transactions
[**TransactionsByLevelRange**](TransactionApi.md#TransactionsByLevelRange) | **Get** /transactions/{startlevel}/{stoplevel} | Get All Transactions for a specific Level-Range


# **GetTransaction**
> Transaction GetTransaction($transactionHash)

Get Transaction

Get a specific Transaction


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionHash** | **string**| The hash of the Transaction to retrieve | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionsRecent**
> []Transaction GetTransactionsRecent()

Returns the last 50 Transactions


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsAll**
> TransactionRange TransactionsAll($page, $order, $limit)

Get All Transactions

Get all Transactions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32**| Pagination, 200 tx per page max | [optional] 
 **order** | **string**| ASC or DESC | [optional] 
 **limit** | **int32**| Results per Page | [optional] 

### Return type

[**TransactionRange**](TransactionRange.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsByLevelRange**
> TransactionRange TransactionsByLevelRange($startlevel, $stoplevel, $page, $order, $limit)

Get All Transactions for a specific Level-Range

Get all Transactions for a specific Level-Range


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startlevel** | **float32**| lowest blocklevel to return | 
 **stoplevel** | **float32**| highest blocklevel to return | 
 **page** | **float32**| Pagination, 200 tx per page max | [optional] 
 **order** | **string**| ASC or DESC | [optional] 
 **limit** | **int32**| Results per Page | [optional] 

### Return type

[**TransactionRange**](TransactionRange.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

