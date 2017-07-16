<template>
  <div class="fluid container">
    <div class="form-group form-group-lg panel panel-default">
      <div class="panel-heading">
        <h3 class="panel-title">Column Draggable</h3>
      </div>
      <div class="panel-body">
        <draggable class="list-group" element="ul" v-model="column" :options="dragOptions" :move="onMove" @start="isDragging=true" @end="isDragging=false">
          <transition-group type="transition" :name="'flip-list'">
            <li class="list-group-item" v-for="(element,index) in titles" :key="element.order" style="float: left" >
              <!--<i  @click=" element.fixed=! element.fixed" aria-hidden="true"></i>-->
              {{element.name}}
              <column v-bind:dataList="dataList" v-bind:title="element.name" v-bind:index="index"/>
              <!--<span class="badge">{{element.order}}</span>-->
            </li>
          </transition-group>
        </draggable>
      </div>
      <TableVue/>
    </div>
  </div>
</template>

<script>
import draggable from 'vuedraggable'
import column from '../components/Column'
import TableVue from '../components/TableVue'
const message = [ 'vue.draggable', 'draggable', 'component', 'for', 'vue.js 2.0', 'based' , 'on', 'Sortablejs' ]
export default {
  name: 'hello',
  components: {
    draggable,
    column,
    TableVue,
  },
  data () {
    return {
      list: message.map( (name,index) => {return {name, order: index+1, fixed: false}; }),
      list2:[],
      editable:true,
      isDragging: false,
      delayedDragging:false,
      dataList:[["1_1","1_2","1_3","1_4","1_5"],["2_1","2_2","2_3","2_4","2_5"],["3_1","3_2","3_3","3_4","3_5"],["4_1","4_2","4_3","4_4","4_5"],["5_1","5_2","5_3","5_4","5_5"]],
      titles:["t1","t2","t3","t4","t5"].map( (name,index) => {return {name,index, order: index+1, fixed: false}; }),
    }
  },
  methods:{
    orderList () {
      this.list = this.list.sort((one,two) =>{return one.order-two.order; })
    },
    onMove ({relatedContext, draggedContext}) {
      const relatedElement = relatedContext.element;
      const draggedElement = draggedContext.element;
      //console.log(relatedElement.index+1," ",draggedElement.index+1,relatedElement)
     console.log(relatedElement)
      /* var i;
      for (i=0;i<this.dataList.length;i++){
          transfer(this.dataList[i],relatedElement.index,draggedElement.index)
      }*/
      return (!relatedElement || !relatedElement.fixed) && !draggedElement.fixed
    }
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
  }
}
function transfer(array,i1,i2) {
  var i;
  var tmp
  for (i=0;i<array.length;i++){
      if(i == i1){
        tmp = array[i1];
        array[i1] = array[i2];
        array[i2] = tmp;
        return
      }
  }
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
