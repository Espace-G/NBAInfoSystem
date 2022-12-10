<template>
  <div class="app-wrapper"> 
    <div class="app-content pt-3 p-md-3 p-lg-4">
      <div class="container-xl">
        <div class="row g-3 mb-4 align-items-center justify-content-between">
          <div class="col-auto">
            <h1 class="app-page-title mb-0">TEAM</h1>
          </div>
          <div class="col-auto" v-show="(this.utype == 0)">
            <button type="button" class="btn app-btn-secondary" @click="myfavor">My Favors</button>
            <button type="button" class="btn app-btn-secondary" @click="all">All Teams</button>
          </div>
          <div class="col-auto" v-show="(this.utype == 1)">
            <button type="button" class="btn app-btn-secondary" @click="create">Create Team</button>
          </div>
          <div class="col-auto">
            <div class="page-utilities">
              <div class="row g-2 justify-content-start justify-content-md-end align-items-center">
                <div class="col-auto">
                  <form class="table-search-form row gx-1 align-items-center">
                    <div class="col-auto">
                      <input v-model="searchText" type="text" id="search-orders" name="searchorders" class="form-control search-orders" placeholder="Team Name">
                    </div>
                    <div class="col-auto">
                      <button @click="search" type="button" class="btn app-btn-secondary">Search</button>
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
                        <th class="cell">Team Name</th>
                        <th class="cell">Team City</th>
                        <th class="cell">Team Zone</th>
                        <th class="cell">Arena</th>
                        <th class="cell">Players</th>
                        <th class="cell" v-if="(utype == 1)">Edit</th>
                        <th class="cell" v-if="(utype == 0)">Follow</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="team in teams.slice(down,up)" key="team.tid">
                        <td class="cell">
                          <img style="height:40px" :src="('./img/team_img/' + team.logo)" alt="team_logo">
                        </td>
                        <td class="cell">{{ team.tname }}</td>
                        <td class="cell">{{ team.tcity }}</td>
                        <td class="cell">{{ team.tzone }}</td>
                        <td class="cell"><span>{{ team.arena }}</span></td>            
                        <td class="cell">
                          <a class="btn-sm app-btn-secondary" @click="toPlayer(team.tid)">Lineup</a>
                        </td>
                        <td class="cell" v-if="(this.$route.query.utype == 0)">
                          <a class="btn-sm app-btn-secondary" @click="follow(team.tid)" v-if="!this.favorTid.includes(team.tid)">
                            Follow
                          </a>
                          <a class="btn-sm app-btn-secondary" @click="unfollow(team.tid)" v-if="this.favorTid.includes(team.tid)">
                            Unfollow
                          </a>
                        </td>
                        <td class="cell" v-if="(this.$route.query.utype == 1)">
                          <a class="btn-sm app-btn-secondary" @click="edit(team.tid)" >
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
        <footer class="app-footer">
          <div class="container text-center py-3">
            <small>
              chen
            </small>
          </div>
        </footer><!--//app-footer-->
      </div>
    </div>
  </div>
  <router-view></router-view>
</template>

<script>
import axios from 'axios'
  export default{
    data(){
      return {
        teams:[],
        up:0,
        down:0,
        totalRecord:0,
        currentPage:1,
        searchText:"",
        Tid:[],
        favorTid:[],
        favorTeam:[],
        utype:this.$route.query.utype
      }
    },
    methods:{
      test(){
        alert(typeof(this.utype))
        alert(Boolean(this.utype))
      },
      getTeams(){
        axios.get("/api/team/query").then(response => {
          this.teams=response.data.obj
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
      initPage(){
        this.totalRecord = this.teams.length
        this.up = this.currentPage * 5 > this.totalRecord ? this.totalRecord : this.currentPage * 5
        this.down = (this.currentPage-1) * 5 <= 0 ? 0 : (this.currentPage - 1) * 5
      },
      search(){
        axios.get("/api/team/query",{
          params:{
            search:this.searchText
          }
        }).then(response => {
          console.log(response.data.obj)
          this.teams = response.data.obj
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
            this.favorTeam = []
          }else{
            //已登录
            uid = response.data.obj.uid
            axios.get("/api/favor/team",{
              params:{
                id:uid       
              }
            }).then(res => {
              this.favorTeam = res.data.obj
              this.favorTid = []
              for(var i = 0;i < this.favorTeam.length;i++){
                this.favorTid[i] = this.favorTeam[i].tid
              }
            })
          }
        })
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
                type:"team",
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
                type:"team",
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
      toPlayer(params){
        this.$router.push({
        path:"/teamPlayer",
        query:{
          tid:params
        }})
      },
      all(){
        this.getTeams()
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
            this.teams = this.favorTeam
            this.initPage()
          }
        })
      },
      edit(params){
        this.$router.push({
          path:"/editTeam",
          query:{
            tid:params
          }
        })
      },
      create(){
        this.$router.push({
          path:"/editTeam",
          query:{
            tid:0
          }
        })
      }
    },
    mounted(){
      this.getTeams()
      this.getFavor()
    },
    watch:{
      $route(){
        this.getTeams()
        this.utype = this.$route.query.utype
      }
    }
    

  }
</script>