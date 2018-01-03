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
      <div class="column col-8 col-mx-auto" id="main_notice">
        <div class="columns">
          <div class="column col-11 col-mx-auto" id="title">请在以下填写您的注册信息</div>
        </div>

        <form class="form-horizontal" style="color: white; font-weight: bold">

          <div class="form-group">
            <div class="col-2">
              <label class="form-label">会员账号：</label>
            </div>
            <div class="col-4">
              <input class="form-input" type="text" v-model="user_info.user_id" placeholder="会员账号">
              <p class="form-input-hint">The name is invalid.</p>
            </div>
            <div class="col-6">
              <label class="form-label tile" >*:账号必须为6-15为数字和字母组合</label>
            </div>
          </div>

          <div class="form-group">
            <div class="col-2">
              <label class="form-label">用户名称：</label>
            </div>
            <div class="col-4">
              <input class="form-input" type="text" v-model="user_info.user_name" placeholder="张三">
              <p class="form-input-hint">The name is invalid.</p>
            </div>
            <div class="col-6">
              <label class="form-label tile">*:姓名必须与提款的银行户口名字一致,否则无法提款</label>
            </div>
          </div>

          <!-- form structure -->
          <div class="form-group">
            <div class="col-2">
              <label class="form-label">登陆密码：</label>
            </div>
            <div class="col-4">
              <input class="form-input" type="password" v-model="user_info.user_password" placeholder="密码">
              <p class="form-input-hint">The name is invalid.</p>
            </div>
            <div class="col-6">
              <label class="form-label tile">*:密码长度要有6-15个字符,必须含有数字和字母组合</label>
            </div>
          </div>

          <div class="form-group">
            <div class="col-2">
              <label class="form-label">确认密码：</label>
            </div>
            <div class="col-4">
              <input class="form-input" type="password" v-model="user_info.user_password" placeholder="密码">
              <p class="form-input-hint">The name is invalid.</p>
            </div>
            <div class="col-6">
              <label class="form-label tile">*:说明</label>
            </div>
          </div>

          <div class="form-group has-success">
            <div class="col-2">
              <label class="form-label">推荐人：</label>
            </div>
            <div class="col-4">
              <input class="form-input" type="text" v-model="user_info.referrer_id" placeholder="888888@qq.com">
              <p class="form-input-hint">The name is invalid.</p>
            </div>
            <div class="col-6">
              <label class="form-label tile">*:email格式：888888@qq.com</label>
            </div>
          </div>

          <!--
          <div class="form-group has-success">
            <div class="col-2">
              <label class="form-label">取款密码：</label>
            </div>
            <div class="col-4">
              <input class="form-input" type="text" v-model="user_info.pay_password" placeholder="取款密码">
              <p class="form-input-hint">The name is invalid.</p>
            </div>
            <div class="col-6">
              <label class="form-label tile">*:说明</label>
            </div>
          </div>
          -->

        </form>

        <div class="columns">
          <div class="columns col-8 col-mr-auto">
            <div class="column col-5 col-ml-auto">
              <button class="btn" @click="_doRegister">确认注册</button>
            </div>

            <div class="column col-4 col-mr-auto">
              <button class="btn btn-primary" style="margin-right: 1.0em" @click="doLogin">用户登录</button>
            </div>

          </div>
        </div>

      </div>
    </div>

    <!-- -->
    <div class="columns col-oneline">
      <div class="column col-12" id="footer">
        {{msg}}
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
    name: 'Register',
    data() {
      return {
        user_info: {
          user_id:"",
          referrer_id: "",
          user_password: "",
          user_name: "",
          real_name: "",
          pay_password: "",
          phone_number: "",
          user_email:""
        }
      }
    },
    methods: {
      _doRegister() {
        this.$axios.post(this.$app.context + '/register', this.user_info)
          .then(response => {
            if (response.data) {
              console.log("register=>", response.data)

              switch (response.data) {
                case 0:
                  alert("注册成功,请登录！");
                  this.$router.push({path: "/"})
                  break
                case -1:
                  alert("数据注册错误");
                  break
                case -2:
                  alert("已经存在此用户");
                  break
                default:
                  break
              }

            } else {
              this.hasError = true
            }
          })
          .catch(error => {
            console.log(error);
          });
      },
      doLogin() {
        this.$router.push({path: "/"})
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

  .form-group {
    margin-bottom: 0em !important;
  }

  #header {
    height: 100px;
    background-image: url("../assets/img/top_0.jpg");
  }

  #notice {
    border-radius: 5em;
    height: 28px;
    color: white;
    background-color: rgba(194, 178, 210, 0.5); /* IE6和部分IE7内核的浏览器(如QQ浏览器)会读懂，但解析为透明 */
  }

  #body_bg {
    height: 600px;
    top: 100px;
    background-repeat: no-repeat;
    /*background-image: url('../assets/img/CSNz-fypatmw5989631.jpg')*/
    background-image: url('../assets/img/ChMkJlbKw0CIDdJjAAH96AIRIG8AALG0AMeCW4AAf4A647.jpg')
  }

  #body_bg #title {
    font-size: 2em;
    /*border-bottom-style: dashed;*/
    color: white;
  }

  #main_notice {
    background-color: white;
    height: 500px;
    margin-top: 3em;
    border-radius: .5em;
    background-color: rgba(194, 178, 210, 0.5);
  }

  #main_notice p {
    margin-bottom: 0em !important;
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
