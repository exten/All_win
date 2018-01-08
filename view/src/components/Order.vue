<template>
  <div class="container " style="">

    <div class="columns m-0 shadow" id="header" style="padding-top:.3rem;height: 6em;">

      <!-- 用户信息区域 -->
      <div class="column col-2 col-mr-auto shadow bg-secondary">
        <div class="tile">
          <div class="tile-icon">
            <div class="example-tile-icon">
              <figure class="avatar avatar-lg">
                <img src="../../node_modules/spectre.css/docs/img/avatar-2.png" alt="Avatar">
              </figure>
            </div>
          </div>
          <div class="tile-content">
            <p class="tile-title">{{user_info.user_id}}</p>
            <p class="tile-subtitle text-gray">积分: {{user_info.calc_integrals}}</p>
          </div>
          <div class="tile-action" style="width: 2rem;">
            <button class="btn btn-primary btn-sm" style="padding-bottom: 0.1rem" @click="showStuff('充值')">充值</button>
            <button class="btn btn-sm" @click="showStuff('提现')">提现</button>
            <button class="btn btn-link btn-sm" @click="loginOut">注销</button>
          </div>
        </div>
      </div>


      <div class="col-1 " style="color: #e06870;" v-for="(item, index) in lottery_num_result">
        <!--<img v-bind:src="['../assets/num_png/bullet-'+item+'.png']" style="width: 5rem">-->
        <figure class="avatar avatar-xl bold shadow" v-bind:data-initial="[item]"></figure>
      </div>
    </div>


    <!-- 第1行下注区 -->
    <div class="columns bg-gray">
      <div class="col-xs-2 rounded card shadow m-1">
        <div class="chip shadow">
          <figure class="avatar avatar-lg shadow" data-initial="1"></figure>
          投注金额:{{lottery_amt_result.location_amt1}}
        </div>
        <div class="col-10 col-mx-auto form-group ">
          <input class="form-input input-sm shadow" type="text" v-model="user_order.location_amt1"
                 @keyup="validNumber"
                 @blur="validNumber"
                 placeholder="请输入投注金额">
        </div>
      </div>
      <div class="column col-lg-1"></div>
      <div class="col-xs-2 rounded card shadow m-1">
        <div class="chip shadow">
          <figure class="avatar avatar-lg shadow" data-initial="2"></figure>
          投注金额:{{lottery_amt_result.location_amt2}}
        </div>
        <div class="col-10 col-mx-auto form-group ">
          <input class="form-input input-sm shadow" type="text" v-model="user_order.location_amt2"
                 @keyup="validNumber"
                 @blur="validNumber"
                 placeholder="请输入投注金额">
        </div>
      </div>
      <div class="column col-lg-1"></div>
      <div class="col-xs-2 rounded card shadow m-1">
        <div class="chip shadow">
          <figure class="avatar avatar-lg shadow" data-initial="3"></figure>
          投注金额:{{lottery_amt_result.location_amt3}}
        </div>
        <div class="col-10 col-mx-auto form-group ">
          <input class="form-input input-sm shadow" type="text" v-model="user_order.location_amt3"
                 @keyup="validNumber"
                 @blur="validNumber"
                 placeholder="请输入投注金额">
        </div>
      </div>
      <div class="column col-lg-1"></div>
      <div class="col-xs-2 rounded card shadow m-1">
        <div class="chip shadow">
          <figure class="avatar avatar-lg shadow" data-initial="4"></figure>
          投注金额:{{lottery_amt_result.location_amt4}}
        </div>
        <div class="col-10 col-mx-auto form-group ">
          <input class="form-input input-sm shadow" type="text" v-model="user_order.location_amt4"
                 @keyup="validNumber"
                 @blur="validNumber"
                 placeholder="请输入投注金额">
        </div>
      </div>
      <div class="column col-lg-1"></div>
      <div class="col-xs-2 rounded card shadow m-1">
        <div class="chip shadow">
          <figure class="avatar avatar-lg shadow" data-initial="5"></figure>
          投注金额:{{lottery_amt_result.location_amt5}}
        </div>
        <div class="col-10 col-mx-auto form-group ">
          <input class="form-input input-sm shadow" type="text" v-model="user_order.location_amt5"
                 @keyup="validNumber"
                 @blur="validNumber"
                 placeholder="请输入投注金额">
        </div>
      </div>
    </div>
    <!-- 第二行下注区 -->
    <div class="columns bg-gray">
      <div class="col-xs-2 rounded card shadow m-1">
        <div class="chip shadow">
          <figure class="avatar avatar-lg shadow" data-initial="6"></figure>
          投注金额:{{lottery_amt_result.location_amt6}}
        </div>
        <div class="col-10 col-mx-auto form-group ">
          <input class="form-input input-sm shadow" type="text" v-model="user_order.location_amt6"
                 @keyup="validNumber"
                 @blur="validNumber"
                 placeholder="请输入投注金额">
        </div>
      </div>
      <div class="column col-lg-1"></div>
      <div class="col-xs-2 rounded card shadow m-1">
        <div class="chip shadow">
          <figure class="avatar avatar-lg shadow" data-initial="7"></figure>
          投注金额:{{lottery_amt_result.location_amt7}}
        </div>
        <div class="col-10 col-mx-auto form-group ">
          <input class="form-input input-sm shadow" type="text" v-model="user_order.location_amt7"
                 @keyup="validNumber"
                 @blur="validNumber"
                 placeholder="请输入投注金额">
        </div>
      </div>
      <div class="column col-lg-1"></div>
      <div class="col-xs-2 rounded card shadow m-1">
        <div class="chip shadow">
          <figure class="avatar avatar-lg shadow" data-initial="8"></figure>
          投注金额:{{lottery_amt_result.location_amt8}}
        </div>
        <div class="col-10 col-mx-auto form-group ">
          <input class="form-input input-sm shadow" type="text" v-model="user_order.location_amt8"
                 @keyup="validNumber"
                 @blur="validNumber"
                 placeholder="请输入投注金额">
        </div>
      </div>
      <div class="column col-lg-1"></div>
      <div class="col-xs-2 rounded card shadow m-1">
        <div class="chip shadow">
          <figure class="avatar avatar-lg shadow" data-initial="9"></figure>
          投注金额:{{lottery_amt_result.location_amt9}}
        </div>
        <div class="col-10 col-mx-auto form-group ">
          <input class="form-input input-sm shadow" type="text" v-model="user_order.location_amt9"
                 @keyup="validNumber"
                 @blur="validNumber"
                 placeholder="请输入投注金额">
        </div>
      </div>
      <div class="column col-lg-1"></div>
      <div class="col-xs-2 rounded card shadow m-1">
        <div class="chip shadow">
          <figure class="avatar avatar-lg shadow" data-initial="10"></figure>
          投注金额:{{lottery_amt_result.location_amt10}}
        </div>
        <div class="col-10 col-mx-auto form-group ">
          <input class="form-input input-sm shadow" type="text" v-model="user_order.location_amt10"
                 @keyup="validNumber"
                 @blur="validNumber"
                 placeholder="请输入投注金额">
        </div>
      </div>
    </div>


    <div class="columns m-2">
      <div class="column col-mx-auto col-4 ">
        <div class="input-group shadow">
          <span class="input-group-addon addon-lg">总投注金额</span>
          <input type="text" class="form-input input-lg col-sm-1" readonly="readonly" placeholder="site name"
                 v-model="total_amt">
          <button class="btn btn-primary input-group-btn btn-lg" @click="userOrderpay">确认购买</button>
        </div>
      </div>
    </div>


    <div class="columns " style="height: 33rem;">
      <div class="column col-mx-auto col-12">
        <table class="table table-striped table-hover shadow ">
          <thead class="bg-secondary">
          <tr>
            <th>开奖期数</th>
            <th>1号投注</th>
            <th>2号投注</th>
            <th>3号投注</th>
            <th>4号投注</th>
            <th>5号投注</th>
            <th>6号投注</th>
            <th>7号投注</th>
            <th>8号投注</th>
            <th>9号投注</th>
            <th>10号投注</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="item in list_lottery_result">
            <td>{{item.lottery_period}}期</td>
            <td>{{item.location_amt1}}</td>
            <td>{{item.location_amt2}}</td>
            <td>{{item.location_amt3}}</td>
            <td>{{item.location_amt4}}</td>
            <td>{{item.location_amt5}}</td>
            <td>{{item.location_amt6}}</td>
            <td>{{item.location_amt7}}</td>
            <td>{{item.location_amt8}}</td>
            <td>{{item.location_amt9}}</td>
            <td>{{item.location_amt10}}</td>
          </tr>
          </tbody>
        </table>
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
            <p>{{page_config.msg}}</p>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-link btn-sm" @click="hideNotic">确认</button>
        </div>
      </div>
    </div>


    <div class="modal active " id="modal-stuff" v-bind:class="[page_config.is_show_stuff]">
      <a href="#close" class="modal-overlay" aria-label="Close"></a>
      <div class="modal-container">
        <div class="modal-header">
          <span class="btn btn-clear float-right" aria-label="Close" @click="hideStuff"></span>

          <div class="modal-title h5"> 提示</div>
        </div>
        <div class="modal-body">
          <div class="content">
            <!-- content here -->
            <div class="columns m-1">
              <div class="column col-mx-auto col-12 ">
                <div class="input-group shadow">
                  <span class="input-group-addon addon-lg">{{page_config.title}}金额</span>
                  <input type="text" class="form-input input-lg col-sm-1" readonly="readonly" placeholder="site name"
                         v-model="total_amt">
                  <button class="btn btn-primary input-group-btn btn-lg" @click="userOrderpay">确认{{page_config.title}}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-link btn-sm" @click="hideStuff"></button>
        </div>
      </div>
    </div>

  </div>


</template>

<script>
  export default {
    name: 'Order',
    data() {
      return {
        //用户当前这期投注金额
        user_order: {},
        //当前这期的总投注金额
        lottery_amt_result: {},
        //当前这期的开奖号码
        lottery_num_result: [],
        //用户当天的中奖纪录
        list_lottery_result: [],
        //用户信息
        user_info: {},
        total_amt: 0,
        page_config: {
          title: "提示",
          msg: "",
          go_path: "",
          is_show_notic: "isHide",
          is_show_stuff: "isHide"
        }
      }
    },
    mounted() {

      this.user_info.calc_integrals = this.user_info.integrals

      console.log("页面加载完成", this.lottery_amt_result, this)

      this.testInitData()
      //调用本组件的方法

      //1.加载用户数据
      this.findUserInfo()

      //2.调用查询开奖记录 lottery_num_result
      this.findCurrNumInfo()

      //3.当期总下注金额 lottery_amt_result
      this.findOrderTotalInfo()


      //3.加载开奖记录列表 list_lottery_result
      this.findOrderUserPayListInfo()


    },

    watch: {
      // 如果 `question` 发生改变，这个函数就会运行
      lottery_num_result: function () {
        console.log("watch lottery_num_result change...")
        this.findCurrNumInfo()
      }
    },

    methods: {
      testInitData() {
        for (let i = 0; i < 29; i++) {
          let o = {
            lottery_period: 1089996,
            location_amt1: 200.0,
            location_amt2: 2.0,
            location_amt3: 200.0,
            location_amt4: 2.0,
            location_amt5: 202.0,
            location_amt6: 62.0,
            location_amt7: 92.0,
            location_amt8: 2.0,
            location_amt9: 2.0,
            location_amt10: 2.0
          }

          this.list_lottery_result.push(o)
        }

        this.lottery_num_result = [
          1,
          2,
          3,
          4,
          5,
          6,
          7,
          8,
          9,
          10
        ];

        this.user_order = {
          location_amt1: 1110.0,
          location_amt2: 2.0,
          location_amt3: "",
          location_amt4: "",
          location_amt5: "",
          location_amt6: "",
          location_amt7: 0.0,
          location_amt8: "",
          location_amt9: "",
          location_amt10: 0.0
        };

        this.lottery_amt_result = {
          location_amt1: 200.0,
          location_amt2: 2.0,
          location_amt3: 200.0,
          location_amt4: 2.0,
          location_amt5: 202.0,
          location_amt6: 62.0,
          location_amt7: 92.0,
          location_amt8: 2.0,
          location_amt9: 2.0,
          location_amt10: 2.0
        };

        this.user_info = {
          user_id: "The S.H.I.E.L.D 12345678901",
          user_name: "The S.H.I.E.L.D.",
          integrals: "7889",
          calc_integrals: 7889,
          trial_integrals: "100"
        }

      },
      goToLogin() {
        this.$router.push({path: "Login"})
        //this.$router.push({path: "/"})
      },

      showNotic(msg, path) {
        this.page_config.is_show_notic = "isShow"
        this.page_config.msg = msg
        this.page_config.go_path = path
      },
      hideNotic() {
        this.page_config.is_show_notic = "isHide"
        if (this.page_config.go_path.length > 0) {
          this.$router.push({path: this.page_config.go_path})
        }
      },
      showStuff(title) {
        this.page_config.title = title
        this.page_config.is_show_stuff = "isShow"
      },
      hideStuff() {
        this.page_config.is_show_stuff = "isHide"
      },

      loginOut(){
        this.$axios.post(this.$app.context + '/loginOut',"")
          .then(response => {
            console.log(response)
            this.$router.push({path: "/"})
          })
          .catch(error => {
            console.log("注销出现异常.", error);
          });
      },

      //用户信息
      findUserInfo() {
        console.log("页面加载完成", "this findUserInfo")
        this.$axios.post(this.$app.context + '/findUserInfo', this.user)
          .then(response => {
            if (response.data == -1) {
              this.showNotic("登录已经过期请重新登录","/")
            } else {
              this.user_info = response.data
              this.user_info.calc_integrals = this.user_info.integrals
            }
          })
          .catch(error => {
            this.showNotic("登录已经过期请重新登录","/")
            console.log(error);
          });
      },
      //用户下注金额
      userOrderpay() {
        console.log("页面加载完成", "this userOrderpay")
        let sumAmt = 0

        for (let n in this.user_order) {
          let value = this.user_order[n].toString().replace(/[^0-9-]+/, '')

          this.user_order[n] = value
          sumAmt += Number(value) * 1
        }
        if (sumAmt <= 0) {
          this.showNotic("余额不足,请充值")
          return
        }

        this.$axios.post(this.$app.context + '/userOrderpay', this.user_order)
          .then(response => {
            if (response.data) {
              console.log(response.data)
              let msgRev = ""
              let goPath = ""

              switch (response.data) {
                case 0:
                  msgRev = "对不起,您余额不足,请充值";
                  break
                case -1:
                  msgRev = "登录已经过期请重新登录,请重新登录!";
                  goPath = "/"
                  break
                case -2:
                  msgRev = "未到下单时间,请稍后购买!";
                  break
                default:
                  msgRev = "恭喜您,下注成功！";
                  break
              }

              this.showNotic(msgRev, goPath)

              for (let i = 1; i < 11; i++) {
                this.user_order["location_amt" + i] = 0
              }

            } else {
              this.hasError = true
            }
          })
          .catch(error => {
            console.log(error);
          });
      },
      //当期开奖号码
      findCurrNumInfo() {
        console.log("页面加载完成", "this findCurrNumInfo")
        this.$axios.get(this.$app.context + '/findCurrNumInfo', this.user)
          .then(response => {
            if (response.data) {
              console.log("页面加载完成", "this findCurrNumInfo", this, response.data)
              this.lottery_num_result = response.data
              this.lottery_num_result.date = new Date()
            } else {
              this.hasError = true
            }
          })
          .catch(error => {
            console.log(error);
          });
      },

      //当期用户下注金额
      findOrderUserPayInfo() {
        console.log("页面加载完成", "this findOrderUserPayInfo")
        this.$axios.post(this.$app.context + '/findOrderUserPayInfo', this.user)
          .then(response => {
            if (response.data) {
              this.lottery_num_result = response.data
            } else {
              this.hasError = true
            }
          })
          .catch(error => {
            console.log(error);
          });
      },

      //总下注金额
      findOrderTotalInfo() {
        console.log("页面加载完成", "this findOrderTotalInfo")
        this.$axios.post(this.$app.context + '/findOrderTotalInfo', this.user)
          .then(response => {
            if (response.data) {
              this.lottery_amt_result = response.data
            } else {
              this.hasError = true
            }
          })
          .catch(error => {
            console.log(error);
          });
      },
      //用户当天下注历史
      findOrderUserPayListInfo() {
        console.log("页面加载完成", "this pay_order_money")
        this.$axios.get(this.$app.context + '/findOrderUserPayListInfo', this.user)
          .then(response => {
            if (response.data) {
              console.log("页面加载完成->list_lottery_result", response.data)
              this.list_lottery_result = response.data
            } else {
              this.hasError = true
            }
          })
          .catch(error => {
            console.log(error);
          });
      },

      validNumber() {
        let sumAmt = 0;

        for (let n in this.user_order) {
          let value = this.user_order[n].toString().replace(/[^0-9-]+/, '')

          this.user_order[n] = value
          sumAmt += Number(value) * 1
        }

        if (this.user_info.integrals < sumAmt) {
          //alert("余额不足请充值");
          this.user_info.calc_integrals = 0
          this.showNotic("余额不足请充值")
        } else {
          this.user_info.calc_integrals = (this.user_info.integrals - sumAmt)
          this.total_amt = sumAmt
        }
      }

    },
    filters: {
      ft_total_amt: function (value) {
        return value;
      },
      filter_img: function (value) {
        return "../assets/num_png/bullet-" + value + ".png"
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

  tbody tr:hover {
    /*font-style: italic;*/
    /*font-size: x-large;*/
    /*font-weight: bolder;*/
  }

  .avatar-lg {
    font-size: 1.8em;
    font-weight: bolder;
  }

  .chip {
    height: 2em;
    font-style: italic;
    font-size: 1em;
    font-weight: bolder;
    color: red;
    background-color: mistyrose;
  }

  .card {
    width: 12.9em;
    height: 6.5em;
    margin-left: .5em;
    margin-top: .5em;
    /*background-color: azure;*/
  }

  .form-group {
    padding-top: 1em
  }

  .form-input {
    font-weight: bolder;
    text-align: right;
  }

  /*以下是table body 添加滚动条*/
  table tbody {
    display: block;
    height: 30rem;
    width: 100%;
    overflow-y: scroll;
  }

  table thead, tbody tr {
    display: table;
    width: 100%;
    table-layout: fixed;
  }

  table thead, tbody {
    width: calc(100% - 0em)
  }

  .avatar.avatar-xl {
    font-size: 4.6rem !important;
    font-weight: bold;
    height: 4.2rem !important;
    width: 4.2rem !important;
  }

  .isShow {
    display: contents;
  }

  .isHide {
    display: none;
  }


</style>
