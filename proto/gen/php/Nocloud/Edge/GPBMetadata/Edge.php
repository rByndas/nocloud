<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pkg/edge/proto/edge.proto

namespace Nocloud\Edge\GPBMetadata;

class Edge
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \Nocloud\States\GPBMetadata\States::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
pkg/edge/proto/edge.protonocloud.edge"
TestRequest"<
TestResponse
result (Rresult
error (	Rerror"
Empty2�
EdgeService=
Test.nocloud.edge.TestRequest.nocloud.edge.TestResponse=
	PostState.nocloud.states.ObjectState.nocloud.edge.EmptyBn
com.nocloud.edgeB	EdgeProtoP�NEX�Nocloud.Edge�Nocloud\\Edge�Nocloud\\Edge\\GPBMetadata�Nocloud::Edgebproto3'
        , true);

        static::$is_initialized = true;
    }
}

