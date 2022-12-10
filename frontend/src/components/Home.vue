<template>
  <head>
    <title>Portal - Bootstrap 5 Admin Dashboard Template For Developers</title>
    
    <!-- Meta -->
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    
    <meta name="description" content="Portal - Bootstrap 5 Admin Dashboard Template For Developers">
    <meta name="author" content="Xiaoying Riley at 3rd Wave Media">    
    <link rel="shortcut icon" href="favicon.ico"> 
    <!-- App CSS -->  
    <link id="theme-style" rel="stylesheet" href="@/assets/css/portal.css">
  </head>
  <body>
  <header class="app-header fixed-top">	   	            
    <div class="app-header-inner">  
      <div class="container-fluid py-2">
        <div class="app-header-content"> 
          <div class="row justify-content-between align-items-center">
          <div class="col-auto">
            <a id="sidepanel-toggler" class="sidepanel-toggler d-inline-block d-xl-none" href="#">
              <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" viewBox="0 0 30 30" role="img"><title>Menu</title><path stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2" d="M4 7h22M4 15h22M4 23h22"></path></svg>
            </a>
          </div><!--//col-->
          <div class="search-mobile-trigger d-sm-none col">
            <svg class="svg-inline--fa fa-magnifying-glass search-mobile-trigger-icon" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="magnifying-glass" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" data-fa-i2svg=""><path fill="currentColor" d="M500.3 443.7l-119.7-119.7c27.22-40.41 40.65-90.9 33.46-144.7C401.8 87.79 326.8 13.32 235.2 1.723C99.01-15.51-15.51 99.01 1.724 235.2c11.6 91.64 86.08 166.7 177.6 178.9c53.8 7.189 104.3-6.236 144.7-33.46l119.7 119.7c15.62 15.62 40.95 15.62 56.57 0C515.9 484.7 515.9 459.3 500.3 443.7zM79.1 208c0-70.58 57.42-128 128-128s128 57.42 128 128c0 70.58-57.42 128-128 128S79.1 278.6 79.1 208z"></path></svg><!-- <i class="search-mobile-trigger-icon fas fa-search"></i> Font Awesome fontawesome.com -->
          </div><!--//col-->
          <div class="app-utilities col-auto">
            <div class="app-utility-item app-user-dropdown dropdown">
              <a class="dropdown-toggle" id="user-dropdown-toggle" data-bs-toggle="dropdown" role="button" aria-expanded="false">
                <img :src="headshot" alt="user profile">
              </a>
              <button class="btn app-btn-primary" @click="login" v-show="!loginUser.uid">Log In</button>
              <button class="btn app-btn-primary" @click="logout" v-show="loginUser.uid">Log Out</button>
            </div><!--//app-user-dropdown--> 
          </div><!--//app-utilities-->
        </div><!--//row-->
        </div><!--//app-header-content-->
      </div><!--//container-fluid-->
    </div><!--//app-header-inner-->
    <div id="app-sidepanel" class="app-sidepanel sidepanel-visible"> 
      <div id="sidepanel-drop" class="sidepanel-drop"></div>
      <div class="sidepanel-inner d-flex flex-column">
        <a href="#" id="sidepanel-close" class="sidepanel-close d-xl-none">×</a>
        <div class="app-branding">
          <a class="app-logo" href="index.html">
            <img style="height:30px" src="/img/NBALogo.png" alt="logo">
            <span class="logo-text">NBA System</span>
          </a>
        </div><!--//app-branding-->  
      <nav id="app-nav-main" class="app-nav app-nav-main flex-grow-1">
        <ul class="app-menu list-unstyled accordion" id="menu-accordion">
          <!-- button one -->
          <li class="nav-item">
              <!--//Bootstrap Icons: https://icons.getbootstrap.com/ -->
            <a class="nav-link" @click="PushPlayer">
              <span class="nav-link-text">Player</span>
            </a><!--//nav-link-->
          </li><!--//nav-item-->
          <!-- button two -->
          <li class="nav-item" v-show="true">
            <!--//Bootstrap Icons: https://icons.getbootstrap.com/ -->
            <a class="nav-link" @click="PushTeam">
              <span class="nav-link-text">Team</span>
            </a><!--//nav-link-->
          </li><!--//nav-item-->
          <!--button three -->
          <li class="nav-item">
            <!--//Bootstrap Icons: https://icons.getbootstrap.com/ -->
            <a class="nav-link" @click="PushUserList" v-if="loginUser.utype">
              <span class="nav-link-text">User List</span>
            </a><!--//nav-link-->
          </li><!--//nav-item-->
        </ul><!--//app-menu-->
      </nav><!--//app-nav-->
      
      </div><!--//sidepanel-inner-->
    </div><!--//app-sidepanel-->
  </header>
  <router-view></router-view>
      
  </body>
  </template>
  
  <script>
  import axios from 'axios'
    export default{
      
      data(){
        return{
          "headshot" : "./img/headshot/stars.jpg",
          "loginUser" : {uid:0}, 
          "errorMsg" : ""
        }
      },
      methods: {  
      login(){
        this.$router.push("/login")
      },
      logout(){
        if(confirm("Do you want to Log Out?")){
          axios.get("/api/user/logout").then(response => {
            console.log(response)
            if(response.data.status == 1){
              //成功退出
              this.loginUser = {uid : 0}
            }else{
              //console.log(response.data.message)
            }
            this.headshot = "./img/headshot/stars.jpg"
            this.getLogin()
            this.$router.push("/")
          })
        }
      },
      getLogin(){
        axios.get("/api/user/getLogining").then(response =>{
          if(response.data.status == 0){
            return
          }
          this.loginUser = response.data.obj
          this.headshot = "/img/headshot/" + response.data.obj.headshot
        })
      },
      PushUserList(){
        this.$router.push("/userList")
      },  
      PushPlayer(){
        this.$router.push({
          path:"/player",
          query:{
            utype:this.loginUser.utype
          }
        })
      },
      PushTeam(){
        this.$router.push({
          path:"/team",
          query:{
            utype:this.loginUser.utype
          }
        })
      }
    },
    mounted(){
      this.getLogin()
    },
    watch:{
      $route(){
        this.getLogin()
      }
    }
    }
  </script>
  <style scoped>
    @import url("@/assets/css/portal.css");
  </style>