// Code generated by protoc-gen-alias. DO NOT EDIT.
package v1beta1

import "istio.io/api/networking/v1alpha3"

type VirtualService = v1alpha3.VirtualService
type Destination = v1alpha3.Destination
type HTTPRoute = v1alpha3.HTTPRoute
type Delegate = v1alpha3.Delegate
type Headers = v1alpha3.Headers
type Headers_HeaderOperations = v1alpha3.Headers_HeaderOperations
type TLSRoute = v1alpha3.TLSRoute
type TCPRoute = v1alpha3.TCPRoute
type HTTPMatchRequest = v1alpha3.HTTPMatchRequest
type HTTPRouteDestination = v1alpha3.HTTPRouteDestination
type RouteDestination = v1alpha3.RouteDestination
type L4MatchAttributes = v1alpha3.L4MatchAttributes
type TLSMatchAttributes = v1alpha3.TLSMatchAttributes
type HTTPRedirect = v1alpha3.HTTPRedirect
type HTTPRedirect_RedirectPortSelection = v1alpha3.HTTPRedirect_RedirectPortSelection

const HTTPRedirect_FROM_PROTOCOL_DEFAULT HTTPRedirect_RedirectPortSelection = v1alpha3.HTTPRedirect_FROM_PROTOCOL_DEFAULT
const HTTPRedirect_FROM_REQUEST_PORT HTTPRedirect_RedirectPortSelection = v1alpha3.HTTPRedirect_FROM_REQUEST_PORT

type HTTPRedirect_Port = v1alpha3.HTTPRedirect_Port
type HTTPRedirect_DerivePort = v1alpha3.HTTPRedirect_DerivePort
type HTTPDirectResponse = v1alpha3.HTTPDirectResponse
type HTTPBody = v1alpha3.HTTPBody
type HTTPBody_String_ = v1alpha3.HTTPBody_String_
type HTTPBody_Bytes = v1alpha3.HTTPBody_Bytes
type HTTPRewrite = v1alpha3.HTTPRewrite
type RegexRewrite = v1alpha3.RegexRewrite
type StringMatch = v1alpha3.StringMatch
type StringMatch_Exact = v1alpha3.StringMatch_Exact
type StringMatch_Prefix = v1alpha3.StringMatch_Prefix
type StringMatch_Regex = v1alpha3.StringMatch_Regex
type HTTPRetry = v1alpha3.HTTPRetry
type CorsPolicy = v1alpha3.CorsPolicy
type CorsPolicy_UnmatchedPreflights = v1alpha3.CorsPolicy_UnmatchedPreflights

const CorsPolicy_UNSPECIFIED CorsPolicy_UnmatchedPreflights = v1alpha3.CorsPolicy_UNSPECIFIED
const CorsPolicy_FORWARD CorsPolicy_UnmatchedPreflights = v1alpha3.CorsPolicy_FORWARD
const CorsPolicy_IGNORE CorsPolicy_UnmatchedPreflights = v1alpha3.CorsPolicy_IGNORE

type HTTPFaultInjection = v1alpha3.HTTPFaultInjection
type HTTPFaultInjection_Delay = v1alpha3.HTTPFaultInjection_Delay
type HTTPFaultInjection_Delay_FixedDelay = v1alpha3.HTTPFaultInjection_Delay_FixedDelay
type HTTPFaultInjection_Delay_ExponentialDelay = v1alpha3.HTTPFaultInjection_Delay_ExponentialDelay
type HTTPFaultInjection_Abort = v1alpha3.HTTPFaultInjection_Abort
type HTTPFaultInjection_Abort_HttpStatus = v1alpha3.HTTPFaultInjection_Abort_HttpStatus
type HTTPFaultInjection_Abort_GrpcStatus = v1alpha3.HTTPFaultInjection_Abort_GrpcStatus
type HTTPFaultInjection_Abort_Http2Error = v1alpha3.HTTPFaultInjection_Abort_Http2Error
type HTTPMirrorPolicy = v1alpha3.HTTPMirrorPolicy
type PortSelector = v1alpha3.PortSelector
type Percent = v1alpha3.Percent
