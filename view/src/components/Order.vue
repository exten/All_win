<template>
  <div class="container" style="ba">
    <div class="columns" id="header" style="height: 6em">
      <div class="column col-3 col-mr-auto" id="logo" style="background-color: blueviolet; line-height: 30px">
        <!--<img src="../assets/img/6c2d6d1289c6.jpg" style="border-radius: 5em; height: 60px"/>-->
        <h3 style="font-size: x-large; color: white">SI HAI YU LE LOG</h3>
      </div>
      <div class="columns col-9 col-mx-auto" id="notice">
        <div class="column col-1 text-center rounded " v-for="(item, index) in lottery_num_result"
             style="height: 5em;width:5em;
             margin-top:0.5em;
             margin-left: 1.6em;
             text-align:center;
             line-height:5.5em;">
          <div class="bg-error" style="height: 1.5em;width:1.5em;font-size: 3.5em;font-weight: bold;color: #2c3e50">
            {{item}}
          </div>
        </div>

      </div>

    </div>

    <div class="columns col-13" style="background-color: #2c3e50" id="nav">

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

    <div class="columns" style="height: 36em;">
      <div class="col-3 col-mr-auto" style="background-color: paleturquoise">
        <div class="column col-9 col-sm-12" style="margin-top: 1em">
          <div class="tile">
            <div class="tile-icon">
              <figure class="avatar avatar-lg">
                <img src="../../node_modules/spectre.css/docs/img/avatar-2.png" alt="Avatar">
              </figure>
            </div>
            <div class="tile-content">
              <p class="tile-title text-bold">{{user_info.user_id}}</p>
              <p class="tile-subtitle text-bold text-gray" style="color: #6e29c0">
                金额: 50000 &nbsp; &nbsp; 积分: {{user_info.calc_integrals}}
              </p>
              <p>
                <button class="btn btn-primary btn-sm">充值</button>
                <button class="btn btn-sm">提现</button>
              </p>
            </div>
          </div>
        </div>

        <div class="columns ">
          <div class="col-10 col-mx-auto">
            <div class="form-group  col-11 col-mx-auto">
              <button class="btn  col-11 ">扩展功能按钮</button>
            </div>
            <div class="form-group  col-11 col-mx-auto">
              <button class="btn  col-11 ">扩展功能按钮</button>
            </div>
            <div class="form-group  col-11 col-mx-auto">
              <button class="btn  col-11 ">扩展功能按钮</button>
            </div>
            <div class="form-group  col-11 col-mx-auto">
              <button class="btn  col-11 ">扩展功能按钮</button>
            </div>
            <li class="breadcrumb-item">
              <a href="#">Change avatar</a>
            </li>
          </div>
        </div>
      </div>


      <div class="col-9 col-ml-auto">
        <div class="col-12 col-mr-auto" style="height: 18em;">
          下注区

          <!-- 第一行下注区 -->
          <div class="columns col-12 col-mx-auto">
            <div class="col-2 rounded card">
              <div class="chip">
                <figure class="avatar avatar-lg" data-initial="1"></figure>
                投注金额:{{lottery_amt_result.location_amt1}}
              </div>
              <div class="col-10 col-mx-auto form-group">
                <input class="form-input input-sm" type="text" v-model="user_order.location_amt1"
                       @keyup="validNumber"
                       @blur="validNumber"
                       placeholder="请输入投注金额">
              </div>
            </div>

            <div class="col-2 rounded card">
              <div class="chip">
                <figure class="avatar avatar-lg" data-initial="2"></figure>
                投注金额:{{lottery_amt_result.location_amt2}}
              </div>
              <div class="col-10 col-mx-auto form-group">
                <input class="form-input input-sm" type="text" v-model="user_order.location_amt2"
                       @keyup="validNumber"
                       @blur="validNumber"
                       placeholder="请输入投注金额">
              </div>
            </div>

            <div class="col-2 rounded card">
              <div class="chip">
                <figure class="avatar avatar-lg" data-initial="3"></figure>
                投注金额:{{lottery_amt_result.location_amt3}}
              </div>
              <div class="col-10 col-mx-auto form-group">
                <input class="form-input input-sm" type="text" v-model="user_order.location_amt3"
                       @keyup="validNumber"
                       @blur="validNumber"
                       placeholder="请输入投注金额">
              </div>
            </div>

            <div class="col-2 rounded card">
              <div class="chip">
                <figure class="avatar avatar-lg" data-initial="4"></figure>
                投注金额:{{lottery_amt_result.location_amt4}}
              </div>
              <div class="col-10 col-mx-auto form-group">
                <input class="form-input input-sm" type="text" v-model="user_order.location_amt4"
                       @keyup="validNumber"
                       @blur="validNumber"
                       placeholder="请输入投注金额">
              </div>
            </div>

            <div class="col-2 rounded card">
              <div class="chip">
                <figure class="avatar avatar-lg" data-initial="5"></figure>
                投注金额:{{lottery_amt_result.location_amt5}}
              </div>
              <div class="col-10 col-mx-auto form-group">
                <input class="form-input input-sm" type="text" v-model="user_order.location_amt5"
                       @keyup="validNumber"
                       @blur="validNumber"
                       placeholder="请输入投注金额">
              </div>
            </div>
          </div>


          <!-- 第二行下注区 -->
          <div class="columns col-12 col-mx-auto">
            <div class="col-2 rounded card">
              <div class="chip">
                <figure class="avatar avatar-lg" data-initial="6"></figure>
                投注金额:{{lottery_amt_result.location_amt6}}
              </div>
              <div class="col-10 col-mx-auto form-group">
                <input class="form-input input-sm" type="text" v-model="user_order.location_amt6"
                       @keyup="validNumber"
                       @blur="validNumber"
                       placeholder="请输入投注金额">
              </div>
            </div>

            <div class="col-2 rounded card">
              <div class="chip">
                <figure class="avatar avatar-lg" data-initial="7"></figure>
                投注金额:{{lottery_amt_result.location_amt7}}
              </div>
              <div class="col-10 col-mx-auto form-group">
                <input class="form-input input-sm" type="text" v-model="user_order.location_amt7"
                       @keyup="validNumber"
                       @blur="validNumber"
                       placeholder="请输入投注金额">
              </div>
            </div>

            <div class="col-2 rounded card">
              <div class="chip">
                <figure class="avatar avatar-lg" data-initial="8"></figure>
                投注金额:{{lottery_amt_result.location_amt8}}
              </div>
              <div class="col-10 col-mx-auto form-group">
                <input class="form-input input-sm" type="text" v-model="user_order.location_amt8"
                       @keyup="validNumber"
                       @blur="validNumber"
                       placeholder="请输入投注金额">
              </div>
            </div>

            <div class="col-2 rounded card">
              <div class="chip">
                <figure class="avatar avatar-lg" data-initial="9"></figure>
                投注金额:{{lottery_amt_result.location_amt9}}
              </div>
              <div class="col-10 col-mx-auto form-group">
                <input class="form-input input-sm" type="text" v-model="user_order.location_amt9"
                       @keyup="validNumber"
                       @blur="validNumber"
                       placeholder="请输入投注金额">
              </div>
            </div>

            <div class="col-2 rounded card">
              <div class="chip">
                <figure class="avatar avatar-lg" data-initial="10"></figure>
                投注金额:{{lottery_amt_result.location_amt10}}
              </div>
              <div class="col-10 col-mx-auto form-group">
                <input class="form-input input-sm" type="text" v-model="user_order.location_amt10"
                       @keyup="validNumber"
                       placeholder="请输入投注金额">
              </div>
            </div>
          </div>

          <!-- 下注汇总提交区 -->
          <div class="columns" style="margin-top: .8em; ">
            <div class="column col-4 col-ml-auto">
              <div class="input-group">
                <span class="input-group-addon addon text-bold" style="width: 10rem;overflow: auto">
                  总投注金额:
                  {{
                  parseInt(this.user_order.location_amt1 || "0", 0)+
                  parseInt(this.user_order.location_amt2 || "0", 0)+
                  parseInt(this.user_order.location_amt3 || "0", 0)+
                  parseInt(this.user_order.location_amt4 || "0", 0)+
                  parseInt(this.user_order.location_amt5 || "0", 0)+
                  parseInt(this.user_order.location_amt6 || "0", 0)+
                  parseInt(this.user_order.location_amt7 || "0", 0)+
                  parseInt(this.user_order.location_amt8 || "0", 0)+
                  parseInt(this.user_order.location_amt9 || "0", 0)+
                  parseInt(this.user_order.location_amt10 || "0", 0)
                  }}
                </span>
                <button class="btn input-group-btn btn" @click="userOrderpay">确认购买</button>
              </div>
            </div>
          </div>

        </div>


        <div class="col-12 col-mx-auto" style="height: 18em;margin-left: .5em">
          开奖区

          <table class="table table-striped table-hover">
            <thead>
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
    </div>


  </div>
</template>

<script>
  export default {
    name: 'Order',
    data() {
      return {
        //用户当前这期投注金额
        user_order: {
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
        },
        //当前这期的总投注金额
        lottery_amt_result: {
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
        },
        //当前这期的开奖号码
        lottery_num_result: [
          2,
          3,
          1,
          5,
          8,
          4,
          6,
          7,
          10,
          9
        ],
        //用户当天的中奖纪录
        list_lottery_result: [{
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
        }],
        //用户信息
        user_info: {
          user_id: "The S.H.I.E.L.D.",
          user_name: "The S.H.I.E.L.D.",
          integrals: "7889",
          calc_integrals: 0,
          trial_integrals: "100"
        },
        total_amt: 0
      }
    },
    mounted() {

      this.user_info.calc_integrals = this.user_info.integrals

      console.log("页面加载完成", this.lottery_amt_result, this)

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
      //用户信息
      findUserInfo() {
        console.log("页面加载完成", "this findUserInfo")
        this.$axios.post(this.$app.context + '/findUserInfo', this.user)
          .then(response => {
            if (response.data) {
              console.log("页面加载完成", "this findUserInfo", response.data)
              if (response.data == -1) {
                alert("登录已经过期请重新登录");
                this.$router.push({path: "Login"})
              } else {
                this.user_info = response.data
              }

            } else {
              this.hasError = true
            }
          })
          .catch(error => {
            alert("登录已经过期请重新登录");
            this.$router.push({path: "/"})
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
          alert("余额不足,请充值");
        }

        this.$axios.post(this.$app.context + '/userOrderpay', this.user_order)
          .then(response => {
            if (response.data) {
              console.log(response.data)

              switch (response.data) {
                case 0:
                  alert("对不起,您余额不足,请充值");
                  break
                case -1:
                  alert("登录已经过期请重新登录,请重新登录!");
                  this.$router.push({path: "/"})
                  break
                case -2:
                  alert("未到下单时间,请稍后购买!");
                  break
                default:
                  alert("恭喜您,下注成功！");
                  break
              }

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
        } else {
          this.user_info.calc_integrals = (this.user_info.integrals - sumAmt)
        }
      }

    },
    filters: {
      ft_total_amt: function (value) {
        return value;
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
    background-color: azure;
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
    height: 10.7rem;
    width: 53.2rem;
    overflow-y: scroll;
  }

  table thead, tbody tr {
    display: table;
    width: 100%;
    table-layout: fixed;
  }

  table thead {
    width: calc(100% - 1em)
  }

  table thead th {
    background: #ccc;
  }
</style>
