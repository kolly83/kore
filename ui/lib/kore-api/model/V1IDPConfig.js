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
import V1GithubIDP from './V1GithubIDP';
import V1GoogleIDP from './V1GoogleIDP';
import V1OIDCIDP from './V1OIDCIDP';
import V1SAMLIDP from './V1SAMLIDP';
import V1StaticOIDCIDP from './V1StaticOIDCIDP';

/**
 * The V1IDPConfig model module.
 * @module model/V1IDPConfig
 * @version 0.0.1
 */
class V1IDPConfig {
    /**
     * Constructs a new <code>V1IDPConfig</code>.
     * @alias module:model/V1IDPConfig
     */
    constructor() { 
        
        V1IDPConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1IDPConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1IDPConfig} obj Optional instance to populate.
     * @return {module:model/V1IDPConfig} The populated <code>V1IDPConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1IDPConfig();

            if (data.hasOwnProperty('github')) {
                obj['github'] = V1GithubIDP.constructFromObject(data['github']);
            }
            if (data.hasOwnProperty('google')) {
                obj['google'] = V1GoogleIDP.constructFromObject(data['google']);
            }
            if (data.hasOwnProperty('oidc')) {
                obj['oidc'] = V1OIDCIDP.constructFromObject(data['oidc']);
            }
            if (data.hasOwnProperty('oidcdirect')) {
                obj['oidcdirect'] = V1StaticOIDCIDP.constructFromObject(data['oidcdirect']);
            }
            if (data.hasOwnProperty('saml')) {
                obj['saml'] = V1SAMLIDP.constructFromObject(data['saml']);
            }
        }
        return obj;
    }

/**
     * @return {module:model/V1GithubIDP}
     */
    getGithub() {
        return this.github;
    }

    /**
     * @param {module:model/V1GithubIDP} github
     */
    setGithub(github) {
        this['github'] = github;
    }
/**
     * @return {module:model/V1GoogleIDP}
     */
    getGoogle() {
        return this.google;
    }

    /**
     * @param {module:model/V1GoogleIDP} google
     */
    setGoogle(google) {
        this['google'] = google;
    }
/**
     * @return {module:model/V1OIDCIDP}
     */
    getOidc() {
        return this.oidc;
    }

    /**
     * @param {module:model/V1OIDCIDP} oidc
     */
    setOidc(oidc) {
        this['oidc'] = oidc;
    }
/**
     * @return {module:model/V1StaticOIDCIDP}
     */
    getOidcdirect() {
        return this.oidcdirect;
    }

    /**
     * @param {module:model/V1StaticOIDCIDP} oidcdirect
     */
    setOidcdirect(oidcdirect) {
        this['oidcdirect'] = oidcdirect;
    }
/**
     * @return {module:model/V1SAMLIDP}
     */
    getSaml() {
        return this.saml;
    }

    /**
     * @param {module:model/V1SAMLIDP} saml
     */
    setSaml(saml) {
        this['saml'] = saml;
    }

}

/**
 * @member {module:model/V1GithubIDP} github
 */
V1IDPConfig.prototype['github'] = undefined;

/**
 * @member {module:model/V1GoogleIDP} google
 */
V1IDPConfig.prototype['google'] = undefined;

/**
 * @member {module:model/V1OIDCIDP} oidc
 */
V1IDPConfig.prototype['oidc'] = undefined;

/**
 * @member {module:model/V1StaticOIDCIDP} oidcdirect
 */
V1IDPConfig.prototype['oidcdirect'] = undefined;

/**
 * @member {module:model/V1SAMLIDP} saml
 */
V1IDPConfig.prototype['saml'] = undefined;






export default V1IDPConfig;

