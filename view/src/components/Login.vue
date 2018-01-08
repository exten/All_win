<template>

  <div class="container">

    <div class="columns" id="header">

      <div class="column col-12 col-mx-auto" id="notice">
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
    </div>

    <div class="column col-12" style="height: 2px"></div>

    <div class="columns " id="body_bg">
      <div class="column col-3 col-mx-auto shadow" style="width: 20rem; border-radius: 0.2em;" id="login">
        <div style="margin:.5rem ;color: white" on-click="_doLogin">
          <span class="h3" style="font-family: 'Georgia', Tahoma, Sans-Serif">用户登录</span>
        </div>

        <div class="form-group " v-bind:class="{ 'has-error': config.hasError }">
          <div class="has-icon-left shadow">
            <input type="text" class="form-input " placeholder="请输入您的账号" v-model="user.user_id">
            <i class="form-icon fa fa-user"></i>
          </div>
          <p class="form-input-hint">{{config.hasMsg}}&nbsp;</p>
        </div>

        <div class="form-group " v-bind:class="{ 'has-error': config.hasError }">
          <div class="has-icon-left shadow">
            <input type="password" class="form-input" placeholder="请输入您的密码" v-model="user.user_password">
            <i class="form-icon fa fa-key fa-fw"></i>
          </div>
          <p class="form-input-hint">{{config.hasMsg}}&nbsp;</p>
        </div>

        <div class="form-group ">
          <div class="col-12 ">
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


    </div>
  </div>

</template>


<script>
  export default {
    name: 'Login',
    data() {
      return {
        user: {
          user_id: "",
          user_password: ""
        },
        config: {
          hasError: false,
          hasMsg: ""
        }

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
              this.config.hasError = true
              this.config.hasMsg = "账号或者密码错误，请核对后登陆"
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

  .form-group {
    margin: 0rem 2rem;
  }

  #header {
    height: 100px;
    background-image: url("../assets/img/top_0.jpg");
  }

  #login {
    height: 300px;
    margin-top: 7em;
    /*background-color: rgba(255, 241, 241, 0.9);*/
    background-color: rgba(100, 241, 241, 0.3);
  }

  #notice {
    /*border-radius: 5em;*/
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
    /*background-image: url('../assets/img/ChMkJlbKw0CIDdJjAAH96AIRIG8AALG0AMeCW4AAf4A647.jpg')*/
    background: url('../assets/img/ChMkJlbKw0CIDdJjAAH96AIRIG8AALG0AMeCW4AAf4A647.jpg') no-repeat -0rem -5rem;
  }

</style>
