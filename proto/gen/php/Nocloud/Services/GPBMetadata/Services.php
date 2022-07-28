<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pkg/services/proto/services.proto

namespace Nocloud\Services\GPBMetadata;

class Services
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Protobuf\Struct::initOnce();
        \Nocloud\Instances\GPBMetadata\Instances::initOnce();
        \Google\Api\GPBMetadata\Annotations::initOnce();
        \Nocloud\Hasher\GPBMetadata\Hasher::initOnce();
        \Nocloud\States\GPBMetadata\States::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
!pkg/services/proto/services.protonocloud.services#pkg/instances/proto/instances.protogoogle/api/annotations.proto pkg/hasher/hasherpb/hasher.protopkg/states/proto/states.proto"�
Service
uuid (	Ruuid
version (	BеRversion
title (	Rtitle7
status (2.nocloud.services.ServiceStatusRstatusF
context (2&.nocloud.services.Service.ContextEntryBеRcontextL
instances_groups (2!.nocloud.instances.InstancesGroupRinstancesGroups
hash (	B��Rhash&
access_level (H RaccessLevel�R
ContextEntry
key (	Rkey,
value (2.google.protobuf.ValueRvalue:8B
_access_level"9
Services-
pool (2.nocloud.services.ServiceRpool"j
TestConfigError
error (	Rerror%
instance_group (	RinstanceGroup
instance (	Rinstance"g
TestConfigResponse
result (Rresult9
errors (2!.nocloud.services.TestConfigErrorRerrors"�
CreateRequest3
service (2.nocloud.services.ServiceRservice
	namespace (	R	namespace\\
deploy_policies (23.nocloud.services.CreateRequest.DeployPoliciesEntryRdeployPoliciesA
DeployPoliciesEntry
key (Rkey
value (	Rvalue:8"#
DeleteRequest
uuid (	Ruuid">
DeleteResponse
result (Rresult
error (	Rerror"
	UpRequest
uuid (	Ruuid"{
UpError7
data (2#.nocloud.services.UpError.DataEntryRdata7
	DataEntry
key (	Rkey
value (	Rvalue:8"?

UpResponse1
errors (2.nocloud.services.UpErrorRerrors"!
DownRequest
uuid (	Ruuid"
DownResponse"�
ListRequest&
show_deleted (	H RshowDeleted�!
	namespace (	HR	namespace�
depth (HRdepth�B
_show_deletedB

_namespaceB
_depth" 

GetRequest
uuid (	Ruuid"&
GetStatesRequest
uuid (	Ruuid"#
StreamRequest
uuid (	Ruuid*]
ServiceStatus
UNSPECIFIED 
INIT
MODIFIED
UP
DOWN
DEL
PROC2�
ServicesServicei

TestConfig.nocloud.services.CreateRequest$.nocloud.services.TestConfigResponse"���"	/services:*Z
Create.nocloud.services.CreateRequest.nocloud.services.Service"���	/services:*[
Update.nocloud.services.Service.nocloud.services.Service"���2/services/{uuid}:*e
Delete.nocloud.services.DeleteRequest .nocloud.services.DeleteResponse"���*/services/{uuid}X
Get.nocloud.services.GetRequest.nocloud.services.Service"���/services/{uuid}T
List.nocloud.services.ListRequest.nocloud.services.Services"���	/services_
Up.nocloud.services.UpRequest.nocloud.services.UpResponse"���"/services/{uuid}/up:*g
Down.nocloud.services.DownRequest.nocloud.services.DownResponse" ���"/services/{uuid}/down:*i
Stream.nocloud.services.StreamRequest.nocloud.states.ObjectState"���/services/{uuid}/stream0B�
com.nocloud.servicesBServicesProtoP�NSX�Nocloud.Services�Nocloud\\Services�Nocloud\\Services\\GPBMetadata�Nocloud::Servicesbproto3'
        , true);

        static::$is_initialized = true;
    }
}

