import{o as d,d as u,j as s,w as r,v as l,n as _,t as p,p as c,k as h,m}from"./vendor.js";import{_ as f}from"./school.js";import{_ as v}from"./main.js";const g={data(){return{name:"",username:"",password:"",confirmPassword:"",status_22:"22_open.png",status_33:"33_open.png",status:"\u6CE8\u518C",statusClass:"safe"}},methods:{close(){this.status_22="22_close.png",this.status_33="33_close.png"},open(){this.status_22="22_open.png",this.status_33="33_open.png"},isSafe(){this.password.length>16?(this.status="\u5BC6\u7801\u8FC7\u957F",this.statusClass="unsafe"):this.password!==this.confirmPassword?(this.status="\u5BC6\u7801\u4E0D\u4E00\u81F4",this.statusClass="unsafe"):(this.status="\u6CE8\u518C",this.statusClass="safe")},checkString(a){var t=/^[a-zA-Z0-9_]{1,}$/;return t.test(a)},checkSafe(){return!(this.statusClass!=="safe"||!this.checkString(this.password))},submit(){this.checkSafe()&&this.$axios.post("/register",{name:this.name,username:this.username,password:this.password}).then(a=>{console.log(a),a.data.code==200?(alert("\u6CE8\u518C\u6210\u529F"),this.$router.push("/login")):(this.status=a.data.msg,alert(this.status),this.statusClass="unsafe")})}}},n=a=>(c("data-v-35f270d6"),a=a(),h(),a),w={id:"register-box"},x=n(()=>s("div",{id:"logo"},[s("img",{src:f,id:"school"})],-1)),b={id:"register"},C=["src"],S=["src"],k=n(()=>s("div",{class:"top"},"\u5B66\u751F\u7CFB\u7EDF\u6CE8\u518C",-1)),y=n(()=>s("div",{class:"split"},null,-1)),V={class:"center"},F={class:"box"},I={class:"search"},P=n(()=>s("label",{for:"username"},"\u59D3\u540D\uFF1A",-1)),U={class:"box"},B={class:"search"},z=n(()=>s("label",{for:"username"},"\u8D26\u53F7\uFF1A",-1)),D={class:"box"},N={class:"search"},R=n(()=>s("label",{for:"password"},"\u5BC6\u7801\uFF1A",-1)),T={class:"box"},j={class:"search"},A=n(()=>s("label",{for:"password"},"\u5BC6\u7801\uFF1A",-1)),E=m(" /> ");function M(a,t,Z,q,o,i){return d(),u("div",w,[x,s("div",b,[s("img",{src:o.status_22,width:"185",class:"left-22"},null,8,C),s("img",{src:o.status_33,width:"185",class:"right-33"},null,8,S),k,y,s("div",V,[s("div",F,[s("div",I,[P,r(s("input",{type:"text",id:"name","onUpdate:modelValue":t[0]||(t[0]=e=>o.name=e),placeholder:"\u8BF7\u8F93\u5165\u59D3\u540D"},null,512),[[l,o.name]])])]),s("div",U,[s("div",B,[z,r(s("input",{type:"text",id:"username","onUpdate:modelValue":t[1]||(t[1]=e=>o.username=e),placeholder:"\u8BF7\u8F93\u5165\u8D26\u53F7"},null,512),[[l,o.username]])])]),s("div",D,[s("div",N,[R,r(s("input",{type:"password",id:"password","onUpdate:modelValue":t[2]||(t[2]=e=>o.password=e),placeholder:"\u8BF7\u8F93\u5165\u5BC6\u7801",onFocusin:t[3]||(t[3]=(...e)=>i.close&&i.close(...e)),onFocusout:t[4]||(t[4]=(...e)=>i.open&&i.open(...e)),onChange:t[5]||(t[5]=e=>i.isSafe())},null,544),[[l,o.password]])])]),s("div",T,[s("div",j,[A,r(s("input",{type:"password",id:"confirm","onUpdate:modelValue":t[6]||(t[6]=e=>o.confirmPassword=e),placeholder:"\u518D\u6B21\u8F93\u5165\u5BC6\u7801",onFocusin:t[7]||(t[7]=(...e)=>i.close&&i.close(...e)),onFocusout:t[8]||(t[8]=(...e)=>i.open&&i.open(...e)),onChange:t[9]||(t[9]=e=>i.isSafe())},null,544),[[l,o.confirmPassword]]),E])]),s("div",{class:_(["clk-button safe",o.statusClass])},[s("button",{onClick:t[10]||(t[10]=e=>i.submit())},[s("p",null,p(o.status),1)])],2)])])])}var K=v(g,[["render",M],["__scopeId","data-v-35f270d6"]]);export{K as default};
