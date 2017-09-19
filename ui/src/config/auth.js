const SERVER_URL = 'http://localhost:8080'
const LOGIN_URL = SERVER_URL + '/login'

export default {
  data: {
    authenticated: false
  },
  login (context, user) {
    context.$http.post(LOGIN_URL, user).then((result) => {
      console.log(result)
      localStorage.setItem('token', result.data.data)
      this.authenticated = true
      this.$router.push('index')
    }, (result) => {
      console.log(result + ',' + result.data.message)
      context.error = result.data.message
    })
  },
  getAuthHeader () {
    return {
      'Authorization': 'Bearer ' + localStorage.getItem('token')
    }
  },
  isAuth () {
    let token = localStorage.getItem('token')
    if (token) {
      this.authenticated = true
    } else {
      this.authenticated = false
    }
  }
}
