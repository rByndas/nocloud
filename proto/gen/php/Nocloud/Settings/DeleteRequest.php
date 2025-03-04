<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pkg/settings/proto/settings.proto

namespace Nocloud\Settings;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>nocloud.settings.DeleteRequest</code>
 */
class DeleteRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string key = 1 [json_name = "key"];</code>
     */
    protected $key = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $key
     * }
     */
    public function __construct($data = NULL) {
        \Nocloud\Settings\GPBMetadata\Settings::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string key = 1 [json_name = "key"];</code>
     * @return string
     */
    public function getKey()
    {
        return $this->key;
    }

    /**
     * Generated from protobuf field <code>string key = 1 [json_name = "key"];</code>
     * @param string $var
     * @return $this
     */
    public function setKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->key = $var;

        return $this;
    }

}

