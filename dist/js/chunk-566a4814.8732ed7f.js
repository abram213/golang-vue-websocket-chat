(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-566a4814"],{"16b7":function(t,e,i){"use strict";var s=i("2b0e");e["a"]=s["a"].extend().extend({name:"delayable",props:{openDelay:{type:[Number,String],default:0},closeDelay:{type:[Number,String],default:0}},data:()=>({openTimeout:void 0,closeTimeout:void 0}),methods:{clearDelay(){clearTimeout(this.openTimeout),clearTimeout(this.closeTimeout)},runDelay(t,e){this.clearDelay();const i=parseInt(this[`${t}Delay`],10);this[`${t}Timeout`]=setTimeout(e||(()=>{this.isActive={open:!0,close:!1}[t]}),i)}}})},"21be":function(t,e,i){"use strict";var s=i("2b0e"),n=i("80d2");e["a"]=s["a"].extend().extend({name:"stackable",data(){return{stackElement:null,stackExclude:null,stackMinZIndex:0,isActive:!1}},computed:{activeZIndex(){if("undefined"===typeof window)return 0;const t=this.stackElement||this.$refs.content,e=this.isActive?this.getMaxZIndex(this.stackExclude||[t])+2:Object(n["m"])(t);return null==e?e:parseInt(e)}},methods:{getMaxZIndex(t=[]){const e=this.$el,i=[this.stackMinZIndex,Object(n["m"])(e)],s=[...document.getElementsByClassName("v-menu__content--active"),...document.getElementsByClassName("v-dialog__content--active")];for(let a=0;a<s.length;a++)t.includes(s[a])||i.push(Object(n["m"])(s[a]));return Math.max(...i)}}})},"2fa4":function(t,e,i){"use strict";i("20f6");var s=i("80d2");e["a"]=Object(s["g"])("spacer","div","v-spacer")},"368e":function(t,e,i){},"3c93":function(t,e,i){},"4ad4":function(t,e,i){"use strict";var s=i("16b7"),n=i("f2e7"),a=i("58df"),o=i("80d2"),r=i("d9bd");const l=Object(a["a"])(s["a"],n["a"]);e["a"]=l.extend({name:"activatable",props:{activator:{default:null,validator:t=>["string","object"].includes(typeof t)},disabled:Boolean,internalActivator:Boolean,openOnHover:Boolean},data:()=>({activatorElement:null,activatorNode:[],events:["click","mouseenter","mouseleave"],listeners:{}}),watch:{activator:"resetActivator",openOnHover:"resetActivator"},mounted(){const t=Object(o["l"])(this,"activator",!0);t&&["v-slot","normal"].includes(t)&&Object(r["b"])('The activator slot must be bound, try \'<template v-slot:activator="{ on }"><v-btn v-on="on">\'',this),this.addActivatorEvents()},beforeDestroy(){this.removeActivatorEvents()},methods:{addActivatorEvents(){if(!this.activator||this.disabled||!this.getActivator())return;this.listeners=this.genActivatorListeners();const t=Object.keys(this.listeners);for(const e of t)this.getActivator().addEventListener(e,this.listeners[e])},genActivator(){const t=Object(o["k"])(this,"activator",Object.assign(this.getValueProxy(),{on:this.genActivatorListeners(),attrs:this.genActivatorAttributes()}))||[];return this.activatorNode=t,t},genActivatorAttributes(){return{role:"button","aria-haspopup":!0,"aria-expanded":String(this.isActive)}},genActivatorListeners(){if(this.disabled)return{};const t={};return this.openOnHover?(t.mouseenter=t=>{this.getActivator(t),this.runDelay("open")},t.mouseleave=t=>{this.getActivator(t),this.runDelay("close")}):t.click=t=>{const e=this.getActivator(t);e&&e.focus(),t.stopPropagation(),this.isActive=!this.isActive},t},getActivator(t){if(this.activatorElement)return this.activatorElement;let e=null;if(this.activator){const t=this.internalActivator?this.$el:document;e="string"===typeof this.activator?t.querySelector(this.activator):this.activator.$el?this.activator.$el:this.activator}else if(1===this.activatorNode.length||this.activatorNode.length&&!t){const t=this.activatorNode[0].componentInstance;e=t&&t.$options.mixins&&t.$options.mixins.some(t=>t.options&&["activatable","menuable"].includes(t.options.name))?t.getActivator():this.activatorNode[0].elm}else t&&(e=t.currentTarget||t.target);return this.activatorElement=e,this.activatorElement},getContentSlot(){return Object(o["k"])(this,"default",this.getValueProxy(),!0)},getValueProxy(){const t=this;return{get value(){return t.isActive},set value(e){t.isActive=e}}},removeActivatorEvents(){if(!this.activator||!this.activatorElement)return;const t=Object.keys(this.listeners);for(const e of t)this.activatorElement.removeEventListener(e,this.listeners[e]);this.listeners={}},resetActivator(){this.removeActivatorEvents(),this.activatorElement=null,this.getActivator(),this.addActivatorEvents()}}})},"4bd4":function(t,e,i){"use strict";var s=i("58df"),n=i("7e2b"),a=i("3206");e["a"]=Object(s["a"])(n["a"],Object(a["b"])("form")).extend({name:"v-form",inheritAttrs:!1,props:{lazyValidation:Boolean,value:Boolean},data:()=>({inputs:[],watchers:[],errorBag:{}}),watch:{errorBag:{handler(t){const e=Object.values(t).includes(!0);this.$emit("input",!e)},deep:!0,immediate:!0}},methods:{watchInput(t){const e=t=>t.$watch("hasError",e=>{this.$set(this.errorBag,t._uid,e)},{immediate:!0}),i={_uid:t._uid,valid:()=>{},shouldValidate:()=>{}};return this.lazyValidation?i.shouldValidate=t.$watch("shouldValidate",s=>{s&&(this.errorBag.hasOwnProperty(t._uid)||(i.valid=e(t)))}):i.valid=e(t),i},validate(){return 0===this.inputs.filter(t=>!t.validate(!0)).length},reset(){this.inputs.forEach(t=>t.reset()),this.resetErrorBag()},resetErrorBag(){this.lazyValidation&&setTimeout(()=>{this.errorBag={}},0)},resetValidation(){this.inputs.forEach(t=>t.resetValidation()),this.resetErrorBag()},register(t){this.inputs.push(t),this.watchers.push(this.watchInput(t))},unregister(t){const e=this.inputs.find(e=>e._uid===t._uid);if(!e)return;const i=this.watchers.find(t=>t._uid===e._uid);i&&(i.valid(),i.shouldValidate()),this.watchers=this.watchers.filter(t=>t._uid!==e._uid),this.inputs=this.inputs.filter(t=>t._uid!==e._uid),this.$delete(this.errorBag,e._uid)}},render(t){return t("form",{staticClass:"v-form",attrs:{novalidate:!0,...this.attrs$},on:{submit:t=>this.$emit("submit",t)}},this.$slots.default)}})},"52e4":function(t,e,i){},"6ca7":function(t,e,i){},"75eb":function(t,e,i){"use strict";var s=i("9d65"),n=i("80d2"),a=i("58df"),o=i("d9bd");function r(t){const e=typeof t;return"boolean"===e||"string"===e||t.nodeType===Node.ELEMENT_NODE}e["a"]=Object(a["a"])(s["a"]).extend({name:"detachable",props:{attach:{default:!1,validator:r},contentClass:{type:String,default:""}},data:()=>({activatorNode:null,hasDetached:!1}),watch:{attach(){this.hasDetached=!1,this.initDetach()},hasContent(){this.$nextTick(this.initDetach)}},beforeMount(){this.$nextTick(()=>{if(this.activatorNode){const t=Array.isArray(this.activatorNode)?this.activatorNode:[this.activatorNode];t.forEach(t=>{if(!t.elm)return;if(!this.$el.parentNode)return;const e=this.$el===this.$el.parentNode.firstChild?this.$el:this.$el.nextSibling;this.$el.parentNode.insertBefore(t.elm,e)})}})},mounted(){this.hasContent&&this.initDetach()},deactivated(){this.isActive=!1},beforeDestroy(){try{if(this.$refs.content&&this.$refs.content.parentNode&&this.$refs.content.parentNode.removeChild(this.$refs.content),this.activatorNode){const t=Array.isArray(this.activatorNode)?this.activatorNode:[this.activatorNode];t.forEach(t=>{t.elm&&t.elm.parentNode&&t.elm.parentNode.removeChild(t.elm)})}}catch(t){console.log(t)}},methods:{getScopeIdAttrs(){const t=Object(n["j"])(this.$vnode,"context.$options._scopeId");return t&&{[t]:""}},initDetach(){if(this._isDestroyed||!this.$refs.content||this.hasDetached||""===this.attach||!0===this.attach||"attach"===this.attach)return;let t;t=!1===this.attach?document.querySelector("[data-app]"):"string"===typeof this.attach?document.querySelector(this.attach):this.attach,t?(t.appendChild(this.$refs.content),this.hasDetached=!0):Object(o["c"])(`Unable to locate target ${this.attach||"[data-app]"}`,this)}}})},"99d9":function(t,e,i){"use strict";i.d(e,"a",(function(){return a})),i.d(e,"b",(function(){return r})),i.d(e,"c",(function(){return l}));var s=i("b0af"),n=i("80d2");const a=Object(n["g"])("v-card__actions"),o=Object(n["g"])("v-card__subtitle"),r=Object(n["g"])("v-card__text"),l=Object(n["g"])("v-card__title");s["a"]},ac11:function(t,e,i){"use strict";i.r(e);var s=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",[i("Toolbar",{scopedSlots:t._u([{key:"title",fn:function(){return[t._v("\n          Chats\n        ")]},proxy:!0},{key:"prepend",fn:function(){return[i("v-btn",{attrs:{icon:""},on:{click:function(e){return t.$router.go(-1)}}},[i("v-icon",[t._v("mdi-arrow-left")])],1)]},proxy:!0},{key:"append",fn:function(){return[i("v-spacer"),i("NewChat"),i("v-btn",{attrs:{icon:""},on:{click:function(t){}}},[i("v-icon",[t._v("mdi-magnify")])],1)]},proxy:!0}])}),t.chats&&t.chats.length?i("v-list",{staticClass:"pt-0"},[t._l(t.chats,(function(e,s){return[0!=s?i("v-divider",{key:s}):t._e(),i("v-list-item",{key:e.title,on:{click:function(i){return t.getChat(e.identifier)}}},[i("v-list-item-avatar",{attrs:{color:e.title.split(" ").length>1?"blue":"success"}},[i("span",{staticClass:"white--text headline"},[t._v(t._s(e.title.charAt(0).toUpperCase()))])]),i("v-list-item-content",[i("v-list-item-title",{domProps:{textContent:t._s(e.title)}})],1)],1)]})),i("v-divider")],2):[i("p",{staticClass:"text-center pt-4 subtitle-1"},[t._v("You dont`t have chats")])]],2)},n=[],a=i("6908"),o=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",{staticClass:"text-center"},[i("v-dialog",{attrs:{width:"600"},scopedSlots:t._u([{key:"activator",fn:function(e){var s=e.on;return[i("v-btn",t._g({attrs:{icon:""}},s),[i("v-icon",[t._v("mdi-plus")])],1)]}}]),model:{value:t.dialog,callback:function(e){t.dialog=e},expression:"dialog"}},[i("v-card",[i("v-form",{ref:"form",on:{submit:function(e){return e.preventDefault(),t.addNewConversation(e)}},model:{value:t.valid,callback:function(e){t.valid=e},expression:"valid"}},[i("v-card-title",{staticClass:"headline grey lighten-2",attrs:{"primary-title":""}},[t._v("\n                  New chat\n                  "),i("v-spacer"),t.newConversation.users.length>0?[t._v("Selected: "+t._s(t.newConversation.users.length))]:t._e()],2),i("v-card-text",[i("v-text-field",{attrs:{placeholder:"Chat name"},model:{value:t.newConversation.title,callback:function(e){t.$set(t.newConversation,"title",e)},expression:"newConversation.title"}}),t.friends&&t.friends.length?i("v-list-item-group",{attrs:{multiple:""},model:{value:t.newConversation.users,callback:function(e){t.$set(t.newConversation,"users",e)},expression:"newConversation.users"}},t._l(t.friends,(function(e){return i("v-list-item",{key:e.id,attrs:{value:e},scopedSlots:t._u([{key:"default",fn:function(s){var n=s.active,a=s.toggle;return[i("v-list-item-avatar",[i("v-img",{attrs:{src:"https://cdn.vuetifyjs.com/images/lists/1.jpg"}})],1),i("v-list-item-content",[i("v-list-item-title",{domProps:{textContent:t._s(e.username)}})],1),i("v-list-item-action",[i("v-checkbox",{attrs:{"input-value":n,"true-value":e,color:"primary"},on:{click:a}})],1)]}}],null,!0)})})),1):[i("p",{staticClass:"text-center pt-4 subtitle-1"},[t._v("\n                          You dont`t have friends to add to the chat."),i("br"),t._v("\n                          Add friends first on "),i("router-link",{attrs:{to:{name:"users"}}},[t._v("users")]),t._v(" page.\n                      ")],1)]],2),i("v-divider"),i("v-card-actions",[i("v-spacer"),i("v-btn",{attrs:{text:""},on:{click:function(e){t.dialog=!1}}},[t._v("Cancel")]),i("v-btn",{attrs:{text:"",type:"submit",disabled:0===t.newConversation.users.length}},[t._v("Create")])],1)],1)],1)],1)],1)},r=[],l=i("bc3a"),c=i.n(l),h={name:"NewChat",data(){return{dialog:!1,valid:!1,newConversation:{title:"",users:[]}}},computed:{friends(){return this.$store.state.user.friends}},methods:{addNewConversation(){this.$refs.form.validate()&&(c.a.post("/chats",JSON.stringify(this.newConversation)).then(t=>{console.log(t.data),this.dialog=!1,this.$store.dispatch("getChats")}).catch(t=>{console.log(t)}),this.reset())}},mounted(){this.$store.dispatch("getFriends")}},d=h,u=(i("e593"),i("2877")),v=i("6544"),p=i.n(v),m=i("8336"),f=i("b0af"),g=i("99d9"),y=i("ac7c"),b=(i("368e"),i("7560")),C=b["a"].extend({name:"v-theme-provider",props:{root:Boolean},computed:{isDark(){return this.root?this.rootIsDark:b["a"].options.computed.isDark.call(this)}},render(){return this.$slots.default&&this.$slots.default.find(t=>!t.isComment&&" "!==t.text)}}),w=i("4ad4"),x=i("b848"),$=i("75eb"),k=(i("3c93"),i("a9ad")),A=i("f2e7"),_=i("58df"),O=Object(_["a"])(k["a"],b["a"],A["a"]).extend({name:"v-overlay",props:{absolute:Boolean,color:{type:String,default:"#212121"},dark:{type:Boolean,default:!0},opacity:{type:[Number,String],default:.46},value:{default:!0},zIndex:{type:[Number,String],default:5}},computed:{__scrim(){const t=this.setBackgroundColor(this.color,{staticClass:"v-overlay__scrim",style:{opacity:this.computedOpacity}});return this.$createElement("div",t)},classes(){return{"v-overlay--absolute":this.absolute,"v-overlay--active":this.isActive,...this.themeClasses}},computedOpacity(){return Number(this.isActive?this.opacity:0)},styles(){return{zIndex:this.zIndex}}},methods:{genContent(){return this.$createElement("div",{staticClass:"v-overlay__content"},this.$slots.default)}},render(t){const e=[this.__scrim];return this.isActive&&e.push(this.genContent()),t("div",{staticClass:"v-overlay",class:this.classes,style:this.styles},e)}}),E=O,I=i("80d2"),V=i("2b0e"),S=V["a"].extend().extend({name:"overlayable",props:{hideOverlay:Boolean,overlayColor:String,overlayOpacity:[Number,String]},data(){return{overlay:null}},watch:{hideOverlay(t){this.isActive&&(t?this.removeOverlay():this.genOverlay())}},beforeDestroy(){this.removeOverlay()},methods:{createOverlay(){const t=new E({propsData:{absolute:this.absolute,value:!1,color:this.overlayColor,opacity:this.overlayOpacity}});t.$mount();const e=this.absolute?this.$el.parentNode:document.querySelector("[data-app]");e&&e.insertBefore(t.$el,e.firstChild),this.overlay=t},genOverlay(){if(this.hideScroll(),!this.hideOverlay)return this.overlay||this.createOverlay(),requestAnimationFrame(()=>{this.overlay&&(void 0!==this.activeZIndex?this.overlay.zIndex=String(this.activeZIndex-1):this.$el&&(this.overlay.zIndex=Object(I["m"])(this.$el)))}),this.overlay&&(this.overlay.value=!0),!0},removeOverlay(t=!0){this.overlay&&(Object(I["a"])(this.overlay.$el,"transitionend",()=>{this.overlay&&this.overlay.$el&&this.overlay.$el.parentNode&&!this.overlay.value&&(this.overlay.$el.parentNode.removeChild(this.overlay.$el),this.overlay.$destroy(),this.overlay=null)}),this.overlay.value=!1),t&&this.showScroll()},scrollListener(t){if("keydown"===t.type){if(["INPUT","TEXTAREA","SELECT"].includes(t.target.tagName)||t.target.isContentEditable)return;const e=[I["o"].up,I["o"].pageup],i=[I["o"].down,I["o"].pagedown];if(e.includes(t.keyCode))t.deltaY=-1;else{if(!i.includes(t.keyCode))return;t.deltaY=1}}(t.target===this.overlay||"keydown"!==t.type&&t.target===document.body||this.checkPath(t))&&t.preventDefault()},hasScrollbar(t){if(!t||t.nodeType!==Node.ELEMENT_NODE)return!1;const e=window.getComputedStyle(t);return["auto","scroll"].includes(e.overflowY)&&t.scrollHeight>t.clientHeight},shouldScroll(t,e){return 0===t.scrollTop&&e<0||t.scrollTop+t.clientHeight===t.scrollHeight&&e>0},isInside(t,e){return t===e||null!==t&&t!==document.body&&this.isInside(t.parentNode,e)},checkPath(t){const e=t.path||this.composedPath(t),i=t.deltaY;if("keydown"===t.type&&e[0]===document.body){const t=this.$refs.dialog,e=window.getSelection().anchorNode;return!(t&&this.hasScrollbar(t)&&this.isInside(e,t))||this.shouldScroll(t,i)}for(let s=0;s<e.length;s++){const t=e[s];if(t===document)return!0;if(t===document.documentElement)return!0;if(t===this.$refs.content)return!0;if(this.hasScrollbar(t))return this.shouldScroll(t,i)}return!0},composedPath(t){if(t.composedPath)return t.composedPath();const e=[];let i=t.target;while(i){if(e.push(i),"HTML"===i.tagName)return e.push(document),e.push(window),e;i=i.parentElement}return e},hideScroll(){this.$vuetify.breakpoint.smAndDown?document.documentElement.classList.add("overflow-y-hidden"):(Object(I["b"])(window,"wheel",this.scrollListener,{passive:!1}),window.addEventListener("keydown",this.scrollListener))},showScroll(){document.documentElement.classList.remove("overflow-y-hidden"),window.removeEventListener("wheel",this.scrollListener),window.removeEventListener("keydown",this.scrollListener)}}}),D=V["a"].extend({name:"returnable",props:{returnValue:null},data:()=>({isActive:!1,originalValue:null}),watch:{isActive(t){t?this.originalValue=this.returnValue:this.$emit("update:return-value",this.originalValue)}},methods:{save(t){this.originalValue=t,setTimeout(()=>{this.isActive=!1})}}}),N=i("21be");function T(){return!1}function B(t,e,i){i.args=i.args||{};const s=i.args.closeConditional||T;if(!t||!1===s(t))return;if("isTrusted"in t&&!t.isTrusted||"pointerType"in t&&!t.pointerType)return;const n=(i.args.include||(()=>[]))();n.push(e),!n.some(e=>e.contains(t.target))&&setTimeout(()=>{s(t)&&i.value&&i.value(t)},0)}const j={inserted(t,e){const i=i=>B(i,t,e),s=document.querySelector("[data-app]")||document.body;s.addEventListener("click",i,!0),t._clickOutside=i},unbind(t){if(!t._clickOutside)return;const e=document.querySelector("[data-app]")||document.body;e&&e.removeEventListener("click",t._clickOutside,!0),delete t._clickOutside}};var L=j,M=i("d9bd");const P=Object(_["a"])(w["a"],x["a"],$["a"],S,D,N["a"],A["a"]);var F=P.extend({name:"v-dialog",directives:{ClickOutside:L},props:{dark:Boolean,disabled:Boolean,fullscreen:Boolean,light:Boolean,maxWidth:{type:[String,Number],default:"none"},noClickAnimation:Boolean,origin:{type:String,default:"center center"},persistent:Boolean,retainFocus:{type:Boolean,default:!0},scrollable:Boolean,transition:{type:[String,Boolean],default:"dialog-transition"},width:{type:[String,Number],default:"auto"}},data(){return{activatedBy:null,animate:!1,animateTimeout:-1,isActive:!!this.value,stackMinZIndex:200}},computed:{classes(){return{[`v-dialog ${this.contentClass}`.trim()]:!0,"v-dialog--active":this.isActive,"v-dialog--persistent":this.persistent,"v-dialog--fullscreen":this.fullscreen,"v-dialog--scrollable":this.scrollable,"v-dialog--animated":this.animate}},contentClasses(){return{"v-dialog__content":!0,"v-dialog__content--active":this.isActive}},hasActivator(){return Boolean(!!this.$slots.activator||!!this.$scopedSlots.activator)}},watch:{isActive(t){t?(this.show(),this.hideScroll()):(this.removeOverlay(),this.unbind())},fullscreen(t){this.isActive&&(t?(this.hideScroll(),this.removeOverlay(!1)):(this.showScroll(),this.genOverlay()))}},created(){this.$attrs.hasOwnProperty("full-width")&&Object(M["d"])("full-width",this)},beforeMount(){this.$nextTick(()=>{this.isBooted=this.isActive,this.isActive&&this.show()})},beforeDestroy(){"undefined"!==typeof window&&this.unbind()},methods:{animateClick(){this.animate=!1,this.$nextTick(()=>{this.animate=!0,window.clearTimeout(this.animateTimeout),this.animateTimeout=window.setTimeout(()=>this.animate=!1,150)})},closeConditional(t){const e=t.target;return!(this._isDestroyed||!this.isActive||this.$refs.content.contains(e)||this.overlay&&e&&!this.overlay.$el.contains(e))&&this.activeZIndex>=this.getMaxZIndex()},hideScroll(){this.fullscreen?document.documentElement.classList.add("overflow-y-hidden"):S.options.methods.hideScroll.call(this)},show(){!this.fullscreen&&!this.hideOverlay&&this.genOverlay(),this.$nextTick(()=>{this.$refs.content.focus(),this.bind()})},bind(){window.addEventListener("focusin",this.onFocusin)},unbind(){window.removeEventListener("focusin",this.onFocusin)},onClickOutside(t){this.$emit("click:outside",t),this.persistent?this.noClickAnimation||this.animateClick():this.isActive=!1},onKeydown(t){if(t.keyCode===I["o"].esc&&!this.getOpenDependents().length)if(this.persistent)this.noClickAnimation||this.animateClick();else{this.isActive=!1;const t=this.getActivator();this.$nextTick(()=>t&&t.focus())}this.$emit("keydown",t)},onFocusin(t){if(!t||!this.retainFocus)return;const e=t.target;if(e&&![document,this.$refs.content].includes(e)&&!this.$refs.content.contains(e)&&this.activeZIndex>=this.getMaxZIndex()&&!this.getOpenDependentElements().some(t=>t.contains(e))){const t=this.$refs.content.querySelectorAll('button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])');t.length&&t[0].focus()}},genContent(){return this.showLazyContent(()=>[this.$createElement(C,{props:{root:!0,light:this.light,dark:this.dark}},[this.$createElement("div",{class:this.contentClasses,attrs:{role:"document",tabindex:this.isActive?0:void 0,...this.getScopeIdAttrs()},on:{keydown:this.onKeydown},style:{zIndex:this.activeZIndex},ref:"content"},[this.genTransition()])])])},genTransition(){const t=this.genInnerContent();return this.transition?this.$createElement("transition",{props:{name:this.transition,origin:this.origin,appear:!0}},[t]):t},genInnerContent(){const t={class:this.classes,ref:"dialog",directives:[{name:"click-outside",value:this.onClickOutside,args:{closeConditional:this.closeConditional,include:this.getOpenDependentElements}},{name:"show",value:this.isActive}],style:{transformOrigin:this.origin}};return this.fullscreen||(t.style={...t.style,maxWidth:"none"===this.maxWidth?void 0:Object(I["f"])(this.maxWidth),width:"auto"===this.width?void 0:Object(I["f"])(this.width)}),this.$createElement("div",t,this.getContentSlot())}},render(t){return t("div",{staticClass:"v-dialog__container",class:{"v-dialog__container--attached":""===this.attach||!0===this.attach||"attach"===this.attach},attrs:{role:"dialog"}},[this.genActivator(),this.genContent()])}}),Z=i("ce7e"),z=i("4bd4"),q=i("132d"),H=i("adda"),Y=i("da13"),K=i("1800"),W=i("8270"),J=i("5d23"),R=i("1baa"),U=i("2fa4"),G=i("8654"),X=Object(u["a"])(d,o,r,!1,null,"d10f3432",null),Q=X.exports;p()(X,{VBtn:m["a"],VCard:f["a"],VCardActions:g["a"],VCardText:g["b"],VCardTitle:g["c"],VCheckbox:y["a"],VDialog:F,VDivider:Z["a"],VForm:z["a"],VIcon:q["a"],VImg:H["a"],VListItem:Y["a"],VListItemAction:K["a"],VListItemAvatar:W["a"],VListItemContent:J["a"],VListItemGroup:R["a"],VListItemTitle:J["c"],VSpacer:U["a"],VTextField:G["a"]});var tt={name:"Chats",components:{Toolbar:a["a"],NewChat:Q},data(){return{colors:["success","blue"]}},computed:{chats(){return this.$store.state.chat.chats}},methods:{getChat(t){this.$store.dispatch("getChat",t)},randomInt(t,e){return t+Math.floor((e-t)*Math.random())}},mounted(){this.$store.dispatch("getChats")}},et=tt,it=i("8860"),st=Object(u["a"])(et,s,n,!1,null,null,null);e["default"]=st.exports;p()(st,{VBtn:m["a"],VDivider:Z["a"],VIcon:q["a"],VList:it["a"],VListItem:Y["a"],VListItemAvatar:W["a"],VListItemContent:J["a"],VListItemTitle:J["c"],VSpacer:U["a"]})},ac7c:function(t,e,i){"use strict";i("6ca7"),i("ec29");var s=i("9d26"),n=i("c37a"),a=i("5607"),o=i("2b0e"),r=o["a"].extend({name:"rippleable",directives:{ripple:a["a"]},props:{ripple:{type:[Boolean,Object],default:!0}},methods:{genRipple(t={}){return this.ripple?(t.staticClass="v-input--selection-controls__ripple",t.directives=t.directives||[],t.directives.push({name:"ripple",value:{center:!0}}),t.on=Object.assign({click:this.onChange},this.$listeners),this.$createElement("div",t)):null},onChange(){}}}),l=i("80d2"),c=o["a"].extend({name:"comparable",props:{valueComparator:{type:Function,default:l["h"]}}}),h=i("58df"),d=Object(h["a"])(n["a"],r,c).extend({name:"selectable",model:{prop:"inputValue",event:"change"},props:{id:String,inputValue:null,falseValue:null,trueValue:null,multiple:{type:Boolean,default:null},label:String},data(){return{hasColor:this.inputValue,lazyValue:this.inputValue}},computed:{computedColor(){if(this.isActive)return this.color?this.color:this.isDark&&!this.appIsDark?"white":"primary"},isMultiple(){return!0===this.multiple||null===this.multiple&&Array.isArray(this.internalValue)},isActive(){const t=this.value,e=this.internalValue;return this.isMultiple?!!Array.isArray(e)&&e.some(e=>this.valueComparator(e,t)):void 0===this.trueValue||void 0===this.falseValue?t?this.valueComparator(t,e):Boolean(e):this.valueComparator(e,this.trueValue)},isDirty(){return this.isActive},rippleState(){return this.disabled||this.validationState?this.validationState:"primary"}},watch:{inputValue(t){this.lazyValue=t,this.hasColor=t}},methods:{genLabel(){const t=n["a"].options.methods.genLabel.call(this);return t?(t.data.on={click:t=>{t.preventDefault(),this.onChange()}},t):t},genInput(t,e){return this.$createElement("input",{attrs:Object.assign({"aria-checked":this.isActive.toString(),disabled:this.isDisabled,id:this.computedId,role:t,type:t},e),domProps:{value:this.value,checked:this.isActive},on:{blur:this.onBlur,change:this.onChange,focus:this.onFocus,keydown:this.onKeydown},ref:"input"})},onBlur(){this.isFocused=!1},onChange(){if(this.isDisabled)return;const t=this.value;let e=this.internalValue;if(this.isMultiple){Array.isArray(e)||(e=[]);const i=e.length;e=e.filter(e=>!this.valueComparator(e,t)),e.length===i&&e.push(t)}else e=void 0!==this.trueValue&&void 0!==this.falseValue?this.valueComparator(e,this.trueValue)?this.falseValue:this.trueValue:t?this.valueComparator(e,t)?null:t:!e;this.validate(!0,e),this.internalValue=e,this.hasColor=e},onFocus(){this.isFocused=!0},onKeydown(t){}}});e["a"]=d.extend({name:"v-checkbox",props:{indeterminate:Boolean,indeterminateIcon:{type:String,default:"$checkboxIndeterminate"},offIcon:{type:String,default:"$checkboxOff"},onIcon:{type:String,default:"$checkboxOn"}},data(){return{inputIndeterminate:this.indeterminate}},computed:{classes(){return{...n["a"].options.computed.classes.call(this),"v-input--selection-controls":!0,"v-input--checkbox":!0,"v-input--indeterminate":this.inputIndeterminate}},computedIcon(){return this.inputIndeterminate?this.indeterminateIcon:this.isActive?this.onIcon:this.offIcon},validationState(){if(!this.disabled||this.inputIndeterminate)return this.hasError&&this.shouldValidate?"error":this.hasSuccess?"success":null!==this.hasColor?this.computedColor:void 0}},watch:{indeterminate(t){this.$nextTick(()=>this.inputIndeterminate=t)},inputIndeterminate(t){this.$emit("update:indeterminate",t)},isActive(){this.indeterminate&&(this.inputIndeterminate=!1)}},methods:{genCheckbox(){return this.$createElement("div",{staticClass:"v-input--selection-controls__input"},[this.$createElement(s["a"],this.setTextColor(this.validationState,{props:{dense:this.dense,dark:this.dark,light:this.light}}),this.computedIcon),this.genInput("checkbox",{...this.attrs$,"aria-checked":this.inputIndeterminate?"mixed":this.isActive.toString()}),this.genRipple(this.setTextColor(this.rippleState))])},genDefaultSlot(){return[this.genCheckbox(),this.genLabel()]}}})},b848:function(t,e,i){"use strict";var s=i("58df");function n(t){const e=[];for(let i=0;i<t.length;i++){const s=t[i];s.isActive&&s.isDependent?e.push(s):e.push(...n(s.$children))}return e}e["a"]=Object(s["a"])().extend({name:"dependent",data(){return{closeDependents:!0,isActive:!1,isDependent:!0}},watch:{isActive(t){if(t)return;const e=this.getOpenDependents();for(let i=0;i<e.length;i++)e[i].isActive=!1}},methods:{getOpenDependents(){return this.closeDependents?n(this.$children):[]},getOpenDependentElements(){const t=[],e=this.getOpenDependents();for(let i=0;i<e.length;i++)t.push(...e[i].getClickableDependentElements());return t},getClickableDependentElements(){const t=[this.$el];return this.$refs.content&&t.push(this.$refs.content),this.overlay&&t.push(this.overlay.$el),t.push(...this.getOpenDependentElements()),t}}})},e593:function(t,e,i){"use strict";var s=i("52e4"),n=i.n(s);n.a},ec29:function(t,e,i){}}]);
//# sourceMappingURL=chunk-566a4814.8732ed7f.js.map