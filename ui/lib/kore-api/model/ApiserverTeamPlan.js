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
import V1PlanSpec from './V1PlanSpec';

/**
 * The ApiserverTeamPlan model module.
 * @module model/ApiserverTeamPlan
 * @version 0.0.1
 */
class ApiserverTeamPlan {
    /**
     * Constructs a new <code>ApiserverTeamPlan</code>.
     * @alias module:model/ApiserverTeamPlan
     * @param parameterEditable {Object.<String, Boolean>} 
     * @param schema {String} 
     */
    constructor(parameterEditable, schema) { 
        
        ApiserverTeamPlan.initialize(this, parameterEditable, schema);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, parameterEditable, schema) { 
        obj['parameterEditable'] = parameterEditable;
        obj['schema'] = schema;
    }

    /**
     * Constructs a <code>ApiserverTeamPlan</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiserverTeamPlan} obj Optional instance to populate.
     * @return {module:model/ApiserverTeamPlan} The populated <code>ApiserverTeamPlan</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiserverTeamPlan();

            if (data.hasOwnProperty('parameterEditable')) {
                obj['parameterEditable'] = ApiClient.convertToType(data['parameterEditable'], {'String': 'Boolean'});
            }
            if (data.hasOwnProperty('plan')) {
                obj['plan'] = V1PlanSpec.constructFromObject(data['plan']);
            }
            if (data.hasOwnProperty('schema')) {
                obj['schema'] = ApiClient.convertToType(data['schema'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {Object.<String, Boolean>}
     */
    getParameterEditable() {
        return this.parameterEditable;
    }

    /**
     * @param {Object.<String, Boolean>} parameterEditable
     */
    setParameterEditable(parameterEditable) {
        this['parameterEditable'] = parameterEditable;
    }
/**
     * @return {module:model/V1PlanSpec}
     */
    getPlan() {
        return this.plan;
    }

    /**
     * @param {module:model/V1PlanSpec} plan
     */
    setPlan(plan) {
        this['plan'] = plan;
    }
/**
     * @return {String}
     */
    getSchema() {
        return this.schema;
    }

    /**
     * @param {String} schema
     */
    setSchema(schema) {
        this['schema'] = schema;
    }

}

/**
 * @member {Object.<String, Boolean>} parameterEditable
 */
ApiserverTeamPlan.prototype['parameterEditable'] = undefined;

/**
 * @member {module:model/V1PlanSpec} plan
 */
ApiserverTeamPlan.prototype['plan'] = undefined;

/**
 * @member {String} schema
 */
ApiserverTeamPlan.prototype['schema'] = undefined;






export default ApiserverTeamPlan;

