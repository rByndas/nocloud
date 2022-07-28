<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pkg/services_providers/proto/services_providers.proto

namespace Nocloud\ServicesProviders;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>nocloud.services_providers.InvokeRequest</code>
 */
class InvokeRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string uuid = 1 [json_name = "uuid"];</code>
     */
    protected $uuid = '';
    /**
     * Generated from protobuf field <code>string method = 2 [json_name = "method"];</code>
     */
    protected $method = '';
    /**
     * Generated from protobuf field <code>map<string, .google.protobuf.Value> params = 3 [json_name = "params"];</code>
     */
    private $params;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $uuid
     *     @type string $method
     *     @type array|\Google\Protobuf\Internal\MapField $params
     * }
     */
    public function __construct($data = NULL) {
        \Nocloud\ServicesProviders\GPBMetadata\ServicesProviders::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string uuid = 1 [json_name = "uuid"];</code>
     * @return string
     */
    public function getUuid()
    {
        return $this->uuid;
    }

    /**
     * Generated from protobuf field <code>string uuid = 1 [json_name = "uuid"];</code>
     * @param string $var
     * @return $this
     */
    public function setUuid($var)
    {
        GPBUtil::checkString($var, True);
        $this->uuid = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string method = 2 [json_name = "method"];</code>
     * @return string
     */
    public function getMethod()
    {
        return $this->method;
    }

    /**
     * Generated from protobuf field <code>string method = 2 [json_name = "method"];</code>
     * @param string $var
     * @return $this
     */
    public function setMethod($var)
    {
        GPBUtil::checkString($var, True);
        $this->method = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>map<string, .google.protobuf.Value> params = 3 [json_name = "params"];</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getParams()
    {
        return $this->params;
    }

    /**
     * Generated from protobuf field <code>map<string, .google.protobuf.Value> params = 3 [json_name = "params"];</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setParams($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::MESSAGE, \Google\Protobuf\Value::class);
        $this->params = $arr;

        return $this;
    }

}

