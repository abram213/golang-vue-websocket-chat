(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-7eba8c30"],{"16b7":function(e,t,s){"use strict";var i=s("2b0e");t["a"]=i["a"].extend().extend({name:"delayable",props:{openDelay:{type:[Number,String],default:0},closeDelay:{type:[Number,String],default:0}},data:()=>({openTimeout:void 0,closeTimeout:void 0}),methods:{clearDelay(){clearTimeout(this.openTimeout),clearTimeout(this.closeTimeout)},runDelay(e,t){this.clearDelay();const s=parseInt(this[`${e}Delay`],10);this[`${e}Timeout`]=setTimeout(t||(()=>{this.isActive={open:!0,close:!1}[e]}),s)}}})},"2fa4":function(e,t,s){"use strict";s("20f6");var i=s("80d2");t["a"]=Object(i["g"])("spacer","div","v-spacer")},dbef:function(e,t,s){"use strict";s.r(t);var i=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("div",[s("Toolbar",{scopedSlots:e._u([{key:"title",fn:function(){return[e._v("\n          All users\n        ")]},proxy:!0},{key:"prepend",fn:function(){return[s("v-btn",{attrs:{icon:""},on:{click:function(t){return e.$router.go(-1)}}},[s("v-icon",[e._v("mdi-arrow-left")])],1)]},proxy:!0},{key:"append",fn:function(){return[s("v-spacer"),s("v-btn",{attrs:{icon:""},on:{click:function(e){}}},[s("v-icon",[e._v("mdi-magnify")])],1)]},proxy:!0}])}),e.users&&e.users.length?s("v-list",{staticClass:"pt-0"},[e._l(e.users,(function(t,i){return[0!=i?s("v-divider",{key:i,attrs:{inset:""}}):e._e(),s("v-list-item",{key:t.id,staticClass:"pr-1"},[s("v-list-item-avatar",[s("v-icon",{attrs:{large:""}},[e._v("mdi-account-circle")])],1),s("v-list-item-content",[s("v-list-item-title",{domProps:{textContent:e._s(t.username)}})],1),s("v-list-item-action",[t.friendship?s("v-hover",{scopedSlots:e._u([{key:"default",fn:function(i){var n=i.hover;return[n?s("v-btn",{attrs:{small:"",text:""},on:{click:function(s){return s.stopPropagation(),e.deleteFriend(t.id)}}},[e._v("\n                  delete\n                  "),s("v-icon",{attrs:{right:"",color:"error"}},[e._v("mdi-trash-can-outline")])],1):s("v-btn",{attrs:{small:"",text:""}},[e._v("\n                  friend\n                  "),s("v-icon",{attrs:{right:"",color:"success"}},[e._v("mdi-check-bold")])],1)]}}],null,!0)}):s("v-btn",{attrs:{small:"",text:""},on:{click:function(s){return s.stopPropagation(),e.addFriend(t.id)}}},[s("v-icon",{attrs:{left:""}},[e._v("mdi-account-plus")]),e._v("\n                  Add to friends\n              ")],1)],1)],1)]})),s("v-divider",{attrs:{inset:""}})],2):[s("p",{staticClass:"text-center pt-4 subtitle-1"},[e._v("No other users")])]],2)},n=[],o=s("6908"),r={name:"Users",components:{Toolbar:o["a"]},data(){return{}},computed:{users(){return this.$store.getters.usersWithFriendship}},methods:{addFriend(e){this.$store.dispatch("addFriend",e).then(t=>{this.users.forEach(t=>{t.id===e&&(t.friendship=!0)}),console.log(t),this.$forceUpdate()}).catch(e=>{console.log(e)})},deleteFriend(e){this.$store.dispatch("deleteFriend",e).then(t=>{this.users.forEach(t=>{t.id===e&&(t.friendship=!1)}),console.log(t),this.$forceUpdate()}).catch(e=>{console.log(e)})}},mounted(){this.$store.dispatch("getUsers"),this.$store.dispatch("getFriends")}},a=r,l=s("2877"),c=s("6544"),d=s.n(c),u=s("8336"),v=s("ce7e"),h=s("16b7"),p=s("f2e7"),f=s("58df"),m=s("d9bd"),b=Object(f["a"])(h["a"],p["a"]).extend({name:"v-hover",props:{disabled:{type:Boolean,default:!1},value:{type:Boolean,default:void 0}},methods:{onMouseEnter(){this.runDelay("open")},onMouseLeave(){this.runDelay("close")}},render(){if(!this.$scopedSlots.default&&void 0===this.value)return Object(m["c"])("v-hover is missing a default scopedSlot or bound value",this),null;let e;return this.$scopedSlots.default&&(e=this.$scopedSlots.default({hover:this.isActive})),Array.isArray(e)&&1===e.length&&(e=e[0]),e&&!Array.isArray(e)&&e.tag?(this.disabled||(e.data=e.data||{},this._g(e.data,{mouseenter:this.onMouseEnter,mouseleave:this.onMouseLeave})),e):(Object(m["c"])("v-hover should only contain a single element",this),e)}}),y=s("132d"),g=s("8860"),_=s("da13"),k=s("1800"),$=s("8270"),x=s("5d23"),T=s("2fa4"),V=Object(l["a"])(a,i,n,!1,null,null,null);t["default"]=V.exports;d()(V,{VBtn:u["a"],VDivider:v["a"],VHover:b,VIcon:y["a"],VList:g["a"],VListItem:_["a"],VListItemAction:k["a"],VListItemAvatar:$["a"],VListItemContent:x["a"],VListItemTitle:x["c"],VSpacer:T["a"]})}}]);
//# sourceMappingURL=chunk-7eba8c30.6d305628.js.map