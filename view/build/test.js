

const axios = require('axios')


function doLogin() {
  axios.get('http://localhost:8081/login')
    .then(function (response) {
      console.log(response.data);
    })
    .catch(function (error) {
      console.log(error);
    });
}

doLogin()
