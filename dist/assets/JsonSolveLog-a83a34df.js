import{u as V,r as a,b as k,x as z,w as o,e as l,h as p,X as M,t as m,z as f,d,A as $,_ as q,m as G,ad as I,ae as O,a0 as X,af as H,a4 as K,ag as Q,ah as U}from"./index-bed6c77c.js";/* empty css                 */import{_ as W}from"./PageContainer-634a4ae5.js";/* empty css                     *//* empty css                    */import"./el-tooltip-4ed993c7.js";import{g as Y,d as Z,f as ee}from"./format-a2d930ed.js";import{i as ae}from"./index-c25176ff.js";/* empty css                   */import"./_plugin-vue_export-helper-c27b6911.js";const te=d("br",null,null,-1),oe=d("br",null,null,-1),le=d("br",null,null,-1),fe={__name:"JsonSolveLog",setup(ne){const v=V(),_=a(!1),h=a([]),c=a(0),E=a(0),x=a([]),y=a([]),b=a([]),L=a([]),g=a(null),C=a([]),s=a({page:1,size:7}),r=async()=>{_.value=!0;const e=await Y(s.value);h.value=e.data.data,c.value=e.data.length,E.value=e.data.total_page,_.value=!1,console.log(c.value),x.value=e.data.data.map(t=>t.total_score.toFixed(2)),y.value=e.data.data.map(t=>t.page_error.toFixed(2)),b.value=e.data.data.map(t=>t.page_load.toFixed(2)),L.value=e.data.data.map(t=>t.page_experience.toFixed(2)),j()};r();const T=e=>{s.value.page=1,s.value.size=e,r()},B=e=>{s.value.page=e,r()},D=async e=>{await $.confirm("此操作将永久删除该记录, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}),await Z(e.file_name),console.log("删除成功"+e.file_name),setTimeout(()=>{r(),console.log("列表刷新成功"+e.file_name),q.success("删除成功"),console.log(c.value)},300)},w=()=>{v.push("/json/solve")},F=e=>{console.log(e.file_name),console.log(typeof e.file_name),v.push({path:"/json/show",query:{file_name:e.file_name}})},j=()=>{if(C.value=[{name:"综合得分",data:x.value},{name:"页面报错得分",data:y.value},{name:"页面加载得分",data:b.value},{name:"页面体验得分",data:L.value}],g.value){const e=ae(g.value),t={title:{text:"最近七次解析反馈数据"},tooltip:{trigger:"axis"},legend:{data:["综合得分","页面报错得分","页面加载得分","页面体验得分"]},grid:{left:"3%",right:"4%",bottom:"0%",containLabel:!0},toolbox:{feature:{saveAsImage:{}}},xAxis:{type:"category",boundaryGap:!1,data:["第一次解析","第二次解析","第三次解析","第四次解析","第五次解析","第六次解析","第七次解析"]},yAxis:{type:"value"},series:C.value.map(i=>({name:i.name,data:i.data,type:"line"}))};e.setOption(t)}};return(e,t)=>{const i=G,S=I,u=O,A=X,J=H,N=W,P=K;return k(),z(N,{title:"解析记录"},{extra:o(()=>[l(i,{type:"primary",onClick:w},{default:o(()=>[p("添加解析")]),_:1})]),default:o(()=>[M((k(),z(A,{data:h.value},{default:o(()=>[l(u,{label:"解析文件名",prop:"title"},{default:o(({row:n})=>[l(S,{type:"primary",underline:!1},{default:o(()=>[p(m(n.file_name),1)]),_:2},1024)]),_:1}),l(u,{label:"综合得分",prop:"title"},{default:o(({row:n})=>[l(S,{type:"primary",underline:!1},{default:o(()=>[p(m((n.total_score*2).toFixed(2)),1)]),_:2},1024)]),_:1}),l(u,{label:"解析时间",prop:"date"},{default:o(({row:n})=>[p(m(f(ee)(n.create_time)),1)]),_:1}),l(u,{label:"操作"},{default:o(({row:n})=>[l(i,{circle:"",plain:"",type:"primary",icon:f(Q),onClick:R=>F(n)},null,8,["icon","onClick"]),l(i,{circle:"",plain:"",type:"danger",icon:f(U),onClick:R=>D(n)},null,8,["icon","onClick"])]),_:1})]),_:1},8,["data"])),[[P,_.value]]),l(J,{"current-page":s.value.page,"page-size":s.value.size,background:!0,layout:"jumper, total, prev, sizes, pager, next, ->",total:c.value,onSizeChange:T,onCurrentChange:B,style:{"margin-top":"20px","justify-content":"flex-end"}},null,8,["current-page","page-size","total"]),te,oe,d("div",{ref_key:"chartRef",ref:g,style:{width:"100%",height:"400px"}},null,512),le]),_:1})}}};export{fe as default};
