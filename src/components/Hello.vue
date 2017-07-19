<template>
  <div class="layout">
    <div class="header">
      <div class="logo">
        可拖动列表
      </div>
      <div class="nav">
      </div>
    </div>
    <div class="content">
      <!--<div class="breadcrumb">
      </div>-->
      <div class="body">
        <!--<mu-sub-header>阳光</mu-sub-header>-->
        <mu-content-block>
          <div class="form-group form-group-lg panel panel-default">
            <div class="panel-heading">
              <h3 class="panel-title"></h3>
            </div>
            <div class="panel-body" style="width: 1500px">
              <draggable class="list-group" element="div" v-model="titles"  :options="dragOptions" :move="onMove" @start="isDragging=true" @end="isDragging=true">
                <transition-group type="transition" :name="'flip-list'" >
                  <div :style="colsStyle" v-for="(element,index) in showCols" :key="element.order"  >
                    <column  v-bind:dataList="dataList" v-bind:title="element.name" v-bind:index="index"/>
                  </div>
                </transition-group>
              </draggable>
            </div>
          </div>
          <div>
            <mu-raised-button label="保存列信息" @click="saveColInfo()"/>
            <mu-raised-button label="选择显示列" @click="toggle()"/>
            <mu-drawer right :open="open" :docked="docked" @close="toggle()">
              <mu-list @itemClick="docked ? '' : toggle()">
                <div v-for="e in titles">
                  <mu-list-item v-bind:title="e.name.colName">
                    <mu-text-field @focus="oldColNamed(e.name.colName)"  @blur="changing($event,e.name.colName)"  label="设置列名"  :maxLength="10" v-model="e.name.alias"/>
                    <mu-checkbox  v-bind:label="showInfo(e.name.show)" v-model="e.name.show" />
                  </mu-list-item>
                </div>
                <mu-list-item v-if="docked" @click.native="open = false" title="关闭"/>
              </mu-list>
            </mu-drawer>
          </div>
        </mu-content-block>
      </div>
    </div>
    <div class="footer">
    </div>
  </div>
</template>

<script>
import draggable from 'vuedraggable'
import column from '../components/Column'

import Vue from 'vue'
import MuseUI from 'muse-ui'
import 'muse-ui/dist/muse-ui.css'
import light from 'muse-ui/dist/theme-default.css'
import dark from 'muse-ui/dist/theme-dark.css'
import carbon from 'muse-ui/dist/theme-carbon.css'
import teal from 'muse-ui/dist/theme-teal.css'

import VueResource from 'vue-resource'

Vue.use(VueResource)
Vue.use(MuseUI)

Vue.http.options.emulateHTTP = true;
Vue.http.options.emulateJSON = true;
export default {
  name: 'hello',
  components: {
    draggable,
    column,
  },
  data () {
    return {
      colsStyle:{
          float:'left',
          width: '5%',
      },
      column:column,
      editable:true,
      isDragging: false,
      delayedDragging:false,
      dataList:[],//[{"col6":"1_6","col4":"1_4","col5":"1_5","col2":"1_2","col3":"1_3","col1":"1_1"},{"col6":"2_6","col4":"2_4","col5":"2_5","col2":"2_2","col3":"2_3","col1":"2_1"},{"col6":"3_6","col4":"3_4","col5":"3_5","col2":"3_2","col3":"3_3","col1":"3_1"},{"col6":"4_6","col4":"4_4","col5":"4_5","col2":"4_2","col3":"4_3","col1":"4_1"},{"col6":"5_6","col4":"5_4","col5":"5_5","col2":"5_2","col3":"5_3","col1":"5_1"},{"col6":"6_6","col4":"6_4","col5":"6_5","col2":"6_2","col3":"6_3","col1":"6_1"},{"col6":"7_6","col4":"7_4","col5":"7_5","col2":"7_2","col3":"7_3","col1":"7_1"}],
      titles:[],//[{"col6":"1_6","col4":"1_4","col5":"1_5","col2":"1_2","col3":"1_3","col1":"1_1"},{"col6":"2_6","col4":"2_4","col5":"2_5","col2":"2_2","col3":"2_3","col1":"2_1"},{"col6":"3_6","col4":"3_4","col5":"3_5","col2":"3_2","col3":"3_3","col1":"3_1"},{"col6":"4_6","col4":"4_4","col5":"4_5","col2":"4_2","col3":"4_3","col1":"4_1"},{"col6":"5_6","col4":"5_4","col5":"5_5","col2":"5_2","col3":"5_3","col1":"5_1"},{"col6":"6_6","col4":"6_4","col5":"6_5","col2":"6_2","col3":"6_3","col1":"6_1"},{"col6":"7_6","col4":"7_4","col5":"7_5","col2":"7_2","col3":"7_3","col1":"7_1"}].map( (name,index) => {return {name,index, order: index+1, fixed: false}; }),
      showTitles:[],
      open: false,
      docked: true,
      activeTab: 'tab1',
      oldColName:"",
    }
  },
  methods:{
    orderList () {
      this.list = this.list.sort((one,two) =>{return one.order-two.order; })
    },
    onMove ({relatedContext, draggedContext}) {
      const relatedElement = relatedContext.element;
      const draggedElement = draggedContext.element;
      //console.log(draggedContext)
      //console.log(relatedContext)
      return (!relatedElement || !relatedElement.fixed) && !draggedElement.fixed
    },
    toggle (flag) {
      this.open = !this.open
      this.docked = !flag
    },
    showInfo (show) {
      if(show){
        return '显示';
      }else {
        return '隐藏'
      }
    },
    saveColInfo(){
      var array = new Array();
      //console.log(this.titles)
      this.titles.forEach(function (t) {
         var object = {};
         object.colName = t.name.colName;
         object.alias = t.name.alias;
         object.show = t.name.show;
         array.push(object)
      })
      console.log(JSON.stringify(array))
      Vue.http.get('http://localhost:8089/saveColInfo',{params: {colInfo: JSON.stringify(array)}}).then((response) => {
        console.log(response);
      }, (response) => {

      });
    },
    handleTabChange (val) {
      this.activeTab = val
    },
    changing (event,colName){
      var alias = event.target.value;
      var title;
      var repick = false;
      var oldColName = this.oldColName
      this.titles.forEach(function (t) {
          //console.log(oldColName,t)
        if(oldColName == String(t.name.colName)){
          title = t;
          //console.log("old col name:",title)
        }
        if(t.name.colName!=colName && t.name.alias==alias){
          //console.log("别名重复",alias,colName,event)
          repick = true;
        }
      });
      if(repick){
        title.name.alias = alias + " new"
      }
    },
    oldColNamed (event){
      //console.log(event)
      this.oldColName = event
    },
  },
  computed: {
    dragOptions () {
      return  {
        animation: 0,
        group: 'description',
        disabled: !this.editable,
        ghostClass: 'ghost'
      };
    },
    listString(){
      return JSON.stringify(this.list, null, 2);
    },
    list2String(){
      return JSON.stringify(this.list2, null, 2);
    },
    showCols(){
      return this.titles.filter(function (t) {
        return t.name.show;
      });
    }
  },
  watch: {
    isDragging (newValue) {
      if (newValue){
        this.delayedDragging= true
        return
      }
      this.$nextTick( () =>{
           this.delayedDragging =false
      })
    }
  },
  beforeCreate () {
    Vue.http.get('http://localhost:8089/colNames',{name: 'jobList'}).then((response) => {
      console.log(response);
      let json = JSON.parse(response.bodyText);
      this.titles = json.map( (name,index) => {return {name,index, order: index+1, fixed: false}; });
      console.log("titles:",json)
    }, (response) => {

    });
    Vue.http.get('http://localhost:8089/dataList', {name: 'jobList'}).then((response) => {
      console.log(response);
      let json = JSON.parse(response.bodyText);
      var sum = this.titles;
      json.forEach(function (d) {
        Object.entries(d).forEach(function (v,i) {
          sum.forEach(function (t) {
            if(t.name.colName == v[0]){
                if(t.name.sum==undefined){
                  t.name.sum = v[1]
                }else {
                  t.name.sum = t.name.sum + v[1]
                }
            }
          })
        })
      });
      this.titles = sum;
      this.dataList = json;
      this.colsStyle.width = String(100/this.dataList.length) + '%';
      console.log(this.colsStyle.width);
    }, (response) => {

    })
  }
}

</script>
<style>
.list-group {
  min-height: 10px;
}

.list-group-item i{
  cursor: pointer;
}

.layout{
  background-color: rgb(236, 236, 236);
}

.header{
  background-color: #7e57c2;
}

.logo{
  font-size: 18px;
  color: white;
  display: inline-block;
  padding: 10px 20px;
}

.nav{
  display: inline-block;
  width: calc(100% - 150px);
  margin: 0 auto;
}

.tab{
  margin: 0 auto;
  width: 400px;
  background-color: rgba(0, 0, 0, 0);
}

.content{
  width: 90%;
  margin: 0 auto;
}

.breadcrumb{
  margin: 10px 0;
}

.body{
  background-color: white;
  border-radius: 5px;
  min-height: 500px;
}

.footer{
  padding: 20px 0;
  text-align: center;
}
</style>
