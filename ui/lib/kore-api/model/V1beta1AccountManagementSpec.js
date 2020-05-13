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
import V1Ownership from './V1Ownership';
import V1beta1AccountsRule from './V1beta1AccountsRule';

/**
 * The V1beta1AccountManagementSpec model module.
 * @module model/V1beta1AccountManagementSpec
 * @version 0.0.1
 */
class V1beta1AccountManagementSpec {
    /**
     * Constructs a new <code>V1beta1AccountManagementSpec</code>.
     * @alias module:model/V1beta1AccountManagementSpec
     * @param provider {String} 
     */
    constructor(provider) { 
        
        V1beta1AccountManagementSpec.initialize(this, provider);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, provider) { 
        obj['provider'] = provider;
    }

    /**
     * Constructs a <code>V1beta1AccountManagementSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1beta1AccountManagementSpec} obj Optional instance to populate.
     * @return {module:model/V1beta1AccountManagementSpec} The populated <code>V1beta1AccountManagementSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1beta1AccountManagementSpec();

            if (data.hasOwnProperty('organization')) {
                obj['organization'] = V1Ownership.constructFromObject(data['organization']);
            }
            if (data.hasOwnProperty('provider')) {
                obj['provider'] = ApiClient.convertToType(data['provider'], 'String');
            }
            if (data.hasOwnProperty('rules')) {
                obj['rules'] = ApiClient.convertToType(data['rules'], [V1beta1AccountsRule]);
            }
        }
        return obj;
    }

/**
     * @return {module:model/V1Ownership}
     */
    getOrganization() {
        return this.organization;
    }

    /**
     * @param {module:model/V1Ownership} organization
     */
    setOrganization(organization) {
        this['organization'] = organization;
    }
/**
     * @return {String}
     */
    getProvider() {
        return this.provider;
    }

    /**
     * @param {String} provider
     */
    setProvider(provider) {
        this['provider'] = provider;
    }
/**
     * @return {Array.<module:model/V1beta1AccountsRule>}
     */
    getRules() {
        return this.rules;
    }

    /**
     * @param {Array.<module:model/V1beta1AccountsRule>} rules
     */
    setRules(rules) {
        this['rules'] = rules;
    }

}

/**
 * @member {module:model/V1Ownership} organization
 */
V1beta1AccountManagementSpec.prototype['organization'] = undefined;

/**
 * @member {String} provider
 */
V1beta1AccountManagementSpec.prototype['provider'] = undefined;

/**
 * @member {Array.<module:model/V1beta1AccountsRule>} rules
 */
V1beta1AccountManagementSpec.prototype['rules'] = undefined;






export default V1beta1AccountManagementSpec;
