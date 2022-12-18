<template>
  <div class="d-flex flex-column align-content-end">
    <div class="app-auth-body mx-auto">	
      <div class="app-auth-branding mb-4">
        <img class="logo-icon me-2" style="height:70px" src="img/NBALogo.png" alt="logo">
      </div> 
      
      <h2 class="auth-heading text-center mb-5">Edit Player Information</h2>
        <div class="auth-form-container text-start">
          <div v-show="(this.$route.query.pid != 0)">
            <span class="sr-only">cn-Name</span>
            <input type="text" v-model="editPlayer.cnname" class="form-control " placeholder="cn-Name">
            <label class="sr-only">en-Name</label>
            <input type="text" v-model="editPlayer.enname" class="form-control " placeholder="en-Name">
            <label class="sr-only" >Height</label>
            <input type="text" v-model="editPlayer.height" class="form-control " placeholder="height">
            <label class="sr-only" >Weight</label>
            <input type="text" v-model="editPlayer.weight" class="form-control " placeholder="Weight">
            <label class="sr-only" >Nation</label>
            <input type="text" v-model="editPlayer.nation" class="form-control " placeholder="Nation">
            <label class="sr-only" >Age</label>
            <input type="text" v-model="editPlayer.age" class="form-control " placeholder="Age">
            <select v-model="editPlayer.tid" class="form-select w-auto" >
              <option :value="team.tid" v-for="team in teamList" key="team.tid" :selected="(team.tid==editPlayer.tid)">{{ team.tname }}</option>
            </select>
            <div class="text-center">
            <button class="btn app-btn-primary w-100 theme-btn mx-auto" @click="update">Update</button>
            <hr>
            <button class="btn app-btn-primary w-100 theme-btn mx-auto" @click="deletePlayer">Delete</button>
          </div>
          </div>
          <div v-show="(this.$route.query.pid == 0)">
            <div class="email mb-3">
              <label class="sr-only" >cn-Name</label>
              <input type="text" v-model="createPlayer.cnname" class="form-control" placeholder="cn-Name">
            </div><!--//form-group-->
            <div class="password mb-3">
              <label class="sr-only" for="signin-password">en-Name</label>
              <input type="text" v-model="createPlayer.enname" class="form-control" placeholder="en-Name">
            </div><!--//form-group-->
            <div class="password mb-3">
              <label class="sr-only" for="signin-password">Nation</label>
              <input type="text" v-model="createPlayer.nation" class="form-control " placeholder="Nation">
            </div><!--//form-group-->
            <div class="password mb-3">
              <label class="sr-only" for="signin-password">Height</label>
              <input type="text" v-model="createPlayer.height" class="form-control " placeholder="Height">
            </div><!--//form-group-->
            <div class="password mb-3">
              <label class="sr-only" for="signin-password">Weight</label>
              <input type="text" v-model="createPlayer.weight" class="form-control " placeholder="Weight">
            </div><!--//form-group-->
            <div class="password mb-3">
              <label class="sr-only" for="signin-password">Age</label>
              <input type="text" v-model="createPlayer.age" name="age" class="form-control " placeholder="Age">
            </div><!--//form-group-->
            <div class="password mb-3">
              <label class="sr-only" for="signin-password">Player Img</label>
              <input type="file"  ref="img" name="img" class="form-control " placeholder="Img" @change="imgChange">
            </div><!--//form-group-->
            <select v-model="createPlayer.tid" class="form-select w-auto" >
              <option :value="team.tid" v-for="team in teamList" key="team.tid">{{ team.tname }}</option>
            </select>
            <div class="btn">
              <input type="submit" class="form-control" value="submit" @click="create">
              <p style="color:red"> {{ errorMsg }}</p>
            </div>
            
          </div>
        </div><!--//auth-form-container-->	
        
    </div><!--//auth-body-->
  </div>
</template>

<script>
import axios from 'axios'
  export default{
    data(){
      return{
        editPlayer:{},
        teamList:[],
        createPlayer:{},
        data:{},
        errorMsg:""
      }
    },
    methods: { 
      create(){
        let form = new FormData()
        let img = this.$refs.img
        if(img.files[0] != null){ form.append("img",img.files[0]) }
        form.append("username","zhangsan")
        form.append("cnname",this.createPlayer.cnname)
        form.append("enname",this.createPlayer.enname)
        form.append("nation",this.createPlayer.nation)
        form.append("height",this.createPlayer.height)
        form.append("weight",this.createPlayer.weight)
        form.append("age",this.createPlayer.age)
        form.append("tid",this.createPlayer.tid)
        let options = {
          url: 'api/player/insertPlayer',
          data: form,
          method: 'post',
          headers: {
              'Content-Type': 'multipart/form-data'
          }
        }
        axios(options).then(response => {
          if(response.data.status == 0){
            this.errorMsg = "Error Data!Please Check Again!"
          }else{ this.$router.go(-1) }
        })
      },  
      getPlayer(){
        axios.get("/api/player/selectPlayerById",{
          params:{
            id:this.$route.query.pid
          }
        }).then(response => {
          //console.log(response.data.obj)
          this.editPlayer = response.data.obj
        })
      },
      getTeamList(){
        axios.get("/api/team/query").then(response => {
          this.teamList = response.data.obj
        })
      },
      update(){
        var params = new URLSearchParams()
        params.append('pid',this.$route.query.pid)
        params.append('cnname', this.editPlayer.cnname)
        params.append('enname', this.editPlayer.enname)
        params.append('height', this.editPlayer.height)
        params.append('weight', this.editPlayer.weight)
        params.append('nation', this.editPlayer.nation)
        params.append('age', this.editPlayer.age)
        params.append("tid",this.editPlayer.tid)
        axios.post("/api/player/update",params).then(this.$router.go(-1))
      },
      deletePlayer(){
        axios.get("/api/player/remove",{
          params:{
            id:this.$route.query.pid
          }
        }).then(
          this.$router.go(-1)
        )
      },
      
      test(){
        this.$router.go(-1)
      }
    },
    mounted(){
      if(this.$route.query.pid != 0){
        this.getPlayer()
        this.getTeamList()
      }else{
        this.getTeamList()
        this.createPlayer = {
          tid:0,
          cnname:"",
          enname:"",
          nation:"United States" ,
          height:"",
          weight:"",
          age:"",
        }
      }
    }

  }
</script>