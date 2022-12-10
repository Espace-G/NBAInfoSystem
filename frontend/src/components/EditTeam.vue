<template>
  <div class="d-flex flex-column align-content-end">
    <div class="app-auth-body mx-auto">	
      <div class="app-auth-branding mb-4">
        <img class="logo-icon me-2" style="height:70px" src="img/NBALogo.png" alt="logo">
      </div> 
      
      <h2 class="auth-heading text-center mb-5" v-show="(this.$route.query.tid != 0)">Edit Team Information</h2>
      <h2 class="auth-heading text-center mb-5" v-show="(this.$route.query.tid == 0)">Create Team</h2>
        <div class="auth-form-container text-start">
          <div v-show="(this.$route.query.tid != 0)">
            <label class="sr-only">Team Name</label>
            <input type="text" v-model="editTeam.tname" class="form-control " placeholder="Team Name">
            <label class="sr-only" >Team City</label>
            <input type="text" v-model="editTeam.tcity" class="form-control " placeholder="Team City">
            <label class="sr-only" >Team Zone</label>
            <input type="text" v-model="editTeam.tzone" class="form-control " placeholder="Team Zone">
            <label class="sr-only" >Arena</label>
            <input type="text" v-model="editTeam.arena" class="form-control " placeholder="Arena">
            <div class="text-center">
              <button class="btn app-btn-primary w-100 theme-btn mx-auto" @click="update">Update</button>
              <hr>
              <button class="btn app-btn-primary w-100 theme-btn mx-auto" @click="deleteTeam">Delete</button>
            </div>
          </div>

          <div v-show="(this.$route.query.tid == 0)">
            <label class="sr-only">Team Name</label>
            <input type="text" v-model="createTeam.tname" class="form-control " placeholder="Team Name">
            <label class="sr-only" >Team City</label>
            <input type="text" v-model="createTeam.tcity" class="form-control " placeholder="Team City">
            <label class="sr-only" >Team Zone</label>
            <input type="text" v-model="createTeam.tzone" class="form-control " placeholder="Team Zone">
            <label class="sr-only" >Arena</label>
            <input type="text" v-model="createTeam.arena" class="form-control " placeholder="Arena">
            <div class="password mb-3">
              <label class="sr-only" for="signin-password">Player Img</label>
              <input type="file"  ref="img" name="logo_img" class="form-control " placeholder="Img" @change="imgChange">
            </div>
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
      return {
        tid:this.$route.query.tid,
        createTeam:{},
        editTeam:{},
        errorMsg:"",
      }
    },
    methods:{
      create(){
        let form = new FormData()
        let img = this.$refs.img
        if(img.files[0] != null){
          form.append("img",img.files[0])
        }
        form.append("username","zhangsan")
        form.append("tname",this.createTeam.tname)
        form.append("tcity",this.createTeam.tcity)
        form.append("tzone",this.createTeam.tzone)
        form.append("arena",this.createTeam.arena)
        let options = {
          url: 'api/team/insertTeam',
          data: form,
          method: 'post',
          headers: {
              'Content-Type': 'multipart/form-data'
          }
        }
        axios(options).then(response => {
          console.log(response.data)
          if(response.data.status == 0){
            //创建失败显示提示信息
            this.errorMsg = "Error Data!Please Check Again!"
          }else{
            this.$router.go(0)
          }
        })
      },
      getTeam(){
        axios.get("/api/team/selectById",{
          params:{
            id:this.$route.query.tid
          }
        }).then(response => {
          this.editTeam = response.data.obj
        })
      },
      update(){
        var params = new URLSearchParams()
        params.append("tname",this.editTeam.tname)
        params.append("tcity",this.editTeam.tcity)
        params.append("tzone",this.editTeam.tzone)
        params.append("arena",this.editTeam.arena)
        params.append("tid",this.$route.query.tid)
        axios.post("api/team/updateTeam",params).then(response=>{
          console.log(response.data)
          this.$router.go(-1)
        })
      },
      deleteTeam(){
        axios.get("/api/team/deleteTeam",{
          params:{
            id:this.$route.query.tid
          }
        }).then(
          this.$router.go(-1)
        )
      }
    },
    mounted(){
      if(this.$route.query.tid != 0){
        //edit
        this.getTeam()
      }else{
        this.createTeam = {
          tname:"",
          tcity:"",
          tzone:"",
          arena:"",
        }
      }
    }
  }
</script>