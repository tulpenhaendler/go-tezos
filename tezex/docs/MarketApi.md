# \MarketApi

All URIs are relative to *http://betaapi.tezex.info/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Candlestick**](MarketApi.md#Candlestick) | **Get** /price/{denominator}/{numerator}/{period} | Candlestick Data
[**Ticker**](MarketApi.md#Ticker) | **Get** /ticker/{numerator} | Get Ticker for a specific Currency


# **Candlestick**
> []Candlestick Candlestick($denominator, $numerator, $period)

Candlestick Data

Returns CandleStick Prices


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **denominator** | **string**| which currency | 
 **numerator** | **string**| to which currency | 
 **period** | **string**| Timeframe of one candle | 

### Return type

[**[]Candlestick**](Candlestick.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ticker**
> Ticker Ticker($numerator)

Get Ticker for a specific Currency

Returns BTC, USD, EUR and CNY Prices


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **numerator** | **string**| The level of the Blocks to retrieve | 

### Return type

[**Ticker**](Ticker.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

