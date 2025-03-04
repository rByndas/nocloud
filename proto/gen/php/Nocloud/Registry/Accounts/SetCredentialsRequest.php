<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pkg/registry/proto/accounts/accounts.proto

namespace Nocloud\Registry\Accounts;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>nocloud.registry.accounts.SetCredentialsRequest</code>
 */
class SetCredentialsRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string account = 1 [json_name = "account"];</code>
     */
    protected $account = '';
    /**
     * Generated from protobuf field <code>.nocloud.registry.accounts.Credentials auth = 2 [json_name = "auth"];</code>
     */
    protected $auth = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $account
     *     @type \Nocloud\Registry\Accounts\Credentials $auth
     * }
     */
    public function __construct($data = NULL) {
        \Nocloud\Registry\Accounts\GPBMetadata\Accounts::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string account = 1 [json_name = "account"];</code>
     * @return string
     */
    public function getAccount()
    {
        return $this->account;
    }

    /**
     * Generated from protobuf field <code>string account = 1 [json_name = "account"];</code>
     * @param string $var
     * @return $this
     */
    public function setAccount($var)
    {
        GPBUtil::checkString($var, True);
        $this->account = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.nocloud.registry.accounts.Credentials auth = 2 [json_name = "auth"];</code>
     * @return \Nocloud\Registry\Accounts\Credentials|null
     */
    public function getAuth()
    {
        return $this->auth;
    }

    public function hasAuth()
    {
        return isset($this->auth);
    }

    public function clearAuth()
    {
        unset($this->auth);
    }

    /**
     * Generated from protobuf field <code>.nocloud.registry.accounts.Credentials auth = 2 [json_name = "auth"];</code>
     * @param \Nocloud\Registry\Accounts\Credentials $var
     * @return $this
     */
    public function setAuth($var)
    {
        GPBUtil::checkMessage($var, \Nocloud\Registry\Accounts\Credentials::class);
        $this->auth = $var;

        return $this;
    }

}

