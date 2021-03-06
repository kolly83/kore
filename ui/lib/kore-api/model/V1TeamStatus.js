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
import V1Condition from './V1Condition';

/**
 * The V1TeamStatus model module.
 * @module model/V1TeamStatus
 * @version 0.0.1
 */
class V1TeamStatus {
    /**
     * Constructs a new <code>V1TeamStatus</code>.
     * @alias module:model/V1TeamStatus
     * @param conditions {Array.<module:model/V1Condition>} 
     * @param status {String} 
     */
    constructor(conditions, status) { 
        
        V1TeamStatus.initialize(this, conditions, status);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, conditions, status) { 
        obj['conditions'] = conditions;
        obj['status'] = status;
    }

    /**
     * Constructs a <code>V1TeamStatus</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1TeamStatus} obj Optional instance to populate.
     * @return {module:model/V1TeamStatus} The populated <code>V1TeamStatus</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1TeamStatus();

            if (data.hasOwnProperty('conditions')) {
                obj['conditions'] = ApiClient.convertToType(data['conditions'], [V1Condition]);
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {Array.<module:model/V1Condition>}
     */
    getConditions() {
        return this.conditions;
    }

    /**
     * @param {Array.<module:model/V1Condition>} conditions
     */
    setConditions(conditions) {
        this['conditions'] = conditions;
    }
/**
     * @return {String}
     */
    getStatus() {
        return this.status;
    }

    /**
     * @param {String} status
     */
    setStatus(status) {
        this['status'] = status;
    }

}

/**
 * @member {Array.<module:model/V1Condition>} conditions
 */
V1TeamStatus.prototype['conditions'] = undefined;

/**
 * @member {String} status
 */
V1TeamStatus.prototype['status'] = undefined;






export default V1TeamStatus;

