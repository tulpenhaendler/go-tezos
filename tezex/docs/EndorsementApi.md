# \EndorsementApi

All URIs are relative to *http://betaapi.tezex.info/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEndorsement**](EndorsementApi.md#GetEndorsement) | **Get** /endorsement/{endorsement_hash} | Get Endorsement
[**GetEndorsementForBlock**](EndorsementApi.md#GetEndorsementForBlock) | **Get** /endorsement/for/{block_hash} | Get Endorsement


# **GetEndorsement**
> Endorsement GetEndorsement($endorsementHash)

Get Endorsement

Get a specific Endorsement


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endorsementHash** | **string**| The hash of the Endorsement to retrieve | 

### Return type

[**Endorsement**](Endorsement.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEndorsementForBlock**
> []Endorsement GetEndorsementForBlock($blockHash)

Get Endorsement

Get a specific Endorsement


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockHash** | **string**| blockhash | 

### Return type

[**[]Endorsement**](Endorsement.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

