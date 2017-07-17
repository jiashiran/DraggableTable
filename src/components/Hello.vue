<template>
  <div class="fluid container">
    <div class="form-group form-group-lg panel panel-default">
      <div class="panel-heading">
        <h3 class="panel-title">Column Draggable</h3>
      </div>
      <div class="panel-body">
        <draggable class="list-group" element="ul" v-model="column" :options="dragOptions" :move="onMove" @start="isDragging=true" @end="isDragging=false">
          <transition-group type="transition" :name="'flip-list'">
            <div v-bind:style="colsStyle" v-for="(element,index) in showCols"
                 :key="element.order"  >
              <column  v-bind:dataList="dataList" v-bind:title="element.name" v-bind:index="index"/>
            </div>
          </transition-group>
        </draggable>
      </div>
    </div>
    <div>
      <div v-for="e in titles">
        <mu-checkbox  v-bind:label="e.name.colName" class="demo-checkbox" c></mu-checkbox>
      </div>
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
Vue.use(MuseUI)
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
          width: '16%',
      },
      editable:true,
      isDragging: false,
      delayedDragging:false,
      dataList:[{"col6":"1_6","col4":"1_4","col5":"1_5","col2":"1_2","col3":"1_3","col1":"1_1"},{"col6":"2_6","col4":"2_4","col5":"2_5","col2":"2_2","col3":"2_3","col1":"2_1"},{"col6":"3_6","col4":"3_4","col5":"3_5","col2":"3_2","col3":"3_3","col1":"3_1"},{"col6":"4_6","col4":"4_4","col5":"4_5","col2":"4_2","col3":"4_3","col1":"4_1"},{"col6":"5_6","col4":"5_4","col5":"5_5","col2":"5_2","col3":"5_3","col1":"5_1"},{"col6":"6_6","col4":"6_4","col5":"6_5","col2":"6_2","col3":"6_3","col1":"6_1"},{"col6":"7_6","col4":"7_4","col5":"7_5","col2":"7_2","col3":"7_3","col1":"7_1"}],
      titles:[{"colName":"col1","show":true},{"colName":"col2","show":true},{"colName":"col3","show":true},{"colName":"col4","show":true},{"colName":"col5","show":true},{"colName":"col6","show":true}].map( (name,index) => {return {name,index, order: index+1, fixed: false}; }),
    }
  },
  methods:{
    orderList () {
      this.list = this.list.sort((one,two) =>{return one.order-two.order; })
    },
    onMove ({relatedContext, draggedContext}) {
      const relatedElement = relatedContext.element;
      const draggedElement = draggedContext.element;
      //console.log(relatedElement)
      return (!relatedElement || !relatedElement.fixed) && !draggedElement.fixed
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
        //console.log(t);
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
}
</script>

<style>
.flip-list-move {
  transition: transform 0.5s;
}

.no-move {
  transition: transform 0s;
}

.ghost {
  opacity: .5;
  background: #C8EBFB;
}

.list-group {
  min-height: 20px;
}

.list-group-item {
  cursor: move;
}

.list-group-item i{
  cursor: pointer;
}
</style>
