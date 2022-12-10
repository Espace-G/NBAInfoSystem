<template>
  <div class="d-flex flex-column align-content-end">
    <div class="app-auth-body mx-auto">	
      <div class="app-auth-branding mb-4">
        <img class="logo-icon me-2" style="height:100px" src="img/NBALogo.png" alt="logo">
      </div> 
      <h2 class="auth-heading text-center mb-5">Log in to NBA Info System</h2>
        <div class="auth-form-container text-start">
             
        <div class="email mb-3">
          <label class="sr-only" for="signin-email">Username</label>
          <input id="signin-email" name="signin-email" type="text" class="form-control " placeholder="Username" 
          v-model="loginUsername" required>
        </div><!--//form-group-->
        <div class="password mb-3">
          <label class="sr-only" for="signin-password">Password</label>
          <input id="signin-password" name="signin-password" type="password" class="form-control signin-password" placeholder="Password" 
          v-model="loginPassword" required>
          <p style="color:red "> {{ errMsg }}</p>
          <div class="extra mt-3 row justify-content-between">
          </div><!--//extra-->
        </div><!--//form-group-->
        <div class="text-center">
          <button class="btn app-btn-primary w-100 theme-btn mx-auto" @click="login">Log In</button>
        </div>
     
      <!-- 注册 -->
      <div class="auth-option text-center pt-5">No Account? Sign up <a class="text-link" @click="signup">here</a>.</div>
    </div><!--//auth-form-container-->	

    </div><!--//auth-body-->
  </div>
</template>


<script>
  import axios from 'axios';
  export default{
    data(){
      return{
      loginUsername:"",
      loginPassword:"",
      errMsg:"",
      }
    },
    methods: {
      signup(){
        this.$router.push("/signup")
      },
      login(){
        var params = new URLSearchParams();
        params.append('username', this.loginUsername);
        params.append('password', this.loginPassword);
        axios.post("/api/user/checkLogin",params).then(response =>{
          //console.log(response)
          if(response.data.status == 1){
            this.$router.push("/")
          }else{
            this.errMsg = response.data.message
          }
        })
      }
    },
    mounted(){
    }
  }
</script>