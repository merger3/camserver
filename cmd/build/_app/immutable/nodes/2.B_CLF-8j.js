import{h as dt,c as q,o as B,a as $,t as V,r as M}from"../chunks/disclose-version.CfGNboRe.js";import{p as ve,h as je,g as h,w as _,v as D,x as qe,H as j,R as Se}from"../chunks/runtime.DZmSjpD9.js";import{e as Oe,s as pt}from"../chunks/render.BZh8DXrS.js";import{i as Me}from"../chunks/lifecycle.0BE9jQ3h.js";import{p as oe,o as ht,a as mt,b as yt,i as wt}from"../chunks/index-client.Csmj3eWD.js";import{s as bt,d as gt}from"../chunks/misc.m7q6wSMb.js";function Et(e,t,n,r){n=n==null?null:n+"";var s=e.__attributes??(e.__attributes={});dt&&(s[t]=e.getAttribute(t)),s[t]!==(s[t]=n)&&(n===null?e.removeAttribute(t):e.setAttribute(t,n))}const Rt=!0,cr=Object.freeze(Object.defineProperty({__proto__:null,prerender:Rt},Symbol.toStringTag,{value:"Module"}));var St=V('<div class="rectangle svelte-1xw8x1p"></div>'),Ot=V('<div style="position: relative;"><!> <!></div>');function Tt(e,t){ve(t,!1);let n=oe(t,"onUpdateRectangle"),r=oe(t,"rectangleStyle"),s=D(!1),o=oe(t,"drawn",4,!1),i=D(!1),c=D(0),f=D(0),u=D(0),l=D(0),d=D();function E(){let w=h(c)-h(u),A=h(f)-h(l);return!(Math.hypot(w,A)<5)}const g=w=>{const{clientX:A,clientY:U}=w;if(!h(d))return;_(s,!0),o(!1),_(i,!1);const T=h(d).getBoundingClientRect();_(c,A-T.left),_(f,U-T.top),_(u,h(c)),_(l,h(f))},p=w=>{const{clientX:A,clientY:U}=w;if((!h(s)||!h(d))&&o())return;E()?_(i,!0):_(i,!1);const T=h(d).getBoundingClientRect();_(u,A-T.left),_(l,U-T.top)},m=()=>{if(!h(s))return;_(s,!1),E()?o(!0):(o(!1),_(i,!1));const w={x:Math.min(h(c),h(u)),y:Math.min(h(f),h(l)),width:Math.abs(h(u)-h(c)),height:Math.abs(h(l)-h(f))};n()(w)};ht(()=>{h(d)&&(h(d).addEventListener("mousedown",g),h(d).addEventListener("mousemove",p),h(d).addEventListener("mouseup",m))}),mt(()=>{h(d)&&(h(d).removeEventListener("mousedown",g),h(d).removeEventListener("mousemove",p),h(d).removeEventListener("mouseup",m))}),Me();var b=Ot();yt(b,w=>_(d,w),()=>h(d));var R=q(b);bt(R,gt(t),{});var O=B(B(R,!0));wt(O,()=>(h(s)||o())&&h(i),w=>{var A=St();const U=j(()=>Math.min(h(c),h(u))),T=j(()=>Math.min(h(f),h(l))),L=j(()=>Math.abs(h(u)-h(c))),X=j(()=>Math.abs(h(l)-h(f))),ut=j(()=>r().border),ft=j(()=>r().backgroundColor);qe(()=>Et(A,"style",`left: ${h(U)}px; top: ${h(T)}px; width: ${h(L)}px; height: ${h(X)}px; border: ${h(ut)}; background-color: ${h(ft)};`)),$(w,A)}),M(b),$(e,b),je()}function He(e,t){return function(){return e.apply(t,arguments)}}const{toString:At}=Object.prototype,{getPrototypeOf:we}=Object,ee=(e=>t=>{const n=At.call(t);return e[n]||(e[n]=n.slice(8,-1).toLowerCase())})(Object.create(null)),N=e=>(e=e.toLowerCase(),t=>ee(t)===e),te=e=>t=>typeof t===e,{isArray:H}=Array,J=te("undefined");function xt(e){return e!==null&&!J(e)&&e.constructor!==null&&!J(e.constructor)&&C(e.constructor.isBuffer)&&e.constructor.isBuffer(e)}const Ie=N("ArrayBuffer");function _t(e){let t;return typeof ArrayBuffer<"u"&&ArrayBuffer.isView?t=ArrayBuffer.isView(e):t=e&&e.buffer&&Ie(e.buffer),t}const Ct=te("string"),C=te("function"),ze=te("number"),ne=e=>e!==null&&typeof e=="object",Pt=e=>e===!0||e===!1,G=e=>{if(ee(e)!=="object")return!1;const t=we(e);return(t===null||t===Object.prototype||Object.getPrototypeOf(t)===null)&&!(Symbol.toStringTag in e)&&!(Symbol.iterator in e)},Nt=N("Date"),Lt=N("File"),Ft=N("Blob"),Dt=N("FileList"),Bt=e=>ne(e)&&C(e.pipe),Ut=e=>{let t;return e&&(typeof FormData=="function"&&e instanceof FormData||C(e.append)&&((t=ee(e))==="formdata"||t==="object"&&C(e.toString)&&e.toString()==="[object FormData]"))},kt=N("URLSearchParams"),[vt,jt,qt,Mt]=["ReadableStream","Request","Response","Headers"].map(N),Ht=e=>e.trim?e.trim():e.replace(/^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g,"");function W(e,t,{allOwnKeys:n=!1}={}){if(e===null||typeof e>"u")return;let r,s;if(typeof e!="object"&&(e=[e]),H(e))for(r=0,s=e.length;r<s;r++)t.call(null,e[r],r,e);else{const o=n?Object.getOwnPropertyNames(e):Object.keys(e),i=o.length;let c;for(r=0;r<i;r++)c=o[r],t.call(null,e[c],c,e)}}function $e(e,t){t=t.toLowerCase();const n=Object.keys(e);let r=n.length,s;for(;r-- >0;)if(s=n[r],t===s.toLowerCase())return s;return null}const Je=typeof globalThis<"u"?globalThis:typeof self<"u"?self:typeof window<"u"?window:global,Ve=e=>!J(e)&&e!==Je;function ue(){const{caseless:e}=Ve(this)&&this||{},t={},n=(r,s)=>{const o=e&&$e(t,s)||s;G(t[o])&&G(r)?t[o]=ue(t[o],r):G(r)?t[o]=ue({},r):H(r)?t[o]=r.slice():t[o]=r};for(let r=0,s=arguments.length;r<s;r++)arguments[r]&&W(arguments[r],n);return t}const It=(e,t,n,{allOwnKeys:r}={})=>(W(t,(s,o)=>{n&&C(s)?e[o]=He(s,n):e[o]=s},{allOwnKeys:r}),e),zt=e=>(e.charCodeAt(0)===65279&&(e=e.slice(1)),e),$t=(e,t,n,r)=>{e.prototype=Object.create(t.prototype,r),e.prototype.constructor=e,Object.defineProperty(e,"super",{value:t.prototype}),n&&Object.assign(e.prototype,n)},Jt=(e,t,n,r)=>{let s,o,i;const c={};if(t=t||{},e==null)return t;do{for(s=Object.getOwnPropertyNames(e),o=s.length;o-- >0;)i=s[o],(!r||r(i,e,t))&&!c[i]&&(t[i]=e[i],c[i]=!0);e=n!==!1&&we(e)}while(e&&(!n||n(e,t))&&e!==Object.prototype);return t},Vt=(e,t,n)=>{e=String(e),(n===void 0||n>e.length)&&(n=e.length),n-=t.length;const r=e.indexOf(t,n);return r!==-1&&r===n},Wt=e=>{if(!e)return null;if(H(e))return e;let t=e.length;if(!ze(t))return null;const n=new Array(t);for(;t-- >0;)n[t]=e[t];return n},Kt=(e=>t=>e&&t instanceof e)(typeof Uint8Array<"u"&&we(Uint8Array)),Xt=(e,t)=>{const r=(e&&e[Symbol.iterator]).call(e);let s;for(;(s=r.next())&&!s.done;){const o=s.value;t.call(e,o[0],o[1])}},Gt=(e,t)=>{let n;const r=[];for(;(n=e.exec(t))!==null;)r.push(n);return r},Yt=N("HTMLFormElement"),Qt=e=>e.toLowerCase().replace(/[-_\s]([a-z\d])(\w*)/g,function(n,r,s){return r.toUpperCase()+s}),Te=(({hasOwnProperty:e})=>(t,n)=>e.call(t,n))(Object.prototype),Zt=N("RegExp"),We=(e,t)=>{const n=Object.getOwnPropertyDescriptors(e),r={};W(n,(s,o)=>{let i;(i=t(s,o,e))!==!1&&(r[o]=i||s)}),Object.defineProperties(e,r)},en=e=>{We(e,(t,n)=>{if(C(e)&&["arguments","caller","callee"].indexOf(n)!==-1)return!1;const r=e[n];if(C(r)){if(t.enumerable=!1,"writable"in t){t.writable=!1;return}t.set||(t.set=()=>{throw Error("Can not rewrite read-only method '"+n+"'")})}})},tn=(e,t)=>{const n={},r=s=>{s.forEach(o=>{n[o]=!0})};return H(e)?r(e):r(String(e).split(t)),n},nn=()=>{},rn=(e,t)=>e!=null&&Number.isFinite(e=+e)?e:t,ie="abcdefghijklmnopqrstuvwxyz",Ae="0123456789",Ke={DIGIT:Ae,ALPHA:ie,ALPHA_DIGIT:ie+ie.toUpperCase()+Ae},sn=(e=16,t=Ke.ALPHA_DIGIT)=>{let n="";const{length:r}=t;for(;e--;)n+=t[Math.random()*r|0];return n};function on(e){return!!(e&&C(e.append)&&e[Symbol.toStringTag]==="FormData"&&e[Symbol.iterator])}const an=e=>{const t=new Array(10),n=(r,s)=>{if(ne(r)){if(t.indexOf(r)>=0)return;if(!("toJSON"in r)){t[s]=r;const o=H(r)?[]:{};return W(r,(i,c)=>{const f=n(i,s+1);!J(f)&&(o[c]=f)}),t[s]=void 0,o}}return r};return n(e,0)},cn=N("AsyncFunction"),ln=e=>e&&(ne(e)||C(e))&&C(e.then)&&C(e.catch),a={isArray:H,isArrayBuffer:Ie,isBuffer:xt,isFormData:Ut,isArrayBufferView:_t,isString:Ct,isNumber:ze,isBoolean:Pt,isObject:ne,isPlainObject:G,isReadableStream:vt,isRequest:jt,isResponse:qt,isHeaders:Mt,isUndefined:J,isDate:Nt,isFile:Lt,isBlob:Ft,isRegExp:Zt,isFunction:C,isStream:Bt,isURLSearchParams:kt,isTypedArray:Kt,isFileList:Dt,forEach:W,merge:ue,extend:It,trim:Ht,stripBOM:zt,inherits:$t,toFlatObject:Jt,kindOf:ee,kindOfTest:N,endsWith:Vt,toArray:Wt,forEachEntry:Xt,matchAll:Gt,isHTMLForm:Yt,hasOwnProperty:Te,hasOwnProp:Te,reduceDescriptors:We,freezeMethods:en,toObjectSet:tn,toCamelCase:Qt,noop:nn,toFiniteNumber:rn,findKey:$e,global:Je,isContextDefined:Ve,ALPHABET:Ke,generateString:sn,isSpecCompliantForm:on,toJSONObject:an,isAsyncFn:cn,isThenable:ln};function y(e,t,n,r,s){Error.call(this),Error.captureStackTrace?Error.captureStackTrace(this,this.constructor):this.stack=new Error().stack,this.message=e,this.name="AxiosError",t&&(this.code=t),n&&(this.config=n),r&&(this.request=r),s&&(this.response=s)}a.inherits(y,Error,{toJSON:function(){return{message:this.message,name:this.name,description:this.description,number:this.number,fileName:this.fileName,lineNumber:this.lineNumber,columnNumber:this.columnNumber,stack:this.stack,config:a.toJSONObject(this.config),code:this.code,status:this.response&&this.response.status?this.response.status:null}}});const Xe=y.prototype,Ge={};["ERR_BAD_OPTION_VALUE","ERR_BAD_OPTION","ECONNABORTED","ETIMEDOUT","ERR_NETWORK","ERR_FR_TOO_MANY_REDIRECTS","ERR_DEPRECATED","ERR_BAD_RESPONSE","ERR_BAD_REQUEST","ERR_CANCELED","ERR_NOT_SUPPORT","ERR_INVALID_URL"].forEach(e=>{Ge[e]={value:e}});Object.defineProperties(y,Ge);Object.defineProperty(Xe,"isAxiosError",{value:!0});y.from=(e,t,n,r,s,o)=>{const i=Object.create(Xe);return a.toFlatObject(e,i,function(f){return f!==Error.prototype},c=>c!=="isAxiosError"),y.call(i,e.message,t,n,r,s),i.cause=e,i.name=e.name,o&&Object.assign(i,o),i};const un=null;function fe(e){return a.isPlainObject(e)||a.isArray(e)}function Ye(e){return a.endsWith(e,"[]")?e.slice(0,-2):e}function xe(e,t,n){return e?e.concat(t).map(function(s,o){return s=Ye(s),!n&&o?"["+s+"]":s}).join(n?".":""):t}function fn(e){return a.isArray(e)&&!e.some(fe)}const dn=a.toFlatObject(a,{},null,function(t){return/^is[A-Z]/.test(t)});function re(e,t,n){if(!a.isObject(e))throw new TypeError("target must be an object");t=t||new FormData,n=a.toFlatObject(n,{metaTokens:!0,dots:!1,indexes:!1},!1,function(m,b){return!a.isUndefined(b[m])});const r=n.metaTokens,s=n.visitor||l,o=n.dots,i=n.indexes,f=(n.Blob||typeof Blob<"u"&&Blob)&&a.isSpecCompliantForm(t);if(!a.isFunction(s))throw new TypeError("visitor must be a function");function u(p){if(p===null)return"";if(a.isDate(p))return p.toISOString();if(!f&&a.isBlob(p))throw new y("Blob is not supported. Use a Buffer instead.");return a.isArrayBuffer(p)||a.isTypedArray(p)?f&&typeof Blob=="function"?new Blob([p]):Buffer.from(p):p}function l(p,m,b){let R=p;if(p&&!b&&typeof p=="object"){if(a.endsWith(m,"{}"))m=r?m:m.slice(0,-2),p=JSON.stringify(p);else if(a.isArray(p)&&fn(p)||(a.isFileList(p)||a.endsWith(m,"[]"))&&(R=a.toArray(p)))return m=Ye(m),R.forEach(function(w,A){!(a.isUndefined(w)||w===null)&&t.append(i===!0?xe([m],A,o):i===null?m:m+"[]",u(w))}),!1}return fe(p)?!0:(t.append(xe(b,m,o),u(p)),!1)}const d=[],E=Object.assign(dn,{defaultVisitor:l,convertValue:u,isVisitable:fe});function g(p,m){if(!a.isUndefined(p)){if(d.indexOf(p)!==-1)throw Error("Circular reference detected in "+m.join("."));d.push(p),a.forEach(p,function(R,O){(!(a.isUndefined(R)||R===null)&&s.call(t,R,a.isString(O)?O.trim():O,m,E))===!0&&g(R,m?m.concat(O):[O])}),d.pop()}}if(!a.isObject(e))throw new TypeError("data must be an object");return g(e),t}function _e(e){const t={"!":"%21","'":"%27","(":"%28",")":"%29","~":"%7E","%20":"+","%00":"\0"};return encodeURIComponent(e).replace(/[!'()~]|%20|%00/g,function(r){return t[r]})}function be(e,t){this._pairs=[],e&&re(e,this,t)}const Qe=be.prototype;Qe.append=function(t,n){this._pairs.push([t,n])};Qe.toString=function(t){const n=t?function(r){return t.call(this,r,_e)}:_e;return this._pairs.map(function(s){return n(s[0])+"="+n(s[1])},"").join("&")};function pn(e){return encodeURIComponent(e).replace(/%3A/gi,":").replace(/%24/g,"$").replace(/%2C/gi,",").replace(/%20/g,"+").replace(/%5B/gi,"[").replace(/%5D/gi,"]")}function Ze(e,t,n){if(!t)return e;const r=n&&n.encode||pn,s=n&&n.serialize;let o;if(s?o=s(t,n):o=a.isURLSearchParams(t)?t.toString():new be(t,n).toString(r),o){const i=e.indexOf("#");i!==-1&&(e=e.slice(0,i)),e+=(e.indexOf("?")===-1?"?":"&")+o}return e}class Ce{constructor(){this.handlers=[]}use(t,n,r){return this.handlers.push({fulfilled:t,rejected:n,synchronous:r?r.synchronous:!1,runWhen:r?r.runWhen:null}),this.handlers.length-1}eject(t){this.handlers[t]&&(this.handlers[t]=null)}clear(){this.handlers&&(this.handlers=[])}forEach(t){a.forEach(this.handlers,function(r){r!==null&&t(r)})}}const et={silentJSONParsing:!0,forcedJSONParsing:!0,clarifyTimeoutError:!1},hn=typeof URLSearchParams<"u"?URLSearchParams:be,mn=typeof FormData<"u"?FormData:null,yn=typeof Blob<"u"?Blob:null,wn={isBrowser:!0,classes:{URLSearchParams:hn,FormData:mn,Blob:yn},protocols:["http","https","file","blob","url","data"]},ge=typeof window<"u"&&typeof document<"u",bn=(e=>ge&&["ReactNative","NativeScript","NS"].indexOf(e)<0)(typeof navigator<"u"&&navigator.product),gn=typeof WorkerGlobalScope<"u"&&self instanceof WorkerGlobalScope&&typeof self.importScripts=="function",En=ge&&window.location.href||"http://localhost",Rn=Object.freeze(Object.defineProperty({__proto__:null,hasBrowserEnv:ge,hasStandardBrowserEnv:bn,hasStandardBrowserWebWorkerEnv:gn,origin:En},Symbol.toStringTag,{value:"Module"})),P={...Rn,...wn};function Sn(e,t){return re(e,new P.classes.URLSearchParams,Object.assign({visitor:function(n,r,s,o){return P.isNode&&a.isBuffer(n)?(this.append(r,n.toString("base64")),!1):o.defaultVisitor.apply(this,arguments)}},t))}function On(e){return a.matchAll(/\w+|\[(\w*)]/g,e).map(t=>t[0]==="[]"?"":t[1]||t[0])}function Tn(e){const t={},n=Object.keys(e);let r;const s=n.length;let o;for(r=0;r<s;r++)o=n[r],t[o]=e[o];return t}function tt(e){function t(n,r,s,o){let i=n[o++];if(i==="__proto__")return!0;const c=Number.isFinite(+i),f=o>=n.length;return i=!i&&a.isArray(s)?s.length:i,f?(a.hasOwnProp(s,i)?s[i]=[s[i],r]:s[i]=r,!c):((!s[i]||!a.isObject(s[i]))&&(s[i]=[]),t(n,r,s[i],o)&&a.isArray(s[i])&&(s[i]=Tn(s[i])),!c)}if(a.isFormData(e)&&a.isFunction(e.entries)){const n={};return a.forEachEntry(e,(r,s)=>{t(On(r),s,n,0)}),n}return null}function An(e,t,n){if(a.isString(e))try{return(t||JSON.parse)(e),a.trim(e)}catch(r){if(r.name!=="SyntaxError")throw r}return(n||JSON.stringify)(e)}const K={transitional:et,adapter:["xhr","http","fetch"],transformRequest:[function(t,n){const r=n.getContentType()||"",s=r.indexOf("application/json")>-1,o=a.isObject(t);if(o&&a.isHTMLForm(t)&&(t=new FormData(t)),a.isFormData(t))return s?JSON.stringify(tt(t)):t;if(a.isArrayBuffer(t)||a.isBuffer(t)||a.isStream(t)||a.isFile(t)||a.isBlob(t)||a.isReadableStream(t))return t;if(a.isArrayBufferView(t))return t.buffer;if(a.isURLSearchParams(t))return n.setContentType("application/x-www-form-urlencoded;charset=utf-8",!1),t.toString();let c;if(o){if(r.indexOf("application/x-www-form-urlencoded")>-1)return Sn(t,this.formSerializer).toString();if((c=a.isFileList(t))||r.indexOf("multipart/form-data")>-1){const f=this.env&&this.env.FormData;return re(c?{"files[]":t}:t,f&&new f,this.formSerializer)}}return o||s?(n.setContentType("application/json",!1),An(t)):t}],transformResponse:[function(t){const n=this.transitional||K.transitional,r=n&&n.forcedJSONParsing,s=this.responseType==="json";if(a.isResponse(t)||a.isReadableStream(t))return t;if(t&&a.isString(t)&&(r&&!this.responseType||s)){const i=!(n&&n.silentJSONParsing)&&s;try{return JSON.parse(t)}catch(c){if(i)throw c.name==="SyntaxError"?y.from(c,y.ERR_BAD_RESPONSE,this,null,this.response):c}}return t}],timeout:0,xsrfCookieName:"XSRF-TOKEN",xsrfHeaderName:"X-XSRF-TOKEN",maxContentLength:-1,maxBodyLength:-1,env:{FormData:P.classes.FormData,Blob:P.classes.Blob},validateStatus:function(t){return t>=200&&t<300},headers:{common:{Accept:"application/json, text/plain, */*","Content-Type":void 0}}};a.forEach(["delete","get","head","post","put","patch"],e=>{K.headers[e]={}});const xn=a.toObjectSet(["age","authorization","content-length","content-type","etag","expires","from","host","if-modified-since","if-unmodified-since","last-modified","location","max-forwards","proxy-authorization","referer","retry-after","user-agent"]),_n=e=>{const t={};let n,r,s;return e&&e.split(`
`).forEach(function(i){s=i.indexOf(":"),n=i.substring(0,s).trim().toLowerCase(),r=i.substring(s+1).trim(),!(!n||t[n]&&xn[n])&&(n==="set-cookie"?t[n]?t[n].push(r):t[n]=[r]:t[n]=t[n]?t[n]+", "+r:r)}),t},Pe=Symbol("internals");function z(e){return e&&String(e).trim().toLowerCase()}function Y(e){return e===!1||e==null?e:a.isArray(e)?e.map(Y):String(e)}function Cn(e){const t=Object.create(null),n=/([^\s,;=]+)\s*(?:=\s*([^,;]+))?/g;let r;for(;r=n.exec(e);)t[r[1]]=r[2];return t}const Pn=e=>/^[-_a-zA-Z0-9^`|~,!#$%&'*+.]+$/.test(e.trim());function ae(e,t,n,r,s){if(a.isFunction(r))return r.call(this,t,n);if(s&&(t=n),!!a.isString(t)){if(a.isString(r))return t.indexOf(r)!==-1;if(a.isRegExp(r))return r.test(t)}}function Nn(e){return e.trim().toLowerCase().replace(/([a-z\d])(\w*)/g,(t,n,r)=>n.toUpperCase()+r)}function Ln(e,t){const n=a.toCamelCase(" "+t);["get","set","has"].forEach(r=>{Object.defineProperty(e,r+n,{value:function(s,o,i){return this[r].call(this,t,s,o,i)},configurable:!0})})}class x{constructor(t){t&&this.set(t)}set(t,n,r){const s=this;function o(c,f,u){const l=z(f);if(!l)throw new Error("header name must be a non-empty string");const d=a.findKey(s,l);(!d||s[d]===void 0||u===!0||u===void 0&&s[d]!==!1)&&(s[d||f]=Y(c))}const i=(c,f)=>a.forEach(c,(u,l)=>o(u,l,f));if(a.isPlainObject(t)||t instanceof this.constructor)i(t,n);else if(a.isString(t)&&(t=t.trim())&&!Pn(t))i(_n(t),n);else if(a.isHeaders(t))for(const[c,f]of t.entries())o(f,c,r);else t!=null&&o(n,t,r);return this}get(t,n){if(t=z(t),t){const r=a.findKey(this,t);if(r){const s=this[r];if(!n)return s;if(n===!0)return Cn(s);if(a.isFunction(n))return n.call(this,s,r);if(a.isRegExp(n))return n.exec(s);throw new TypeError("parser must be boolean|regexp|function")}}}has(t,n){if(t=z(t),t){const r=a.findKey(this,t);return!!(r&&this[r]!==void 0&&(!n||ae(this,this[r],r,n)))}return!1}delete(t,n){const r=this;let s=!1;function o(i){if(i=z(i),i){const c=a.findKey(r,i);c&&(!n||ae(r,r[c],c,n))&&(delete r[c],s=!0)}}return a.isArray(t)?t.forEach(o):o(t),s}clear(t){const n=Object.keys(this);let r=n.length,s=!1;for(;r--;){const o=n[r];(!t||ae(this,this[o],o,t,!0))&&(delete this[o],s=!0)}return s}normalize(t){const n=this,r={};return a.forEach(this,(s,o)=>{const i=a.findKey(r,o);if(i){n[i]=Y(s),delete n[o];return}const c=t?Nn(o):String(o).trim();c!==o&&delete n[o],n[c]=Y(s),r[c]=!0}),this}concat(...t){return this.constructor.concat(this,...t)}toJSON(t){const n=Object.create(null);return a.forEach(this,(r,s)=>{r!=null&&r!==!1&&(n[s]=t&&a.isArray(r)?r.join(", "):r)}),n}[Symbol.iterator](){return Object.entries(this.toJSON())[Symbol.iterator]()}toString(){return Object.entries(this.toJSON()).map(([t,n])=>t+": "+n).join(`
`)}get[Symbol.toStringTag](){return"AxiosHeaders"}static from(t){return t instanceof this?t:new this(t)}static concat(t,...n){const r=new this(t);return n.forEach(s=>r.set(s)),r}static accessor(t){const r=(this[Pe]=this[Pe]={accessors:{}}).accessors,s=this.prototype;function o(i){const c=z(i);r[c]||(Ln(s,i),r[c]=!0)}return a.isArray(t)?t.forEach(o):o(t),this}}x.accessor(["Content-Type","Content-Length","Accept","Accept-Encoding","User-Agent","Authorization"]);a.reduceDescriptors(x.prototype,({value:e},t)=>{let n=t[0].toUpperCase()+t.slice(1);return{get:()=>e,set(r){this[n]=r}}});a.freezeMethods(x);function ce(e,t){const n=this||K,r=t||n,s=x.from(r.headers);let o=r.data;return a.forEach(e,function(c){o=c.call(n,o,s.normalize(),t?t.status:void 0)}),s.normalize(),o}function nt(e){return!!(e&&e.__CANCEL__)}function I(e,t,n){y.call(this,e??"canceled",y.ERR_CANCELED,t,n),this.name="CanceledError"}a.inherits(I,y,{__CANCEL__:!0});function rt(e,t,n){const r=n.config.validateStatus;!n.status||!r||r(n.status)?e(n):t(new y("Request failed with status code "+n.status,[y.ERR_BAD_REQUEST,y.ERR_BAD_RESPONSE][Math.floor(n.status/100)-4],n.config,n.request,n))}function Fn(e){const t=/^([-+\w]{1,25})(:?\/\/|:)/.exec(e);return t&&t[1]||""}function Dn(e,t){e=e||10;const n=new Array(e),r=new Array(e);let s=0,o=0,i;return t=t!==void 0?t:1e3,function(f){const u=Date.now(),l=r[o];i||(i=u),n[s]=f,r[s]=u;let d=o,E=0;for(;d!==s;)E+=n[d++],d=d%e;if(s=(s+1)%e,s===o&&(o=(o+1)%e),u-i<t)return;const g=l&&u-l;return g?Math.round(E*1e3/g):void 0}}function Bn(e,t){let n=0;const r=1e3/t;let s=null;return function(){const i=this===!0,c=Date.now();if(i||c-n>r)return s&&(clearTimeout(s),s=null),n=c,e.apply(null,arguments);s||(s=setTimeout(()=>(s=null,n=Date.now(),e.apply(null,arguments)),r-(c-n)))}}const Q=(e,t,n=3)=>{let r=0;const s=Dn(50,250);return Bn(o=>{const i=o.loaded,c=o.lengthComputable?o.total:void 0,f=i-r,u=s(f),l=i<=c;r=i;const d={loaded:i,total:c,progress:c?i/c:void 0,bytes:f,rate:u||void 0,estimated:u&&c&&l?(c-i)/u:void 0,event:o,lengthComputable:c!=null};d[t?"download":"upload"]=!0,e(d)},n)},Un=P.hasStandardBrowserEnv?function(){const t=/(msie|trident)/i.test(navigator.userAgent),n=document.createElement("a");let r;function s(o){let i=o;return t&&(n.setAttribute("href",i),i=n.href),n.setAttribute("href",i),{href:n.href,protocol:n.protocol?n.protocol.replace(/:$/,""):"",host:n.host,search:n.search?n.search.replace(/^\?/,""):"",hash:n.hash?n.hash.replace(/^#/,""):"",hostname:n.hostname,port:n.port,pathname:n.pathname.charAt(0)==="/"?n.pathname:"/"+n.pathname}}return r=s(window.location.href),function(i){const c=a.isString(i)?s(i):i;return c.protocol===r.protocol&&c.host===r.host}}():function(){return function(){return!0}}(),kn=P.hasStandardBrowserEnv?{write(e,t,n,r,s,o){const i=[e+"="+encodeURIComponent(t)];a.isNumber(n)&&i.push("expires="+new Date(n).toGMTString()),a.isString(r)&&i.push("path="+r),a.isString(s)&&i.push("domain="+s),o===!0&&i.push("secure"),document.cookie=i.join("; ")},read(e){const t=document.cookie.match(new RegExp("(^|;\\s*)("+e+")=([^;]*)"));return t?decodeURIComponent(t[3]):null},remove(e){this.write(e,"",Date.now()-864e5)}}:{write(){},read(){return null},remove(){}};function vn(e){return/^([a-z][a-z\d+\-.]*:)?\/\//i.test(e)}function jn(e,t){return t?e.replace(/\/?\/$/,"")+"/"+t.replace(/^\/+/,""):e}function st(e,t){return e&&!vn(t)?jn(e,t):t}const Ne=e=>e instanceof x?{...e}:e;function v(e,t){t=t||{};const n={};function r(u,l,d){return a.isPlainObject(u)&&a.isPlainObject(l)?a.merge.call({caseless:d},u,l):a.isPlainObject(l)?a.merge({},l):a.isArray(l)?l.slice():l}function s(u,l,d){if(a.isUndefined(l)){if(!a.isUndefined(u))return r(void 0,u,d)}else return r(u,l,d)}function o(u,l){if(!a.isUndefined(l))return r(void 0,l)}function i(u,l){if(a.isUndefined(l)){if(!a.isUndefined(u))return r(void 0,u)}else return r(void 0,l)}function c(u,l,d){if(d in t)return r(u,l);if(d in e)return r(void 0,u)}const f={url:o,method:o,data:o,baseURL:i,transformRequest:i,transformResponse:i,paramsSerializer:i,timeout:i,timeoutMessage:i,withCredentials:i,withXSRFToken:i,adapter:i,responseType:i,xsrfCookieName:i,xsrfHeaderName:i,onUploadProgress:i,onDownloadProgress:i,decompress:i,maxContentLength:i,maxBodyLength:i,beforeRedirect:i,transport:i,httpAgent:i,httpsAgent:i,cancelToken:i,socketPath:i,responseEncoding:i,validateStatus:c,headers:(u,l)=>s(Ne(u),Ne(l),!0)};return a.forEach(Object.keys(Object.assign({},e,t)),function(l){const d=f[l]||s,E=d(e[l],t[l],l);a.isUndefined(E)&&d!==c||(n[l]=E)}),n}const ot=e=>{const t=v({},e);let{data:n,withXSRFToken:r,xsrfHeaderName:s,xsrfCookieName:o,headers:i,auth:c}=t;t.headers=i=x.from(i),t.url=Ze(st(t.baseURL,t.url),e.params,e.paramsSerializer),c&&i.set("Authorization","Basic "+btoa((c.username||"")+":"+(c.password?unescape(encodeURIComponent(c.password)):"")));let f;if(a.isFormData(n)){if(P.hasStandardBrowserEnv||P.hasStandardBrowserWebWorkerEnv)i.setContentType(void 0);else if((f=i.getContentType())!==!1){const[u,...l]=f?f.split(";").map(d=>d.trim()).filter(Boolean):[];i.setContentType([u||"multipart/form-data",...l].join("; "))}}if(P.hasStandardBrowserEnv&&(r&&a.isFunction(r)&&(r=r(t)),r||r!==!1&&Un(t.url))){const u=s&&o&&kn.read(o);u&&i.set(s,u)}return t},qn=typeof XMLHttpRequest<"u",Mn=qn&&function(e){return new Promise(function(n,r){const s=ot(e);let o=s.data;const i=x.from(s.headers).normalize();let{responseType:c}=s,f;function u(){s.cancelToken&&s.cancelToken.unsubscribe(f),s.signal&&s.signal.removeEventListener("abort",f)}let l=new XMLHttpRequest;l.open(s.method.toUpperCase(),s.url,!0),l.timeout=s.timeout;function d(){if(!l)return;const g=x.from("getAllResponseHeaders"in l&&l.getAllResponseHeaders()),m={data:!c||c==="text"||c==="json"?l.responseText:l.response,status:l.status,statusText:l.statusText,headers:g,config:e,request:l};rt(function(R){n(R),u()},function(R){r(R),u()},m),l=null}"onloadend"in l?l.onloadend=d:l.onreadystatechange=function(){!l||l.readyState!==4||l.status===0&&!(l.responseURL&&l.responseURL.indexOf("file:")===0)||setTimeout(d)},l.onabort=function(){l&&(r(new y("Request aborted",y.ECONNABORTED,s,l)),l=null)},l.onerror=function(){r(new y("Network Error",y.ERR_NETWORK,s,l)),l=null},l.ontimeout=function(){let p=s.timeout?"timeout of "+s.timeout+"ms exceeded":"timeout exceeded";const m=s.transitional||et;s.timeoutErrorMessage&&(p=s.timeoutErrorMessage),r(new y(p,m.clarifyTimeoutError?y.ETIMEDOUT:y.ECONNABORTED,s,l)),l=null},o===void 0&&i.setContentType(null),"setRequestHeader"in l&&a.forEach(i.toJSON(),function(p,m){l.setRequestHeader(m,p)}),a.isUndefined(s.withCredentials)||(l.withCredentials=!!s.withCredentials),c&&c!=="json"&&(l.responseType=s.responseType),typeof s.onDownloadProgress=="function"&&l.addEventListener("progress",Q(s.onDownloadProgress,!0)),typeof s.onUploadProgress=="function"&&l.upload&&l.upload.addEventListener("progress",Q(s.onUploadProgress)),(s.cancelToken||s.signal)&&(f=g=>{l&&(r(!g||g.type?new I(null,e,l):g),l.abort(),l=null)},s.cancelToken&&s.cancelToken.subscribe(f),s.signal&&(s.signal.aborted?f():s.signal.addEventListener("abort",f)));const E=Fn(s.url);if(E&&P.protocols.indexOf(E)===-1){r(new y("Unsupported protocol "+E+":",y.ERR_BAD_REQUEST,e));return}l.send(o||null)})},Hn=(e,t)=>{let n=new AbortController,r;const s=function(f){if(!r){r=!0,i();const u=f instanceof Error?f:this.reason;n.abort(u instanceof y?u:new I(u instanceof Error?u.message:u))}};let o=t&&setTimeout(()=>{s(new y(`timeout ${t} of ms exceeded`,y.ETIMEDOUT))},t);const i=()=>{e&&(o&&clearTimeout(o),o=null,e.forEach(f=>{f&&(f.removeEventListener?f.removeEventListener("abort",s):f.unsubscribe(s))}),e=null)};e.forEach(f=>f&&f.addEventListener&&f.addEventListener("abort",s));const{signal:c}=n;return c.unsubscribe=i,[c,()=>{o&&clearTimeout(o),o=null}]},In=function*(e,t){let n=e.byteLength;if(!t||n<t){yield e;return}let r=0,s;for(;r<n;)s=r+t,yield e.slice(r,s),r=s},zn=async function*(e,t,n){for await(const r of e)yield*In(ArrayBuffer.isView(r)?r:await n(String(r)),t)},Le=(e,t,n,r,s)=>{const o=zn(e,t,s);let i=0;return new ReadableStream({type:"bytes",async pull(c){const{done:f,value:u}=await o.next();if(f){c.close(),r();return}let l=u.byteLength;n&&n(i+=l),c.enqueue(new Uint8Array(u))},cancel(c){return r(c),o.return()}},{highWaterMark:2})},Fe=(e,t)=>{const n=e!=null;return r=>setTimeout(()=>t({lengthComputable:n,total:e,loaded:r}))},se=typeof fetch=="function"&&typeof Request=="function"&&typeof Response=="function",it=se&&typeof ReadableStream=="function",de=se&&(typeof TextEncoder=="function"?(e=>t=>e.encode(t))(new TextEncoder):async e=>new Uint8Array(await new Response(e).arrayBuffer())),$n=it&&(()=>{let e=!1;const t=new Request(P.origin,{body:new ReadableStream,method:"POST",get duplex(){return e=!0,"half"}}).headers.has("Content-Type");return e&&!t})(),De=64*1024,pe=it&&!!(()=>{try{return a.isReadableStream(new Response("").body)}catch{}})(),Z={stream:pe&&(e=>e.body)};se&&(e=>{["text","arrayBuffer","blob","formData","stream"].forEach(t=>{!Z[t]&&(Z[t]=a.isFunction(e[t])?n=>n[t]():(n,r)=>{throw new y(`Response type '${t}' is not supported`,y.ERR_NOT_SUPPORT,r)})})})(new Response);const Jn=async e=>{if(e==null)return 0;if(a.isBlob(e))return e.size;if(a.isSpecCompliantForm(e))return(await new Request(e).arrayBuffer()).byteLength;if(a.isArrayBufferView(e))return e.byteLength;if(a.isURLSearchParams(e)&&(e=e+""),a.isString(e))return(await de(e)).byteLength},Vn=async(e,t)=>{const n=a.toFiniteNumber(e.getContentLength());return n??Jn(t)},Wn=se&&(async e=>{let{url:t,method:n,data:r,signal:s,cancelToken:o,timeout:i,onDownloadProgress:c,onUploadProgress:f,responseType:u,headers:l,withCredentials:d="same-origin",fetchOptions:E}=ot(e);u=u?(u+"").toLowerCase():"text";let[g,p]=s||o||i?Hn([s,o],i):[],m,b;const R=()=>{!m&&setTimeout(()=>{g&&g.unsubscribe()}),m=!0};let O;try{if(f&&$n&&n!=="get"&&n!=="head"&&(O=await Vn(l,r))!==0){let T=new Request(t,{method:"POST",body:r,duplex:"half"}),L;a.isFormData(r)&&(L=T.headers.get("content-type"))&&l.setContentType(L),T.body&&(r=Le(T.body,De,Fe(O,Q(f)),null,de))}a.isString(d)||(d=d?"cors":"omit"),b=new Request(t,{...E,signal:g,method:n.toUpperCase(),headers:l.normalize().toJSON(),body:r,duplex:"half",withCredentials:d});let w=await fetch(b);const A=pe&&(u==="stream"||u==="response");if(pe&&(c||A)){const T={};["status","statusText","headers"].forEach(X=>{T[X]=w[X]});const L=a.toFiniteNumber(w.headers.get("content-length"));w=new Response(Le(w.body,De,c&&Fe(L,Q(c,!0)),A&&R,de),T)}u=u||"text";let U=await Z[a.findKey(Z,u)||"text"](w,e);return!A&&R(),p&&p(),await new Promise((T,L)=>{rt(T,L,{data:U,headers:x.from(w.headers),status:w.status,statusText:w.statusText,config:e,request:b})})}catch(w){throw R(),w&&w.name==="TypeError"&&/fetch/i.test(w.message)?Object.assign(new y("Network Error",y.ERR_NETWORK,e,b),{cause:w.cause||w}):y.from(w,w&&w.code,e,b)}}),he={http:un,xhr:Mn,fetch:Wn};a.forEach(he,(e,t)=>{if(e){try{Object.defineProperty(e,"name",{value:t})}catch{}Object.defineProperty(e,"adapterName",{value:t})}});const Be=e=>`- ${e}`,Kn=e=>a.isFunction(e)||e===null||e===!1,at={getAdapter:e=>{e=a.isArray(e)?e:[e];const{length:t}=e;let n,r;const s={};for(let o=0;o<t;o++){n=e[o];let i;if(r=n,!Kn(n)&&(r=he[(i=String(n)).toLowerCase()],r===void 0))throw new y(`Unknown adapter '${i}'`);if(r)break;s[i||"#"+o]=r}if(!r){const o=Object.entries(s).map(([c,f])=>`adapter ${c} `+(f===!1?"is not supported by the environment":"is not available in the build"));let i=t?o.length>1?`since :
`+o.map(Be).join(`
`):" "+Be(o[0]):"as no adapter specified";throw new y("There is no suitable adapter to dispatch the request "+i,"ERR_NOT_SUPPORT")}return r},adapters:he};function le(e){if(e.cancelToken&&e.cancelToken.throwIfRequested(),e.signal&&e.signal.aborted)throw new I(null,e)}function Ue(e){return le(e),e.headers=x.from(e.headers),e.data=ce.call(e,e.transformRequest),["post","put","patch"].indexOf(e.method)!==-1&&e.headers.setContentType("application/x-www-form-urlencoded",!1),at.getAdapter(e.adapter||K.adapter)(e).then(function(r){return le(e),r.data=ce.call(e,e.transformResponse,r),r.headers=x.from(r.headers),r},function(r){return nt(r)||(le(e),r&&r.response&&(r.response.data=ce.call(e,e.transformResponse,r.response),r.response.headers=x.from(r.response.headers))),Promise.reject(r)})}const ct="1.7.2",Ee={};["object","boolean","number","function","string","symbol"].forEach((e,t)=>{Ee[e]=function(r){return typeof r===e||"a"+(t<1?"n ":" ")+e}});const ke={};Ee.transitional=function(t,n,r){function s(o,i){return"[Axios v"+ct+"] Transitional option '"+o+"'"+i+(r?". "+r:"")}return(o,i,c)=>{if(t===!1)throw new y(s(i," has been removed"+(n?" in "+n:"")),y.ERR_DEPRECATED);return n&&!ke[i]&&(ke[i]=!0,console.warn(s(i," has been deprecated since v"+n+" and will be removed in the near future"))),t?t(o,i,c):!0}};function Xn(e,t,n){if(typeof e!="object")throw new y("options must be an object",y.ERR_BAD_OPTION_VALUE);const r=Object.keys(e);let s=r.length;for(;s-- >0;){const o=r[s],i=t[o];if(i){const c=e[o],f=c===void 0||i(c,o,e);if(f!==!0)throw new y("option "+o+" must be "+f,y.ERR_BAD_OPTION_VALUE);continue}if(n!==!0)throw new y("Unknown option "+o,y.ERR_BAD_OPTION)}}const me={assertOptions:Xn,validators:Ee},F=me.validators;class k{constructor(t){this.defaults=t,this.interceptors={request:new Ce,response:new Ce}}async request(t,n){try{return await this._request(t,n)}catch(r){if(r instanceof Error){let s;Error.captureStackTrace?Error.captureStackTrace(s={}):s=new Error;const o=s.stack?s.stack.replace(/^.+\n/,""):"";try{r.stack?o&&!String(r.stack).endsWith(o.replace(/^.+\n.+\n/,""))&&(r.stack+=`
`+o):r.stack=o}catch{}}throw r}}_request(t,n){typeof t=="string"?(n=n||{},n.url=t):n=t||{},n=v(this.defaults,n);const{transitional:r,paramsSerializer:s,headers:o}=n;r!==void 0&&me.assertOptions(r,{silentJSONParsing:F.transitional(F.boolean),forcedJSONParsing:F.transitional(F.boolean),clarifyTimeoutError:F.transitional(F.boolean)},!1),s!=null&&(a.isFunction(s)?n.paramsSerializer={serialize:s}:me.assertOptions(s,{encode:F.function,serialize:F.function},!0)),n.method=(n.method||this.defaults.method||"get").toLowerCase();let i=o&&a.merge(o.common,o[n.method]);o&&a.forEach(["delete","get","head","post","put","patch","common"],p=>{delete o[p]}),n.headers=x.concat(i,o);const c=[];let f=!0;this.interceptors.request.forEach(function(m){typeof m.runWhen=="function"&&m.runWhen(n)===!1||(f=f&&m.synchronous,c.unshift(m.fulfilled,m.rejected))});const u=[];this.interceptors.response.forEach(function(m){u.push(m.fulfilled,m.rejected)});let l,d=0,E;if(!f){const p=[Ue.bind(this),void 0];for(p.unshift.apply(p,c),p.push.apply(p,u),E=p.length,l=Promise.resolve(n);d<E;)l=l.then(p[d++],p[d++]);return l}E=c.length;let g=n;for(d=0;d<E;){const p=c[d++],m=c[d++];try{g=p(g)}catch(b){m.call(this,b);break}}try{l=Ue.call(this,g)}catch(p){return Promise.reject(p)}for(d=0,E=u.length;d<E;)l=l.then(u[d++],u[d++]);return l}getUri(t){t=v(this.defaults,t);const n=st(t.baseURL,t.url);return Ze(n,t.params,t.paramsSerializer)}}a.forEach(["delete","get","head","options"],function(t){k.prototype[t]=function(n,r){return this.request(v(r||{},{method:t,url:n,data:(r||{}).data}))}});a.forEach(["post","put","patch"],function(t){function n(r){return function(o,i,c){return this.request(v(c||{},{method:t,headers:r?{"Content-Type":"multipart/form-data"}:{},url:o,data:i}))}}k.prototype[t]=n(),k.prototype[t+"Form"]=n(!0)});class Re{constructor(t){if(typeof t!="function")throw new TypeError("executor must be a function.");let n;this.promise=new Promise(function(o){n=o});const r=this;this.promise.then(s=>{if(!r._listeners)return;let o=r._listeners.length;for(;o-- >0;)r._listeners[o](s);r._listeners=null}),this.promise.then=s=>{let o;const i=new Promise(c=>{r.subscribe(c),o=c}).then(s);return i.cancel=function(){r.unsubscribe(o)},i},t(function(o,i,c){r.reason||(r.reason=new I(o,i,c),n(r.reason))})}throwIfRequested(){if(this.reason)throw this.reason}subscribe(t){if(this.reason){t(this.reason);return}this._listeners?this._listeners.push(t):this._listeners=[t]}unsubscribe(t){if(!this._listeners)return;const n=this._listeners.indexOf(t);n!==-1&&this._listeners.splice(n,1)}static source(){let t;return{token:new Re(function(s){t=s}),cancel:t}}}function Gn(e){return function(n){return e.apply(null,n)}}function Yn(e){return a.isObject(e)&&e.isAxiosError===!0}const ye={Continue:100,SwitchingProtocols:101,Processing:102,EarlyHints:103,Ok:200,Created:201,Accepted:202,NonAuthoritativeInformation:203,NoContent:204,ResetContent:205,PartialContent:206,MultiStatus:207,AlreadyReported:208,ImUsed:226,MultipleChoices:300,MovedPermanently:301,Found:302,SeeOther:303,NotModified:304,UseProxy:305,Unused:306,TemporaryRedirect:307,PermanentRedirect:308,BadRequest:400,Unauthorized:401,PaymentRequired:402,Forbidden:403,NotFound:404,MethodNotAllowed:405,NotAcceptable:406,ProxyAuthenticationRequired:407,RequestTimeout:408,Conflict:409,Gone:410,LengthRequired:411,PreconditionFailed:412,PayloadTooLarge:413,UriTooLong:414,UnsupportedMediaType:415,RangeNotSatisfiable:416,ExpectationFailed:417,ImATeapot:418,MisdirectedRequest:421,UnprocessableEntity:422,Locked:423,FailedDependency:424,TooEarly:425,UpgradeRequired:426,PreconditionRequired:428,TooManyRequests:429,RequestHeaderFieldsTooLarge:431,UnavailableForLegalReasons:451,InternalServerError:500,NotImplemented:501,BadGateway:502,ServiceUnavailable:503,GatewayTimeout:504,HttpVersionNotSupported:505,VariantAlsoNegotiates:506,InsufficientStorage:507,LoopDetected:508,NotExtended:510,NetworkAuthenticationRequired:511};Object.entries(ye).forEach(([e,t])=>{ye[t]=e});function lt(e){const t=new k(e),n=He(k.prototype.request,t);return a.extend(n,k.prototype,t,{allOwnKeys:!0}),a.extend(n,t,null,{allOwnKeys:!0}),n.create=function(s){return lt(v(e,s))},n}const S=lt(K);S.Axios=k;S.CanceledError=I;S.CancelToken=Re;S.isCancel=nt;S.VERSION=ct;S.toFormData=re;S.AxiosError=y;S.Cancel=S.CanceledError;S.all=function(t){return Promise.all(t)};S.spread=Gn;S.isAxiosError=Yn;S.mergeConfig=v;S.AxiosHeaders=x;S.formToJSON=e=>tt(a.isHTMLForm(e)?new FormData(e):e);S.getAdapter=at.getAdapter;S.HttpStatusCode=ye;S.default=S;var Qn=V('<div id="wrapper" class="svelte-dwbzmc"><iframe title="da cameras" id="cams" src="http://74.208.238.87:8889/ptz-alv/?controls=0" frameborder="0" allow="autoplay; fullscreen" allowfullscreen class="svelte-dwbzmc"></iframe></div>'),Zn=V('<div id="video-container" class="svelte-dwbzmc"><span class="svelte-dwbzmc"> </span> <div><button>Get Data</button></div> <!> <h1 class="svelte-dwbzmc">Vite + Svelte</h1></div>');function er(e,t){ve(t,!1);let n=D({x:0,y:0});function r(b){var R=b.target.getBoundingClientRect();Se(n,h(n).x=Math.round(b.clientX-R.left)),Se(n,h(n).y=Math.round(b.clientY-R.top))}let s=!1,o;const i={border:"2px solid red",backgroundColor:"rgba(255, 0, 0, 0.5)"};function c(b){o=b}async function f(){o&&S.post("/draw",{x:o.x,y:o.y,width:o.width,height:o.height}).then(function(b){console.log(b)}).catch(function(b){console.log(b)})}Me();var u=Zn(),l=q(u),d=q(l);M(l);var E=B(B(l,!0)),g=q(E);M(E);var p=B(E,!0);p.nodeValue="  ";var m=B(p);Tt(m,{onUpdateRectangle:c,rectangleStyle:i,drawn:s,children:(b,R)=>{var O=Qn();q(O),M(O),Oe("click",O,r,!1),$(b,O)},$$slots:{default:!0},$$legacy:!0}),B(B(m,!0)),M(u),qe(()=>pt(d,`${h(n).x??""} x ${h(n).y??""}`)),Oe("click",g,f,!1),$(e,u),je()}var tr=V("<main><!></main>");function lr(e){var t=tr(),n=q(t);er(n,{$$legacy:!0}),M(t),$(e,t)}export{lr as component,cr as universal};
