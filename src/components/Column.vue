<template>
  <div>
    <div>
      <mu-table :showCheckbox="false" ref="table">
        <mu-thead >
          <mu-tr>
            <mu-th>
              {{title.alias}}
            </mu-th>
          </mu-tr>
        </mu-thead>
        <mu-tbody>
          <mu-tr  v-for="(val,i) in dataList">
            <div v-for="(va,key) in val" >
              <mu-td v-if="key==title.colName">
                {{va}}
              </mu-td>
            </div>
          </mu-tr>
          <mu-tr>
            <mu-td>{{sum}}</mu-td>
          </mu-tr>
        </mu-tbody>
      </mu-table>
    </div>
    <div>
    </div>
  </div>
</template>
<script>
  import Vue from 'vue'
  import MuseUI from 'muse-ui'
  import 'muse-ui/dist/muse-ui.css'
  import light from 'muse-ui/dist/theme-default.css'
  import dark from 'muse-ui/dist/theme-dark.css'
  import carbon from 'muse-ui/dist/theme-carbon.css'
  import teal from 'muse-ui/dist/theme-teal.css'
  Vue.use(MuseUI)
  export default {
    name:'column',
    props: ['dataList',"title","index"],
    data () {
          return{
            theme: 'teal',
            themes: {
              light,
              dark,
              carbon,
              teal
            },
            list1:[],
            titleEdit:false,
            sum:0,
          }
    },
    methods:{
      toData (list) {
            return list.map( (name,index) => {return {name, index}; })
       },
      changeTheme (theme) {
        this.theme = theme
        const styleEl = this.getThemeStyle()
        styleEl.innerHTML = this.themes[theme] || ''
      },
      getThemeStyle () {
        const themeId = 'muse-theme'
        let styleEl = document.getElementById(themeId)
        if (styleEl) return styleEl
        styleEl = document.createElement('style')
        styleEl.id = themeId
        document.body.appendChild(styleEl)
        return styleEl
      },
      Click () {
        // `this` 在方法里指当前 Vue 实例
        console.log(event)
        alert('Hello ' + this.name + '!')
        // `event` 是原生 DOM 事件
        if (event) {
          alert(event.target.tagName)
        }
      },
      sumCal (data) {
        this.sum = this.sum + data.length;
        return data;
      },
      addStr (str){
          return '（' + str + '）'
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
  .demo-list {
    width: 200px
  }
</style>
