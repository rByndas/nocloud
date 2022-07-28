<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pkg/instances/proto/instances.proto

namespace Nocloud\Instances;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>nocloud.instances.TestInstancesGroupConfigRequest</code>
 */
class TestInstancesGroupConfigRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.nocloud.instances.InstancesGroup group = 1 [json_name = "group"];</code>
     */
    protected $group = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Nocloud\Instances\InstancesGroup $group
     * }
     */
    public function __construct($data = NULL) {
        \Nocloud\Instances\GPBMetadata\Instances::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.nocloud.instances.InstancesGroup group = 1 [json_name = "group"];</code>
     * @return \Nocloud\Instances\InstancesGroup|null
     */
    public function getGroup()
    {
        return $this->group;
    }

    public function hasGroup()
    {
        return isset($this->group);
    }

    public function clearGroup()
    {
        unset($this->group);
    }

    /**
     * Generated from protobuf field <code>.nocloud.instances.InstancesGroup group = 1 [json_name = "group"];</code>
     * @param \Nocloud\Instances\InstancesGroup $var
     * @return $this
     */
    public function setGroup($var)
    {
        GPBUtil::checkMessage($var, \Nocloud\Instances\InstancesGroup::class);
        $this->group = $var;

        return $this;
    }

}

