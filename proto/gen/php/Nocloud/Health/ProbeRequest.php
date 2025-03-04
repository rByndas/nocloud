<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pkg/health/proto/health.proto

namespace Nocloud\Health;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>nocloud.health.ProbeRequest</code>
 */
class ProbeRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string probe_type = 1 [json_name = "probeType"];</code>
     */
    protected $probe_type = '';
    /**
     * Generated from protobuf field <code>string payload = 2 [json_name = "payload"];</code>
     */
    protected $payload = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $probe_type
     *     @type string $payload
     * }
     */
    public function __construct($data = NULL) {
        \Nocloud\Health\GPBMetadata\Health::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string probe_type = 1 [json_name = "probeType"];</code>
     * @return string
     */
    public function getProbeType()
    {
        return $this->probe_type;
    }

    /**
     * Generated from protobuf field <code>string probe_type = 1 [json_name = "probeType"];</code>
     * @param string $var
     * @return $this
     */
    public function setProbeType($var)
    {
        GPBUtil::checkString($var, True);
        $this->probe_type = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string payload = 2 [json_name = "payload"];</code>
     * @return string
     */
    public function getPayload()
    {
        return $this->payload;
    }

    /**
     * Generated from protobuf field <code>string payload = 2 [json_name = "payload"];</code>
     * @param string $var
     * @return $this
     */
    public function setPayload($var)
    {
        GPBUtil::checkString($var, True);
        $this->payload = $var;

        return $this;
    }

}

