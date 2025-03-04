<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pkg/services/proto/services.proto

namespace Nocloud\Services;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>nocloud.services.TestConfigResponse</code>
 */
class TestConfigResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>bool result = 1 [json_name = "result"];</code>
     */
    protected $result = false;
    /**
     * Generated from protobuf field <code>repeated .nocloud.services.TestConfigError errors = 2 [json_name = "errors"];</code>
     */
    private $errors;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type bool $result
     *     @type array<\Nocloud\Services\TestConfigError>|\Google\Protobuf\Internal\RepeatedField $errors
     * }
     */
    public function __construct($data = NULL) {
        \Nocloud\Services\GPBMetadata\Services::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>bool result = 1 [json_name = "result"];</code>
     * @return bool
     */
    public function getResult()
    {
        return $this->result;
    }

    /**
     * Generated from protobuf field <code>bool result = 1 [json_name = "result"];</code>
     * @param bool $var
     * @return $this
     */
    public function setResult($var)
    {
        GPBUtil::checkBool($var);
        $this->result = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .nocloud.services.TestConfigError errors = 2 [json_name = "errors"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getErrors()
    {
        return $this->errors;
    }

    /**
     * Generated from protobuf field <code>repeated .nocloud.services.TestConfigError errors = 2 [json_name = "errors"];</code>
     * @param array<\Nocloud\Services\TestConfigError>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setErrors($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Nocloud\Services\TestConfigError::class);
        $this->errors = $arr;

        return $this;
    }

}

