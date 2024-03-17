function base64urlToBuffer(e){const r="==".slice(0,(4-e.length%4)%4);const t=e.replace(/-/g,"+").replace(/_/g,"/")+r;const n=atob(t);const i=new ArrayBuffer(n.length);const o=new Uint8Array(i);for(let e=0;e<n.length;e++)o[e]=n.charCodeAt(e);return i}function bufferToBase64url(e){const r=new Uint8Array(e);let t="";for(const e of r)t+=String.fromCharCode(e);const n=btoa(t);const i=n.replace(/\+/g,"-").replace(/\//g,"_").replace(/=/g,"");return i}var e="copy";var r="convert";function convert(t,n,i){if(n===e)return i;if(n===r)return t(i);if(n instanceof Array)return i.map((e=>convert(t,n[0],e)));if(n instanceof Object){const e={};for(const[r,o]of Object.entries(n)){if(o.derive){const e=o.derive(i);void 0!==e&&(i[r]=e)}if(r in i)null!=i[r]?e[r]=convert(t,o.schema,i[r]):e[r]=null;else if(o.required)throw new Error(`Missing key: ${r}`)}return e}}function derived(e,r){return{required:true,schema:e,derive:r}}function required(e){return{required:true,schema:e}}function optional(e){return{required:false,schema:e}}var t={type:required(e),id:required(r),transports:optional(e)};var n={appid:optional(e),appidExclude:optional(e),credProps:optional(e)};var i={appid:optional(e),appidExclude:optional(e),credProps:optional(e)};var o={publicKey:required({rp:required(e),user:required({id:required(r),name:required(e),displayName:required(e)}),challenge:required(r),pubKeyCredParams:required(e),timeout:optional(e),excludeCredentials:optional([t]),authenticatorSelection:optional(e),attestation:optional(e),extensions:optional(n)}),signal:optional(e)};var a={type:required(e),id:required(e),rawId:required(r),authenticatorAttachment:optional(e),response:required({clientDataJSON:required(r),attestationObject:required(r),transports:derived(e,(e=>{var r;return(null==(r=e.getTransports)?void 0:r.call(e))||[]}))}),clientExtensionResults:derived(i,(e=>e.getClientExtensionResults()))};var u={mediation:optional(e),publicKey:required({challenge:required(r),timeout:optional(e),rpId:optional(e),allowCredentials:optional([t]),userVerification:optional(e),extensions:optional(n)}),signal:optional(e)};var s={type:required(e),id:required(e),rawId:required(r),authenticatorAttachment:optional(e),response:required({clientDataJSON:required(r),authenticatorData:required(r),signature:required(r),userHandle:required(r)}),clientExtensionResults:derived(i,(e=>e.getClientExtensionResults()))};var c={credentialCreationOptions:o,publicKeyCredentialWithAttestation:a,credentialRequestOptions:u,publicKeyCredentialWithAssertion:s};function createRequestFromJSON(e){return convert(base64urlToBuffer,o,e)}function createResponseToJSON(e){return convert(bufferToBase64url,a,e)}async function create(e){const r=await navigator.credentials.create(createRequestFromJSON(e));return createResponseToJSON(r)}function getRequestFromJSON(e){return convert(base64urlToBuffer,u,e)}function getResponseToJSON(e){return convert(bufferToBase64url,s,e)}async function get(e){const r=await navigator.credentials.get(getRequestFromJSON(e));return getResponseToJSON(r)}function supported(){return!!(navigator.credentials&&navigator.credentials.create&&navigator.credentials.get&&window.PublicKeyCredential)}export{create,get,c as schema,supported};
