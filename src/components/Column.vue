<template>
  <div>
    <ul >
      <div v-for="(val,i) in dataList" ><!---->
        title={{title}}
          <div v-for="(va,ind) in val"  >
            <li v-if="ind==index" style="width:100px;height:30px;float:left;text-align:center;border-right:1px solid #000;border-bottom:1px solid #000;">
              {{va}}
            </li>
          </div>
      </div>
    </ul>
  </div>
</template>
<script>
  export default {
    name:'column',
    props: ['dataList',"title","index"],
    data () {
          return{
              list1:[],
          }
    },
    methods:{
        toData (list) {
            return list.map( (name,index) => {return {name, index}; })
        }
    }
  }
</script>
<style>
  li {
    list-style-type:none;
  }
  td1 {
    border-collapse:collapse;
    border: 1px solid black;
  }
</style>
