/**
 * Appvia Kore API
 * Kore API provides the frontend API for the Appvia Kore (kore.appvia.io)
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: info@appvia.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1ClusterUser model module.
 * @module model/V1ClusterUser
 * @version 0.0.1
 */
class V1ClusterUser {
    /**
     * Constructs a new <code>V1ClusterUser</code>.
     * @alias module:model/V1ClusterUser
     * @param roles {Array.<String>} 
     * @param username {String} 
     */
    constructor(roles, username) { 
        
        V1ClusterUser.initialize(this, roles, username);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, roles, username) { 
        obj['roles'] = roles;
        obj['username'] = username;
    }

    /**
     * Constructs a <code>V1ClusterUser</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1ClusterUser} obj Optional instance to populate.
     * @return {module:model/V1ClusterUser} The populated <code>V1ClusterUser</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1ClusterUser();

            if (data.hasOwnProperty('roles')) {
                obj['roles'] = ApiClient.convertToType(data['roles'], ['String']);
            }
            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {Array.<String>}
     */
    getRoles() {
        return this.roles;
    }

    /**
     * @param {Array.<String>} roles
     */
    setRoles(roles) {
        this['roles'] = roles;
    }
/**
     * @return {String}
     */
    getUsername() {
        return this.username;
    }

    /**
     * @param {String} username
     */
    setUsername(username) {
        this['username'] = username;
    }

}

/**
 * @member {Array.<String>} roles
 */
V1ClusterUser.prototype['roles'] = undefined;

/**
 * @member {String} username
 */
V1ClusterUser.prototype['username'] = undefined;






export default V1ClusterUser;

