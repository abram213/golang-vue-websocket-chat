(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-608f7f58"],{"2fa4":function(t,e,n){"use strict";n("20f6");var i=n("80d2");e["a"]=Object(i["g"])("spacer","div","v-spacer")},d318:function(t,e,n){"use strict";n.r(e);var i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("Toolbar",{scopedSlots:t._u([{key:"title",fn:function(){return[t._v("\n          Friends\n        ")]},proxy:!0},{key:"prepend",fn:function(){return[n("v-btn",{attrs:{icon:""},on:{click:function(e){return t.$router.go(-1)}}},[n("v-icon",[t._v("mdi-arrow-left")])],1)]},proxy:!0},{key:"append",fn:function(){return[n("v-spacer"),n("v-btn",{attrs:{icon:""},on:{click:function(t){}}},[n("v-icon",[t._v("mdi-magnify")])],1)]},proxy:!0}])}),t.friends&&t.friends.length?n("v-list",{staticClass:"pt-0"},[t._l(t.friends,(function(e,i){return[0!=i?n("v-divider",{key:i,attrs:{inset:""}}):t._e(),n("v-list-item",{key:e.id,staticClass:"pr-1"},[n("v-list-item-avatar",[n("v-icon",{attrs:{large:""}},[t._v("mdi-account-circle")])],1),n("v-list-item-content",[n("v-list-item-title",{domProps:{textContent:t._s(e.username)}})],1),n("v-list-item-action",[n("v-btn",{attrs:{small:"",text:""},on:{click:function(n){return n.stopPropagation(),t.deleteFriend(e.id,i)}}},[t._v("\n                    delete\n                    "),n("v-icon",{attrs:{right:"",color:"error"}},[t._v("mdi-trash-can-outline")])],1)],1)],1)]})),n("v-divider",{attrs:{inset:""}})],2):[n("p",{staticClass:"text-center pt-4 subtitle-1"},[t._v("\n          You dont`t have friends."),n("br"),t._v("\n          Add friends first on "),n("router-link",{attrs:{to:{name:"users"}}},[t._v("users")]),t._v(" page.\n        ")],1)]],2)},r=[],s=n("6908"),o={name:"Friends",components:{Toolbar:s["a"]},data(){return{}},computed:{friends(){return this.$store.state.user.friends}},methods:{deleteFriend(t,e){this.$store.dispatch("deleteFriend",t).then(t=>{this.friends.splice(e,1),console.log(t)}).catch(t=>{console.log(t)})}},mounted(){this.$store.dispatch("getFriends")}},a=o,c=n("2877"),d=n("6544"),l=n.n(d),u=n("8336"),v=n("ce7e"),f=n("132d"),p=n("8860"),m=n("da13"),h=n("1800"),_=n("8270"),k=n("5d23"),b=n("2fa4"),g=Object(c["a"])(a,i,r,!1,null,null,null);e["default"]=g.exports;l()(g,{VBtn:u["a"],VDivider:v["a"],VIcon:f["a"],VList:p["a"],VListItem:m["a"],VListItemAction:h["a"],VListItemAvatar:_["a"],VListItemContent:k["a"],VListItemTitle:k["c"],VSpacer:b["a"]})}}]);
//# sourceMappingURL=chunk-608f7f58.4cbffa9a.js.map