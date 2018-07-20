# \BlockApi

All URIs are relative to *http://betaapi.tezex.info/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlocksAll**](BlockApi.md#BlocksAll) | **Get** /blocks/all | Get All Blocks 
[**BlocksByLevel**](BlockApi.md#BlocksByLevel) | **Get** /blocks/{level} | Get All Blocks for a specific Level
[**BlocksByLevelRange**](BlockApi.md#BlocksByLevelRange) | **Get** /blocks/{startlevel}/{stoplevel} | Get All Blocks for a specific Level-Range
[**GetBlock**](BlockApi.md#GetBlock) | **Get** /block/{blockhash} | Get Block By Blockhash
[**GetBlockDelegations**](BlockApi.md#GetBlockDelegations) | **Get** /block/{blockhash}/operations/delegations | Get Delegations of a Block
[**GetBlockEndorsements**](BlockApi.md#GetBlockEndorsements) | **Get** /block/{blockhash}/operations/endorsements | Get Endorsements of a Block
[**GetBlockOperationsSorted**](BlockApi.md#GetBlockOperationsSorted) | **Get** /block/{blockhash}/operations | Get operations of a block, sorted
[**GetBlockOriginations**](BlockApi.md#GetBlockOriginations) | **Get** /block/{blockhash}/operations/originations | Get Originations of a Block
[**GetBlockTransaction**](BlockApi.md#GetBlockTransaction) | **Get** /block/{blockhash}/operations/transactions | Get Transactions of Block
[**RecentBlocks**](BlockApi.md#RecentBlocks) | **Get** /blocks/recent | returns the last 25 blocks


# **BlocksAll**
> BlocksAll BlocksAll($page, $order, $limit)

Get All Blocks 

Get all Blocks


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32**| Pagination, 200 tx per page max | [optional] 
 **order** | **string**| ASC or DESC | [optional] 
 **limit** | **int32**| Results per Page | [optional] 

### Return type

[**BlocksAll**](BlocksAll.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlocksByLevel**
> []Block BlocksByLevel($level)

Get All Blocks for a specific Level

Get all Blocks for a specific Level


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | **float32**| The level of the Blocks to retrieve, includes abandoned | 

### Return type

[**[]Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlocksByLevelRange**
> BlockRange BlocksByLevelRange($startlevel, $stoplevel)

Get All Blocks for a specific Level-Range

Get all Blocks for a specific Level-Range


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startlevel** | **float32**| lowest blocklevel to return | 
 **stoplevel** | **float32**| highest blocklevel to return | 

### Return type

[**BlockRange**](BlockRange.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlock**
> Block GetBlock($blockhash)

Get Block By Blockhash

Get a block by its hash


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockhash** | **string**| The hash of the Block to retrieve | 

### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockDelegations**
> []Delegation GetBlockDelegations($blockhash)

Get Delegations of a Block

Get all Delegations of a specific Block


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockhash** | **string**| Blockhash | 

### Return type

[**[]Delegation**](Delegation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockEndorsements**
> []Endorsement GetBlockEndorsements($blockhash)

Get Endorsements of a Block

Get all Endorsements of a specific Block


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockhash** | **string**| Blockhash | 

### Return type

[**[]Endorsement**](Endorsement.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockOperationsSorted**
> BlockOperationsSorted GetBlockOperationsSorted($blockhash)

Get operations of a block, sorted

Get the maximum Level we have seen, Blocks at this level may become abandoned Blocks later on


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockhash** | **string**| The hash of the Block to retrieve | 

### Return type

[**BlockOperationsSorted**](BlockOperationsSorted.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockOriginations**
> []Origination GetBlockOriginations($blockhash)

Get Originations of a Block

Get all Originations of a spcific Block


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockhash** | **string**| Blockhash | 

### Return type

[**[]Origination**](Origination.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockTransaction**
> []Transaction GetBlockTransaction($blockhash)

Get Transactions of Block

Get all Transactions of a spcific Block


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockhash** | **string**| Blockhash | 

### Return type

[**[]Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecentBlocks**
> []Block RecentBlocks()

returns the last 25 blocks

Get all Blocks for a specific Level


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

