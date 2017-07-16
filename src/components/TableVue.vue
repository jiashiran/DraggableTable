<template>
  <div id="people" class="container">
    <div>
      <v-client-table :data="tableData" :columns="columns" :options="options"></v-client-table>
    </div>
    <div>
      <div>
        <draggable v-model="columns" :options="{group:'people'}" @start="drag=true" @end="drag=false">
             <div v-for="element in columns">{{element}}</div>
        </draggable>
      </div>
    </div>
  </div>
</template>
<script>
import Vue from 'vue'
import {ServerTable, ClientTable, Event} from 'vue-tables-2';
import draggable from 'vuedraggable'
Vue.use(ClientTable, {
  compileTemplates: true,
  //highlightMatches: true,
  //pagination: {
  // dropdown:true
  // chunk:5
  // },
  filterByColumn: true,
  texts: {
    filter: "Search:"
  },
 /* datepickerOptions: {
    showDropdowns: true
  }*/
  //skin:''
});

export default({
  element: "#people",
  components: {
    ClientTable,
    draggable,
  },
  data () {
    return {
    columns: ['name', 'age','id'],
    //tableTitles: this.columns.map( (name,index) => {return {name, order: index+1, fixed: false}; }),
    tableData: [
      {id: 1, name: "John", age: "20"},
      {id: 2, name: "Jane", age: "24"},
      {id: 3, name: "Susan", age: "16"},
      {id: 4, name: "Chris", age: "55"},
      {id: 5, name: "Dan", age: "40"}
    ],
    options: {
      filterByColumn: true,
      listColumns: {
        animal: [
          { id:1, text: 'Dog' },
          { id:2, text: 'Cat' },
          { id:3, text: 'Tiger' },
          { id:4, text: 'Bear' }
          ]
        }
      }
    }
  }
})

// Courtesy of Tomasz Nurkiewicz (http://stackoverflow.com/questions/9035627/elegant-method-to-generate-array-of-random-dates-within-two-dates)

function randomDate(start, end) {
    return start.getTime();
  //return moment(start.getTime() + Math.random() * (end.getTime() - start.getTime()));
}

function randomOption() {
  return Math.floor(Math.random() * 5);
}

</script>
<style>
  #people {
    text-align: center;
    width: 95%;
    margin: 0 auto;
  }

  h2 {
    margin-bottom: 30px;
  }

  th,
  td {
    text-align: left;
  }

  th:nth-child(n+2),
  td:nth-child(n+2) {
    text-align: center;
  }

  thead tr:nth-child(2) th {
    font-weight: normal;
  }
</style>
