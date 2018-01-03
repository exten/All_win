<template>
  <div class="container">
    <div class="columns" id="header">
      <div class="column col-3 col-mr-auto" id="logo" style="background-color: blueviolet; line-height: 30px">
        <!--<img src="../assets/img/6c2d6d1289c6.jpg" style="border-radius: 5em; height: 60px"/>-->
        <h3 style="font-size: x-large; color: white">SI HAI YU LE LOG</h3>
      </div>
      <div class="column col-10 col-mx-auto" id="notice">
        <div class="column col-1 col-mr-auto">
          <span style="position: absolute;margin-left: -26px; margin-top: 3px">系统公告:</span>
        </div>
        <marquee direction=left
                 style="height: 28px;margin-left: 70px; margin-top: 3px; width: 88%; ">
          用户登录:Add a container element with the menu class. And add child elements with the menu-item class.
          You can separate menu items with a divider. Spectre.css does not include JavaScript code, you will need to
          implement your JS to interact with the menus.
        </marquee>
      </div>
      
      <div class="column col-8 col-ml-auto" id="nav">
        
        <ul class="breadcrumb">
          <li class="breadcrumb-item">
            <a href="#">Home</a>
          </li>
          <li class="breadcrumb-item">
            <a href="#">Settings</a>
          </li>
          <li class="breadcrumb-item">
            <a href="#">Change avatar</a>
          </li>
          <li class="breadcrumb-item">
            <a href="#">Change avatar</a>
          </li>
          <li class="breadcrumb-item">
            <a href="#">Change avatar</a>
          </li>
          <li class="">
            <a href="#"></a>
          </li>
        </ul>
      
      </div>
    
    </div>
    
    <div class="column col-12" style="height: 2px"></div>
    
    <div class="columns" id="body_bg">
      <div class="column col-2 col-mx-auto" style="width: 270px; border-radius: 0.2em;" id="login">
        <div style="margin-top: 10px" on-click="_doLogin">
          <span class="h3">用户登录</span>
        </div>
        
        <div class="form-group ">
          <div class="has-icon-left">
            <input type="text" class="form-input " placeholder="请输入您的账号" v-model="user.user_id">
            <i class="form-icon icon icon-people"></i>
          </div>
          <p class="form-input-hint">The name is invalid.</p>
        </div>
        
        <div class="form-group " v-bind:class="{ 'has-error': hasError }">
          <div class="has-icon-left">
            <input type="password" class="form-input" placeholder="请输入您的密码" v-model="user.user_password">
            <i class="form-icon icon icon-mail"></i>
          </div>
          <p class="form-input-hint">账号或者密码错误，请您再次核对.</p>
        </div>
        
        <div class="form-group ">
          <div class="col-12">
            <button class="btn" style="width: 110px; margin-left: -1px" @click="_doLogin">
              <span class="h6">登&nbsp;&nbsp;&nbsp;&nbsp;陆</span>
            </button>
            <button class="btn btn-primary" style="width: 115px; margin-left: 23px"
                    @click="_doRegister">
              <span class="h6">注&nbsp;&nbsp;&nbsp;&nbsp;册</span>
            </button>
          </div>
        </div>
        
        <div class="form-group" style="margin-top: 1em">
          <div class="btn-group-block">
            <button class="btn" style="width:252px; border-radius: 5em">游客参观</button>
          </div>
        </div>
      
      </div>
      
      <div class="column col-6 col-mx-auto" id="main_notice">
        <span>系统公告</span>
      </div>
    
    </div>
    
    <!-- -->
    <div class="columns col-oneline">
      <div class="column col-12" id="footer">
        <!--<div class="column col-8">CopyRight 2002-2025 永利国际娱乐 All Rights Reserved{{msg}}</div>-->
        <!--&#59906;-->
        <span class="iconfont">&#59907;</span>
        <span class="iconfont">&#59906;</span>
        <span class="iconfont">&#60096;</span>
        <span class="iconfont">&#60080;</span>
        <span class="iconfont">&#60048;</span>
      
      
      </div>
    </div>
  
  
  </div>

</template>


<script>
  export default {
    name: 'Login',
    data() {
      return {
        user:{
          user_id:"",
          user_password:""
        },
        hasError: false
      }
    },
    methods: {
      _doLogin() {
        this.msg = "_doLogin";
        this.$axios.post(this.$app.context + '/login', this.user)
          .then(response => {
            if (response.data) {
              this.$router.push({path: "Order"})
            } else {
              this.hasError = true
            }
          })
          .catch(error => {
            console.log(error);
          });
      },
      _doRegister() {
        this.$router.push({path: "Register"})
      }
    }
  }
</script>


<style scoped>
  
  .form-input-hint {
    margin-bottom: 0em !important;
  }
  
  .form-group:not(:last-child) {
    margin-bottom: .0rem !important;
  }
  
  #header {
    height: 100px;
    background-image: url("../assets/img/top_0.jpg");
  }
  
  #login {
    background-color: gainsboro;
    height: 300px;
    margin-top: 7em;
    /*color: white;*/
    background-color: rgba(255, 241, 241, 0.6);
  }
  
  #notice {
    border-radius: 5em;
    height: 28px;
    /*opacity: 0.65;*/
    /*filter: Alpha(opacity=50);*/
    color: white;
    background-color: rgba(194, 178, 210, 0.5); /* IE6和部分IE7内核的浏览器(如QQ浏览器)会读懂，但解析为透明 */
  }
  
  #body_bg {
    height: 600px;
    top: 100px;
    /*background-color: red;*/
    /*background-size: 100% auto;*/
    background-repeat: no-repeat;
    /*background-image: url('../assets/img/CSNz-fypatmw5989631.jpg')*/
    background-image: url('../assets/img/ChMkJlbKw0CIDdJjAAH96AIRIG8AALG0AMeCW4AAf4A647.jpg')
  }
  
  #main_notice {
    /*background-color: aquamarine;*/
    margin-top: 3em;
    height: 500px;
    border-radius: .5em;
    background-color: rgba(194, 178, 210, 0.5);
  }
  
  #footer {
    background-color: cadetblue;
    height: 100px;
  }
  
  li:hover {
    /*background-color: red;*/
  }
  
  .iconfont {
    font-family: "w-iconfont" !important;
    font-size: 46px;
    font-style: normal;
    -webkit-font-smoothing: antialiased;
    -webkit-text-stroke-width: 0.2px;
    -moz-osx-font-smoothing: grayscale;
    color: #6e29c0;
  }

</style>
