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
 * The V1AuditEventSpec model module.
 * @module model/V1AuditEventSpec
 * @version 0.0.1
 */
class V1AuditEventSpec {
    /**
     * Constructs a new <code>V1AuditEventSpec</code>.
     * @alias module:model/V1AuditEventSpec
     */
    constructor() { 
        
        V1AuditEventSpec.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1AuditEventSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1AuditEventSpec} obj Optional instance to populate.
     * @return {module:model/V1AuditEventSpec} The populated <code>V1AuditEventSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1AuditEventSpec();

            if (data.hasOwnProperty('apiVersion')) {
                obj['apiVersion'] = ApiClient.convertToType(data['apiVersion'], 'String');
            }
            if (data.hasOwnProperty('completedAt')) {
                obj['completedAt'] = ApiClient.convertToType(data['completedAt'], 'String');
            }
            if (data.hasOwnProperty('createdAt')) {
                obj['createdAt'] = ApiClient.convertToType(data['createdAt'], 'String');
            }
            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('message')) {
                obj['message'] = ApiClient.convertToType(data['message'], 'String');
            }
            if (data.hasOwnProperty('operation')) {
                obj['operation'] = ApiClient.convertToType(data['operation'], 'String');
            }
            if (data.hasOwnProperty('resource')) {
                obj['resource'] = ApiClient.convertToType(data['resource'], 'String');
            }
            if (data.hasOwnProperty('resourceURI')) {
                obj['resourceURI'] = ApiClient.convertToType(data['resourceURI'], 'String');
            }
            if (data.hasOwnProperty('responseCode')) {
                obj['responseCode'] = ApiClient.convertToType(data['responseCode'], 'Number');
            }
            if (data.hasOwnProperty('startedAt')) {
                obj['startedAt'] = ApiClient.convertToType(data['startedAt'], 'String');
            }
            if (data.hasOwnProperty('team')) {
                obj['team'] = ApiClient.convertToType(data['team'], 'String');
            }
            if (data.hasOwnProperty('user')) {
                obj['user'] = ApiClient.convertToType(data['user'], 'String');
            }
            if (data.hasOwnProperty('verb')) {
                obj['verb'] = ApiClient.convertToType(data['verb'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {String}
     */
    getApiVersion() {
        return this.apiVersion;
    }

    /**
     * @param {String} apiVersion
     */
    setApiVersion(apiVersion) {
        this['apiVersion'] = apiVersion;
    }
/**
     * @return {String}
     */
    getCompletedAt() {
        return this.completedAt;
    }

    /**
     * @param {String} completedAt
     */
    setCompletedAt(completedAt) {
        this['completedAt'] = completedAt;
    }
/**
     * @return {String}
     */
    getCreatedAt() {
        return this.createdAt;
    }

    /**
     * @param {String} createdAt
     */
    setCreatedAt(createdAt) {
        this['createdAt'] = createdAt;
    }
/**
     * @return {Number}
     */
    getId() {
        return this.id;
    }

    /**
     * @param {Number} id
     */
    setId(id) {
        this['id'] = id;
    }
/**
     * @return {String}
     */
    getMessage() {
        return this.message;
    }

    /**
     * @param {String} message
     */
    setMessage(message) {
        this['message'] = message;
    }
/**
     * @return {String}
     */
    getOperation() {
        return this.operation;
    }

    /**
     * @param {String} operation
     */
    setOperation(operation) {
        this['operation'] = operation;
    }
/**
     * @return {String}
     */
    getResource() {
        return this.resource;
    }

    /**
     * @param {String} resource
     */
    setResource(resource) {
        this['resource'] = resource;
    }
/**
     * @return {String}
     */
    getResourceURI() {
        return this.resourceURI;
    }

    /**
     * @param {String} resourceURI
     */
    setResourceURI(resourceURI) {
        this['resourceURI'] = resourceURI;
    }
/**
     * @return {Number}
     */
    getResponseCode() {
        return this.responseCode;
    }

    /**
     * @param {Number} responseCode
     */
    setResponseCode(responseCode) {
        this['responseCode'] = responseCode;
    }
/**
     * @return {String}
     */
    getStartedAt() {
        return this.startedAt;
    }

    /**
     * @param {String} startedAt
     */
    setStartedAt(startedAt) {
        this['startedAt'] = startedAt;
    }
/**
     * @return {String}
     */
    getTeam() {
        return this.team;
    }

    /**
     * @param {String} team
     */
    setTeam(team) {
        this['team'] = team;
    }
/**
     * @return {String}
     */
    getUser() {
        return this.user;
    }

    /**
     * @param {String} user
     */
    setUser(user) {
        this['user'] = user;
    }
/**
     * @return {String}
     */
    getVerb() {
        return this.verb;
    }

    /**
     * @param {String} verb
     */
    setVerb(verb) {
        this['verb'] = verb;
    }

}

/**
 * @member {String} apiVersion
 */
V1AuditEventSpec.prototype['apiVersion'] = undefined;

/**
 * @member {String} completedAt
 */
V1AuditEventSpec.prototype['completedAt'] = undefined;

/**
 * @member {String} createdAt
 */
V1AuditEventSpec.prototype['createdAt'] = undefined;

/**
 * @member {Number} id
 */
V1AuditEventSpec.prototype['id'] = undefined;

/**
 * @member {String} message
 */
V1AuditEventSpec.prototype['message'] = undefined;

/**
 * @member {String} operation
 */
V1AuditEventSpec.prototype['operation'] = undefined;

/**
 * @member {String} resource
 */
V1AuditEventSpec.prototype['resource'] = undefined;

/**
 * @member {String} resourceURI
 */
V1AuditEventSpec.prototype['resourceURI'] = undefined;

/**
 * @member {Number} responseCode
 */
V1AuditEventSpec.prototype['responseCode'] = undefined;

/**
 * @member {String} startedAt
 */
V1AuditEventSpec.prototype['startedAt'] = undefined;

/**
 * @member {String} team
 */
V1AuditEventSpec.prototype['team'] = undefined;

/**
 * @member {String} user
 */
V1AuditEventSpec.prototype['user'] = undefined;

/**
 * @member {String} verb
 */
V1AuditEventSpec.prototype['verb'] = undefined;






export default V1AuditEventSpec;
