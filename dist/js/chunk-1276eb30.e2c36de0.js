(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-1276eb30"],{"0798":function(t,e,s){"use strict";s("0c18");var r=s("10d2"),i=s("afdd"),a=s("9d26"),o=s("f2e7"),n=s("7560"),l=s("2b0e"),d=l["a"].extend({name:"transitionable",props:{mode:String,origin:String,transition:String}}),c=s("58df"),h=s("d9bd");e["a"]=Object(c["a"])(r["a"],o["a"],d).extend({name:"v-alert",props:{border:{type:String,validator(t){return["top","right","bottom","left"].includes(t)}},closeLabel:{type:String,default:"$vuetify.close"},coloredBorder:Boolean,dense:Boolean,dismissible:Boolean,icon:{default:"",type:[Boolean,String],validator(t){return"string"===typeof t||!1===t}},outlined:Boolean,prominent:Boolean,text:Boolean,type:{type:String,validator(t){return["info","error","success","warning"].includes(t)}},value:{type:Boolean,default:!0}},computed:{__cachedBorder(){if(!this.border)return null;let t={staticClass:"v-alert__border",class:{[`v-alert__border--${this.border}`]:!0}};return this.coloredBorder&&(t=this.setBackgroundColor(this.computedColor,t),t.class["v-alert__border--has-color"]=!0),this.$createElement("div",t)},__cachedDismissible(){if(!this.dismissible)return null;const t=this.iconColor;return this.$createElement(i["a"],{staticClass:"v-alert__dismissible",props:{color:t,icon:!0,small:!0},attrs:{"aria-label":this.$vuetify.lang.t(this.closeLabel)},on:{click:()=>this.isActive=!1}},[this.$createElement(a["a"],{props:{color:t}},"$cancel")])},__cachedIcon(){return this.computedIcon?this.$createElement(a["a"],{staticClass:"v-alert__icon",props:{color:this.iconColor}},this.computedIcon):null},classes(){const t={...r["a"].options.computed.classes.call(this),"v-alert--border":Boolean(this.border),"v-alert--dense":this.dense,"v-alert--outlined":this.outlined,"v-alert--prominent":this.prominent,"v-alert--text":this.text};return this.border&&(t[`v-alert--border-${this.border}`]=!0),t},computedColor(){return this.color||this.type},computedIcon(){return!1!==this.icon&&("string"===typeof this.icon&&this.icon?this.icon:!!["error","info","success","warning"].includes(this.type)&&`$${this.type}`)},hasColoredIcon(){return this.hasText||Boolean(this.border)&&this.coloredBorder},hasText(){return this.text||this.outlined},iconColor(){return this.hasColoredIcon?this.computedColor:void 0},isDark(){return!(!this.type||this.coloredBorder||this.outlined)||n["a"].options.computed.isDark.call(this)}},created(){this.$attrs.hasOwnProperty("outline")&&Object(h["a"])("outline","outlined",this)},methods:{genWrapper(){const t=[this.$slots.prepend||this.__cachedIcon,this.genContent(),this.__cachedBorder,this.$slots.append,this.$scopedSlots.close?this.$scopedSlots.close({toggle:this.toggle}):this.__cachedDismissible],e={staticClass:"v-alert__wrapper"};return this.$createElement("div",e,t)},genContent(){return this.$createElement("div",{staticClass:"v-alert__content"},this.$slots.default)},genAlert(){let t={staticClass:"v-alert",attrs:{role:"alert"},class:this.classes,style:this.styles,directives:[{name:"show",value:this.isActive}]};if(!this.coloredBorder){const e=this.hasText?this.setTextColor:this.setBackgroundColor;t=e(this.computedColor,t)}return this.$createElement("div",t,[this.genWrapper()])},toggle(){this.isActive=!this.isActive}},render(t){const e=this.genAlert();return this.transition?t("transition",{props:{name:this.transition,origin:this.origin,mode:this.mode}},[e]):e}})},"0c18":function(t,e,s){},"4bd4":function(t,e,s){"use strict";var r=s("58df"),i=s("7e2b"),a=s("3206");e["a"]=Object(r["a"])(i["a"],Object(a["b"])("form")).extend({name:"v-form",inheritAttrs:!1,props:{lazyValidation:Boolean,value:Boolean},data:()=>({inputs:[],watchers:[],errorBag:{}}),watch:{errorBag:{handler(t){const e=Object.values(t).includes(!0);this.$emit("input",!e)},deep:!0,immediate:!0}},methods:{watchInput(t){const e=t=>t.$watch("hasError",e=>{this.$set(this.errorBag,t._uid,e)},{immediate:!0}),s={_uid:t._uid,valid:()=>{},shouldValidate:()=>{}};return this.lazyValidation?s.shouldValidate=t.$watch("shouldValidate",r=>{r&&(this.errorBag.hasOwnProperty(t._uid)||(s.valid=e(t)))}):s.valid=e(t),s},validate(){return 0===this.inputs.filter(t=>!t.validate(!0)).length},reset(){this.inputs.forEach(t=>t.reset()),this.resetErrorBag()},resetErrorBag(){this.lazyValidation&&setTimeout(()=>{this.errorBag={}},0)},resetValidation(){this.inputs.forEach(t=>t.resetValidation()),this.resetErrorBag()},register(t){this.inputs.push(t),this.watchers.push(this.watchInput(t))},unregister(t){const e=this.inputs.find(e=>e._uid===t._uid);if(!e)return;const s=this.watchers.find(t=>t._uid===e._uid);s&&(s.valid(),s.shouldValidate()),this.watchers=this.watchers.filter(t=>t._uid!==e._uid),this.inputs=this.inputs.filter(t=>t._uid!==e._uid),this.$delete(this.errorBag,e._uid)}},render(t){return t("form",{staticClass:"v-form",attrs:{novalidate:!0,...this.attrs$},on:{submit:t=>this.$emit("submit",t)}},this.$slots.default)}})},"5c9c":function(t,e,s){"use strict";s.r(e);var r=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("v-container",{staticClass:"fill-height",attrs:{fluid:""}},[s("v-row",{attrs:{align:"center",justify:"center"}},[s("v-col",{attrs:{cols:"12",sm:"10",md:"8",lg:"6",xl:"4"}},[s("v-card",{staticClass:"elevation-12"},[s("v-form",{ref:"form",on:{submit:function(e){return e.preventDefault(),t.signUp(e)}},model:{value:t.valid,callback:function(e){t.valid=e},expression:"valid"}},[s("v-toolbar",{staticClass:"white--text",attrs:{color:"grey darken-3"}},[s("v-toolbar-title",[t._v("Sign up")])],1),s("v-card-text",[s("v-text-field",{attrs:{"prepend-icon":"mdi-account",label:"Username*",rules:[t.rules.required,t.rules.minMaxUsername],required:""},model:{value:t.newUser.username,callback:function(e){t.$set(t.newUser,"username",e)},expression:"newUser.username"}}),s("v-text-field",{attrs:{"append-icon":t.showPass?"mdi-eye":"mdi-eye-off","prepend-icon":"mdi-lock",label:"Password*",type:t.showPass?"text":"password",rules:[t.rules.required,t.rules.minMaxPass],required:""},on:{"click:append":function(e){t.showPass=!t.showPass}},model:{value:t.newUser.password,callback:function(e){t.$set(t.newUser,"password",e)},expression:"newUser.password"}}),s("v-text-field",{attrs:{"append-icon":t.showPassConfirm?"mdi-eye":"mdi-eye-off","prepend-icon":"mdi-lock",label:"Repeat password*",type:t.showPassConfirm?"text":"password",rules:[t.rules.required,t.rules.passwordMatch],required:""},on:{"click:append":function(e){t.showPassConfirm=!t.showPassConfirm}},model:{value:t.passwordConfirm,callback:function(e){t.passwordConfirm=e},expression:"passwordConfirm"}}),"error"===t.authStatus?s("v-alert",{attrs:{dense:"",outlined:"",type:"error"}},[s("strong",[t._v(t._s(t.authError))])]):t._e()],1),s("v-card-actions",[s("v-spacer"),s("v-btn",{staticClass:"white--text",attrs:{type:"submit",color:"grey darken-3"}},[t._v("Sign up")])],1),s("v-divider")],1),s("p",{staticClass:"text-center font-weight-light"},[t._v("Already have an account? "),s("router-link",{staticClass:"font-weight-bold",attrs:{to:"/sign_in"}},[t._v("Sign in")])],1)],1)],1)],1)],1)},i=[],a={data(){return{newUser:{username:"",password:""},passwordConfirm:"",showPass:!1,showPassConfirm:!1,valid:!1,rules:{required:t=>!!t||"Required field",minMaxPass:t=>t.length>=8&&t.length<=30||"From 8 to 30 characters",minMaxUsername:t=>t.length>=5&&t.length<=30||"From 5 to 30 characters",passwordMatch:t=>t===this.newUser.password||"Passwords don't not match"}}},methods:{signUp(){this.$refs.form.validate()&&this.$store.dispatch("signUp",this.newUser).then(t=>{this.$store.dispatch("getUserCredentials"),this.$router.push("/")}).catch(t=>{console.log(t)})}},computed:{authStatus(){return this.$store.getters.authStatus},authError(){return this.$store.getters.authError}},beforeRouteLeave(t,e,s){this.$store.state.auth.authStatus="none",s()}},o=a,n=s("2877"),l=s("6544"),d=s.n(l),c=s("0798"),h=s("8336"),u=s("b0af"),p=s("99d9"),m=s("62ad"),v=s("a523"),f=s("ce7e"),g=s("4bd4"),w=s("0fd9"),b=s("2fa4"),_=s("8654"),C=s("71d9"),$=s("2a7f"),y=Object(n["a"])(o,r,i,!1,null,null,null);e["default"]=y.exports;d()(y,{VAlert:c["a"],VBtn:h["a"],VCard:u["a"],VCardActions:p["a"],VCardText:p["b"],VCol:m["a"],VContainer:v["a"],VDivider:f["a"],VForm:g["a"],VRow:w["a"],VSpacer:b["a"],VTextField:_["a"],VToolbar:C["a"],VToolbarTitle:$["a"]})},afdd:function(t,e,s){"use strict";var r=s("8336");e["a"]=r["a"]}}]);
//# sourceMappingURL=chunk-1276eb30.e2c36de0.js.map