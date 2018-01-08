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

    <div class="column col-12 " style="height: 2px"></div>

    <div class="columns " id="body_bg">
      <div class="column col-8 col-mx-auto shadow" id="main_register">
        <div class="columns shadow"
             style="margin-top: .5rem;background: #4b48d6;font-weight: bold;background-color: rgba(75,72,214, 0.7)">
          <div class="column col-11 col-mx-auto" id="title">
            <span class="h3" style="font-family: 'Georgia', Tahoma, Sans-Serif">请填写您的注册信息</span>
          </div>
        </div>

        <form class="form-horizontal" style="color: white; font-weight: bold">

          <div class="form-group " v-bind:class="{ 'has-error': page_config.uid_err }">
            <div class="col-2" style="text-align: right">
              <label class="form-label">会员账号：</label>
            </div>
            <div class="col-4">
              <input class="form-input shadow" type="text" v-model="user_info.user_id" placeholder="会员账号">
              <p class="form-input-hint">{{page_config.uid_err_msg}}&nbsp;</p>
            </div>
            <div class="col-6">
              <label class="form-label tile">*:账号必须为6-15为数字和字母组合</label>
            </div>
          </div>

          <div class="form-group" v-bind:class="{ 'has-error': page_config.unm_err }">
            <div class="col-2" style="text-align: right">
              <label class="form-label">用户名称：</label>
            </div>
            <div class="col-4">
              <input class="form-input shadow" type="text" v-model="user_info.user_name" placeholder="张三">
              <p class="form-input-hint">{{page_config.unm_err_msg}}&nbsp;</p>
            </div>
            <div class="col-6">
              <label class="form-label tile">*:姓名必须与提款的银行户口名字一致,否则无法提款</label>
            </div>
          </div>

          <!-- form structure -->
          <div class="form-group" v-bind:class="{ 'has-error': page_config.upw_err }">
            <div class="col-2" style="text-align: right">
              <label class="form-label">登陆密码：</label>
            </div>
            <div class="col-4">
              <input class="form-input shadow" type="password" v-model="user_info.user_password" placeholder="密码">
              <p class="form-input-hint">{{page_config.upw_err_msg}}&nbsp;</p>
            </div>
            <div class="col-6">
              <label class="form-label tile">*:密码长度要有6-15个字符,必须含有数字和字母组合</label>
            </div>
          </div>

          <div class="form-group" v-bind:class="{ 'has-error': page_config.upw_err }">
            <div class="col-2" style="text-align: right">
              <label class="form-label">确认密码：</label>
            </div>
            <div class="col-4">
              <input class="form-input shadow" type="password" v-model="user_info.r_user_password" placeholder="密码">
              <p class="form-input-hint">{{page_config.upw_err_msg}}&nbsp;</p>
            </div>
            <div class="col-6">
              <label class="form-label tile">*:说明,两次密码必须输入相同</label>
            </div>
          </div>

          <div class="form-group " v-bind:class="{ 'has-error': page_config.ureferrer_id_err }">
            <div class="col-2 " style="text-align: right">
              <label class="form-label">推荐人：</label>
            </div>
            <div class="col-4 ">
              <input class="form-input shadow" type="text" v-model="user_info.referrer_id" placeholder="推荐人的账号">
              <p class="form-input-hint">{{page_config.ureferrer_id_err_msg}}&nbsp;</p>
            </div>
            <div class="col-6">
              <label class="form-label tile">*:推荐人的账号</label>
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
              <button class="btn btn-primary" style="margin-right: 1.0em" @click="doLogin">返回登录</button>
            </div>

          </div>
        </div>

      </div>
    </div>

    <div class="modal active " id="modal-notic" v-bind:class="[page_config.is_show_notic]">
      <a href="#close" class="modal-overlay" aria-label="Close"></a>
      <div class="modal-container">
        <div class="modal-header">
          <span href="#" class="btn btn-clear float-right" aria-label="Close" @click="hideNotic"></span>
          <div class="modal-title h5"> 提示</div>
        </div>
        <div class="modal-body">
          <div class="content">
            <!-- content here -->
            <p>{{page_config.show_notic_msg}}</p>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-link btn-sm" @click="hideNotic">确认</button>
        </div>
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
          user_id: "",
          referrer_id: "",
          user_password: "",
          r_user_password: "",
          user_name: "",
          real_name: "",
          pay_password: "",
          phone_number: "",
          user_email: ""
        },
        page_config: {
          uid_err: false,
          uid_err_msg: "",
          unm_err: false,
          unm_err_msg: "",
          upw_err: false,
          upw_err_msg: "",
          ureferrer_id_err: false,
          ureferrer_id_err_msg: "",
          is_show_notic: "isHide",
          show_notic_msg: "",
          show_notic_go: ""
        }
      }
    },
    methods: {
      showNotic(msg, path) {
        this.page_config.is_show_notic = "isShow"
        this.page_config.show_notic_msg = msg
        this.page_config.show_notic_go = path
      },
      hideNotic() {
        this.page_config.is_show_notic = "isHide"
        this.page_config.show_notic_msg = ""
        if (this.page_config.show_notic_go) {
          this.$router.push({path: this.page_config.show_notic_go})
        }
      },
      _doRegister() {
        let isOk = true

        //验证数据
        this.page_config.uid_err = this.user_info.user_id.match(/[a-zA-Z0-9]{6,16}/) && true
        if (!this.page_config.uid_err) {
          this.page_config.uid_err_msg = "用户名不符合要求，请核对信息"
          isOk = false
        } else {
          this.page_config.uid_err = false
          this.page_config.uid_err_msg = ""
        }

        this.page_config.upw_err = (this.user_info.r_user_password == this.user_info.user_password) && true
        if (!this.page_config.upw_err) {
          this.page_config.upw_err_msg = "两次输入的密码不相同，请核对信息"
          isOk = false
        } else {
          this.page_config.upw_err = this.user_info.user_password.match(/[a-zA-Z0-9]{6,16}/) && true
          if (!this.page_config.upw_err) {
            this.page_config.upw_err_msg = "密码长度必须是6位以上,并包含数字和字母"
            isOk = false
          } else {
            this.page_config.upw_err = false
            this.page_config.upw_err_msg = ""
          }
        }

        if (!isOk) {
          console.log("数据不正确")
          return
        }



        this.$axios.post(this.$app.context + '/register', this.user_info)
          .then(response => {
            console.log("register=>", response.data)
            switch (response.data) {
              case 0:
                this.showNotic("注册成功,跳转到登陆界面,请登录", "/")
                break
              case -1:
                this.showNotic("数据错误请联系管理员")
                break
              case -2:
                this.page_config.uid_err = true
                this.page_config.uid_err_msg = "已经存在此用户，请核对信息"
                break
              case -3:
                this.page_config.ureferrer_id_err = true
                this.page_config.ureferrer_id_err_msg = "推荐人账号不存在,请核对信息"
                break
              default:
                break
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
    height: 28px;
    color: white;
    background-color: rgba(194, 178, 210, 0.5); /* IE6和部分IE7内核的浏览器(如QQ浏览器)会读懂，但解析为透明 */
  }

  #body_bg {
    height: 600px;
    top: 100px;
    background-repeat: no-repeat;
    background: url('../assets/img/ChMkJlbKw0CIDdJjAAH96AIRIG8AALG0AMeCW4AAf4A647.jpg') no-repeat -0rem -5rem;
  }

  #body_bg #title {
    font-size: 2em;
    /*border-bottom-style: dashed;*/
    color: white;
  }

  #main_register {
    height: 500px;
    margin-top: 3em;
    border-radius: .5em;
    background-color: rgba(100, 241, 241, 0.3);
  }

  #main_register p {
    margin-bottom: 0em !important;
  }

  .form-input-hint {
    color: red;
  }

  .isShow {
    display: contents;
  }

  .isHide {
    display: none;
  }


</style>
