/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.9
 * Contact: nurmi@anchore.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type VulnerabilitiesApiService service


/* VulnerabilitiesApiService Get vulnerabilities by type
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param imageDigest 
 @param vtype 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "forceRefresh" (bool) 
     @param "vendorOnly" (bool) 
 @return VulnerabilityResponse*/
func (a *VulnerabilitiesApiService) GetImageVulnerabilitiesByType(ctx context.Context, imageDigest string, vtype string, localVarOptionals map[string]interface{}) (VulnerabilityResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  VulnerabilityResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/images/{imageDigest}/vuln/{vtype}"
	localVarPath = strings.Replace(localVarPath, "{"+"imageDigest"+"}", fmt.Sprintf("%v", imageDigest), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vtype"+"}", fmt.Sprintf("%v", vtype), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["forceRefresh"], "bool", "forceRefresh"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["vendorOnly"], "bool", "vendorOnly"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["forceRefresh"].(bool); localVarOk {
		localVarQueryParams.Add("force_refresh", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["vendorOnly"].(bool); localVarOk {
		localVarQueryParams.Add("vendor_only", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* VulnerabilitiesApiService Get vulnerabilities by type
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param imageId 
 @param vtype 
 @return VulnerabilityResponse*/
func (a *VulnerabilitiesApiService) GetImageVulnerabilitiesByTypeImageId(ctx context.Context, imageId string, vtype string) (VulnerabilityResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  VulnerabilityResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/images/by_id/{imageId}/vuln/{vtype}"
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vtype"+"}", fmt.Sprintf("%v", vtype), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* VulnerabilitiesApiService Get vulnerability types
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param imageDigest 
 @return []string*/
func (a *VulnerabilitiesApiService) GetImageVulnerabilityTypes(ctx context.Context, imageDigest string) ([]string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/images/{imageDigest}/vuln"
	localVarPath = strings.Replace(localVarPath, "{"+"imageDigest"+"}", fmt.Sprintf("%v", imageDigest), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* VulnerabilitiesApiService Get vulnerability types
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param imageId 
 @return []string*/
func (a *VulnerabilitiesApiService) GetImageVulnerabilityTypesByImageId(ctx context.Context, imageId string) ([]string,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/images/by_id/{imageId}/vuln"
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* VulnerabilitiesApiService List images vulnerable to the specific vulnerability ID.
 Returns a listing of images and their respective packages vulnerable to the given vulnerability ID
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param vulnerabilityId The ID of the vulnerability to search for within all images stored in anchore-engine (e.g. CVE-1999-0001)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "namespace" (string) Filter results to images within the given vulnerability namespace (e.g. debian:8, ubuntu:14.04)
     @param "affectedPackage" (string) Filter results to images with vulnable packages with the given package name (e.g. libssl)
     @param "severity" (string) Filter results to vulnerable package/vulnerability with the given severity
     @param "vendorOnly" (bool) Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data
     @param "page" (int32) The page of results to fetch. Pages start at 1
     @param "limit" (int32) Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page
 @return PaginatedVulnerableImageList*/
func (a *VulnerabilitiesApiService) QueryImagesByVulnerability(ctx context.Context, vulnerabilityId string, localVarOptionals map[string]interface{}) (PaginatedVulnerableImageList,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PaginatedVulnerableImageList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/query/images/by_vulnerability"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["namespace"], "string", "namespace"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["affectedPackage"], "string", "affectedPackage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["severity"], "string", "severity"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["vendorOnly"], "bool", "vendorOnly"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["page"], "int32", "page"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("vulnerability_id", parameterToString(vulnerabilityId, ""))
	if localVarTempParam, localVarOk := localVarOptionals["namespace"].(string); localVarOk {
		localVarQueryParams.Add("namespace", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["affectedPackage"].(string); localVarOk {
		localVarQueryParams.Add("affected_package", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["severity"].(string); localVarOk {
		localVarQueryParams.Add("severity", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["vendorOnly"].(bool); localVarOk {
		localVarQueryParams.Add("vendor_only", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["page"].(int32); localVarOk {
		localVarQueryParams.Add("page", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* VulnerabilitiesApiService Listing information about given vulnerability
 List (w/filters) vulnerability records known by the system, with affected packages information if present
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id The ID of the vulnerability (e.g. CVE-1999-0001)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "affectedPackage" (string) Filter results by specified package name (e.g. sed)
     @param "affectedPackageVersion" (string) Filter results by specified package version (e.g. 4.4-1)
     @param "page" (string) The page of results to fetch. Pages start at 1
     @param "limit" (int32) Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page
 @return PaginatedVulnerabilityList*/
func (a *VulnerabilitiesApiService) QueryVulnerabilities(ctx context.Context, id string, localVarOptionals map[string]interface{}) (PaginatedVulnerabilityList,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PaginatedVulnerabilityList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/query/vulnerabilities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["affectedPackage"], "string", "affectedPackage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["affectedPackageVersion"], "string", "affectedPackageVersion"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["page"], "string", "page"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("id", parameterToString(id, ""))
	if localVarTempParam, localVarOk := localVarOptionals["affectedPackage"].(string); localVarOk {
		localVarQueryParams.Add("affected_package", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["affectedPackageVersion"].(string); localVarOk {
		localVarQueryParams.Add("affected_package_version", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["page"].(string); localVarOk {
		localVarQueryParams.Add("page", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

