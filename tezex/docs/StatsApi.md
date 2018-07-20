# \StatsApi

All URIs are relative to *http://betaapi.tezex.info/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatistics**](StatsApi.md#GetStatistics) | **Get** /stats/{group}/{stat}/{period} | Get Statistics
[**GetStatsOverview**](StatsApi.md#GetStatsOverview) | **Get** /stats/overview | Returns some basic Info


# **GetStatistics**
> []Stats GetStatistics($group, $stat, $period, $startTime, $endTime)

Get Statistics

Get Statistics


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| Block, Transaction, etc | 
 **stat** | **string**|  | 
 **period** | **string**|  | 
 **startTime** | **time.Time**|  | [optional] 
 **endTime** | **time.Time**|  | [optional] 

### Return type

[**[]Stats**](Stats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatsOverview**
> StatsOverview GetStatsOverview()

Returns some basic Info


### Parameters
This endpoint does not need any parameter.

### Return type

[**StatsOverview**](StatsOverview.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

