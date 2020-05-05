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
 * The V1SecurityRuleSpec model module.
 * @module model/V1SecurityRuleSpec
 * @version 0.0.1
 */
class V1SecurityRuleSpec {
    /**
     * Constructs a new <code>V1SecurityRuleSpec</code>.
     * @alias module:model/V1SecurityRuleSpec
     */
    constructor() { 
        
        V1SecurityRuleSpec.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1SecurityRuleSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1SecurityRuleSpec} obj Optional instance to populate.
     * @return {module:model/V1SecurityRuleSpec} The populated <code>V1SecurityRuleSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1SecurityRuleSpec();

            if (data.hasOwnProperty('appliesTo')) {
                obj['appliesTo'] = ApiClient.convertToType(data['appliesTo'], ['String']);
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {Array.<String>}
     */
    getAppliesTo() {
        return this.appliesTo;
    }

    /**
     * @param {Array.<String>} appliesTo
     */
    setAppliesTo(appliesTo) {
        this['appliesTo'] = appliesTo;
    }
/**
     * @return {String}
     */
    getCode() {
        return this.code;
    }

    /**
     * @param {String} code
     */
    setCode(code) {
        this['code'] = code;
    }
/**
     * @return {String}
     */
    getDescription() {
        return this.description;
    }

    /**
     * @param {String} description
     */
    setDescription(description) {
        this['description'] = description;
    }
/**
     * @return {String}
     */
    getName() {
        return this.name;
    }

    /**
     * @param {String} name
     */
    setName(name) {
        this['name'] = name;
    }

}

/**
 * @member {Array.<String>} appliesTo
 */
V1SecurityRuleSpec.prototype['appliesTo'] = undefined;

/**
 * @member {String} code
 */
V1SecurityRuleSpec.prototype['code'] = undefined;

/**
 * @member {String} description
 */
V1SecurityRuleSpec.prototype['description'] = undefined;

/**
 * @member {String} name
 */
V1SecurityRuleSpec.prototype['name'] = undefined;






export default V1SecurityRuleSpec;

