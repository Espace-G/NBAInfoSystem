<template>
  <div class="app-wrapper"> 
    <div class="app-content pt-3 p-md-3 p-lg-4">
      <div class="container-xl">
        <div class="row g-3 mb-4 align-items-center justify-content-between">
          <div class="col-auto">
            <h1 class="app-page-title mb-0">PLAYER</h1>
          </div>
          <div class="col-auto">
            <div class="page-utilities">
              <div class="row g-2 justify-content-start justify-content-md-end align-items-center">
                <div class="col-auto">
                  <form class="table-search-form row gx-1 align-items-center">
                    <div class="col-auto" v-show="(this.utype == 0)">
                      <button type="button" class="btn app-btn-secondary" @click="myfavor">My Favors</button>
                    </div>
                    <div class="col-auto" v-show="(this.utype == 1)">
                      <button type="button" class="btn app-btn-secondary" 
                      @click="create">Create Player</button>
                    </div>
                    <div class="col-auto">
                      <button type="button" class="btn app-btn-secondary" @click="order('age')">Order by Age</button>
                    </div>
                    <div class="col-auto">
                      <button type="button" class="btn app-btn-secondary" @click="order('height')">Order by Height</button>
                    </div>
                    <div class="col-auto">
                      <button type="button" class="btn app-btn-secondary" @click="order('weight')">Order by Weight</button>
                    </div>
                    <div class="col-auto">
                      <input v-model="searchText" type="text" class="form-control" placeholder="Player Name">
                    </div>
                    <div class="col-auto">
                      <button type="button" class="btn app-btn-secondary" @click="search">Search</button>
                    </div>
                  </form>               
                </div><!--//col-->
                
                
              </div><!--//row-->
            </div><!--//table-utilities-->
          </div><!--//col-auto-->
        </div><!--//row-->
       
        <div class="tab-content" id="orders-table-tab-content">
          <div class="tab-pane fade show active" id="orders-all" role="tabpanel" aria-labelledby="orders-all-tab">
            <div class="app-card app-card-orders-table shadow-sm mb-5">
              <div class="app-card-body">
                <div class="table-responsive">
                  <table class="table app-table-hover mb-0 text-left">
                    <thead>
                      <tr>
                        <th class="cell"></th>
                        <th class="cell">Cn Name</th>
                        <th class="cell">En Name</th>
                        <th class="cell">Height</th>
                        <th class="cell">Weight</th>
                        <th class="cell">Nation</th>
                        <th class="cell">Age</th>
                        <th class="cell" v-show="(this.$route.query.utype == 0)">Follow</th>
                        <th class="cell" v-show="(this.$route.query.utype == 1)">Edit</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="player in (players).slice(down,up)" :key="player.pid">
                        <td class="cell">
                          <img style="height:40px" :src="('./img/player_img/' + player.imgpath)" 
                          alt="player_img">
                        </td>
                        <td class="cell">{{ player.cnname }}</td>
                        <td class="cell">{{ player.enname }}</td>
                        <td class="cell">{{ player.height }}</td>
                        <td class="cell">{{ player.weight }}</td>
                        <td class="cell">{{ player.nation }}</td>
                        <td class="cell">{{ player.age }}</td>            
                        <!-- 收藏操作 -->
                        <td class="cell" v-show="(this.utype == 0)">
                          <a class="btn-sm app-btn-secondary" @click="follow(player.pid)" 
                          v-if="!this.favorPid.includes(player.pid)">
                            Follow
                          </a>
                          <a class="btn-sm app-btn-secondary" @click="unfollow(player.pid)" 
                          v-if="this.favorPid.includes(player.pid)">
                            Unfollow
                          </a>
                        </td>
                        <!-- 更新信息操作 -->
                        <td class="cell" v-show="(this.utype == 1)">
                          <a class="btn-sm app-btn-secondary" @click="edit(player.pid)" >
                            Edit
                          </a>
                        </td>
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
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
  export default{
    data(){
      return {
        players:[],
        favorPlayers:[],
        currentPage:1,
        totalRecord:0,
        up:0,
        down:0,
        searchText:"",
        favorPid:[],
        utype:this.$route.query.utype
      }
    },
    methods:{
      getPlayers(){
        axios.get("/api/player/selectPlayers").then(response => {
          this.players = response.data.obj
          this.currentPage = 1
          this.initPage()
        })
      },
      pageAdd(){
        if(this.up == this.totalRecord){
          alert("It is Last Page")
        }else{
          this.currentPage++
          this.initPage()
        }
      },
      pageSub(){
        if(this.down == 0){
          alert("It is First Page")
        }else{
          this.currentPage--
          this.initPage()
        }
      },
      search(){
        axios.get("/api/player/selectPlayers",{
          params:{
            search:this.searchText
          }
        }).then(response => {
          console.log(response.data.obj)
          this.players = response.data.obj
          this.currentPage = 1
          this.initPage()
        })
      },
      order(params){
        axios.get("/api/player/selectPlayers",{
          params:{
            orderby:params
          }
        }).then(response => {
          console.log(response.data)
          console.log(response.data.obj)
          this.players = response.data.obj
          this.currentPage = 1
          this.initPage()
        })
      },
      getFavor(){
        var uid
        axios.get("/api/user/getLogining").then(response => {
          //console.log(response.data.obj)
          if(response.data.obj == null){
            //未登录
            this.favorPlayers = []
          }else{
            //已登录
            uid = response.data.obj.uid
            axios.get("/api/favor/player",{
              params:{
                id:uid       
              }
            }).then(res => {
              this.favorPlayers = res.data.obj
              this.favorPid = []
              for(var i = 0;i < this.favorPlayers.length;i++){
                this.favorPid[i] = this.favorPlayers[i].pid
              }
            })
          }
        })
      },
      myfavor(){
        var uid
        axios.get("/api/user/getLogining").then(response => {
          //console.log(response.data.obj)
          if(response.data.obj == null){
            //未登录
            if(confirm("Please Login First.")){
              this.$router.push("/login")
            }
          }else{
            //已登录，将players替换为favorPlayers
            this.players = this.favorPlayers
            this.initPage()
          }
        })
      },
      //每次players或currentPage有变化时调用，更新列表信息
      initPage(){
        this.totalRecord = this.players.length
        this.up = this.currentPage * 10 > this.totalRecord ? this.totalRecord : this.currentPage * 10
        this.down = (this.currentPage-1) * 10 <= 0 ? 0 : (this.currentPage - 1) * 10
      },
      follow(params){
        //alert(params)
        var uid
        axios.get("/api/user/getLogining").then(response => {
          if(response.data.obj == null){
            //未登录
            if(confirm("Please Login First.")){
              this.$router.push("login")
            }
          }else{
            uid = response.data.obj.uid
            axios.get("/api/favor/create",{
              params:{
                type:"player",
                uid:uid,
                id:params
              }
            }).then(res => {
              //console.log(res.data)
              this.getFavor()
            })
          }
        })
      },
      unfollow(params){
        var uid
        axios.get("/api/user/getLogining").then(response => {
          if(response.data.obj == null){
            //未登录
            if(confirm("Please Login First.")){
              this.$router.push("login")
            }
          }else{
            uid = response.data.obj.uid
            axios.get("/api/favor/undo",{
              params:{
                type:"player",
                uid:uid,
                id:params
              }
            }).then(res => {
              //console.log(res.data)
              this.getFavor()
            })
          }
        })
      },
      edit(params){
        this.$router.push({
          path:"/editPlayer",
          query:{
            pid:params
          }
        })
      },
      create(){
        this.$router.push({
          path:"/editPlayer",
          query:{
            pid:0
          }
        })
      }
    },
    mounted() {
      this.getPlayers()
      this.getFavor()
    },
    watch:{
      $route(){
        this.getPlayers()
        this.utype = this.$route.query.utype
      }
    }
    
  }
</script>