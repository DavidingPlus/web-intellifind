import{ao as t,ap as e}from"./index-bed6c77c.js";const n=o=>t.get("/core/show-result/once",{params:{file_name:o}}),r=o=>t.get("/core/show-history/total",{params:o}),c=o=>{const a={file_name:o};t.delete("/core/delete-history/once",{data:a},{headers:{"Content-Type":"application/json"}})},l=o=>e(o).format("YYYY年MM月DD日 HH:mm:ss");export{n as a,c as d,l as f,r as g};
