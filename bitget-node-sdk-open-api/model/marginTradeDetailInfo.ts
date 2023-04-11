/**
 * Bitget Open API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 2.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class MarginTradeDetailInfo {
    'ctime'?: string;
    'feeCcy'?: string;
    'fees'?: string;
    'fillId'?: string;
    'fillPrice'?: string;
    'fillQuantity'?: string;
    'fillTotalAmount'?: string;
    'orderId'?: string;
    'orderType'?: string;
    'side'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "ctime",
            "baseName": "ctime",
            "type": "string"
        },
        {
            "name": "feeCcy",
            "baseName": "feeCcy",
            "type": "string"
        },
        {
            "name": "fees",
            "baseName": "fees",
            "type": "string"
        },
        {
            "name": "fillId",
            "baseName": "fillId",
            "type": "string"
        },
        {
            "name": "fillPrice",
            "baseName": "fillPrice",
            "type": "string"
        },
        {
            "name": "fillQuantity",
            "baseName": "fillQuantity",
            "type": "string"
        },
        {
            "name": "fillTotalAmount",
            "baseName": "fillTotalAmount",
            "type": "string"
        },
        {
            "name": "orderId",
            "baseName": "orderId",
            "type": "string"
        },
        {
            "name": "orderType",
            "baseName": "orderType",
            "type": "string"
        },
        {
            "name": "side",
            "baseName": "side",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return MarginTradeDetailInfo.attributeTypeMap;
    }
}
