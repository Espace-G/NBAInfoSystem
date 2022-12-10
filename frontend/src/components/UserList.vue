<template>
  <div class="tab-content" id="orders-table-tab-content">
          <div class="tab-pane fade show active" id="orders-all" role="tabpanel" aria-labelledby="orders-all-tab">
            <div class="app-card app-card-orders-table shadow-sm mb-5">
              <div class="app-card-body">
                <div class="table-responsive">
                  <table class="table app-table-hover mb-0 text-left">
                    <thead>
                      <tr>
                        <th class="cell"></th>
                        <th class="cell">Username</th>
                        <th class="cell">Usertype</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="user in (userList).slice(down,up)" :key="user.uid">
                        <td class="cell">
                          <img style="height:40px" :src="('./img/headshot/' + user.headshot)" alt="user_headshot">
                        </td>
                        <td class="cell">{{ user.username }}</td>
                        <td class="cell">{{ user.utype }}</td>
                      </tr>
                    </tbody>
                  </table>
                </div><!--//table-responsive-->
                    
              </div><!--//app-card-body-->		
            </div><!--//app-card-->
            <nav class="app-pagination">
              <button class="btn app-btn-secondary" @click="pageSub">previous</button>
              <button class="btn app-btn-secondary" @click="pageAdd">next</button>
            </nav><!--//app-pagination-->
            
          </div><!--//tab-pane-->
        </div>
</template>


<script>
import axios from 'axios'
  export default{
    data(){
      return {
        userList:[],
        currentPage:1,
        totalRecord:0,
        up:0,
        down:0,
      }
    },
    methods:{
      initPage(){
        this.totalRecord = this.userList.length
        this.up = this.currentPage * 10 > this.totalRecord ? this.totalRecord : this.currentPage * 10
        this.down = (this.currentPage-1) * 10 <= 0 ? 0 : (this.currentPage - 1) * 10
      },
      getList(){
        axios.get("/api/user/selectAll").then(response => {
          this.userList = response.data.obj
          this.initPage()
        })
      },
      watch:{
        $route(){
          this.getList()
        }
      }
    },
    mounted(){
      this.getList()
    }

  }
</script>