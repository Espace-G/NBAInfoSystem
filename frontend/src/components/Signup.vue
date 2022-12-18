<template>
  <div class="d-flex flex-column align-content-end">
    <div class="app-auth-body mx-auto">	
      <div class="app-auth-branding mb-4">
        <img class="logo-icon me-2" style="height:100px" src="img/NBALogo.png" alt="logo">
      </div> 
      <h2 class="auth-heading text-center mb-5"> &nbsp; &nbsp;&nbsp; Sign Up &nbsp;&nbsp;&nbsp;</h2>
        <div class="auth-form-container text-start">
             
        <div class="email mb-3">
          <label class="sr-only" for="signin-email">Username</label>
          <input id="signin-email" name="signin-email" type="text" class="form-control " placeholder="Username" 
          v-model="username" required>
        </div><!--//form-group-->
        <div class="password mb-3">
          <label class="sr-only" for="signin-password">Password</label>
          <input id="signin-password" name="signin-password" type="password" class="form-control signin-password" placeholder="Password" 
          v-model="password" required>
        </div><!--//form-group-->
        <div class="password mb-3">
          <label class="sr-only" for="signin-password">Head Shot</label>
          <input type="file"  ref="img" name="img" class="form-control " placeholder="Img">
        </div><!--//form-group-->
        <div style="color:red">{{ errorMsg }}</div>
        <div class="text-center">
          <button class="btn app-btn-primary w-100 theme-btn mx-auto" @click="signup">Sign Up</button>
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
        username:"",
        password:"",
        errorMsg:"",
      }
    },
    methods:{
      signup(){
        let form = new FormData()
        if(this.$refs.img.files[0] != null){
          form.append("headshot",this.$refs.img.files[0])
        }
        form.append("username",this.username)
        form.append("password",this.password)
        let options = {
          url: 'api/user/register',
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
            var msg = response.data.message
            if(msg.indexOf("UNIQUENAME") != -1){
              this.errorMsg = "Username has been registered"
            }else{
              this.errorMsg = response.data.message
            }
          }else{
            this.$router.push("/login")
          }
        })
      }
    }

  }
</script>